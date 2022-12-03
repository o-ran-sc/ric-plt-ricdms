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

package deploy

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

type DeploymentManager struct {
	settings *cli.EnvSettings
}

const (
	HELM_DRIVER        = "HELM_DRIVER"
	CHART_NAME_FORMAT  = "chart-%s-%s.tgz"
	RELESE_NAME_FORMAT = "ricdms-%s-rel-%s"
)

func NewDeploymentManager() IDeploy {
	return &DeploymentManager{
		settings: cli.New(),
	}
}

func (d *DeploymentManager) install(chartPath, appName, version, namesapce string) error {
	conf := action.Configuration{}
	err := conf.Init(d.settings.RESTClientGetter(), namesapce, os.Getenv(HELM_DRIVER), ricdms.Logger.Debug)

	if err != nil {
		ricdms.Logger.Error("not able to prepare install configuration: %v", err)
		return err
	}

	install := action.NewInstall(&conf)
	install.ReleaseName = fmt.Sprintf(RELESE_NAME_FORMAT, appName, strings.ReplaceAll(version, ".", ""))
	install.Namespace = namesapce

	cp, err := install.ChartPathOptions.LocateChart(chartPath, d.settings)
	if err != nil {
		ricdms.Logger.Error("Not able to locate charts on: %s", chartPath)
	}

	chart, err := loader.Load(cp)
	if err != nil {
		ricdms.Logger.Error("Not able to load charts : %v", err)
		return err
	}

	release, err := install.Run(chart, map[string]interface{}{})
	if err != nil {
		ricdms.Logger.Error("Not able to install the xApp : %v", err)
		return err
	}

	ricdms.Logger.Info("chart is installed with following details : %v", release)
	return nil
}

func (d *DeploymentManager) writeToFile(reader io.ReadCloser, appname, version string) error {
	if reader != nil {
		outfile, err := os.Create(fmt.Sprintf(CHART_NAME_FORMAT, appname, version))

		if err != nil {
			ricdms.Logger.Error("outfile can't be created")
			return err
		}

		defer outfile.Close()

		_, err = io.Copy(outfile, reader)
		if err != nil {
			ricdms.Logger.Error("Error while creating chart tar: %v", err)
		}
		return nil
	}
	return fmt.Errorf("reader is nil")
}

func (d *DeploymentManager) Deploy(reader io.ReadCloser, appname, version, namespace string) error {
	err := d.writeToFile(reader, appname, version)
	if err != nil {
		return err
	}

	err = d.install(fmt.Sprintf(CHART_NAME_FORMAT, appname, version), appname, version, namespace)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeploymentManager) Uninstall(appname, version, namespace string) error {
	conf := action.Configuration{}

	err := conf.Init(d.settings.RESTClientGetter(), namespace, os.Getenv(HELM_DRIVER), ricdms.Logger.Debug)
	if err != nil {
		ricdms.Logger.Error("Not able to prepare uninstall configuration: %v", err)
		return err
	}

	uninstall := action.NewUninstall(&conf)
	uninstall.Wait = true

	resp, err := uninstall.Run(fmt.Sprintf(RELESE_NAME_FORMAT, appname, strings.ReplaceAll(version, ".", "")))
	if err != nil {
		ricdms.Logger.Error("Error while uninstalling deployment: %v", err)
		return err
	}

	ricdms.Logger.Info("deployment uninstallation comlete : %", resp.Info)
	return nil
}
