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
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
)

func memStat(ns []string, swag_url string) (*plugin.PluginMetricType, error) {
	mem_type := ns[2]
	switch {
	case regexp.MustCompile(`^/osv/memory/free`).MatchString(joinNamespace(ns)):
		metric, err := getMemStat(swag_url, mem_type)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatUint(metric, 10),
			Timestamp_: time.Now(),
		}, nil

	case regexp.MustCompile(`^/osv/memory/total`).MatchString(joinNamespace(ns)):
		metric, err := getMemStat(swag_url, mem_type)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatUint(metric, 10),
			Timestamp_: time.Now(),
		}, nil

	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getMemoryMetricTypes() ([]plugin.PluginMetricType, error) {
	mts := make([]plugin.PluginMetricType, 0)
	for _, metric_type := range mem_metrics {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "memory", metric_type}})
	}
	return mts, nil
}

func getMemStat(swag_url string, mem_type string) (uint64, error) {
	path := fmt.Sprintf("os/memory/%s", mem_type)
	response, err := osvRestGet(swag_url, path)
	if err != nil {
		return 0, err
	}
	metric, err := strconv.ParseUint(string(response), 10, 0)
	if err != nil {
		return 0, err
	}

	return metric, nil
}
