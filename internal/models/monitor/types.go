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

package monitor

type MonitorTargets struct {
	Targets []struct {
		Instance string `json:"instance"`
		Status   string `json:"status"`
	} `json:"targets"`
}

type NodeInfo struct {
	Instance      string `json:"instance"`
	Arch          string `json:"arch"`
	KernelVersion string `json:"kernal_version"`
	OSVersion     string `json:"os_version"`
}

type NodeStatus struct {
	Instance  string `json:"instance"`
	Container struct {
		Running int `json:"running"`
		Total   int `json:"total"`
	} `json:"container"`
	Cpu struct {
		Usage     float64 `json:"usage"`
		UsedCore  float64 `json:"used_core"`
		TotalCore float64 `json:"total_core"`
		NumberCores float64 `json:"number_cores"`
	} `json:"cpu"`
	Memory struct {
		Usage      float64 `json:"usage"`
		UsedBytes  float64 `json:"used_bytes"`
		TotalBytes float64 `json:"total_bytes"`
	} `json:"memory"`
	Disk struct {
		Usage      float64 `json:"usage"`
		UsedBytes  float64 `json:"used_bytes"`
		TotalBytes float64 `json:"total_bytes"`
	} `json:"disk"`
}
type NodeResource struct {
}

type NodeContainerStatus struct {
}
