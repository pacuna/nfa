package stack

type Stack struct {
	operators []rune
	size      int
}

func New() *Stack {
	return &Stack{
		operators: []rune{},
	}
}

func (s *Stack) Append(r rune) {
	s.operators = append(s.operators, r)
}

func (s *Stack) Pop() rune {
	r := s.operators[len(s.operators)-1]
	s.operators = s.operators[:len(s.operators)-1]
	return r
}

func (s *Stack) Top() rune {
	return s.operators[len(s.operators)-1]
}

func (s *Stack) Empty() bool {
	return len(s.operators) == 0
}
