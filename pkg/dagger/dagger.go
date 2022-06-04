// Copyright 2019 Dagger Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dagger

import (
	"cuelang.org/go/cue"

	"os/exec"
	"os"
	"bytes"
)

type Result struct {
	Stdout string
	Success bool `json:success`
	Error string
}

func Shell(config *cue.Struct) (Result, error) {
       script, err := config.FieldByName("script", false)
       if err != nil {
	       return Result{}, err
       }
       scriptStr, err := script.Value.String()
       if err != nil {
	       return Result{}, err
       }
	cmd := exec.Command("sh", "-c", scriptStr)
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr
	var result Result
	if err := cmd.Run(); err != nil {
		result.Success = false
		result.Error = err.Error()
	} else {
		result.Success = true
	}
	result.Stdout = stdout.String()
	return result, err
}
