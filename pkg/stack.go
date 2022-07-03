package pkg

type Stack struct {
	storage *[]interface{}
}

func (s *Stack) Push(data interface{}) {
	if s.storage == nil {
		s.storage = &[]interface{}{}
	}
	*s.storage = append(*s.storage, data)
}

func (s *Stack) Pop() interface{} {
	t := (*s.storage)[len(*s.storage)-1]
	*s.storage = (*s.storage)[:len(*s.storage)-1]
	return t
}
