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
	"strconv"
	"testing"

	"github.com/intelsdi-x/snap-plugin-collector-osv/osv/httpmock"

	"github.com/intelsdi-x/snap/core"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCpuPlugin(t *testing.T) {
	httpmock.Mock = true

	Convey("getcpuTime Should return cputime value", t, func() {

		defer httpmock.ResetResponders()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count", `{"time_ms": 144123232, "list": []}`, 200)

		cpuTime, err := getCPUTime("http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(cpuTime, 10), ShouldResemble, "144123232")

	})
	Convey("CpuStat Should return pluginMetricType Data", t, func() {

		defer httpmock.ResetResponders()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count", `{"time_ms": 144123232, "list": []}`, 200)

		ns := core.NewNamespace("intel", "osv", "cpu", "cputime")
		cpuTime, err := cpuStat(ns, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(cpuTime.Namespace(), ShouldResemble, ns)
		So(cpuTime.Data_, ShouldResemble, "144123232")

	})
}
