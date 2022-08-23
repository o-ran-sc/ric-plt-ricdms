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

package resthooks

import (
	ph "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
	"github.com/go-openapi/runtime/middleware"
)

func NewResthook(h ph.IHealthChecker) *Resthook {
	return &Resthook{
		HealthChecker: h,
	}
}

func (rh *Resthook) GetDMSHealth() (resp middleware.Responder) {
	ricdms.Logger.Debug("healthchecker : %v\n", rh.HealthChecker)
	return health.NewGetHealthCheckOK().WithPayload(rh.HealthChecker.GetStatus())
}
