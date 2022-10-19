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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
)

type ChartMgr struct {
}

type IChartMgr interface {
	GetCharts() (map[string]interface{}, error)
	DownloadChart(string, string) (io.ReadCloser, error)
	GetChartsByName(name string) ([]map[string]interface{}, error)
	GetChartsByNameAndVersion(name, version string) (map[string]interface{}, error)
}

func NewChartmgr() IChartMgr {
	return &ChartMgr{}
}

func (c *ChartMgr) GetCharts() (map[string]interface{}, error) {
	ricdms.Logger.Debug("GetCharts invoked")

	resp, err := http.Get(ricdms.Config.GetChartsURL)
	if err != nil {
		ricdms.Logger.Debug("Error in getting charts : %+v", err)
		return make(map[string]interface{}, 0), err
	}

	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ricdms.Logger.Debug("error in response: %+v", respByte)
		return make(map[string]interface{}, 0), err
	}

	ricdms.Logger.Debug("response : %+v", string(respByte))

	v := make(map[string]interface{}, 0)
	json.Unmarshal(respByte, &v)
	return v, nil
}

func (c *ChartMgr) DownloadChart(chartName string, version string) (io.ReadCloser, error) {
	ricdms.Logger.Debug("Download Charts invoked")

	if chartName == "" || version == "" {
		return nil, fmt.Errorf("chartname or version is empty")
	}

	ChartURL := fmt.Sprintf(ricdms.Config.DownloadChartURLFormat, chartName, version)

	resp, err := http.Get(ChartURL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *ChartMgr) GetChartsByName(name string) ([]map[string]interface{}, error) {
	ricdms.Logger.Debug("Get Chart by xApp name is invoked")

	if name == "" {
		return make([]map[string]interface{}, 0), fmt.Errorf("xAppname is empty")
	}

	URL := fmt.Sprintf(ricdms.Config.GetChartsByxAppNameURL, name)

	response, err := http.Get(URL)
	if err != nil {
		ricdms.Logger.Error("error: %v", err)
		return make([]map[string]interface{}, 0), err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ricdms.Logger.Debug("Reading response failed with err : %v", err)
		return make([]map[string]interface{}, 0), err
	}

	v := make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &v)
	if err != nil {
		ricdms.Logger.Debug("Error while parsing res: %v", err)
		return make([]map[string]interface{}, 0), err
	}
	return v, nil
}

func (c *ChartMgr) GetChartsByNameAndVersion(name, version string) (map[string]interface{}, error) {
	ricdms.Logger.Debug("Get Charts by name and version is invoked")

	if name == "" || version == "" {
		return make(map[string]interface{}, 0), fmt.Errorf("name or version is not provided")
	}

	URL := fmt.Sprintf(ricdms.Config.GetChartsByNameAndVersionURL, name, version)

	response, err := http.Get(URL)
	if err != nil {
		ricdms.Logger.Debug("error in retrieving chart: %v", err)
		return make(map[string]interface{}, 0), err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ricdms.Logger.Debug("error in reading response: %v", err)
		return make(map[string]interface{}, 0), err
	}

	v := make(map[string]interface{}, 0)
	err = json.Unmarshal(data, &v)
	if err != nil {
		ricdms.Logger.Debug("error in parsing response: %v", err)
		return make(map[string]interface{}, 0), err
	}

	return v, nil
}
