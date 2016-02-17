// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/kr/pretty"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type Config struct {
	*viper.Viper
}

func New() Config {
	cfg := Config{viper.New()}

	///// INIT and RUN
	// Used for "init" command.
	// "earthquake-container" ignores this.
	// e.g. "init.sh"
	cfg.SetDefault("init", "")

	// Used for "run" command.
	// "earthquake-container" ignores this.
	// e.g. "run.sh"
	cfg.SetDefault("run", "")

	// Used for "run" command.
	// "earthquake-container" ignores this.
	// e.g. "clean.sh"
	cfg.SetDefault("clean", "")

	// Used for "run" command.
	// "earthquake-container" ignores this.
	// e.g. "validate.sh"
	cfg.SetDefault("validate", "")

	// Used for something deprecated?
	// if true, skip clean.sh when validate.sh failed.
	// "earthquake-container" ignores this.
	cfg.SetDefault("notCleanIfValidationFail", false)

	///// STORAGE
	// "earthquake-container" ignores this.
	// Used for "run" command
	cfg.SetDefault("storageType", "naive")

	///// INSPECTOR HANDLER ENDPOINT
	// "earthquake-container" ignores these values.
	// Used for PB inspector handler (used by Java and C inspector)
	// if non-positive, PB inspector handler is disabled
	// e.g. 10000
	cfg.SetDefault("pbPort", 0)

	// Used for REST inspector handler
	// if non-positive, REST inspector handler is disabled
	// e.g. 10080
	cfg.SetDefault("restPort", 0)

	///// EXPLORATION POLICY
	// "earthquake-container" also uses these params
	cfg.SetDefault("explorePolicy", "random")
	cfg.SetDefault("explorePolicyParam", map[string]interface{}{})

	///// CONTAINER
	// used only in "earthquake-container"
	cfg.SetDefault("container", map[string]interface{}{
		"enableEthernetInspector": false,
		"enableProcInspector":     true,
		"ethernetNFQNumber":       42,
		"procWatchInterval":       time.Second,
	})
	return cfg
}

func NewFromString(s string, typ string) (Config, error) {
	cfg := New()
	cfg.SetConfigType(typ)
	err := cfg.ReadConfig(strings.NewReader(s))
	return cfg, err
}

func NewFromFile(filePath string) (Config, error) {
	// viper supports JSON, YAML, and TOML
	cfg := New()
	cfg.SetConfigFile(filePath)
	err := cfg.ReadInConfig()
	return cfg, err
}

func (cfg Config) String() string {
	m := make(map[string]interface{})
	err := cfg.Unmarshal(&m)
	if err != nil {
		panic(err)
	}
	return pretty.Sprintf("Config{%# v}", m)
}