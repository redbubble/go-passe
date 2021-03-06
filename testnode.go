package main

import "strings"

type TestState int

const (
	Unknown TestState = iota + 1
	Passed
	Failed
)

type testNode struct {
	State          TestState
	Output         []string
	ChildrenByName map[string]*testNode
}

func newTestNode() *testNode {
	return &testNode{
		State:          Unknown,
		ChildrenByName: make(map[string]*testNode),
	}
}

func (n *testNode) AppendOutput(output string) {
	trimmed := strings.TrimSpace(output)

	if trimmed == "FAIL" {
		return
	}
	if strings.HasPrefix(trimmed, "exit status ") {
		return
	}
	if strings.HasPrefix(trimmed, "coverage: ") && strings.HasSuffix(trimmed, " of statements") {
		return
	}
	if strings.HasPrefix(trimmed, "FAIL\t") {
		return
	}
	if strings.HasPrefix(trimmed, "PASS") {
		return
	}
	if strings.HasPrefix(trimmed, "=== RUN   ") {
		return
	}
	if strings.HasPrefix(trimmed, "--- FAIL: ") {
		return
	}
	if strings.HasPrefix(trimmed, "--- PASS: ") {
		return
	}
	if strings.HasPrefix(trimmed, "?   ") && strings.HasSuffix(trimmed, "\t[no test files]") {
		return
	}

	n.Output = append(n.Output, trimmed)
}

func (n *testNode) Get(name string) *testNode {
	next, rest := pathStep(name)
	if next == "" {
		return n
	}

	return n.childForNextPathStep(next).Get(rest)
}

func (n *testNode) MarkFailed(name string) {
	n.State = Failed

	next, rest := pathStep(name)
	if next == "" {
		return
	}

	n.childForNextPathStep(next).MarkFailed(rest)
}

func (n *testNode) MarkPassed(name string) {
	next, rest := pathStep(name)
	if next == "" {
		n.State = Passed
		return
	}

	n.childForNextPathStep(next).MarkPassed(rest)
}

func (n *testNode) childForNextPathStep(next string) *testNode {
	child, ok := n.ChildrenByName[next]
	if !ok {
		child = newTestNode()
		n.ChildrenByName[next] = child
	}

	return child
}

func pathStep(nodeName string) (next, rest string) {
	separatorIndex := strings.IndexRune(nodeName, '/')
	if separatorIndex < 0 {
		return nodeName, ""
	}

	return nodeName[:separatorIndex], nodeName[separatorIndex+1:]
}
