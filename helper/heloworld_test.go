package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestSubTest(t *testing.T) {
	t.Run("Adam", func(t *testing.T) {
		result := HelloWorld("Adam")
		require.Equal(t, "Hello Adam", result, "Result must be 'Hello Adam'")
	})

	// if want to execute specific sub test, type "go test -v -run TestName/SubTestName"
	// if want to execute all sub test, type "go test -v -run /SubTestName"
	t.Run("Eve", func(t *testing.T) {
		result := HelloWorld("Eve")
		require.Equal(t, "Hello Eve", result, "Result must be 'Hello Adam'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unite Test")
	// eg: connection db, configure etc
	m.Run() // only run in one package
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows") // skipping test
	}
	result := HelloWorld("Adam")
	require.Equal(t, "Hello Adam", result, "Result must be 'Hello Adam'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adam")
	require.Equal(t, "Hello Adam", result, "Result must be 'Hello Adam'")
	fmt.Println("TestHelloWorld done")
}

// for assert testing, using testify package

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Adam")
	assert.Equal(t, "Hello Adam", result, "Result must be 'Hello Adam'")
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Adam")
	if result != "Hello Adam" {
		//panic("Result is not Hello Adam") // don't use panic for failing case the test
		//t.Fail() // no arg
		t.Error("Result is not Hello Adam") // with arg message
		// if use t.Fatal() it won't execute the code below
	}
	fmt.Println("TestHelloWorld done") // if use t.Fatal() this line will not execute
	// for running the test, type "go test" in terminal
	// for detail function, type "go test -v"
	// if want to test specific function, type "go test -v -run 'function name'" *without single quote
	// if want to test from all package, type "go test -v ./..."

	//==============================================================

	// for failing the test
	// there's function Fail(), FailNow(), and Fatal()
	// open the package doc for details
	//t.Fail()
	//t.FailNow()
	//t.Error()
	//t.Fatal()
}

func TestUserGradeA(t *testing.T) {
	// if want to test specific function with grouping specific name, type "go test -v -run='function name %'" || eg: go test -v -run=TestUser | means, every naming function with TestUser in front word will execute
	//D:\workspace\src\golang-unit-test\helper>go test -v -run=TestUser
	//=== RUN   TestUserGradeA
	//--- PASS: TestUserGradeA (0.00s)
	//=== RUN   TestUserGradeB
	//--- PASS: TestUserGradeB (0.00s)
	//PASS
	//ok      golang-unit-test/helper 0.166s

}
