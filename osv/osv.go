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
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

type Osv struct {
}

func NewOsvCollector() *Osv {
	return &Osv{}

}

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func (p *Osv) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	cpure := regexp.MustCompile(`^/osv/cpu/cputime`)
	memre := regexp.MustCompile(`^/osv/memory/.*`)
	tracere := regexp.MustCompile(`^/osv/trace/.*`)
	metrics := make([]plugin.PluginMetricType, len(mts))

	swag_ip := mts[0].Config().Table()["swag_ip"].(ctypes.ConfigValueStr).Value
	swag_port := mts[0].Config().Table()["swag_port"].(ctypes.ConfigValueInt).Value
	swag_url := osvRestUrl(swag_ip, swag_port)

	for i, p := range mts {

		ns := joinNamespace(p.Namespace())
		switch {
		case memre.MatchString(ns):
			metric, err := memStat(p.Namespace(), swag_url)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		case cpure.MatchString(ns):
			metric, err := cpuStat(p.Namespace(), swag_url)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case tracere.MatchString(ns):
			metric, err := traceStat(p.Namespace(), swag_url)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		}
		metrics[i].Source_ = swag_ip

	}
	return metrics, nil
}

func (p *Osv) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	swag_ip, err := cpolicy.NewStringRule("swag_ip", true)
	handleErr(err)
	swag_ip.Description = "Osv ip address"
	config.Add(swag_ip)
	swag_port, err := cpolicy.NewIntegerRule("swag_port", false, 8000)
	handleErr(err)
	swag_port.Description = "Swagger port / default 8000"
	config.Add(swag_port)

	cp.Add([]string{""}, config)
	return cp, nil

}

func (p *Osv) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	metrics := make([]plugin.PluginMetricType, 0)
	counter_mts, err := getCounterMetricTypes()
	if err != nil {
		handleErr(err)
	}
	memory_mts, err := getMemoryMetricTypes()
	if err != nil {
		handleErr(err)
	}
	cpu_mts, err := getCpuMetricTypes()
	if err != nil {
		handleErr(err)
	}
	metrics = append(metrics, counter_mts...)
	metrics = append(metrics, cpu_mts...)
	metrics = append(metrics, memory_mts...)
	return metrics, nil
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
