package main

import "strings"

type testNode struct {
	Passed         *bool
	Output         []string
	ChildrenByName map[string]*testNode
}

func newTestNode() *testNode {
	return &testNode{
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

	child, ok := n.ChildrenByName[next]
	if !ok {
		child = newTestNode()
		n.ChildrenByName[next] = child
	}

	return child.Get(rest)
}

func (n *testNode) MarkFailed(name string) {
	passed := false
	n.Passed = &passed

	next, rest := pathStep(name)
	if next == "" {
		return
	}

	child, ok := n.ChildrenByName[next]
	if !ok {
		child = newTestNode()
		n.ChildrenByName[next] = child
	}

	child.MarkFailed(rest)
}

func pathStep(nodeName string) (next, rest string) {
	separatorIndex := strings.IndexRune(nodeName, '/')
	if separatorIndex < 0 {
		return nodeName, ""
	}

	return nodeName[:separatorIndex], nodeName[separatorIndex+1:]
}
