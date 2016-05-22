package functional_test

import "github.com/revel/revel/testing"

// AppTest is a test for the app
type AppTest struct {
	testing.TestSuite
}

// Before sets up tests
func (t *AppTest) Before() {
	println("Set up")
}

// TestThatIndexPageWorks does a basic test on the index template
func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After cleans up from tests
func (t *AppTest) After() {
	println("Tear down")
}
