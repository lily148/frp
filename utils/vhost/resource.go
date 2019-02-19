// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatedier/frp/utils/version"
)

const (
	NotFound = `<!DOCTYPE html>
<html>
<head>
<title>Smart Gateway | IBEYOND SMART SDN BHD</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1><a href='https://ibeyondsmart.com'><img src='https://ibeyondsmart.com/wp-content/uploads/2018/10/Logo-Light_01f800340.png' data-at2x='https://ibeyondsmart.com/wp-content/uploads/2018/10/Logo-Light_01f800340@2x.png' alt /></a></h1>
<p>The address might be invalid or Smart Gateway is not connected to internet.<br/>
Please make sure enter the correct Smart Gateway link and it is connected to internet.<br />
Then try again.</p>
<p>Powered by <a href="https://ibeyondsmart.com">IBEYOND SMART SDN BHD</a></p>
<p><em>Smart Gateway from IBEYOND SMART SDN BHD</em></p>
</body>

</html>
`
)

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")
	res := &http.Response{
		Status:     "Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(strings.NewReader(NotFound)),
	}
	return res
}
