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

package health

import (
	"fmt"
	"net/http"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/models"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
)

var (
	HEALTHY = "Service is running healthy"
)

type IHealthChecker interface {
	GetStatus() *models.Status
	GetxAppStatus(appName, namespace string) *models.Status
}

type HealthChecker struct {
}

func NewHealthChecker() IHealthChecker {
	return &HealthChecker{}
}
func (h *HealthChecker) GetStatus() *models.Status {
	return &models.Status{
		Status: &HEALTHY,
	}
}

func (h *HealthChecker) GetxAppStatus(appName, namespace string) *models.Status {
	resp, err := http.Get(fmt.Sprintf(ricdms.Config.GETxAPPHealthURL, appName, namespace))
	if err != nil {
		ricdms.Logger.Error("Received error while fetching health info: %v", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		ricdms.Logger.Error("xApp is not healthy (http status=%s)", resp.Status)
		return nil
	}

	return &models.Status{
		Status: &HEALTHY,
	}
}
