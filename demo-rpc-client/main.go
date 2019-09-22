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
	"errors"
	"github.com/go-spring/app-starter"
	"github.com/didi/go-spring/spring-rpc"
	"github.com/didi/go-spring/spring-core"
	"github.com/didi/go-spring/spring-rpc-client"
	"github.com/go-spring/go-spring-boot/spring-boot"
	_ "github.com/go-spring/go-spring-boot-starter/spring-boot-starter-rpc-json"
	_ "github.com/go-spring/go-spring-boot-starter/spring-boot-starter-rpc-client"
)

func init() {
	SpringBoot.RegisterModule(func(springContext SpringCore.SpringContext) {
		springContext.RegisterBean(new(DemoController))
	})
}

type DemoController struct {
	RpcServiceMap *SpringRpcClient.RpcServiceMap `autowire:""`
}

func (controller *DemoController) InitRpcBean(c SpringRpc.RpcContainer) {
	c.Register("/test", "", controller.Test)
}

func (controller *DemoController) Test(ctx SpringRpc.SpringRpcContext) interface{} {
	if controller.RpcServiceMap != nil {
		if service, ok := controller.RpcServiceMap.GetService("OpenApi"); ok {
			if resp, err := service.Call(); err != nil {
				panic(err)
			} else {
				return resp
			}
		}
	}
	panic(errors.New("service not found"))
}

func main() {
	AppStarter.Run(SpringBoot.NewSpringBootApplication("config/"))
}
