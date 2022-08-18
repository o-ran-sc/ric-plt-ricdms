//==================================================================================
//  Copyright (c) 2022 Samsung
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//
//   This source code is part of the near-RT RIC (RAN Intelligent Controller)
//   platform project (RICP).
//==================================================================================
//
package restful

import (
	"log"
	"os"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations"
	"github.com/go-openapi/loads"
)

func NewRestful() *Restful {
	r := &Restful{}
	r.setupHandler()
	return r
}

func (r *Restful) setupHandler() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		os.Exit(1)
	}

	api := operations.NewRICDMSAPI(swaggerSpec)
	r.api = api
}

func (r *Restful) Run() {
	server := restapi.NewServer(r.api)
	defer server.Shutdown()
	server.Port = 8000
	server.Host = "0.0.0.0"
	if err := server.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
