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

package ricdms

import (
	"fmt"
	"os"

	mdclog "gerrit.o-ran-sc.org/r/com/golog"
)

type ricdms struct {
}

var Logger *mdclog.MdcLogger

func Init() {
	var err error
	Logger, err = mdclog.InitLogger("ricdms")
	if err != nil {
		fmt.Println("Logger not initialized !!")
		return
	}

	configFile := os.Getenv("RIC_DMS_CONFIG_FILE")

	if configFile != "" {
		Logger.ParseFileContent(configFile)
		Logger.Info("Logger is initialized with config file : %s", configFile)
	} else {
		Logger.LevelSet(mdclog.INFO)
		Logger.Info("Logger is initialized without config file.")
	}
}
