// Copyright 2021 Tetrate
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

//+build integration

package e2e

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Positive test for whole workflow.
// Make sure all the ENV are exported.
func TestGCPProvider(t *testing.T) {
	casCAName := os.Getenv("GOOGLE_CAS_CA_NAME")

	cmd := exec.Command("./getmesh", "gen-ca", "--provider=gcp",
		fmt.Sprintf("--cas-ca-name=%s", casCAName),
		"--validity-days=100",
		"--secret-file-path=/tmp/gcp-secret.yaml",
	)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	cmd.Run()

	actual := buf.String()
	assert.Contains(t, actual, "Kubernetes Secret YAML created successfully in /tmp/gcp-secret.yaml")

	_, err := ioutil.ReadFile("/tmp/gcp-secret.yaml")
	assert.NoError(t, err)
}
