package module

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

var content = []byte(`module github.com/ormanli/gomodctl

go 1.13

require github.com/stretchr/testify v1.1.1`)

type UpdateTestSuite struct {
	suite.Suite
	cnl      context.CancelFunc
	ctx      context.Context
	tempFile string
	tempDir  string
}

func TestUpdateTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}

func (s *UpdateTestSuite) SetupTest() {
	s.ctx, s.cnl = context.WithCancel(context.Background())

	tmpddir, err := ioutil.TempDir("", "test")
	s.NoError(err)
	s.tempDir = tmpddir

	s.tempFile = filepath.Join(s.tempDir, "go.mod")

	err = ioutil.WriteFile(s.tempFile, content, 0666)
	s.NoError(err)
}

func (s *UpdateTestSuite) TearDownTest() {
	s.cnl()
	os.Remove(s.tempFile)
	os.RemoveAll(s.tempDir)
}

func (s *UpdateTestSuite) Test_Update() {
	update, err := Update(s.ctx, s.tempDir)

	s.NoError(err)
	s.NotEmpty(update)
	s.True(update["github.com/stretchr/testify"].LocalVersion.LessThan(update["github.com/stretchr/testify"].LatestVersion))

	file, err := ioutil.ReadFile(filepath.Join(s.tempDir, "go.mod"))
	s.NoError(err)
	s.NotEqual(content, string(file))
}
