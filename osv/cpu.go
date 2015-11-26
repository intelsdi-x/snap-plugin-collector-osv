/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package osv

import (
	"strconv"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
)

func cpuStat(ns []string, swag_url string) (*plugin.PluginMetricType, error) {
	metric, err := getCpuTime(swag_url)
	if err != nil {
		return nil, err
	}
	return &plugin.PluginMetricType{
		Namespace_: ns,
		Data_:      strconv.FormatUint(metric, 10),
		Timestamp_: time.Now(),
	}, nil

}

func getCpuMetricTypes() ([]plugin.PluginMetricType, error) {
	mts := make([]plugin.PluginMetricType, 0)
	for _, metric_type := range cpu_metrics {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "cpu", metric_type}})
	}
	return mts, nil
}

func getCpuTime(swag_url string) (uint64, error) {
	path := "trace/count"
	response, err := osvRestGet(swag_url, path)
	if err != nil {
		return 0, err
	}
	counters, err := osvRestUnmarshall(response)
	if err != nil {
		return 0, err
	}

	return counters.Time_ms, nil
}
