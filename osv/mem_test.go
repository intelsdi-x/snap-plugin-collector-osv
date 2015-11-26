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

func TestMemPlugin(t *testing.T) {
	Convey("getMemstat Should return memory amount value", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/free",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "20000")
				return resp, nil

			},
		)
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/total",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "10000")
				return resp, nil

			},
		)

		memFree, err := getMemStat("http://192.168.192.200:8000", "free")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(memFree, 10), ShouldResemble, "20000")
		memTotal, err := getMemStat("http://192.168.192.200:8000", "total")
		So(err, ShouldBeNil)
		So(strconv.FormatUint(memTotal, 10), ShouldResemble, "10000")

	})
	Convey("MemStat Should return pluginMetricType Data", t, func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/free",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "20000")
				return resp, nil

			},
		)
		httpmock.RegisterResponder("GET", "http://192.168.192.200:8000/os/memory/total",
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, "10000")
				return resp, nil

			},
		)

		ns := []string{"osv", "memory", "free"}
		ns2 := []string{"osv", "memory", "total"}
		memFree, err := memStat(ns, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(memFree.Namespace_, ShouldResemble, ns)
		So(memFree.Data_, ShouldResemble, "20000")
		memTotal, err := memStat(ns2, "http://192.168.192.200:8000")
		So(err, ShouldBeNil)
		So(memTotal.Namespace_, ShouldResemble, ns2)
		So(memTotal.Data_, ShouldResemble, "10000")

	})
}
