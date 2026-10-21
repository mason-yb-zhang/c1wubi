// 五笔 86 版编码表（一级简码 + 高频字）。
package wubi

import "strings"

// Entry 一字多码：Full 为全码，Short 为简码（可空）。
type Entry struct {
	Char  string
	Full  string
	Short string
	Level int // 1=一级简码 2=二级/高频 3=进阶
}

// Table 内置练习字库（五笔 86）。
var Table = []Entry{
	{Char: "一", Full: "ggll", Short: "ggg", Level: 1},
	{Char: "地", Full: "fbn", Short: "f", Level: 1},
	{Char: "在", Full: "dhfd", Short: "d", Level: 1},
	{Char: "要", Full: "svf", Short: "s", Level: 1},
	{Char: "工", Full: "aaaa", Short: "a", Level: 1},
	{Char: "上", Full: "hhgg", Short: "h", Level: 1},
	{Char: "是", Full: "jghu", Short: "j", Level: 1},
	{Char: "中", Full: "khk", Short: "k", Level: 1},
	{Char: "国", Full: "lgyi", Short: "l", Level: 1},
	{Char: "同", Full: "mgkd", Short: "m", Level: 1},
	{Char: "和", Full: "tkg", Short: "t", Level: 1},
	{Char: "的", Full: "rqyy", Short: "r", Level: 1},
	{Char: "有", Full: "def", Short: "e", Level: 1},
	{Char: "人", Full: "wwww", Short: "w", Level: 1},
	{Char: "我", Full: "trnt", Short: "q", Level: 1},
	{Char: "主", Full: "ygd", Short: "y", Level: 1},
	{Char: "产", Full: "ute", Short: "u", Level: 1},
	{Char: "不", Full: "gii", Short: "i", Level: 1},
	{Char: "为", Full: "ylly", Short: "o", Level: 1},
	{Char: "这", Full: "ypi", Short: "p", Level: 1},
	{Char: "民", Full: "nav", Short: "n", Level: 1},
	{Char: "了", Full: "bnh", Short: "b", Level: 1},
	{Char: "发", Full: "ntcy", Short: "v", Level: 1},
	{Char: "以", Full: "nywy", Short: "c", Level: 1},
	{Char: "经", Full: "xcag", Short: "x", Level: 1},
	{Char: "你", Full: "wqiy", Short: "wq", Level: 2},
	{Char: "好", Full: "vbg", Short: "vb", Level: 2},
	{Char: "他", Full: "wbn", Short: "wb", Level: 2},
	{Char: "她", Full: "vbn", Short: "ve", Level: 2},
	{Char: "们", Full: "wun", Short: "wu", Level: 2},
	{Char: "来", Full: "goi", Short: "go", Level: 2},
	{Char: "去", Full: "fcu", Short: "fc", Level: 2},
	{Char: "说", Full: "yukq", Short: "yu", Level: 2},
	{Char: "会", Full: "wfcu", Short: "wf", Level: 2},
	{Char: "就", Full: "yidn", Short: "yi", Level: 2},
	{Char: "时", Full: "jfy", Short: "jf", Level: 2},
	{Char: "年", Full: "rhfk", Short: "rh", Level: 2},
	{Char: "生", Full: "tgd", Short: "tg", Level: 2},
	{Char: "可", Full: "skd", Short: "sk", Level: 2},
	{Char: "也", Full: "bnhn", Short: "bn", Level: 2},
	{Char: "出", Full: "bmk", Short: "bm", Level: 2},
	{Char: "到", Full: "gcfj", Short: "gc", Level: 2},
	{Char: "对", Full: "cfy", Short: "cf", Level: 2},
	{Char: "能", Full: "cexx", Short: "ce", Level: 2},
	{Char: "都", Full: "ftjb", Short: "ft", Level: 2},
	{Char: "没", Full: "imcy", Short: "im", Level: 2},
	{Char: "看", Full: "rhf", Short: "rh", Level: 2},
	{Char: "那", Full: "vfbh", Short: "vf", Level: 2},
	{Char: "很", Full: "tvey", Short: "tv", Level: 2},
	{Char: "还", Full: "gipi", Short: "gi", Level: 2},
	{Char: "得", Full: "tjgf", Short: "tj", Level: 2},
	{Char: "于", Full: "gfk", Short: "gf", Level: 2},
	{Char: "下", Full: "ghi", Short: "gh", Level: 2},
	{Char: "自", Full: "thd", Short: "th", Level: 2},
	{Char: "作", Full: "wthf", Short: "wt", Level: 2},
	{Char: "然", Full: "qdou", Short: "qd", Level: 2},
	{Char: "所", Full: "rnrh", Short: "rn", Level: 2},
	{Char: "家", Full: "peu", Short: "pe", Level: 2},
	{Char: "学", Full: "ipbf", Short: "ip", Level: 2},
	{Char: "小", Full: "ihty", Short: "ih", Level: 2},
	{Char: "心", Full: "nyny", Short: "ny", Level: 2},
	{Char: "多", Full: "qqu", Short: "qq", Level: 2},
	{Char: "大", Full: "dddd", Short: "dd", Level: 2},
	{Char: "天", Full: "gdi", Short: "gd", Level: 2},
	{Char: "开", Full: "gak", Short: "ga", Level: 2},
	{Char: "面", Full: "dmjd", Short: "dm", Level: 2},
	{Char: "前", Full: "uejj", Short: "ue", Level: 2},
	{Char: "后", Full: "rgkd", Short: "rg", Level: 2},
	{Char: "东", Full: "aii", Short: "ai", Level: 2},
	{Char: "西", Full: "sghg", Short: "sg", Level: 2},
	{Char: "北", Full: "uxn", Short: "ux", Level: 2},
	{Char: "京", Full: "yiu", Short: "yi", Level: 2},
	{Char: "分", Full: "wvb", Short: "wv", Level: 2},
	{Char: "子", Full: "bbbb", Short: "bb", Level: 2},
	{Char: "女", Full: "vvvv", Short: "vv", Level: 2},
	{Char: "又", Full: "cccc", Short: "cc", Level: 2},
	{Char: "月", Full: "eeee", Short: "ee", Level: 2},
	{Char: "日", Full: "jjjj", Short: "jj", Level: 2},
	{Char: "水", Full: "iiii", Short: "ii", Level: 2},
	{Char: "火", Full: "oooo", Short: "oo", Level: 2},
	{Char: "山", Full: "mmmm", Short: "mm", Level: 2},
	{Char: "石", Full: "dgtg", Short: "dg", Level: 2},
	{Char: "田", Full: "llll", Short: "ll", Level: 2},
	{Char: "口", Full: "kkkk", Short: "kk", Level: 2},
	{Char: "手", Full: "rtgh", Short: "rt", Level: 2},
	{Char: "力", Full: "ltn", Short: "lt", Level: 2},
	{Char: "文", Full: "yygy", Short: "yy", Level: 2},
	{Char: "方", Full: "yygn", Short: "yy", Level: 2},
	{Char: "门", Full: "uyhn", Short: "uy", Level: 2},
	{Char: "车", Full: "lgnh", Short: "lg", Level: 2},
	{Char: "马", Full: "cnng", Short: "cn", Level: 2},
	{Char: "白", Full: "rrrr", Short: "rr", Level: 2},
	{Char: "用", Full: "etnh", Short: "et", Level: 2},
	{Char: "明", Full: "jeg", Short: "je", Level: 2},
	{Char: "样", Full: "sudh", Short: "su", Level: 2},
	{Char: "种", Full: "tkhh", Short: "tk", Level: 2},
	{Char: "些", Full: "hxff", Short: "hx", Level: 2},
	{Char: "现", Full: "gmqn", Short: "gm", Level: 2},
	{Char: "点", Full: "hkou", Short: "hk", Level: 2},
	{Char: "正", Full: "ghd", Short: "gh", Level: 2},
	{Char: "第", Full: "txht", Short: "tx", Level: 2},
	{Char: "从", Full: "wwy", Short: "ww", Level: 2},
	{Char: "只", Full: "kwu", Short: "kw", Level: 2},
	{Char: "见", Full: "mqb", Short: "mq", Level: 2},
	{Char: "把", Full: "rcn", Short: "rc", Level: 2},
	{Char: "两", Full: "gmww", Short: "gm", Level: 2},
	{Char: "问", Full: "ukd", Short: "uk", Level: 2},
	{Char: "题", Full: "jghm", Short: "jt", Level: 3},
	{Char: "实", Full: "pudu", Short: "pu", Level: 2},
	{Char: "什", Full: "wfh", Short: "wf", Level: 2},
	{Char: "么", Full: "tcu", Short: "tc", Level: 2},
	{Char: "高", Full: "ymkf", Short: "ym", Level: 2},
	{Char: "长", Full: "tayi", Short: "ta", Level: 2},
	{Char: "手", Full: "rtgh", Short: "rt", Level: 2},
	{Char: "电", Full: "jnv", Short: "jn", Level: 2},
	{Char: "话", Full: "ytdg", Short: "yt", Level: 2},
	{Char: "打", Full: "rsh", Short: "rs", Level: 2},
	{Char: "字", Full: "pbf", Short: "pb", Level: 2},
	{Char: "五", Full: "gghg", Short: "gg", Level: 2},
	{Char: "笔", Full: "ttfn", Short: "tt", Level: 2},
	{Char: "练", Full: "xanw", Short: "xa", Level: 2},
	{Char: "习", Full: "nud", Short: "nu", Level: 2},
	{Char: "键", Full: "qvfp", Short: "qv", Level: 3},
	{Char: "盘", Full: "telf", Short: "te", Level: 3},
	{Char: "速", Full: "gkip", Short: "sk", Level: 3},
	{Char: "度", Full: "yaci", Short: "ya", Level: 2},
	{Char: "成", Full: "dnnt", Short: "dn", Level: 2},
	{Char: "功", Full: "aln", Short: "al", Level: 2},
	{Char: "新", Full: "usrh", Short: "us", Level: 2},
	{Char: "老", Full: "ftxb", Short: "ft", Level: 2},
	{Char: "师", Full: "jgmh", Short: "jg", Level: 2},
	{Char: "学", Full: "ipbf", Short: "ip", Level: 2},
	{Char: "习", Full: "nud", Short: "nu", Level: 2},
	{Char: "快", Full: "nnwy", Short: "nn", Level: 2},
	{Char: "乐", Full: "qii", Short: "qi", Level: 2},
	{Char: "跑", Full: "khqn", Short: "kh", Level: 2},
	{Char: "吃", Full: "ktnn", Short: "kt", Level: 2},
	{Char: "喝", Full: "kjqn", Short: "kj", Level: 2},
	{Char: "走", Full: "fhu", Short: "fh", Level: 2},
	{Char: "飞", Full: "nui", Short: "nu", Level: 2},
	{Char: "机", Full: "smn", Short: "sm", Level: 2},
	{Char: "买", Full: "nudu", Short: "nu", Level: 2},
	{Char: "卖", Full: "fnud", Short: "fn", Level: 2},
	{Char: "钱", Full: "qgt", Short: "qg", Level: 2},
	{Char: "元", Full: "fqb", Short: "fq", Level: 2},
	{Char: "早", Full: "jnhh", Short: "jn", Level: 2},
	{Char: "晚", Full: "jqkq", Short: "jq", Level: 2},
	{Char: "休", Full: "wsy", Short: "ws", Level: 2},
	{Char: "息", Full: "thnu", Short: "th", Level: 2},
	{Char: "再", Full: "gmfd", Short: "gm", Level: 2},
	{Char: "次", Full: "uqwy", Short: "uq", Level: 2},
	{Char: "更", Full: "gjqi", Short: "gj", Level: 2},
	{Char: "最", Full: "jbcu", Short: "jb", Level: 2},
	{Char: "真", Full: "fhwu", Short: "fh", Level: 2},
	{Char: "假", Full: "wnhc", Short: "wn", Level: 3},
	{Char: "信", Full: "wyg", Short: "wy", Level: 2},
	{Char: "息", Full: "thnu", Short: "th", Level: 2},
	{Char: "谢", Full: "ytmf", Short: "yt", Level: 3},
	{Char: "请", Full: "ygeg", Short: "yg", Level: 2},
	{Char: "帮", Full: "dtnh", Short: "dt", Level: 2},
	{Char: "助", Full: "egln", Short: "eg", Level: 2},
	{Char: "安", Full: "pvf", Short: "pv", Level: 2},
	{Char: "全", Full: "wgf", Short: "wg", Level: 2},
	{Char: "己", Full: "nngn", Short: "nn", Level: 2},
	{Char: "已", Full: "nnnn", Short: "nn", Level: 2},
	{Char: "几", Full: "mtn", Short: "mt", Level: 2},
	{Char: "每", Full: "txgu", Short: "tx", Level: 2},
	{Char: "次", Full: "uqwy", Short: "uq", Level: 2},
}

// Levels 返回难度分组下的字。
func Levels(level int) []Entry {
	var out []Entry
	seen := map[string]bool{}
	for _, e := range Table {
		if e.Char == "" || e.Full == "" {
			continue
		}
		if e.Level != level {
			continue
		}
		if seen[e.Char] {
			continue
		}
		seen[e.Char] = true
		out = append(out, e)
	}
	return out
}

// All 去重后的全部可练条目。
func All() []Entry {
	var out []Entry
	seen := map[string]bool{}
	for _, e := range Table {
		if e.Char == "" || e.Full == "" || seen[e.Char] {
			continue
		}
		seen[e.Char] = true
		out = append(out, e)
	}
	return out
}

// Accept 报告输入是否命中该字（简码或全码，忽略大小写）。
func (e Entry) Accept(input string) bool {
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		return false
	}
	if strings.ToLower(e.Short) != "" && input == strings.ToLower(e.Short) {
		return true
	}
	return input == strings.ToLower(e.Full)
}

// Hint 显示推荐编码（简码优先）。
func (e Entry) Hint() string {
	if e.Short != "" {
		return e.Short
	}
	return e.Full
}
