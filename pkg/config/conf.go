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

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	LogLevel                     string `yaml:"log-level"`
	OnboardURL                   string `yaml:"onborder-url"`
	GetChartsURL                 string `yaml:"getCharts-url"`
	GetChartsByxAppNameURL       string `yaml:"getCharts-by-name-url"`
	GetChartsByNameAndVersionURL string `yaml:"getCharts-by-name-and-version-url"`
	DownloadChartURLFormat       string `yaml:"download-charts-url-format"`
	MockServer                   string `yaml:"mock-server"`
	GETxAPPHealthURL             string `yaml:"getXappHealth-url"`
}

func ReadConfiguration(configFile string) (c *Conf, err error) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("error in resolving config file : %+v\n", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("Unmarshal error : %+v\n", err)
		return nil, err
	}

	return c, err
}
