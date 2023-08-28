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

package generic

import (
	"cpds/cpds-analyzer/pkg/utils/net"
	"fmt"

	"github.com/spf13/pflag"
)

type Options struct {
	BindAddress string `json:"bindAddress,omitempty" yaml:"bindAddress,omitempty"`
	Port        int    `json:"port,omitempty" yaml:"port,omitempty"`
}

func NewGenericOptions() *Options {
	return &Options{
		BindAddress: "0.0.0.0",
		Port:        19091,
	}
}

func (s *Options) Validate() []error {
	errs := []error{}

	if !net.IsValidIPAdress(s.BindAddress) {
		errs = append(errs, fmt.Errorf("wrong IP Address format: %s", s.BindAddress))
	}

	if !net.IsValidPort(s.Port) {
		errs = append(errs, fmt.Errorf("invalid port number: %d, should be 0 - 65535", s.Port))
	}

	return errs
}

func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.BindAddress, "bind-address", c.BindAddress, "server bind address")
	fs.IntVar(&s.Port, "port", c.Port, "insecure port number")
}
