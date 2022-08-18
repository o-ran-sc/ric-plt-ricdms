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
	"os"
	"path"
	"testing"

	"gerrit.o-ran-sc.org/r/com/golog"
	"github.com/stretchr/testify/assert"
)

func TestLoggerWithConfigFile(t *testing.T) {
	p, _ := os.Getwd()
	p = path.Join(p, "../../config/config-test.yaml")
	os.Setenv("RIC_DMS_CONFIG_FILE", p)
	Init()
	assert.Equal(t, Logger.LevelGet(), golog.Level(4))
}

func TestLoggerWithoutConfigFile(t *testing.T) {
	os.Unsetenv("RIC_DMS_CONFIG_FILE")
	Init()
	assert.Equal(t, Logger.LevelGet(), golog.Level(3))
}
