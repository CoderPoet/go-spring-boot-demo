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

package main

import (
	"fmt"
	"time"
	"github.com/didi/go-spring/spring-rpc"
	"github.com/didi/go-spring/spring-rpc-json"
)

func sayHello(ctx SpringRpc.SpringRpcContext) interface{} {

	var data string
	ctx.Bind(&data)

	fmt.Println("request:", data)

	return "Jim"
}

func main() {

	s := &SpringRpcJson.SpringRpcJsonContainer{}
	s.Register("", "sayHello", sayHello)
	go s.Start(":8080")

	time.Sleep(time.Second)

	var data string
	SpringRpcJson.CallService("", "sayHello", "hello", &data)

	fmt.Println("response:", data)
}
