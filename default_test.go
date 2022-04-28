package logger

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DefaultLoggerTestSuite struct {
	suite.Suite
}

func (suite *DefaultLoggerTestSuite) SetupTest() {
}

func (suite *DefaultLoggerTestSuite) TestDebug() {
	DevMode = true
	DebugT()
}

func (suite *DefaultLoggerTestSuite) TestInfo() {
	InfoT()
}

func (suite *DefaultLoggerTestSuite) TestError() {
	ErrorT()
}

func (suite *DefaultLoggerTestSuite) TestWith() {
	WithT()
}

func TestDefaultLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultLoggerTestSuite))
}

func DebugT() {
	Debug("This is debug", Field{"nonce", "001"})
}

func InfoT() {
	nonce := 100
	Info("", Field{"nonce", nonce})
}

func ErrorT() {
	Error("This is error", Field{"request id", "0xxxxxx"})
}

func WithT() {
	With(Field{"with-key", "with key value"}).Info("This is info with")
}
