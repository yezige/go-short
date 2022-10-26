package logx

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yezige/go-short/config"
)

func TestSetLogLevel(t *testing.T) {
	log := logrus.New()

	err := SetLogLevel(log, "debug")
	assert.Nil(t, err)

	err = SetLogLevel(log, "invalid")
	assert.Equal(t, "not a valid logrus Level: \"invalid\"", err.Error())
}

func TestSetLogOut(t *testing.T) {
	log := logrus.New()

	err := SetLogOut(log, "stdout")
	assert.Nil(t, err)

	err = SetLogOut(log, "stderr")
	assert.Nil(t, err)

	err = SetLogOut(log, "log/access.log")
	assert.Nil(t, err)

	// missing create logs folder.
	err = SetLogOut(log, "logs/access.log")
	assert.NotNil(t, err)
}

func TestInitDefaultLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	// no errors on default config
	assert.Nil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))

	cfg.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))

	isTerm = true

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestAccessLevel(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestErrorLevel(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.ErrorLevel = "invalid"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestAccessLogPath(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.AccessLog = "logs/access.log"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestErrorLogPath(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.ErrorLog = "logs/error.log"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestHideToken(t *testing.T) {
	assert.Equal(t, "", hideToken("", 2))
	assert.Equal(t, "**345678**", hideToken("1234567890", 2))
	assert.Equal(t, "*****", hideToken("12345", 10))
}

func TestLogPushEntry(t *testing.T) {
	in := InputLog{}

	in.Error = errors.New("error")
	assert.Equal(t, "error", GetLogPushEntry(&in).Error)

	in.Token = "1234567890"
	in.HideToken = true
	assert.Equal(t, "**********", GetLogPushEntry(&in).Token)
}

func TestLogPush(t *testing.T) {
	in := InputLog{}
	isTerm = true

	in.Format = "json"
	in.Status = "succeeded"
	assert.Equal(t, "succeeded", LogPush(&in).Type)

	in.Format = ""
	in.Message = "success"
	assert.Equal(t, "success", LogPush(&in).Message)

	in.Status = "failed"
	in.Message = "failed"
	assert.Equal(t, "failed", LogPush(&in).Message)
}
