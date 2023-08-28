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

package logger

import "github.com/spf13/pflag"

type Options struct {
	FileName   string `json:"fileName,omitempty" yaml:"fileName,omitempty"`
	Level      string `json:"level,omitempty" yaml:"level,omitempty"`
	MaxAge     int    `json:"maxAge,omitempty" yaml:"fileName,omitempty"`
	MaxBackups int    `json:"maxBackups,omitempty" yaml:"maxBackups,omitempty"`
	MaxSize    int    `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	LocalTime  bool   `json:"localTime,omitempty" yaml:"localTime,omitempty"`
	Compress   bool   `json:"compress,omitempty" yaml:"compress,omitempty"`
}

func NewLoggerOptions() *Options {
	return &Options{
		FileName:   "cpds-analyzer.log",
		Level:      "info",
		MaxAge:     15,
		MaxBackups: 7,
		MaxSize:    100,
		LocalTime:  true,
		Compress:   true,
	}
}

func (s *Options) Validate() []error {
	errs := []error{}

	return errs
}

func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.FileName, "log-file", c.FileName, "log file name")
	fs.StringVar(&s.Level, "log-level", c.Level, `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	fs.IntVar(&s.MaxAge, "log-max-age", c.MaxAge, "log max age")
	fs.IntVar(&s.MaxBackups, "log-max-backups", c.MaxBackups, "log max backups")
	fs.IntVar(&s.MaxSize, "log-max-size", c.MaxSize, "log max size")
}
