// 五笔 86 字根表：键位 → 字根，供字根认键练习。
package wubi

import "math/rand"

// RootSet 一个键上的字根（首字为键名）。
type RootSet struct {
	Key    byte   // 小写字母
	Zone   int    // 1横 2竖 3撇 4捺 5折
	Phrase string // 助记（简）
	Roots  []rune // 该键字根（键名在前）
}

// RootKeys 25 键完整字根。
var RootKeys = []RootSet{
	{Key: 'g', Zone: 1, Phrase: "王旁青头戋五一", Roots: []rune{'王', '一', '五', '戋', '丅'}},
	{Key: 'f', Zone: 1, Phrase: "土士二干十寸雨", Roots: []rune{'土', '士', '二', '干', '十', '寸', '雨'}},
	{Key: 'd', Zone: 1, Phrase: "大犬三羊古石厂", Roots: []rune{'大', '犬', '三', '羊', '古', '石', '厂'}},
	{Key: 's', Zone: 1, Phrase: "木丁西", Roots: []rune{'木', '丁', '西'}},
	{Key: 'a', Zone: 1, Phrase: "工戈草头右框七", Roots: []rune{'工', '戈', '艹', '匚', '七'}},
	{Key: 'h', Zone: 2, Phrase: "目具上止卜虎皮", Roots: []rune{'目', '上', '止', '卜', '虎', '皮'}},
	{Key: 'j', Zone: 2, Phrase: "日早两竖与虫依", Roots: []rune{'日', '早', '虫'}},
	{Key: 'k', Zone: 2, Phrase: "口与川，字根稀", Roots: []rune{'口', '川'}},
	{Key: 'l', Zone: 2, Phrase: "田甲方框四车力", Roots: []rune{'田', '甲', '四', '车', '力'}},
	{Key: 'm', Zone: 2, Phrase: "山由贝，下框几", Roots: []rune{'山', '由', '贝', '几'}},
	{Key: 't', Zone: 3, Phrase: "禾竹一撇双人立", Roots: []rune{'禾', '竹'}},
	{Key: 'r', Zone: 3, Phrase: "白手看头三二斤", Roots: []rune{'白', '手', '斤'}},
	{Key: 'e', Zone: 3, Phrase: "月彡乃用家衣底", Roots: []rune{'月', '彡', '用'}},
	{Key: 'w', Zone: 3, Phrase: "人和八，登祭头", Roots: []rune{'人', '八'}},
	{Key: 'q', Zone: 3, Phrase: "金勺缺点无尾鱼", Roots: []rune{'金', '鱼'}},
	{Key: 'y', Zone: 4, Phrase: "言文方广在四一", Roots: []rune{'言', '文', '方', '广'}},
	{Key: 'u', Zone: 4, Phrase: "立辛两点六门病", Roots: []rune{'立', '辛', '六', '门'}},
	{Key: 'i', Zone: 4, Phrase: "水旁兴头小倒立", Roots: []rune{'水', '小'}},
	{Key: 'o', Zone: 4, Phrase: "火业头，四点米", Roots: []rune{'火', '米'}},
	{Key: 'p', Zone: 4, Phrase: "之字军盖建道底", Roots: []rune{'之', '军'}},
	{Key: 'n', Zone: 5, Phrase: "已半巳满不出己", Roots: []rune{'已', '己', '尸', '心'}},
	{Key: 'b', Zone: 5, Phrase: "子耳了也框向上", Roots: []rune{'子', '耳', '了', '也'}},
	{Key: 'v', Zone: 5, Phrase: "女刀九臼山朝西", Roots: []rune{'女', '刀', '九'}},
	{Key: 'c', Zone: 5, Phrase: "又巴马，丢矢矣", Roots: []rune{'又', '巴', '马'}},
	{Key: 'x', Zone: 5, Phrase: "纟母无心弓和匕", Roots: []rune{'纟', '弓', '匕'}},
}

// RootDrill 字根题：显示字根，回答键位。
type RootDrill struct {
	Rune rune
	Key  byte
	Note string // 助记
}

// AllRoots 展平为可练清单。
func AllRoots() []RootDrill {
	var out []RootDrill
	for _, ks := range RootKeys {
		for i, r := range ks.Roots {
			note := ks.Phrase
			if i == 0 {
				note = string(ks.Key) + " 键名 " + note
			}
			out = append(out, RootDrill{Rune: r, Key: ks.Key, Note: note})
		}
	}
	return out
}

// ZoneName 区名。
func ZoneName(z int) string {
	switch z {
	case 1:
		return "横区"
	case 2:
		return "竖区"
	case 3:
		return "撇区"
	case 4:
		return "捺区"
	case 5:
		return "折区"
	}
	return ""
}

// RootSession 字根练习会话。
type RootSession struct {
	Items   []RootDrill
	Index   int
	Input   byte
	Correct int
	Wrong   int
	ShowAns bool
}

// NewRootSession 新建（seed=0 不打乱，否则打乱）。
func NewRootSession(seed int64) *RootSession {
	items := AllRoots()
	if seed != 0 {
		r := rand.New(rand.NewSource(seed))
		r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	}
	return &RootSession{Items: items}
}

func (s *RootSession) Current() *RootDrill {
	if s.Index < 0 || s.Index >= len(s.Items) {
		return nil
	}
	return &s.Items[s.Index]
}

func (s *RootSession) Done() bool { return s.Index >= len(s.Items) }

// Answer 按键作答。
func (s *RootSession) Answer(key rune) {
	cur := s.Current()
	if cur == nil || s.ShowAns {
		return
	}
	if key >= 'A' && key <= 'Z' {
		key += 'a' - 'A'
	}
	if key < 'a' || key > 'z' {
		return
	}
	s.Input = byte(key)
	if byte(key) == cur.Key {
		s.Correct++
		s.Index++
		s.ShowAns = false
		s.Input = 0
	} else {
		s.Wrong++
		s.ShowAns = true
	}
}

// Continue 看答案后下一题。
func (s *RootSession) Continue() {
	if s.ShowAns {
		s.Index++
		s.ShowAns = false
		s.Input = 0
	}
}


// KeyMapRows 物理键盘 QWERTY 三行（含未用键，未用为 0）。
var KeyMapRows = [][]byte{
	{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
	{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 0},
	{'z', 0, 'x', 'c', 'v', 'b', 'n', 'm', 0, 0},
}

// ZoneGlyph 区的图标：横竖撇捺折。
func ZoneGlyph(z int) string {
	switch z {
	case 1:
		return "一"
	case 2:
		return "丨"
	case 3:
		return "丿"
	case 4:
		return "丶"
	case 5:
		return "乙"
	}
	return ""
}

// KeyName 键名字符。
func KeyName(key byte) string {
	for _, ks := range RootKeys {
		if ks.Key == key {
			if len(ks.Roots) > 0 {
				return string(ks.Roots[0])
			}
		}
	}
	return "?"
}

// RootsString 该键字根串（不含助记）。
func RootsString(key byte) string {
	for _, ks := range RootKeys {
		if ks.Key == key {
			return string(ks.Roots)
		}
	}
	return ""
}

// LookupKey 找字根所在键。
func LookupKey(r rune) (byte, bool) {
	for _, ks := range RootKeys {
		for _, rr := range ks.Roots {
			if rr == r {
				return ks.Key, true
			}
		}
	}
	return 0, false
}


// ZoneKeys 该区全部键位（按口诀顺序）。
func ZoneKeys(zone int) []RootSet {
	var out []RootSet
	for _, ks := range RootKeys {
		if ks.Zone == zone {
			out = append(out, ks)
		}
	}
	return out
}
