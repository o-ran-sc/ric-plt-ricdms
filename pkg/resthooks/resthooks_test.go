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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	ch "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/charts"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/models"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/onboard"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/charts"
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
		Onboarder:     onboard.NewOnboarder(),
		ChartMgr:      ch.NewChartmgr(),
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

func TestOnboard(t *testing.T) {
	xApp := &models.Descriptor{
		Config: "SAMPLE_CONFIGURATION",
		Schema: "SAMPLE_SCHEMA",
	}

	svr := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var d map[string]interface{}
		reqBytes, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		err := json.Unmarshal(reqBytes, &d)

		ricdms.Logger.Debug("after unmarshal : %+v body: %+v", d, string(reqBytes))

		if err != nil {
			assert.Fail(t, "Not able to parse the request body")
		}

		assert.Equal(t, xApp.Config, d["config-file.json"])
		assert.Equal(t, xApp.Schema, d["controls-schema.json"])
		fmt.Fprintf(w, "SAMPLE_RESPONSE")
	}))
	svr.Listener.Close()
	svr.Listener, _ = net.Listen("tcp", ricdms.Config.MockServer)

	svr.Start()
	defer svr.Close()

	resp := rh.OnBoard(xApp)
	assert.NotEqual(t, nil, resp)
}

func TestGetCharts(t *testing.T) {

	svr := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ricdms.Logger.Debug("Mock server running")
		fmt.Fprintf(w, "SAMPLE_RESPONSE")
	}))
	svr.Listener.Close()
	svr.Listener, _ = net.Listen("tcp", ricdms.Config.MockServer)

	svr.Start()
	defer svr.Close()

	resp := rh.GetCharts()
	assert.NotEqual(t, nil, resp)

	successResp := resp.(*charts.GetChartsListOK)
	assert.Equal(t, "SAMPLE_RESPONSE", successResp.Payload)
}

type HealthCheckerMock struct {
	mock.Mock
}

func (h HealthCheckerMock) GetStatus() *models.Status {
	return successStatus
}
