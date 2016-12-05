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
	"github.com/intelsdi-x/snap/core"
)

func memStat(ns core.Namespace, swagURL string) (*plugin.MetricType, error) {
	memType := ns.Strings()[3]
	switch {
	case regexp.MustCompile(`^/` + Vendor + `/` + Name + `/memory/free`).MatchString(ns.String()):
		metric, err := getMemStat(swagURL, memType)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      metric,
			Timestamp_: time.Now(),
		}, nil

	case regexp.MustCompile(`^/` + Vendor + `/` + Name + `/memory/total`).MatchString(ns.String()):
		metric, err := getMemStat(swagURL, memType)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      metric,
			Timestamp_: time.Now(),
		}, nil

	}

	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getMemoryMetricTypes() ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	for _, metricType := range memMetrics {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace(Vendor, Name, "memory", metricType)})
	}
	return mts, nil
}

func getMemStat(swagURL string, memType string) (uint64, error) {
	path := fmt.Sprintf("os/memory/%s", memType)
	response, err := osvRestGet(swagURL, path)
	if err != nil {
		return 0, err
	}
	metric, err := strconv.ParseUint(string(response), 10, 0)
	if err != nil {
		return 0, err
	}

	return metric, nil
}
