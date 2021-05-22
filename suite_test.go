package suite_test

import (
	"testing"

	"github.com/ibllex/suite"
	"github.com/stretchr/testify/assert"
)

type MockSuite struct {
	SetupSuiteExecutions    int
	TearDownSuiteExecutions int
	SetupTestExecutions     int
	TearDownTestExecutions  int
	TestExecutions          int
}

func (s *MockSuite) SetupSuite(t *testing.T) {
	s.SetupSuiteExecutions += 1
}

func (s *MockSuite) SetupTest(t *testing.T) {
	s.SetupTestExecutions += 1
}

func (s *MockSuite) TearDownSuite(t *testing.T) {
	s.TearDownSuiteExecutions += 1
}

func (s *MockSuite) TearDownTest(t *testing.T) {
	s.TearDownTestExecutions += 1
}

func (s *MockSuite) TestFunc1(t *testing.T) {
	s.TestExecutions += 1
}

func (s *MockSuite) TestFunc2(t *testing.T) {
	s.TestExecutions += 1
}

func TestSuite(t *testing.T) {
	s := &MockSuite{}
	suite.Run(t, s)

	assert.Equal(t, 1, s.SetupSuiteExecutions)
	assert.Equal(t, 1, s.TearDownSuiteExecutions)
	assert.Equal(t, 2, s.SetupTestExecutions)
	assert.Equal(t, 2, s.TearDownTestExecutions)
	assert.Equal(t, 2, s.TestExecutions)
}
