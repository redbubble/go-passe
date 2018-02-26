package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/logrusorgru/aurora"
)

func printIndent(level int) {
	for i := 0; i <= level; i++ {
		fmt.Print("  ")
	}
}

func printFailedTests(name string, tests *testNode, level int) {
	if tests.Passed {
		return
	}

	if level == 0 {
		fmt.Println()
	}

	printIndent(level)
	fmt.Println(aurora.Bold(aurora.Red("✗")), name)

	if tests.Output != nil {
		fmt.Println()
		for _, line := range tests.Output {
			printIndent(level + 1)
			fmt.Println(aurora.Red(line))
		}
		fmt.Println()
	}

	for childName, child := range tests.ChildrenByName {
		printFailedTests(childName, child, level+1)
	}
}

func main() {
	jsonDecoder := json.NewDecoder(os.Stdin)

	var actionPattern = regexp.MustCompile(`^\S+`)
	var testSuite = newTestSuite()
	var allPassed = true

	for {
		var event TestEvent
		err := jsonDecoder.Decode(&event)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			break
		}

		action := actionPattern.FindString(event.Action)

		switch action {
		case "pass":
			if event.Test != "" {
				fmt.Printf("%s %s %s %gs\n", aurora.Bold(aurora.Green("✓")), event.Package, event.Test, event.Elapsed)
			} else {
				fmt.Printf("%s %s %gs\n", aurora.Bold(aurora.Green("✓")), event.Package, event.Elapsed)
			}
			testSuite.Get(event.TestID).Passed = true

		case "fail":
			if event.Test != "" {
				fmt.Printf("%s %s %s %gs\n", aurora.Bold(aurora.Red("✗")), event.Package, event.Test, event.Elapsed)
			} else {
				fmt.Printf("%s %s %gs\n", aurora.Bold(aurora.Red("✗")), event.Package, event.Elapsed)
			}
			testSuite.Get(event.TestID).Passed = false
			allPassed = false

		case "output":
			testSuite.Get(event.TestID).AppendOutput(event.Output)

		default:
		}
	}

	if allPassed {
		fmt.Println(aurora.Bold(aurora.Green("\nAll tests passed")))
	} else {
		fmt.Println(aurora.Bold(aurora.Red("\nTest failures:")))
		for packageName, tests := range testSuite.TestsByPackage {
			printFailedTests(packageName, tests, 0)
		}

		os.Exit(1)
	}
}
