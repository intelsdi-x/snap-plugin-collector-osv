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
	"net/http"
	"strconv"
	"testing"

	"github.com/intelsdi-x/snap/core"

	"github.com/jarcoal/httpmock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCpuPlugin(t *testing.T) {
	Convey("getcpuTime Should return cputime value", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, `{"time_ms": 144123232, "list": []}`)
				return resp, nil

			},
		)

		cpuTime, err := getCPUTime("http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(cpuTime, 10), ShouldResemble, "144123232")

	})
	Convey("CpuStat Should return pluginMetricType Data", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, `{"time_ms": 144123232, "list": []}`)
				return resp, nil

			},
		)

		ns := core.NewNamespace("intel", "osv", "cpu", "cputime")
		cpuTime, err := cpuStat(ns, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(cpuTime.Namespace(), ShouldResemble, ns)
		So(cpuTime.Data_, ShouldResemble, "144123232")

	})
}
