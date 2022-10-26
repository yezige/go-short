package config

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "test"
	_, err := LoadConf(filename)

	assert.NotNil(t, err)
}

func TestEmptyConfig(t *testing.T) {
	conf, err := LoadConf("testdata/empty.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}

	assert.Equal(t, uint(100), conf.Core.Port)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfGorushDefault *ConfYaml
	ConfGorush        *ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	var err error
	suite.ConfGorushDefault, err = LoadConf()
	if err != nil {
		panic("failed to load default config.yml")
	}
	suite.ConfGorush, err = LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// Core
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.Address)
	assert.Equal(suite.T(), "8088", suite.ConfGorushDefault.Core.Port)
	assert.Equal(suite.T(), int64(30), suite.ConfGorushDefault.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	assert.Equal(suite.T(), "release", suite.ConfGorushDefault.Core.Mode)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Enabled)
	assert.Equal(suite.T(), "app.pid", suite.ConfGorushDefault.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.PID.Override)

	// Api
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorushDefault.API.StatGoURI)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorushDefault.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorushDefault.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorushDefault.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorushDefault.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorushDefault.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Log.HideToken)

}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	assert.Equal(suite.T(), "8088", suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), int64(30), suite.ConfGorush.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorush.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.PID.Override)

	// Api
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorush.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorush.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorush.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorush.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorush.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorush.Log.HideToken)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("GORUSH_CORE_PORT", "9001")
	os.Setenv("GORUSH_GRPC_ENABLED", "true")
	os.Setenv("GORUSH_CORE_MAX_NOTIFICATION", "200")
	os.Setenv("GORUSH_IOS_KEY_ID", "ABC123DEFG")
	os.Setenv("GORUSH_IOS_TEAM_ID", "DEF123GHIJ")
	os.Setenv("GORUSH_API_HEALTH_URI", "/healthz")
	ConfGorush, err := LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
	assert.Equal(t, "9001", ConfGorush.Core.Port)
}

func TestLoadWrongDefaultYAMLConfig(t *testing.T) {
	defaultConf = []byte(`a`)
	_, err := LoadConf()
	assert.Error(t, err)
}
