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
	"os"
	"testing"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/models"
	h "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rh *Resthook
var successStatus *models.Status

func TestMain(m *testing.M) {

	successStatus = &models.Status{
		Status: &health.HEALTHY,
	}
	ricdms.Init()
	rh = &Resthook{
		HealthChecker: HealthCheckerMock{},
	}
	code := m.Run()
	os.Exit(code)
}

func TestHealth(t *testing.T) {
	resp := rh.GetDMSHealth()
	switch resp.(type) {
	case *h.GetHealthCheckOK:
		assert.Equal(t, successStatus, resp.(*h.GetHealthCheckOK).Payload)

	case *h.GetHealthCheckInternalServerError:
		assert.Fail(t, "Internal Server generated: %v", resp)

	default:
		assert.Fail(t, "Unknown type of resp : %v", resp)
	}
}

type HealthCheckerMock struct {
	mock.Mock
}

func (h HealthCheckerMock) GetStatus() *models.Status {
	return successStatus
}
