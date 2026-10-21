// 五笔练习：主菜单 + 学习（规则/对照表/背字根）+ 练习。
package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"c1device"
	"c1wubi/wubi"
)

//go:embed assets/pkg-font.bin
var fontData []byte

type view int

const (
	viewMenu view = iota
	viewRules
	viewChart
	viewLearn
	viewTable
	viewRoots
	viewChars
)

type menuItem struct {
	title string
	to    view
}

var menuItems = []menuItem{
	{"1 打字规则", viewRules},
	{"2 字根对照表", viewChart},
	{"3 背字根", viewLearn},
	{"4 字根表练习", viewTable},
	{"5 无表认键", viewRoots},
	{"6 单字练习", viewChars},
}

type app struct {
	face  *c1device.Face
	view  view
	menuI int

	level  int
	streak int
	stats  stats

	roots  *wubi.RootSession
	sess   *wubi.Session
	learnI int
	chartZ int
	ruleI  int
	ruleY  int

	message string
}

func main() {
	face, err := c1device.NewBitmapFace(fontData, 16)
	if err != nil {
		fmt.Fprintln(os.Stderr, "c1wubi: font:", err)
		os.Exit(1)
	}
	defer face.Close()
	_ = os.MkdirAll(dataDir(), 0o755)

	a := &app{face: face, level: 1, stats: loadStats(), chartZ: 1}
	platform, err := c1device.OpenPlatform()
	if err != nil {
		fmt.Fprintln(os.Stderr, "c1wubi: platform:", err)
		os.Exit(1)
	}
	defer platform.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	draw := func() { _ = platform.Draw(a.render(), true) }
	draw()

	for {
		select {
		case <-ctx.Done():
			return
		case ev, ok := <-platform.Events():
			if !ok {
				return
			}
			if a.handle(ev, cancel) {
				return
			}
			draw()
		case <-time.After(time.Second):
		}
	}
}

func (a *app) enter(v view) {
	a.view = v
	a.message = ""
	a.streak = 0
	switch v {
	case viewRules:
		a.ruleI = 0
		a.ruleY = 0
	case viewChart:
		if a.chartZ < 1 || a.chartZ > 5 {
			a.chartZ = 1
		}
	case viewLearn:
		a.learnI = 0
	case viewTable, viewRoots:
		a.roots = wubi.NewRootSession(time.Now().UnixNano())
		a.sess = nil
	case viewChars:
		items := wubi.Levels(a.level)
		if len(items) == 0 {
			items = wubi.All()
		}
		a.sess = wubi.NewSession(items, time.Now().UnixNano())
		a.roots = nil
	}
}

func (a *app) handle(ev c1device.Event, cancel context.CancelFunc) bool {
	switch a.view {
	case viewMenu:
		return a.handleMenu(ev, cancel)
	case viewRules:
		a.handleRules(ev)
	case viewChart:
		a.handleChart(ev)
	case viewLearn:
		a.handleLearn(ev)
	case viewTable, viewRoots:
		a.handleRoots(ev)
	case viewChars:
		a.handleChars(ev)
	}
	return false
}

func (a *app) handleMenu(ev c1device.Event, cancel context.CancelFunc) bool {
	switch ev.Key {
	case c1device.KeyBack:
		cancel()
		return true
	case c1device.KeyUp:
		if a.menuI > 0 {
			a.menuI--
		}
	case c1device.KeyDown:
		if a.menuI+1 < len(menuItems) {
			a.menuI++
		}
	case c1device.KeyOK, c1device.KeyRight:
		a.enter(menuItems[a.menuI].to)
	case c1device.KeyRune:
		if ev.Rune >= '1' && ev.Rune <= '6' {
			idx := int(ev.Rune - '1')
			if idx < len(menuItems) {
				a.enter(menuItems[idx].to)
			}
		}
	}
	return false
}

func (a *app) handleRules(ev c1device.Event) {
	n := len(wubi.Rules)
	page := wubi.Rules[a.ruleI]
	// 正文区大约到 y=130，行高 18，从 y=30 起
	const maxRows = 5
	scrollable := len(page.Lines) > maxRows
	switch ev.Key {
	case c1device.KeyBack, c1device.KeyLeft:
		a.view = viewMenu
	case c1device.KeyUp, c1device.KeyVolumeUp:
		if scrollable && a.ruleY > 0 {
			a.ruleY--
		} else if a.ruleI > 0 {
			a.ruleI--
			a.ruleY = 0
		}
	case c1device.KeyDown, c1device.KeyVolumeDown:
		if scrollable && a.ruleY+maxRows < len(page.Lines) {
			a.ruleY++
		} else if a.ruleI+1 < n {
			a.ruleI++
			a.ruleY = 0
		}
	case c1device.KeyOK:
		if a.ruleI+1 < n {
			a.ruleI++
			a.ruleY = 0
		}
	}
}

func (a *app) handleChart(ev c1device.Event) {
	switch ev.Key {
	case c1device.KeyBack, c1device.KeyLeft:
		a.view = viewMenu
	case c1device.KeyUp, c1device.KeyVolumeUp:
		if a.chartZ > 1 {
			a.chartZ--
		}
	case c1device.KeyDown, c1device.KeyVolumeDown, c1device.KeyOK:
		if a.chartZ < 5 {
			a.chartZ++
		} else {
			a.chartZ = 1
		}
	}
}

func (a *app) handleLearn(ev c1device.Event) {
	switch ev.Key {
	case c1device.KeyBack, c1device.KeyLeft:
		a.view = viewMenu
	case c1device.KeyUp, c1device.KeyVolumeUp:
		if a.learnI > 0 {
			a.learnI--
		}
	case c1device.KeyDown, c1device.KeyRight, c1device.KeyOK, c1device.KeyVolumeDown:
		if a.learnI < len(wubi.RootKeys)-1 {
			a.learnI++
		}
	case c1device.KeyRune:
		if ks := a.learnKey(); ks != nil {
			r := ev.Rune
			if r >= 'A' && r <= 'Z' {
				r += 'a' - 'A'
			}
			if byte(r) == ks.Key {
				if a.learnI < len(wubi.RootKeys)-1 {
					a.learnI++
				}
			}
		}
	}
}

func (a *app) handleRoots(ev c1device.Event) {
	if a.roots == nil {
		return
	}
	switch ev.Key {
	case c1device.KeyBack, c1device.KeyLeft:
		if a.roots.ShowAns {
			a.roots.Continue()
			return
		}
		a.view = viewMenu
	case c1device.KeyOK:
		if a.roots.Done() {
			a.enter(a.view)
			return
		}
		if a.roots.ShowAns {
			a.roots.Continue()
		}
	case c1device.KeyRune:
		switch ev.Rune {
		case 'n', 'N':
			a.enter(a.view)
		default:
			a.roots.Answer(ev.Rune)
			if a.roots.ShowAns {
				a.streak = 0
			} else {
				a.streak++
			}
		}
	}
}

func (a *app) handleChars(ev c1device.Event) {
	if a.sess == nil {
		return
	}
	switch ev.Key {
	case c1device.KeyBack, c1device.KeyLeft:
		if a.sess.ShowAns {
			a.sess.Continue()
			return
		}
		a.view = viewMenu
	case c1device.KeyOK:
		if a.sess.Done() {
			a.enter(viewChars)
			return
		}
		if a.sess.ShowAns {
			a.sess.Continue()
			return
		}
		if a.sess.Submit() {
			a.streak++
			a.stats.TotalOK++
			if a.streak > a.stats.BestStreak {
				a.stats.BestStreak = a.streak
			}
		} else {
			a.streak = 0
		}
		a.stats.TotalTry++
		_ = saveStats(a.stats)
	case c1device.KeyRune:
		switch ev.Rune {
		case 'n', 'N':
			a.enter(viewChars)
		case 'l', 'L':
			a.level++
			if a.level > 3 {
				a.level = 1
			}
			a.enter(viewChars)
		default:
			a.sess.Type(ev.Rune)
		}
	}
}

func (a *app) learnKey() *wubi.RootSet {
	if a.learnI < 0 || a.learnI >= len(wubi.RootKeys) {
		return nil
	}
	return &wubi.RootKeys[a.learnI]
}
