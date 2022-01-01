package module

import (
	"context"
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckTestSuite struct {
	suite.Suite
	cnl context.CancelFunc
	ctx context.Context
}

func TestCheckTestSuite(t *testing.T) {
	suite.Run(t, new(CheckTestSuite))
}

func (s *CheckTestSuite) SetupTest() {
	s.ctx, s.cnl = context.WithCancel(context.Background())
}

func (s *CheckTestSuite) TearDownTest() {
	s.cnl()
}

func (s *CheckTestSuite) Test_SelfCheck() {
	result, err := Check(s.ctx, "../..")
	s.NoError(err)
	s.NotEmpty(result)
}

func (s *CheckTestSuite) Test_SelfCheck_CancelContextBefore() {
	s.cnl()

	result, err := Check(s.ctx, "../..")
	s.EqualError(err, "context canceled")
	s.Empty(result)
}

//go:embed go.mod.input
var goModContent []byte

func (s *CheckTestSuite) Test_CustomModFile() {
	tempDir, err := ioutil.TempDir("", "test")
	s.NoError(err)

	tempFile := filepath.Join(tempDir, "go.mod")

	err = ioutil.WriteFile(tempFile, goModContent, 0666)
	s.NoError(err)

	result, err := Check(s.ctx, tempDir)
	s.NoError(err)
	s.NotEmpty(result)

	s.NoError(os.Remove(tempFile))
	s.NoError(os.RemoveAll(tempDir))
}
