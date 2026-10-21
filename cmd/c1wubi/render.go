// 界面：主菜单 / 规则 / 对照表 / 背字根 / 键位图练习 / 单字。
package main

import (
	"fmt"
	"image"
	"strings"

	"c1device"
	"c1wubi/wubi"
)

func (a *app) render() c1device.Frame {
	c := c1device.NewCanvas()
	c.Clear()
	switch a.view {
	case viewMenu:
		a.renderMenu(c)
	case viewRules:
		a.renderRules(c)
	case viewChart:
		a.renderChart(c)
	case viewLearn:
		a.renderLearn(c)
	case viewTable:
		a.renderTableMode(c)
	case viewRoots:
		a.renderRoots(c)
	case viewChars:
		a.renderChars(c)
	}
	return c.Frame(128)
}

func fitFace(face *c1device.Face, s string, w int) string {
	if face.Measure(s) <= w {
		return s
	}
	r := []rune(s)
	for len(r) > 0 && face.Measure(string(r)+"…") > w {
		r = r[:len(r)-1]
	}
	if len(r) == 0 {
		return "…"
	}
	return string(r) + "…"
}

func (a *app) renderMenu(c *c1device.Canvas) {
	c.DrawText(a.face, 6, 2, "五笔练习")
	c.DrawTextRight(a.face, 290, 2, fmt.Sprintf("连纪录 %d", a.stats.BestStreak))
	c.FillRect(image.Rect(4, 21, 292, 23))
	// 6 项，一页可显示
	for i, it := range menuItems {
		y := 28 + i*17
		if i == a.menuI {
			// 左对齐，仅高亮底色
			c.DrawTextInverted(a.face, image.Rect(4, y-2, 292, y+15), 12, y, it.title)
		} else {
			c.DrawText(a.face, 12, y, it.title)
		}
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152),
		"↑↓选择 OK进入 1-6快选 Back退出")
}

func (a *app) renderRules(c *c1device.Canvas) {
	page := wubi.Rules[a.ruleI]
	c.DrawText(a.face, 6, 2, page.Title)
	c.DrawTextRight(a.face, 290, 2, fmt.Sprintf("%d/%d", a.ruleI+1, len(wubi.Rules)))
	c.FillRect(image.Rect(4, 21, 292, 23))
	// 正文区：y=28 起，最多 5 行，禁止压到页脚 132
	const maxRows = 5
	const lineH = 18
	y := 28
	end := a.ruleY + maxRows
	if end > len(page.Lines) {
		end = len(page.Lines)
	}
	for _, ln := range page.Lines[a.ruleY:end] {
		c.DrawText(a.face, 8, y, fitFace(a.face, ln, 280))
		y += lineH
	}
	hint := "↑↓滚动  OK下一页  Back菜单"
	if len(page.Lines) > maxRows {
		hint = fmt.Sprintf("↑↓滚动 %d-%d/%d  OK下一页  Back", a.ruleY+1, end, len(page.Lines))
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), fitFace(a.face, hint, 280))
}

func (a *app) renderChart(c *c1device.Canvas) {
	z := a.chartZ
	if z < 1 || z > 5 {
		z = 1
	}
	keys := wubi.ZoneKeys(z)
	c.DrawText(a.face, 6, 2,
		fmt.Sprintf("对照表 %s %s %d/5", wubi.ZoneGlyph(z), wubi.ZoneName(z), z))
	c.FillRect(image.Rect(4, 21, 292, 23))
	y := 26
	for _, ks := range keys {
		key := strings.ToUpper(string([]byte{ks.Key}))
		line := fmt.Sprintf("%s %s %s", key, wubi.KeyName(ks.Key), wubi.RootsString(ks.Key))
		c.DrawText(a.face, 8, y, fitFace(a.face, line, 280))
		y += 16
	}
	if len(keys) > 0 {
		c.DrawText(a.face, 8, 112, fitFace(a.face, "口诀 "+keys[0].Phrase, 280))
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152),
		"↑↓/音量 翻区  OK下一区  Back菜单")
}

func drawKeyMap(c *c1device.Canvas, face *c1device.Face, y0 int, highlight byte) {
	offsets := []int{8, 18, 36}
	rowH := 28
	for row, keys := range wubi.KeyMapRows {
		y := y0 + row*rowH
		x := offsets[row]
		for _, k := range keys {
			if k == 0 {
				x += 26
				continue
			}
			box := image.Rect(x, y, x+24, y+rowH-2)
			if k == highlight {
				c.FillRect(box)
			} else {
				c.DrawRect(box)
			}
			name := wubi.KeyName(k)
			letter := strings.ToUpper(string([]byte{k}))
			cx := x + 12
			c.DrawTextCentered(face, cx, y+2, name)
			c.DrawTextCentered(face, cx, y+15, letter)
			x += 28
		}
	}
}

func (a *app) renderLearn(c *c1device.Canvas) {
	ks := a.learnKey()
	if ks == nil {
		return
	}
	drawKeyMap(c, a.face, 26, ks.Key)
	key := strings.ToUpper(string([]byte{ks.Key}))
	info := fmt.Sprintf("%s %s %s  %s", wubi.ZoneGlyph(ks.Zone), key, wubi.KeyName(ks.Key), ks.Phrase)
	c.DrawInvertedTextBar(a.face, image.Rect(0, 110, 296, 130), fitFace(a.face, info, 280))
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152),
		fitFace(a.face, fmt.Sprintf("第%d/%d键 ←→翻 Back菜单", a.learnI+1, len(wubi.RootKeys)), 280))
}

func (a *app) renderTableMode(c *c1device.Canvas) {
	hl := byte(0)
	if a.roots != nil && a.roots.ShowAns && a.roots.Current() != nil {
		hl = a.roots.Current().Key
	}
	drawKeyMap(c, a.face, 24, hl)
	if a.roots == nil {
		return
	}
	if a.roots.Done() {
		c.DrawInvertedTextBar(a.face, image.Rect(0, 110, 296, 130),
			fmt.Sprintf("完成 对%d 错%d", a.roots.Correct, a.roots.Wrong))
		c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), "OK再来  Back菜单")
		return
	}
	cur := a.roots.Current()
	if cur == nil {
		return
	}
	msg := fmt.Sprintf("找「%s」的键并按下", string(cur.Rune))
	if a.roots.ShowAns {
		msg = fmt.Sprintf("「%s」→ %s %s", string(cur.Rune),
			strings.ToUpper(string([]byte{cur.Key})), wubi.KeyName(cur.Key))
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 110, 296, 130), fitFace(a.face, msg, 280))
	footer := "对照键位图按字母  Back菜单"
	if a.roots.ShowAns {
		footer = "看高亮键，OK继续"
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), fitFace(a.face, footer, 280))
}

func (a *app) renderRoots(c *c1device.Canvas) {
	if a.roots == nil {
		return
	}
	if a.roots.Done() {
		c.DrawTextCentered(a.face, 148, 48, "完成")
		c.DrawTextCentered(a.face, 148, 76,
			fmt.Sprintf("对 %d  错 %d", a.roots.Correct, a.roots.Wrong))
		c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), "OK再来  Back菜单")
		return
	}
	cur := a.roots.Current()
	if cur == nil {
		return
	}
	c.DrawTextCentered(a.face, 148, 34, string(cur.Rune))
	c.DrawTextCentered(a.face, 148, 62, wubi.ZoneName(zoneOf(cur.Key)))
	in := "？"
	if a.roots.Input != 0 {
		in = strings.ToUpper(string([]byte{a.roots.Input}))
	}
	c.DrawTextCentered(a.face, 148, 92, in)
	if a.roots.ShowAns {
		c.DrawTextCentered(a.face, 148, 116,
			fmt.Sprintf("应按 %s  %s", strings.ToUpper(string([]byte{cur.Key})), cur.Note))
	} else {
		c.DrawTextCentered(a.face, 148, 116,
			fmt.Sprintf("%d/%d  按键位字母", a.roots.Index+1, len(a.roots.Items)))
	}
	footer := "按键位字母  OK继续  Back菜单"
	if a.roots.ShowAns {
		footer = "答错了，OK看下一题"
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), footer)
}

func (a *app) renderChars(c *c1device.Canvas) {
	if a.sess == nil {
		return
	}
	if a.sess.Done() {
		c.DrawTextCentered(a.face, 148, 48, "本轮完成")
		c.DrawTextCentered(a.face, 148, 76,
			fmt.Sprintf("对 %d 跳 %d 正确率 %d%%", a.sess.Correct, a.sess.Skipped, a.sess.Accuracy()))
		c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), "OK再来 L难度 Back菜单")
		return
	}
	cur := a.sess.Current()
	if cur == nil {
		return
	}
	c.DrawTextCentered(a.face, 148, 36, cur.Char)
	input := a.sess.Input
	if input == "" {
		input = "—"
	}
	c.DrawTextCentered(a.face, 148, 72, input)
	c.DrawRect(image.Rect(40, 90, 256, 106))
	if a.sess.ShowAns {
		c.DrawTextCentered(a.face, 148, 112,
			fmt.Sprintf("答案 %s / %s", cur.Hint(), cur.Full))
	} else {
		c.DrawTextCentered(a.face, 148, 112,
			fmt.Sprintf("第 %d/%d 题", a.sess.Index+1, len(a.sess.Items)))
	}
	footer := "字母输码 OK提交 L难度 Back菜单"
	if a.sess.ShowAns {
		footer = "答错了，看答案后 OK 继续"
	}
	c.DrawInvertedTextBar(a.face, image.Rect(0, 132, 296, 152), footer)
}

func zoneOf(key byte) int {
	for _, ks := range wubi.RootKeys {
		if ks.Key == key {
			return ks.Zone
		}
	}
	return 0
}
