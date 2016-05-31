//
// +build unit

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-collector-osv/osv/httpmock"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"
	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPlugin(t *testing.T) {
	httpmock.Mock = true

	Convey("Meta should return metadata for the plugin", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, Name)
		So(meta.Version, ShouldResemble, Version)
		So(meta.Type, ShouldResemble, plugin.CollectorPluginType)
	})

	Convey("Create Osv Collector", t, func() {
		osvCol := NewOsvCollector()
		Convey("So psCol should not be nil", func() {
			So(osvCol, ShouldNotBeNil)
		})
		Convey("So psCol should be of Osv type", func() {
			So(osvCol, ShouldHaveSameTypeAs, &Osv{})
		})
		Convey("osvCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := osvCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
	})
	Convey("Get URI ", t, func() {
		Convey("So should return 10.1.0.1:8000", func() {
			swagIP := "10.1.0.1"
			swagPort := 8000
			uri := osvRestURL(swagIP, swagPort)
			So("http://10.1.0.1:8000", ShouldResemble, uri)
		})
	})
	Convey("Get Metrics ", t, func() {
		osvCol := NewOsvCollector()
		cfgNode := cdata.NewNode()
		var cfg = plugin.ConfigType{
			ConfigDataNode: cfgNode,
		}
		Convey("So should return 187 types of metrics", func() {
			metrics, err := osvCol.GetMetricTypes(cfg)
			So(187, ShouldResemble, len(metrics))
			So(err, ShouldBeNil)
		})
		Convey("So should check namespace", func() {
			metrics, err := osvCol.GetMetricTypes(cfg)
			waitNamespace := metrics[0].Namespace().String()
			wait := regexp.MustCompile(`^/intel/osv/trace/virtio/virtio_wait_for_queue`)
			So(true, ShouldEqual, wait.MatchString(waitNamespace))
			So(err, ShouldBeNil)

		})

	})
	Convey("Get Metrics ", t, func() {
		osvCol := NewOsvCollector()
		cfgNode := cdata.NewNode()
		cfgNode.AddItem("swagIP", ctypes.ConfigValueStr{Value: "192.168.192.200"})
		cfgNode.AddItem("swagPort", ctypes.ConfigValueInt{Value: 8000})

		defer httpmock.ResetResponders()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/free", "20000", 200)
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count",
			`{"time_ms": 144123232, "list": [{"name": "waitqueue_wake_one", "count": 1000}]}`, 200)

		Convey("So should get memory metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "osv", "memory", "free"),
				Config_:    cfgNode,
			}}
			collect, err := osvCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(len(collect), ShouldResemble, 1)

		})
		Convey("So should get cpu metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "osv", "cpu", "cputime"),
				Config_:    cfgNode,
			}}
			collect, err := osvCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(collect[0].Data_, ShouldResemble, "144123232")
			So(len(collect), ShouldResemble, 1)

		})
		Convey("So should get trace metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "osv", "trace", "wait", "waitqueue_wake_one"),
				Config_:    cfgNode,
			}}
			collect, err := osvCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(collect[0].Data_, ShouldResemble, "1000")
			So(len(collect), ShouldResemble, 1)

		})

	})
}
