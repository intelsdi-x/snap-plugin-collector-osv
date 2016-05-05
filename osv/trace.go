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
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
)

// Counter struct for unmarshalled json stucture
type Counter struct {
	Name  string
	Count uint64
}

// Counters struct for unmarshalled json stucture
type Counters struct {
	TimeMs uint64 `json:"time_ms"`
	List   []Counter
}

func getCounterMetricTypes() ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	for _, counter := range virtioCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "virtio", counter.Value)})
	}
	for _, counter := range netCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "net", counter.Value)})
	}
	for _, counter := range memoryCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "memory", counter.Value)})
	}
	for _, counter := range calloutCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "callout", counter.Value)})
	}
	for _, counter := range waitCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "wait", counter.Value)})
	}
	for _, counter := range asyncCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "async", counter.Value)})
	}
	for _, counter := range vfsCounters {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("osv", "trace", "vfs", counter.Value)})
	}
	return mts, nil
}

func traceStat(ns core.Namespace, swagURL string) (*plugin.MetricType, error) {
	trace := ns[3].Value
	metric, err := getTrace(trace, swagURL)
	if err != nil {
		return nil, err
	}
	return &plugin.MetricType{
		Namespace_: ns,
		Data_:      strconv.FormatUint(metric, 10),
		Timestamp_: time.Now(),
	}, nil
}

func osvRestUnmarshall(response []byte) (Counters, error) {
	var counters Counters
	if err := json.Unmarshal(response, &counters); err != nil {
		return counters, err
	}
	return counters, nil
}

func parseResult(counters Counters, trace string) (uint64, error) {
	for _, counter := range counters.List {
		if counter.Name == trace {
			return counter.Count, nil
		}
	}
	return 0, fmt.Errorf("Can't find %s in trace list", trace)

}

func osvRestCall(trace string, swagURL string, recovery bool) (uint64, error) {
	path := "trace/count"
	if recovery {
		recoveryPath := fmt.Sprintf("%s/%s?enabled=true", path, trace)
		err := osvRestPost(swagURL, recoveryPath)
		if err != nil {
			return 0, err
		}
	}
	response, err := osvRestGet(swagURL, path)
	if err != nil {
		return 0, err
	}

	counters, err := osvRestUnmarshall(response)
	if err != nil {
		return 0, err
	}
	metric, err := parseResult(counters, trace)
	if err != nil {
		return 0, err
	}
	return metric, nil

}

func getTrace(trace string, swagURL string) (uint64, error) {

	metric, err := osvRestCall(trace, swagURL, false)
	if err != nil {
		metric, err = osvRestCall(trace, swagURL, true)
		if err != nil {
			return 0, err
		}
		return metric, err
	}
	return metric, nil
}
