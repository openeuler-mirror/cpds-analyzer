/* 
 *  Copyright 2023 CPDS Author
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *       https://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package options

import (
	analyzer "cpds/cpds-analyzer/pkg/cpds-analyzer"
	"cpds/cpds-analyzer/pkg/cpds-analyzer/config"

	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	ConfigFile string
	*config.Config

	DebugMode bool
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Config: config.New(),
	}
}

func (s *ServerRunOptions) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("generic", pflag.ExitOnError)
	fs.StringVar(&s.ConfigFile, "config", config.DefaultConfigurationPath, "Directory where configuration files are stored")
	fs.BoolVar(&s.DebugMode, "debug", false, "Don't enable this if you don't know what it means.")

	return fs
}

func (s *ServerRunOptions) NewAnalyzer() (*analyzer.Analyzer, error) {
	analyzer := &analyzer.Analyzer{
		Config: s.Config,
		Debug:  s.DebugMode,
	}
	return analyzer, nil
}
