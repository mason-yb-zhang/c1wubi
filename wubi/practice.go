// 练习会话：出题、判分、统计。
package wubi

import "math/rand"

// Session 一次练习（顺序出题，可跳过）。
type Session struct {
	Items   []Entry
	Index   int
	Input   string
	Correct int
	Wrong   int
	Skipped int
	ShowAns bool
}

// NewSession 用给定题库开一局。
func NewSession(items []Entry, seed int64) *Session {
	if len(items) == 0 {
		items = All()
	}
	cp := append([]Entry(nil), items...)
	if seed != 0 {
		r := rand.New(rand.NewSource(seed))
		r.Shuffle(len(cp), func(i, j int) { cp[i], cp[j] = cp[j], cp[i] })
	}
	return &Session{Items: cp}
}

// Current 当前题目；结束时返回 nil。
func (s *Session) Current() *Entry {
	if s.Index < 0 || s.Index >= len(s.Items) {
		return nil
	}
	return &s.Items[s.Index]
}

// Done 是否做完。
func (s *Session) Done() bool {
	return s.Index >= len(s.Items)
}

// Type 输入字母/退格。
func (s *Session) Type(r rune) {
	if s.Done() || s.ShowAns {
		return
	}
	if r == '\b' {
		if s.Input != "" {
			s.Input = s.Input[:len(s.Input)-1]
		}
		return
	}
	if r >= 'A' && r <= 'Z' {
		r += 'a' - 'A'
	}
	if r < 'a' || r > 'z' {
		return
	}
	if len(s.Input) < 8 {
		s.Input += string(r)
	}
}

// Submit 提交答案；true=进入下一题。
func (s *Session) Submit() bool {
	cur := s.Current()
	if cur == nil {
		return false
	}
	if cur.Accept(s.Input) {
		s.Correct++
		s.next()
		return true
	}
	if s.Input == "" {
		return false
	}
	s.Wrong++
	s.ShowAns = true
	return false
}

// Skip 跳过当前。
func (s *Session) Skip() {
	if s.Done() {
		return
	}
	s.Skipped++
	s.next()
}

// Continue 显示答案后进入下一题。
func (s *Session) Continue() {
	if s.ShowAns {
		s.ShowAns = false
		s.next()
	}
}

func (s *Session) next() {
	s.Index++
	s.Input = ""
	s.ShowAns = false
}

// Accuracy 正确率（百分比）。
func (s *Session) Accuracy() int {
	total := s.Correct + s.Wrong
	if total == 0 {
		return 0
	}
	return s.Correct * 100 / total
}
