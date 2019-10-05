/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package example

import (
	"fmt"

	"github.com/didi/go-spring/spring-utils"
	"github.com/didi/go-spring/spring-rpc-http"
)

func RunTest() {

	var respGet string
	reqGet := &GetReq{Key: "a"}
	err := SpringHttpRpc.CallService("store", "get", reqGet, &respGet)
	fmt.Println("err:", SpringUtils.String(err), "||", "resp:", respGet)

	var respSet string
	reqSet := &SetReq{"a": "1"}
	err = SpringHttpRpc.CallService("store", "set", reqSet, &respSet)
	fmt.Println("err:", SpringUtils.String(err), "||", "resp:", respSet)

	err = SpringHttpRpc.CallService("store", "get", reqGet, &respGet)
	fmt.Println("err:", SpringUtils.String(err), "||", "resp:", respGet)

	var respPanic string
	err = SpringHttpRpc.CallService("store", "panic", nil, &respPanic)
	fmt.Println("err:", SpringUtils.String(err), "||", "resp:", respPanic)
}
