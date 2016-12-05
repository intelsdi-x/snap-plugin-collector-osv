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
	"regexp"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// Parts of returned namespace
	Vendor = "intel"

	// Name of plugin
	Name = "osv"

	// Version of plugin
	Version = 4

	// Type of plugin
	Type = plugin.CollectorPluginType
)

// Meta returns plugin meta data info
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

// Osv struct
type Osv struct {
}

// NewOsvCollector returns new Collector instance
func NewOsvCollector() *Osv {
	return &Osv{}

}

// CollectMetrics returns collected metrics
func (p *Osv) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	cpure := regexp.MustCompile(`^/` + Vendor + `/` + Name + `/cpu/cputime`)
	memre := regexp.MustCompile(`^/` + Vendor + `/` + Name + `/memory/.*`)
	tracere := regexp.MustCompile(`^/` + Vendor + `/` + Name + `/trace/.*`)
	metrics := make([]plugin.MetricType, len(mts))

	swagIP := mts[0].Config().Table()["swagIP"].(ctypes.ConfigValueStr).Value
	swagPort := mts[0].Config().Table()["swagPort"].(ctypes.ConfigValueInt).Value
	swagURL := osvRestURL(swagIP, swagPort)

	for i, p := range mts {

		ns := p.Namespace().String()
		switch {
		case memre.MatchString(ns):
			metric, err := memStat(p.Namespace(), swagURL)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		case cpure.MatchString(ns):
			metric, err := cpuStat(p.Namespace(), swagURL)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case tracere.MatchString(ns):
			metric, err := traceStat(p.Namespace(), swagURL)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		}

		metrics[i].Tags_ = map[string]string{core.STD_TAG_PLUGIN_RUNNING_ON: swagIP}
	}
	return metrics, nil
}

// GetConfigPolicy returns a config policy
func (p *Osv) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	swagIP, err := cpolicy.NewStringRule("swagIP", true)
	handleErr(err)
	swagIP.Description = "Osv ip address"
	config.Add(swagIP)
	swagPort, err := cpolicy.NewIntegerRule("swagPort", false, 8000)
	handleErr(err)
	swagPort.Description = "Swagger port / default 8000"
	config.Add(swagPort)

	cp.Add([]string{Vendor, Name}, config)
	return cp, nil

}

// GetMetricTypes returns metric types that can be collected
func (p *Osv) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	var metrics []plugin.MetricType
	counterMts, err := getCounterMetricTypes()
	if err != nil {
		handleErr(err)
	}
	memoryMts, err := getMemoryMetricTypes()
	if err != nil {
		handleErr(err)
	}
	cpuMts, err := getCPUMetricTypes()
	if err != nil {
		handleErr(err)
	}
	metrics = append(metrics, counterMts...)
	metrics = append(metrics, cpuMts...)
	metrics = append(metrics, memoryMts...)
	return metrics, nil
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
