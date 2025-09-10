package problems

// Stack byte: đủ dùng cho bài toán ngoặc
type Stack struct {
	data []byte
}

func (s *Stack) Push(b byte) {
	s.data = append(s.data, b)
}

func (s *Stack) Pop() (byte, bool) {
	n := len(s.data)
	if n == 0 {
		return 0, false
	}
	v := s.data[n-1]
	s.data = s.data[:n-1]
	return v, true
}

func (s *Stack) Top() (byte, bool) {
	n := len(s.data)
	if n == 0 {
		return 0, false
	}
	return s.data[n-1], true
}

func (s *Stack) Len() int { return len(s.data) }

// Kiểm tra chuỗi ngoặc hợp lệ
func validParentheses(str string) bool {
	pair := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var st Stack

	for i := 0; i < len(str); i++ {
		c := str[i]
		switch c {
		case '(', '[', '{':
			st.Push(c)
		case ')', ']', '}':
			top, ok := st.Pop()
			if !ok || top != pair[c] {
				return false
			}
			// default: // nếu muốn bỏ qua ký tự khác, để trống;
			// hoặc return false nếu không chấp nhận ký tự khác ngoặc
		}
	}
	return st.Len() == 0
}
