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
	"errors"
	"fmt"

	"github.com/didi/go-spring/spring-rpc"
	"github.com/didi/go-spring/spring-core"
	"github.com/didi/go-spring/spring-utils"
	"github.com/go-spring/go-spring-boot/spring-boot"
)

func init() {
	SpringBoot.RegisterModule(func(ctx SpringCore.SpringContext) {
		ctx.RegisterBean(new(Controller))
	})
}

type Controller struct {
}

func (controller *Controller) InitRpcBean(c SpringRpc.RpcContainer) {
	c.Register("store", "get", controller.StoreGet)
	c.Register("store", "set", controller.StoreSet)
	c.Register("store", "panic", controller.StorePanic)
}

var store = make(map[string]string)

type GetReq struct {
	Key string `form:"key" json:"key"`
}

func (controller *Controller) StoreGet(ctx SpringRpc.RpcContext) interface{} {

	var param GetReq
	ctx.Bind(&param)
	fmt.Println("/get", "key=", param.Key)

	val := store[param.Key]
	fmt.Println("/get", "val=", val)

	return val
}

type SetReq map[string]string

func (controller *Controller) StoreSet(ctx SpringRpc.RpcContext) interface{} {

	var param SetReq
	ctx.Bind(&param)
	fmt.Println("/set", "param="+SpringUtils.ToJson(param))

	for k, v := range param {
		store[k] = v
	}

	return "ok"
}

func (controller *Controller) StorePanic(ctx SpringRpc.RpcContext) interface{} {

	err := errors.New("this is a panic")
	SpringRpc.ERROR.Panic(err).When(err != nil)

	return "success"
}
