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

	"github.com/jarcoal/httpmock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTracePlugin(t *testing.T) {
	Convey("getMemstat Should return memory amount value", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, `{"time_ms": 144123232, "list": [{"name": "waitqueue_wake_one", "count": 1000}]}`)
				return resp, nil

			},
		)
		httpmock.RegisterResponder("POST", "http://192.168.192.200:8000/trace/count/waitqueue_wake_one?enabled=True",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "")
				return resp, nil

			},
		)

		trace, err := getTrace("waitqueue_wake_one", "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(trace, 10), ShouldResemble, "1000")

	})
	Convey("MemStat Should return pluginMetricType Data", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/trace/count",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, `{"time_ms": 144123232, "list": [{"name": "waitqueue_wake_one", "count": 1000}]}`)
				return resp, nil

			},
		)

		ns := []string{"osv", "trace", "wait", "waitqueue_wake_one"}
		memFree, err := traceStat(ns, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(memFree.Namespace_, ShouldResemble, ns)
		So(memFree.Data_, ShouldResemble, "1000")

	})
	Convey("osvCallRest should return nil", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("POST", "http://192.168.192.200:8000/trace/count/waitqueue_wake_one?enabled=True",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "")
				return resp, nil

			},
		)

		resp := osvRestPost("http://192.168.192.200:8000", "trace/count/waitqueue_wake_one?enabled=True")
		So(resp, ShouldBeNil)

	})
}
