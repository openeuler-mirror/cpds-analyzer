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

package time

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsValidDuration(input string) bool {
	// 定义时间单位正则表达式
	re := regexp.MustCompile(`^\d+([a-z]+)$`)

	// 判断是否符合时间格式
	if !re.MatchString(strings.ToLower(strings.TrimSpace(input))) {
		return false
	}

	// 检查时间单位是否有效
	validUnits := []string{"ns", "us", "µs", "ms", "s", "m", "h", "d", "w", "mo", "y"}
	unit := input[len(input)-1:]
	for _, u := range validUnits {
		if u == unit {
			return true
		}
	}

	return false
}

func ParseDuration(input string) (time.Duration, error) {
	// 定义所有可能的时间单位及其对应的时间间隔
	units := map[string]time.Duration{
		"ns": time.Nanosecond,
		"us": time.Microsecond,
		"µs": time.Microsecond,
		"ms": time.Millisecond,
		"s":  time.Second,
		"m":  time.Minute,
		"h":  time.Hour,
		"d":  time.Hour * 24,
		"w":  time.Hour * 24 * 7,
		"mo": time.Hour * 24 * 30,
		"y":  time.Hour * 24 * 365,
	}

	// 编译正则表达式，以从输入字符串中提取数字和单位
	re := regexp.MustCompile(`(\d+)(\D+)`)
	matches := re.FindStringSubmatch(strings.ToLower(input))
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid duration format: %s", input)
	}

	// 解析数字和单位
	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	unit, ok := units[matches[2]]
	if !ok {
		return 0, fmt.Errorf("invalid duration unit: %s", matches[2])
	}

	return time.Duration(number) * unit, nil
}

func IsTimestamp(timestamp int64) bool {
	min := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC).Unix()

	if timestamp >= min && timestamp <= max {
		return true
	}

	return false
}
