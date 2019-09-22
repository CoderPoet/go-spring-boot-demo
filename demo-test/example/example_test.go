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
	"os"
	"fmt"
	"time"
	"testing"
	"net/http"
	"github.com/go-spring/go-spring-boot/spring-boot"
)

func TestMain(m *testing.M) {

	app := SpringBoot.NewSpringBootApplication("../config/")
	app.Start()

	time.Sleep(200 * time.Millisecond)
	code := m.Run()

	app.ShutDown()
	os.Exit(code)
}

func TestController_GetA(t *testing.T) {

	resp, err := http.Get("http://127.0.0.1:8080/getA")
	fmt.Println(err)

	b := make([]byte, resp.ContentLength)
	resp.Body.Read(b)

	fmt.Println(string(b))
}

func TestController_GetB(t *testing.T) {

	resp, err := http.Get("http://127.0.0.1:8080/getB")
	fmt.Println(err)

	b := make([]byte, resp.ContentLength)
	resp.Body.Read(b)

	if string(b) != "bbb" {
		t.Error("opus!")
	}
}
