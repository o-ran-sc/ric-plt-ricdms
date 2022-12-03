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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	ch "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/charts"
	dm "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/deploy"
	ph "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/health"
	po "gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/onboard"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/charts"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/deploy"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/health"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/restapi/operations/onboard"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/resthooks"
	"gerrit.o-ran-sc.org/r/ric-plt/ricdms/pkg/ricdms"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func NewRestful() *Restful {
	r := &Restful{
		rh: resthooks.NewResthook(
			ph.NewHealthChecker(),
			po.NewOnboarder(),
			ch.NewChartmgr(),
			dm.NewDeploymentManager(),
		),
	}
	r.setupHandler()
	return r
}

func (r *Restful) setupHandler() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		os.Exit(1)
	}

	api := operations.NewRICDMSAPI(swaggerSpec)

	api.HealthGetHealthCheckHandler = health.GetHealthCheckHandlerFunc(func(ghcp health.GetHealthCheckParams) middleware.Responder {
		ricdms.Logger.Debug("==> HealthCheck API invoked.")
		resp := r.rh.GetDMSHealth()
		return resp
	})

	api.HealthGetHealthcheckXAppXAppNameNamespaceNamespaceHandler = health.GetHealthcheckXAppXAppNameNamespaceNamespaceHandlerFunc(func(param health.GetHealthcheckXAppXAppNameNamespaceNamespaceParams) middleware.Responder {
		ricdms.Logger.Debug("==> Healthcheck for xApp is invoked")
		resp := r.rh.GetxAppHealth(param.XAppName, param.Namespace)
		return resp
	})

	api.OnboardPostOnboardxAppsHandler = onboard.PostOnboardxAppsHandlerFunc(func(poap onboard.PostOnboardxAppsParams) middleware.Responder {
		ricdms.Logger.Debug("==> onboard API invoked.")
		resp := r.rh.OnBoard(poap.Body)
		return resp
	})

	api.ChartsGetChartsListHandler = charts.GetChartsListHandlerFunc(func(param charts.GetChartsListParams) middleware.Responder {
		ricdms.Logger.Debug("==> GetChartList")
		resp := r.rh.GetCharts()
		return resp
	})

	api.ChartsDownloadHelmChartHandler = charts.DownloadHelmChartHandlerFunc(func(param charts.DownloadHelmChartParams) middleware.Responder {
		ricdms.Logger.Debug("==> Download helm chart")
		resp := r.rh.DownloadChart(param.XAppName, param.Version)
		return resp
	})

	api.ChartsGetChartHandler = charts.GetChartHandlerFunc(func(param charts.GetChartParams) middleware.Responder {
		ricdms.Logger.Debug("==> Get Charts by name is invoked")
		resp := r.rh.GetChartsByName(param.XAppName)
		return resp
	})

	api.ChartsGetChartsFetcherHandler = charts.GetChartsFetcherHandlerFunc(func(param charts.GetChartsFetcherParams) middleware.Responder {
		ricdms.Logger.Debug("==> Get Charts by name and version is invoked")
		resp := r.rh.GetChartByNameAndVersion(param.XAppName, param.Version)
		return resp
	})

	api.DeployPostDeployHandler = deploy.PostDeployHandlerFunc(func(param deploy.PostDeployParams) middleware.Responder {
		ricdms.Logger.Debug("==> deployment of xApp")
		resp := r.rh.DownloadAndInstallChart(param.Body.XAppname, param.Body.Version, *param.Body.Namespace)
		return resp
	})

	api.DeployDeleteDeployHandler = deploy.DeleteDeployHandlerFunc(func(param deploy.DeleteDeployParams) middleware.Responder {
		ricdms.Logger.Debug("==> undeploy xApp")
		resp := r.rh.UninstallChart(*param.Body.XAppname, *param.Body.Version, param.Body.Namespace)
		return resp
	})
	api.ApplicationZipProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		if zp, ok := data.(io.ReadCloser); ok {
			defer zp.Close()
			b, err := ioutil.ReadAll(zp)

			if err != nil {
				ricdms.Logger.Error("error: %v", err)
				return err
			}
			_, err = w.Write(b)

			if err != nil {
				ricdms.Logger.Error("error: %v", err)
				return err
			}
			return nil
		}
		return fmt.Errorf("not support")
	})
	r.api = api
}

func (r *Restful) Run() {
	server := restapi.NewServer(r.api)
	defer server.Shutdown()
	server.Port = 8000
	server.Host = "0.0.0.0"
	ricdms.Logger.Info("Starting server at : %s:%d", server.Host, server.Port)
	if err := server.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
