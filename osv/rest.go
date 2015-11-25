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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func osvRestGet(swag_url string, path string) ([]byte, error) {

	call_url := fmt.Sprintf("%s/%s", swag_url, path)
	resp, err := http.Get(call_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func osvRestPost(swag_url string, path string) error {

	call_url := fmt.Sprintf("%s/%s", swag_url, path)
	fmt.Println(call_url)
	_, err := http.PostForm(call_url, url.Values{})
	if err != nil {
		return err
	}
	return nil
}

func osvRestUrl(ip string, port int) string {
	url := fmt.Sprintf("http://%s:%d", ip, port)
	return url
}
