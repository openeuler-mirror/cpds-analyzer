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

package detector

import (
	"errors"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func SendRuleUpdatedRequset(detecotrHost string, detectorPort int) error {
	urlStr := fmt.Sprintf("http://%s:%d/api/v1/rule_updated", detecotrHost, detectorPort)
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := jsoniter.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	} else if data["status"] != float64(200) {
		return errors.New("cannot send rule updated requset to detector")
	}

	return nil
}
