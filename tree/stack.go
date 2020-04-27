package tree

type stack struct {
	list []Node
	len  int
	ptr  int
}

func NewStack() *stack {
	return &stack{
		list: make([]Node, 0),
		len:  0,
		ptr:  -1,
	}

}

func (s *stack) Push(v Node) {
	if len(s.list)-1 == s.ptr {
		s.list = append(s.list, v)
		s.ptr++
	} else {
		s.ptr++
		s.list[s.ptr] = v
	}
}

func (s *stack) Pop() Node {
	if s.ptr >= 0 {
		s.ptr--
		return s.list[s.ptr+1]
	}
	return nil
}
