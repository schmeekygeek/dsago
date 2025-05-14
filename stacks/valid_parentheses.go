package stacks

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	stack := new(Stack)
	for _, v := range s {
		val := string(v)
		switch val {
		case "(":
			stack.Push(val)
		case "[":
			stack.Push(val)
		case "{":
			stack.Push(val)
		case ")":
			if stack.Pop() != "(" {
				return false
			}
		case "]":
			if stack.Pop() != "[" {
				return false
			}
		case "}":
			if stack.Pop() != "{" {
				return false
			}
		}
	}
	if len(stack.Items) != 0 {
		return false
	}
	return true
}

type Stack struct {
	Items []string
}

func (s *Stack) Push(item string) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() string {
	if len(s.Items) == 0 {
		return ""
	}
	result := s.Items[len(s.Items)-1]
	newStack := s.Items[:len(s.Items)-1]
	s.Items = newStack
	return result
}

func (s *Stack) Peek() string {
	if len(s.Items) != 0 {
		return s.Items[len(s.Items)-1]
	}
	return ""
}
