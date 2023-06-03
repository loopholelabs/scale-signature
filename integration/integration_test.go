//go:build integration

/*
	Copyright 2023 Loophole Labs
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		   http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package integration

import (
	"github.com/loopholelabs/scale-signature/pkg/generator/golang"
	"github.com/loopholelabs/scale-signature/pkg/generator/rust"
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"testing"
)

func TestGolangToGolang(t *testing.T) {
	g, err := golang.New()
	require.NoError(t, err)

	s := new(schema.Schema)
	err = s.Decode([]byte(schema.MasterTestingSchema))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	const golangDir = "./golang_tests"

	formatted, err := g.Generate(s, "golang_tests", "v0.1.0")
	require.NoError(t, err)

	err = os.WriteFile(golangDir+"/generated.go", formatted, 0644)
	require.NoError(t, err)

	cmd := exec.Command("go", "test", "./...", "-v", "--tags=integration,golang", "-run", "TestOutput")
	cmd.Dir = golangDir
	out, err := cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))

	cmd = exec.Command("go", "test", "./...", "-v", "--tags=integration,golang", "-run", "TestInput")
	cmd.Dir = golangDir
	out, err = cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))
}

func TestGolangToRust(t *testing.T) {
	g, err := golang.New()
	require.NoError(t, err)

	r, err := rust.New()
	require.NoError(t, err)

	s := new(schema.Schema)
	err = s.Decode([]byte(schema.MasterTestingSchema))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	const golangDir = "./golang_tests"

	formatted, err := g.Generate(s, "golang_tests", "v0.1.0")
	require.NoError(t, err)

	err = os.WriteFile(golangDir+"/generated.go", formatted, 0644)
	require.NoError(t, err)

	const rustDir = "./rust_tests"

	formatted, err = r.Generate(s, "rust_tests", "v0.1.0")
	require.NoError(t, err)

	err = os.WriteFile(rustDir+"/generated.rs", formatted, 0644)
	require.NoError(t, err)

	cmd := exec.Command("go", "test", "./...", "-v", "--tags=integration,golang", "-run", "TestOutput")
	cmd.Dir = golangDir
	out, err := cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))

	cmd = exec.Command("cargo", "test", "test_input")
	cmd.Dir = rustDir
	out, err = cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))
}

func TestRustToGolang(t *testing.T) {
	g, err := golang.New()
	require.NoError(t, err)

	r, err := rust.New()
	require.NoError(t, err)

	s := new(schema.Schema)
	err = s.Decode([]byte(schema.MasterTestingSchema))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	const golangDir = "./golang_tests"

	formatted, err := g.Generate(s, "golang_tests", "v0.1.0")
	require.NoError(t, err)

	err = os.WriteFile(golangDir+"/generated.go", formatted, 0644)
	require.NoError(t, err)

	const rustDir = "./rust_tests"

	formatted, err = r.Generate(s, "rust_tests", "v0.1.0")
	require.NoError(t, err)

	err = os.WriteFile(rustDir+"/generated.rs", formatted, 0644)
	require.NoError(t, err)

	cmd := exec.Command("cargo", "test", "test_output")
	cmd.Dir = rustDir
	out, err := cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))

	cmd = exec.Command("go", "test", "./...", "-v", "--tags=integration,golang", "-run", "TestInput")
	cmd.Dir = golangDir
	out, err = cmd.CombinedOutput()
	assert.NoError(t, err)
	t.Log(string(out))
}
