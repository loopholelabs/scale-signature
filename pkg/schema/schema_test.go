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

package schema

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema(t *testing.T) {
	s := new(Schema)
	err := s.Decode([]byte(`
name = "testName"
tag = "1testTag"
model testModel {
	description = "this is a test model"
    string testString {
		default = "asdfsa"
	    regexValidator {
			expression = ".*"
		}
		lengthValidator {
			min = 1
			max = 3
		}
	}
}

model testModel2 {
	model "myTest" {
		reference = "testModel"
	}
}
`))
	require.NoError(t, err)

	require.NoError(t, s.Validate())

	assert.Equal(t, "testModel", s.Models[0].Name)
	assert.Equal(t, "testModel2", s.Models[1].Name)
	assert.Equal(t, "myTest", s.Models[1].Models[0].Name)
	assert.Equal(t, "testModel", s.Models[1].Models[0].Reference)

	assert.Equal(t, "testName", s.Name)
	assert.Equal(t, "1testTag", s.Tag)

	assert.Equal(t, "this is a test model", s.Models[0].Description)
	assert.Equal(t, "testString", s.Models[0].Strings[0].Name)
	assert.True(t, *s.Models[0].Strings[0].Accessor)
}
