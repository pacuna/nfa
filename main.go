package main

import (
	"fmt"
	stack2 "re/stack"
	"unicode"
)

func Infix2Postfix(input string, ops map[rune]int) string {
	var output string
	stack := stack2.New()
	for _, r := range input {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			output += string(r)
		} else if r == '(' {
			stack.Append(r)
		} else if r == ')' {
			for !stack.Empty() && stack.Top() != '(' {
				t := stack.Pop()
				output += string(t)
			}
			stack.Pop()
		} else {
			for !stack.Empty() && ops[stack.Top()] >= ops[r] {
				output += string(stack.Pop())
			}
			stack.Append(r)
		}
	}
	for !stack.Empty() {
		output += string(stack.Pop())
	}
	return output
}

const (
	MATCH = 256
	SPLIT = 257
)

type State struct {
	c int32
	out, out1 *State
	lastList int
}

func NewState(c int32, out *State, out1 *State) *State {
	return &State{c: c, out: out, out1: out1}
}

type Frag struct {
	/*
	Start points at the start state for the fragment,
	and out is a list of pointers to State* pointers that are not yet connected to anything.
	These are the dangling arrows in the NFA fragment.
	 */
	start *State
	out []**State
}

var (
	matchState = State{c: MATCH}
)

func post2nfa(postfix string) *State {
	var stack []Frag
	var e1, e2, e Frag
	var s *State

	for i:=0;i<len(postfix);i++{
		switch postfix[i] {
		default:
			s = NewState(int32(postfix[i]), nil, nil)
			stack = append(stack, Frag{s, []**State{&s.out}})
		case '.':
			e2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			e1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			patch(e1.out, e2.start)
			stack = append(stack, Frag{
				start: e1.start,
				out:   e2.out,
			})
		case '|':
			e2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			e1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = NewState(SPLIT, e1.start, e2.start)
			stack = append(stack, Frag{
				start: s,
				out:   append(e1.out, e2.out...),
			})
		case '?':
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = NewState(SPLIT, e.start, nil)
			stack = append(stack, Frag{
				start: s,
				out:   append(e.out, &s.out1),
			})
		case '*':
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = NewState(SPLIT, e.start, nil)
			patch(e.out, s)
			stack = append(stack, Frag{
				start: s,
				out:   []**State{&s.out1},
			})
		case '+':
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = NewState(SPLIT, e.start, nil)
			patch(e.out, s)
			stack = append(stack, Frag{
				start: e.start,
				out:   []**State{&s.out1},
			})
		}
	}

	e, stack = stack[len(stack)-1], stack[:len(stack)-1]
	patch(e.out, &matchState)
	return e.start
}

func patch(l []**State, s *State) {
	for i := range l {
		*l[i] = s
	}
}

func main(){
	start := post2nfa("abb.+.a.")
	fmt.Println(start)
}
