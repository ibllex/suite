# Suite

Simple and lightweight test suite for Go.

## Features

1. Support test group.
2. Setup and TearDown hooks for each suite and test case.
3. Use the SUITE_RUN environment variable to run the specified test.

## Quick Start

```go
package suite_test

import (
	"testing"

	"github.com/ibllex/suite"
)

// UserTest is a valid test suite.
// You don’t need to implement all the suite hooks.
// This example is just for completeness.
// So, just implement what you need.
type UserTest struct{}

// SetupSuite will run once before all test cases are run
func (s *UserTest) SetupSuite(t *testing.T) {
	//
}

// TearDownSuite will run once after all test cases are run
func (s *UserTest) TearDownSuite(t *testing.T) {
	//
}

// SetupTest will run before each test case is run
func (s *UserTest) SetupTest(t *testing.T) {
	//
}

// TearDownTest will run after each test case is run
func (s *UserTest) TearDownTest(t *testing.T) {
	//
}

// Every test case function in a suite should start with `Test` prefix
func (s *UserTest) TestCreateUser(t *testing.T) {
	//
}

func (s *UserTest) TestDeleteUser(t *testing.T) {
	//
}

func (s *UserTest) TestUpdateUser(t *testing.T) {
	//
}

// TestUser run the suite
func TestUser(t *testing.T) {
	suite.Run(t, &UserTest{})
}
```

## License

This library is under the [MIT](https://github.com/ibllex/suite/blob/main/LICENSE) license.
