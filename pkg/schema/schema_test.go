/*
	Copyright 2022 Loophole Labs

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

package schema

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema(t *testing.T) {
	s := new(Schema)
	file, diags := hclsyntax.ParseConfig([]byte(`
name = "testName"
tag = "1testTag"
model testModel {
	description = "this is a test model"
    string testString {
		default = "asdfsa"
		accessor = true
	    regexValidator {
			expression = ".*"
		}
		lengthValidator {
			min = 1
			max = 0
		}
	}
}

model testModel2 {
	model "myTest" {
		reference = "testModel"
	}
}
`), "", hcl.Pos{Line: 1, Column: 1})
	require.False(t, diags.HasErrors())

	diags = gohcl.DecodeBody(file.Body, nil, s)
	require.False(t, diags.HasErrors())

	t.Logf("%s", spew.Sdump(s))
}
