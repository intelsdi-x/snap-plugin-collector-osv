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
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
)

type Counter struct {
	Name  string
	Count uint64
}
type Counters struct {
	Time_ms uint64
	List    []Counter
}

func getCounterMetricTypes() ([]plugin.PluginMetricType, error) {
	mts := make([]plugin.PluginMetricType, 0)
	for _, counter := range virtio_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "virtio", counter}})
	}
	for _, counter := range net_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "net", counter}})
	}
	for _, counter := range memory_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "memory", counter}})
	}
	for _, counter := range callout_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "callout", counter}})
	}
	for _, counter := range wait_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "wait", counter}})
	}
	for _, counter := range async_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "async", counter}})
	}
	for _, counter := range vfs_counters {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"osv", "trace", "vfs", counter}})
	}
	return mts, nil
}

func traceStat(ns []string, swag_url string) (*plugin.PluginMetricType, error) {
	trace := ns[3]
	metric, err := getTrace(trace, swag_url)
	if err != nil {
		return nil, err
	}
	return &plugin.PluginMetricType{
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
	fmt.Println(trace, counters.List)
	trc_error := fmt.Sprintf("Can't find %s in trace list", trace)
	return 0, errors.New(trc_error)

}

func osvRestCall(trace string, swag_url string, recovery bool) (uint64, error) {
	path := "trace/count"
	if recovery {
		recovery_path := fmt.Sprintf("%s/%s?enabled=true", path, trace)
		err := osvRestPost(swag_url, recovery_path)
		if err != nil {
			return 0, err
		}
	}
	response, err := osvRestGet(swag_url, path)
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

func getTrace(trace string, swag_url string) (uint64, error) {

	metric, err := osvRestCall(trace, swag_url, false)
	if err != nil {
		metric, err = osvRestCall(trace, swag_url, true)
		if err != nil {
			return 0, err
		}
		return metric, err
	}
	return metric, nil
}
