package suite

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	// SetupSuite is called before any tests runs in a test suite.
	SetupSuite = "SetupSuite"
	// TearDownSuite is called after all tests are completed in a test suite.
	TearDownSuite = "TearDownSuite"
	// SetupTest is called before each test runs.
	SetupTest = "SetupTest"
	// TearDownTest is called after each test runs.
	TearDownTest = "TearDownTest"
	// TestPrefix every test case function in a suite should start with TestPrefix
	TestPrefix = "Test"
)

func getTestFunc(t *testing.T, sv reflect.Value, name string) func(*testing.T) {
	if m := sv.MethodByName(name); m.IsValid() {
		if f, ok := m.Interface().(func(*testing.T)); ok {
			return f
		}
		// Method exists but has the wrong type signature.
		t.Fatalf("suite: function %v has unexpected signature (%T)", name, m.Interface())
	}
	return func(*testing.T) {}
}

// Run runs all "Test___" functions and test suite hook of a suite
// which is implements Suite interface
func Run(t *testing.T, suite interface{}) {
	st := reflect.TypeOf(suite)
	sv := reflect.ValueOf(suite)
	run := os.Getenv("SUITE_RUN")

	getTestFunc(t, sv, SetupSuite)(t)

	for i := 0; i < st.NumMethod(); i++ {
		methodName := st.Method(i).Name
		if !strings.HasPrefix(methodName, TestPrefix) {
			continue
		}

		if run != "" && methodName != run {
			continue
		}

		tfunc := getTestFunc(t, sv, methodName)
		t.Run(strings.TrimPrefix(methodName, TestPrefix), func(t *testing.T) {
			getTestFunc(t, sv, SetupTest)(t)
			tfunc(t)
			getTestFunc(t, sv, TearDownTest)(t)
		})
	}

	getTestFunc(t, sv, TearDownSuite)(t)
}
