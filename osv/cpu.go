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
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
)

func cpuStat(ns core.Namespace, swagURL string) (*plugin.MetricType, error) {
	metric, err := getCPUTime(swagURL)
	if err != nil {
		return nil, err
	}
	return &plugin.MetricType{
		Namespace_: ns,
		Data_:      metric,
		Timestamp_: time.Now(),
	}, nil
}

func getCPUMetricTypes() ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	for _, metricType := range cpuMetrics {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace(Vendor, Name, "cpu", metricType)})
	}
	return mts, nil
}

func getCPUTime(swagURL string) (uint64, error) {
	path := "trace/count"
	response, err := osvRestGet(swagURL, path)
	if err != nil {
		return 0, err
	}
	counters, err := osvRestUnmarshall(response)
	if err != nil {
		return 0, err
	}

	return counters.TimeMs, nil
}
