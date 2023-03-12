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

package generator

import (
	"bytes"
	"github.com/loopholelabs/scale-signature/pkg/generator/templates"
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"github.com/stretchr/testify/require"
	"testing"
	"text/template"
)

func TestGenerator(t *testing.T) {
	m := template.FuncMap{
		"IsPrimitive": schema.ValidPrimitiveType,
		"Deref":       func(i *bool) bool { return *i },
		"LowerFirst":  func(s string) string { return string(s[0]+32) + s[1:] },
	}
	templ, err := template.New("").Funcs(m).ParseFS(templates.FS, "*.templ")
	require.NoError(t, err)

	s := new(schema.Schema)
	err = s.Decode([]byte(`
name = "testName"
tag = "1testTag"
model testModel {
	description = "this is a test model"
    string testString {
		default = "asdfsa"
	    regex_validator {
			expression = ".*"
		}
		length_validator {
			min = 1
			max = 3
		}
	}
}

model testModel2 {
	model "myTest" {
		reference = "testModel"
	}
	string_map testMap {
		value = "testModel"
	}
	string_map testMap2 {
		value = "testModel"
		accessor = true
	}
	string_map testMap3 {
		value = "string"
	}
	string_array testArray {}

	model_map testModelMap {
		value = "testModel"
		reference = "testModel2"
	}

	model_array testModelArray {
		reference = "testModel2"
	}

	enum testEnum {
		default = "test"
		values = ["test", "test2"]
	}

	enum_array testEnumArray {
		values = ["test", "test2"]
	}

	enum_map testEnumMap {
		value = "string"
		values = ["test", "test2"]
	}
		
}
`))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	buf := new(bytes.Buffer)
	err = templ.ExecuteTemplate(buf, "types.go.templ", map[string]any{
		"schema":  s,
		"version": "v0.1.0",
		"package": "types",
	})
	require.NoError(t, err)

	t.Log(buf.String())

}
