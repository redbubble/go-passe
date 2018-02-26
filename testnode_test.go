package main

import (
	"testing"
)

func TestTestNode(t *testing.T) {

	t.Run(".AppendOutput()", func(t *testing.T) {

		t.Run("strips leading and trailing whitespace", func(t *testing.T) {
			n := newTestNode()

			n.AppendOutput("  \t  some output")
			n.AppendOutput("other output  ")

			if actual, expected := len(n.Output), 2; actual != expected {
				t.Fatalf("Expected %d lines of output but got %d", expected, actual)
			}

			if actual, expected := n.Output[0], "some output"; actual != expected {
				t.Errorf("Expected first line of output to be '%s' but was '%s'", expected, actual)
			}
			if actual, expected := n.Output[1], "other output"; actual != expected {
				t.Errorf("Expected second line of output to be '%s' but was '%s'", expected, actual)
			}
		})

		t.Run("ignores lines generated by 'go test' command", func(t *testing.T) {
			n := newTestNode()

			n.AppendOutput("    FAIL")
			n.AppendOutput("    FAIL\tsome/package/test\t0.103s")
			n.AppendOutput("  === RUN   SomeTest")
			n.AppendOutput("  --- FAIL: SomeTest (0.00s)")
			n.AppendOutput("    exit status 1")

			if actual, expected := len(n.Output), 0; actual != expected {
				t.Errorf("Expected all lines of output to be filtered out but got %#v", n.Output)
			}
		})
	})

	t.Run(".Get()", func(t *testing.T) {

		t.Run("creates children for each path step", func(t *testing.T) {
			n := newTestNode()

			child := n.Get("apple/banana/pear")

			apple := n.ChildrenByName["apple"]
			if apple == nil {
				t.Fatalf("Expected a child to be created for path step 'apple' but was nil")
			}

			banana := apple.ChildrenByName["banana"]
			if banana == nil {
				t.Fatalf("Expected a child to be created for path step 'banana' but was nil")
			}

			pear := banana.ChildrenByName["pear"]
			if pear == nil {
				t.Fatalf("Expected a child to be created for path step 'pear' but was nil")
			}

			if pear != child {
				t.Errorf("Expected final child to be returned but was different")
			}
		})

		t.Run("returns an existing child addressed by the path", func(t *testing.T) {
			n := newTestNode()

			child1 := n.Get("apple/banana")
			child2 := n.Get("apple/banana")

			if child1 != child2 {
				t.Fatalf("Expected same child to be returned for same path but was different")
			}
		})
	})

	t.Run("pathStep()", func(t *testing.T) {

		t.Run("separates a node name into the next path step and the remainder", func(t *testing.T) {
			next, rest := pathStep("apple/banana/pear")

			if expected := "apple"; next != expected {
				t.Errorf("Expected next path step to be '%s' but got '%s'", expected, next)
			}
			if expected := "banana/pear"; rest != expected {
				t.Errorf("Expected rest of path to be '%s' but got '%s'", expected, rest)
			}
		})

		t.Run("returns empty string as the remainder if there are no more remaining steps", func(t *testing.T) {
			next, rest := pathStep("apple")

			if expected := "apple"; next != expected {
				t.Errorf("Expected next path step to be '%s' but got '%s'", expected, next)
			}
			if expected := ""; rest != expected {
				t.Errorf("Expected rest of path to be '%s' but got '%s'", expected, rest)
			}
		})

		t.Run("returns empty string as the next step if there are no more steps", func(t *testing.T) {
			next, rest := pathStep("")

			if expected := ""; next != expected {
				t.Errorf("Expected next path step to be '%s' but got '%s'", expected, next)
			}
			if expected := ""; rest != expected {
				t.Errorf("Expected rest of path to be '%s' but got '%s'", expected, rest)
			}
		})
	})
}