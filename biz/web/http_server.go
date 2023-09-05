// Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
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

package web

import (
	"agent/biz/docker"
	"agent/biz/web/routers"
	"agent/config"

	"github.com/gin-gonic/gin"

	// docs are generated by Swag CLI, you have to import it.
	_ "agent/docs"
)

// Start
// @contact.name wenchao
// @contact.url http://www.swagger.io/support
// @contact.email wenchao@iscas.ac.cn
func Start() {

	gin.SetMode(gin.ReleaseMode)
	if config.Config.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	externalWebServer := routers.ExternalWebServer{Router: routers.ExternalRouter()}
	internalWebServer := routers.InternalWebServer{Router: routers.InternalRouter()}

	go externalWebServer.Start()

	docker.SubscribeAsyncDockerNetwork(func(status int) {
		// start internal web server
		go internalWebServer.Start()
	})
}