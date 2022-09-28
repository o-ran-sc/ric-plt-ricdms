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
package onboard

import (
	"bytes"
	"encoding/json"
	"net/http"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/models"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/onboard"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
	"github.com/go-openapi/runtime/middleware"
)

type IOnboarder interface {
	Onboard(descriptor *models.Descriptor) middleware.Responder
}

type Onboarder struct {
}

func NewOnboarder() IOnboarder {
	return &Onboarder{}
}

func (o *Onboarder) Onboard(descriptor *models.Descriptor) middleware.Responder {
	ricdms.Logger.Debug("onboarder invoked : %+v", descriptor)

	// validate if the required patameter is available
	if descriptor.Schema == nil || descriptor.Config == nil {
		return onboard.NewPostOnboardxAppsBadRequest()
	}

	body := map[string]interface{}{
		"config-file.json":     descriptor.Config,
		"controls-schema.json": descriptor.Schema,
	}

	bodyBytes, _ := json.Marshal(body)

	// resp, err := http.Post("http://172.17.0.1:8888/api/v1/onboard", "application/json", bytes.NewReader(bodyBytes))
	ricdms.Logger.Info("config : %+v", ricdms.Config)
	resp, err := http.Post(ricdms.Config.OnboardURL, "application/json", bytes.NewReader(bodyBytes))

	if err == nil {
		ricdms.Logger.Info("no error response: %+v", resp)
	} else {
		ricdms.Logger.Error("err : (%v)", err)
	}
	return onboard.NewPostOnboardxAppsCreated()
}
