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

package charts

import (
	"io/ioutil"
	"net/http"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
)

type ChartMgr struct {
}

type IChartMgr interface {
	GetCharts() (string, error)
}

func NewChartmgr() IChartMgr {
	return &ChartMgr{}
}

func (c *ChartMgr) GetCharts() (string, error) {
	ricdms.Logger.Debug("GetCharts invoked")

	resp, err := http.Get(ricdms.Config.GetChartsURL)
	if err != nil {
		ricdms.Logger.Debug("Error in getting charts : %+v", err)
		return "", err
	}

	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ricdms.Logger.Debug("error in response: %+v", respByte)
		return "", err
	}

	ricdms.Logger.Debug("response : %+v", string(respByte))
	return string(respByte), nil
}
