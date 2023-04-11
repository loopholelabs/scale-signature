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
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	g, err := New()
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

	int32 testInt32 {
		default = 0
	}

	enum testEnum {
		default = "test"
		accessor = true
		values = ["test", "test2"]
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
	string_array testArray {
		initial_size = 0
	}

	model_map testModelMap {
		value = "testModel"
		reference = "testModel2"
	}

	model_map testModelMap2 {
		value = "string"
		accessor = true
		reference = "testModel2"
	}

	model_map testModelMap3 {
		value = "testModel2"
		reference = "testModel"
	}

	model_array testModelArray {
		reference = "testModel2"
		initial_size = 0
	}

	enum testEnum {
		default = "test"
		values = ["test", "test2"]
	}

	enum_array testEnumArray {
		values = ["test", "test2"]
		initial_size = 0
	}

	enum_map testEnumMap {
		value = "string"
		values = ["test", "test2"]
	}
		
}
`))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	formatted, err := g.Generate(s, "v0.1.0")
	require.NoError(t, err)
	t.Log(string(formatted))
}
