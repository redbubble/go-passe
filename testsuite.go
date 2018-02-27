package main

type testSuite struct {
	TestsByPackage map[string]*testNode
}

func newTestSuite() *testSuite {
	return &testSuite{
		TestsByPackage: make(map[string]*testNode),
	}
}

func (s *testSuite) Get(id TestID) *testNode {
	return s.testsForPackage(id.Package).Get(id.Test)
}

func (s *testSuite) MarkFailed(id TestID) {
	s.testsForPackage(id.Package).MarkFailed(id.Test)
}

func (s *testSuite) MarkPassed(id TestID) {
	s.testsForPackage(id.Package).MarkPassed(id.Test)
}

func (s *testSuite) testsForPackage(packageName string) *testNode {
	testNode, ok := s.TestsByPackage[packageName]
	if !ok {
		testNode = newTestNode()
		s.TestsByPackage[packageName] = testNode
	}

	return testNode
}
