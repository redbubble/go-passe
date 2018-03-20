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

func printTests(name string, tests *testNode, level int, states ...TestState) {

	for _, s := range states {

		if tests.State == s {
			if level == 0 {
				fmt.Println()
			}

			printIndent(level)

			switch tests.State {
			case Unknown:
				fmt.Println(aurora.Bold(aurora.Cyan("?")), name)

			case Passed:
				fmt.Println(aurora.Bold(aurora.Green("✓")), name)

			case Failed:
				fmt.Println(aurora.Bold(aurora.Red("✗")), name)

				if tests.Output != nil {
					fmt.Println()
					for _, line := range tests.Output {
						printIndent(level + 1)
						fmt.Println(aurora.Red(line))
					}
					fmt.Println()
				}
			}

			break
		}
	}

	for childName, child := range tests.ChildrenByName {
		printTests(childName, child, level+1, states...)
	}
}

func main() {
	jsonDecoder := json.NewDecoder(os.Stdin)

	var actionPattern = regexp.MustCompile(`^\S+`)
	var testSuite = newTestSuite()
	var allPassed = true
	var eventCount = 0

	for {
		var event TestEvent
		err := jsonDecoder.Decode(&event)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error: %v\n\nForgot to pass -json to 'go test'?\n", err)
				os.Exit(2)
			}
			if eventCount == 0 {
				fmt.Fprintln(os.Stderr, "Error: No Go test events detected")
				os.Exit(3)
			}
			break
		}

		eventCount++

		action := actionPattern.FindString(event.Action)

		switch action {
		case "pass":
			if event.Test != "" {
				fmt.Printf("%s %s %s %gs\n", aurora.Bold(aurora.Green("✓")), event.Package, event.Test, event.Elapsed)
			} else {
				fmt.Printf("%s %s %gs\n", aurora.Bold(aurora.Green("✓")), event.Package, event.Elapsed)
			}
			testSuite.MarkPassed(event.TestID)

		case "fail":
			if event.Test != "" {
				fmt.Printf("%s %s %s %gs\n", aurora.Bold(aurora.Red("✗")), event.Package, event.Test, event.Elapsed)
			} else {
				fmt.Printf("%s %s %gs\n", aurora.Bold(aurora.Red("✗")), event.Package, event.Elapsed)
			}
			testSuite.MarkFailed(event.TestID)
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
			printTests(packageName, tests, 0, Failed)
		}

		os.Exit(1)
	}
}
