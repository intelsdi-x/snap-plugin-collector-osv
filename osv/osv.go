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
	"strings"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// Name of plugin
	Name = "osv"
	// Version of plugin
	Version = 3
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

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

// CollectMetrics returns collected metrics
func (p *Osv) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	cpure := regexp.MustCompile(`^/osv/cpu/cputime`)
	memre := regexp.MustCompile(`^/osv/memory/.*`)
	tracere := regexp.MustCompile(`^/osv/trace/.*`)
	metrics := make([]plugin.PluginMetricType, len(mts))

	swagIP := mts[0].Config().Table()["swagIP"].(ctypes.ConfigValueStr).Value
	swagPort := mts[0].Config().Table()["swagPort"].(ctypes.ConfigValueInt).Value
	swagURL := osvRestURL(swagIP, swagPort)

	for i, p := range mts {

		ns := joinNamespace(p.Namespace())
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
		metrics[i].Source_ = swagIP

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

	cp.Add([]string{""}, config)
	return cp, nil

}

// GetMetricTypes returns metric types that can be collected
func (p *Osv) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	var metrics []plugin.PluginMetricType
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
