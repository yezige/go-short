package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  enabled: true # enable httpd server
  address: "" # ip address to bind (default: any)
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0 # default worker number is runtime.NumCPU()
  mode: "release" # release mode or debug mode
  shutdown_timeout: 30 # default is 30 second
  pid:
    enabled: true
    path: "app.pid"
    override: true

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "debug"
  hide_token: true

redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 0

`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core  SectionCore  `yaml:"core" mapstructure:"core" json:"core"`
	Log   SectionLog   `yaml:"log" mapstructure:"log" json:"log"`
	Redis SectionRedis `yaml:"redis" mapstructure:"redis" json:"redis"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Enabled         bool       `yaml:"enabled" mapstructure:"enabled" json:"enabled"`
	Address         string     `yaml:"address" mapstructure:"address" json:"address"`
	Port            string     `yaml:"port" mapstructure:"port" json:"port"`
	WorkerNum       int64      `yaml:"worker_num" mapstructure:"worker_num" json:"worker_num"`
	Mode            string     `yaml:"mode" mapstructure:"mode" json:"mode"`
	ShutdownTimeout int64      `yaml:"shutdown_timeout" mapstructure:"shutdown_timeout" json:"shutdown_timeout"`
	PID             SectionPID `yaml:"pid" mapstructure:"pid" json:"pid"`
}

type SectionPID struct {
	Enabled  bool   `yaml:"enabled"`
	Path     string `yaml:"path"`
	Override bool   `yaml:"override"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format      string `yaml:"format" mapstructure:"format" json:"format"`
	AccessLog   string `yaml:"access_log" mapstructure:"access_log" json:"access_log"`
	AccessLevel string `yaml:"access_level" mapstructure:"access_level" json:"access_level"`
	ErrorLog    string `yaml:"error_log" mapstructure:"error_log" json:"error_log"`
	ErrorLevel  string `yaml:"error_level" mapstructure:"error_level" json:"error_level"`
	HideToken   bool   `yaml:"hide_token" mapstructure:"hide_token" json:"hide_token"`
}

type SectionRedis struct {
	Addr     string `yaml:"addr" mapstructure:"addr" json:"addr"`
	Password string `yaml:"password" mapstructure:"password" json:"password"`
	DB       int    `yaml:"db" mapstructure:"db" json:"db"`
}

func setDefault() {
}

var conf *ConfYaml

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath ...string) (*ConfYaml, error) {
	conf = &ConfYaml{}

	// load default values
	setDefault()

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()        // read in environment variables that match
	viper.SetEnvPrefix("short") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if len(confPath) > 0 && confPath[0] != "" {
		content, err := ioutil.ReadFile(confPath[0])
		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
		fmt.Println("Using config file:", confPath)
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		viper.AddConfigPath("/etc/goshort/")
		viper.AddConfigPath("$HOME/.goshort")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
			// load default config
			return conf, err
		}
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	if conf.Core.WorkerNum == int64(0) {
		conf.Core.WorkerNum = int64(runtime.NumCPU())
	}

	return conf, nil
}

func GetConfig() (*ConfYaml, error) {
	return conf, nil
}
