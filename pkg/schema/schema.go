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
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
)

const (
	V1AlphaVersion = "v1alpha"
)

var (
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidTag  = errors.New("invalid tag")
)

var (
	ValidLabel    = regexp.MustCompile(`^[A-Za-z0-9]*$`)
	InvalidString = regexp.MustCompile(`[^A-Za-z0-9-.]`)
)

var (
	TitleCaser = cases.Title(language.Und, cases.NoLower)
)

type Schema struct {
	Version string         `hcl:"version,attr"`
	Name    string         `hcl:"name,attr"`
	Tag     string         `hcl:"tag,attr"`
	Enums   []*EnumSchema  `hcl:"enum,block"`
	Models  []*ModelSchema `hcl:"model,block"`
}

func ReadSchema(path string) (*Schema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	schema := new(Schema)
	return schema, schema.Decode(data)
}

func (s *Schema) Decode(data []byte) error {
	file, diag := hclsyntax.ParseConfig(data, "", hcl.Pos{Line: 1, Column: 1})
	if diag.HasErrors() {
		return diag.Errs()[0]
	}

	diag = gohcl.DecodeBody(file.Body, nil, s)
	if diag.HasErrors() {
		return diag.Errs()[0]
	}

	return nil
}

func (s *Schema) Validate() error {
	switch s.Version {
	case V1AlphaVersion:
		if !ValidLabel.MatchString(s.Name) {
			return ErrInvalidName
		}

		if InvalidString.MatchString(s.Tag) {
			return ErrInvalidTag
		}

		// Transform all model names and references to TitleCase (e.g. "myModel" -> "MyModel")
		for _, model := range s.Models {
			model.Normalize()
		}

		// Transform all model names and references to TitleCase (e.g. "myModel" -> "MyModel")
		for _, enum := range s.Enums {
			enum.Normalize()
		}

		// Validate all models
		knownModels := make(map[string]struct{})
		for _, model := range s.Models {
			err := model.Validate(knownModels, s.Enums)
			if err != nil {
				return err
			}
		}

		// Validate all enums
		knownEnums := make(map[string]struct{})
		for _, enum := range s.Enums {
			err := enum.Validate(knownEnums)
			if err != nil {
				return err
			}
		}

		// Ensure all model and enum references are valid
		for _, model := range s.Models {
			for _, modelReference := range model.Models {
				if _, ok := knownModels[modelReference.Reference]; !ok {
					return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, modelReference.Name, modelReference.Reference)
				}
			}

			for _, modelReferenceArray := range model.ModelArrays {
				if _, ok := knownModels[modelReferenceArray.Reference]; !ok {
					return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, modelReferenceArray.Name, modelReferenceArray.Reference)
				}
			}

			for _, strMap := range model.StringMaps {
				if !ValidPrimitiveType(strMap.Value) {
					if _, ok := knownModels[strMap.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, strMap.Name, strMap.Value)
					}
				}
			}

			for _, i32Map := range model.Int32Maps {
				if !ValidPrimitiveType(i32Map.Value) {
					if _, ok := knownModels[i32Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, i32Map.Name, i32Map.Value)
					}
				}
			}

			for _, i64Map := range model.Int64Maps {
				if !ValidPrimitiveType(i64Map.Value) {
					if _, ok := knownModels[i64Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, i64Map.Name, i64Map.Value)
					}
				}
			}

			for _, u32Map := range model.Uint32Maps {
				if !ValidPrimitiveType(u32Map.Value) {
					if _, ok := knownModels[u32Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, u32Map.Name, u32Map.Value)
					}
				}
			}

			for _, u64Map := range model.Uint64Maps {
				if !ValidPrimitiveType(u64Map.Value) {
					if _, ok := knownModels[u64Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, u64Map.Name, u64Map.Value)
					}
				}
			}

			for _, f32Map := range model.Float32Maps {
				if !ValidPrimitiveType(f32Map.Value) {
					if _, ok := knownModels[f32Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, f32Map.Name, f32Map.Value)
					}
				}
			}

			for _, f64Map := range model.Float64Maps {
				if !ValidPrimitiveType(f64Map.Value) {
					if _, ok := knownModels[f64Map.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, f64Map.Name, f64Map.Value)
					}
				}
			}

			for _, enumReference := range model.Enums {
				if _, ok := knownEnums[enumReference.Reference]; !ok {
					return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, enumReference.Name, enumReference.Reference)
				}
			}

			for _, enumReferenceArray := range model.EnumArrays {
				if _, ok := knownEnums[enumReferenceArray.Reference]; !ok {
					return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, enumReferenceArray.Name, enumReferenceArray.Reference)
				}
			}

			for _, enumReferenceMap := range model.EnumMaps {
				if _, ok := knownEnums[enumReferenceMap.Reference]; !ok {
					return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, enumReferenceMap.Name, enumReferenceMap.Reference)
				}

				if !ValidPrimitiveType(enumReferenceMap.Value) {
					if _, ok := knownModels[enumReferenceMap.Value]; !ok {
						return fmt.Errorf("unknown %s.%s.value: %s", model.Name, enumReferenceMap.Name, enumReferenceMap.Value)
					}
				}
			}
		}

		return nil
	default:
		return fmt.Errorf("unknown schema version: %s", s.Version)
	}
}

func ValidPrimitiveType(t string) bool {
	switch t {
	case "string", "int32", "int64", "uint32", "uint64", "float32", "float64", "bool", "bytes":
		return true
	default:
		return false
	}
}

const MasterTestingSchema = `
version = "v1alpha"
name = "MasterSchema"
tag = "MasterSchemaTag"

model EmptyModel {}

model EmptyModelWithDescription {
	description = "Test Description"
}

model ModelWithSingleStringField {
	string StringField {
		default = "DefaultValue"
	}
}

model ModelWithSingleStringFieldAndDescription {
	description = "Test Description"

	string StringField {
		default = "DefaultValue"
	}
}

model ModelWithSingleInt32Field {
	int32 Int32Field {
		default = 32
	}
}

model ModelWithSingleInt32FieldAndDescription {
	description = "Test Description"

	int32 Int32Field {
		default = 32
	}
}

model ModelWithMultipleFields {
	string StringField {
		default = "DefaultValue"
	}

	int32 Int32Field {
		default = 32
	}
}

model ModelWithMultipleFieldsAndDescription {
	description = "Test Description"
	
	string StringField {
		default = "DefaultValue"
	}

	int32 Int32Field {
		default = 32
	}
}

enum GenericEnum {
	values = ["FirstValue", "SecondValue", "DefaultValue"]
}

model ModelWithEnum {
	enum EnumField {
		default = "DefaultValue"
		reference = "GenericEnum"
	}
}

model ModelWithEnumAndDescription {
	description = "Test Description"

	enum EnumField {
		default = "DefaultValue"
		reference = "GenericEnum"
	}
}

model ModelWithEnumAccessor {
	enum EnumField {
		default = "DefaultValue"
		reference = "GenericEnum"
		accessor = true
	}
}

model ModelWithEnumAccessorAndDescription {
	description = "Test Description"

	enum EnumField {
		default = "DefaultValue"
		reference = "GenericEnum"
		accessor = true
	}
}

model ModelWithMultipleFieldsAccessor {
	string StringField {
		default = "DefaultValue"
		accessor = true
		regex_validator {
			expression = "^[a-zA-Z0-9]*$"
		}
		length_validator {
			min = 1
			max= 20
		}
		case_modifier {
			kind = "upper"
		}
	}

	int32 Int32Field {
		default = 32
		accessor = true
		limit_validator {
			min = 0
			max = 100
		}
	}
}

model ModelWithMultipleFieldsAccessorAndDescription {
	description = "Test Description"

	string StringField {
		default = "DefaultValue"
		accessor = true
	}

	int32 Int32Field {
		default = 32
		accessor = true
	}
}

model ModelWithEmbeddedModels {
	model EmbeddedEmptyModel {
		reference = "EmptyModel"
	}

	model_array EmbeddedModelArrayWithMultipleFieldsAccessor {
		reference = "ModelWithMultipleFieldsAccessor"
		initial_size = 64
	}
}

model ModelWithEmbeddedModelsAndDescription {
	description = "Test Description"

	model EmbeddedEmptyModel {
		reference = "EmptyModel"
	}		

	model_array EmbeddedModelArrayWithMultipleFieldsAccessor {
		reference = "ModelWithMultipleFieldsAccessor"
		initial_size = 0
	}
}

model ModelWithEmbeddedModelsAccessor {
	model EmbeddedEmptyModel {
		reference = "EmptyModel"
		accessor = true
	}

	model_array EmbeddedModelArrayWithMultipleFieldsAccessor {
		reference = "ModelWithMultipleFieldsAccessor"
		initial_size = 0
		accessor = true
	}
}

model ModelWithEmbeddedModelsAccessorAndDescription {
	description = "Test Description"

	model EmbeddedEmptyModel {
		reference = "EmptyModel"
		accessor = true
	}

	model_array EmbeddedModelArrayWithMultipleFieldsAccessor {
		reference = "ModelWithMultipleFieldsAccessor"
		initial_size = 0
		accessor = true
	}
}

model ModelWithAllFieldTypes {
	string StringField {
		default = "DefaultValue"
	}

	string_array StringArrayField {
		initial_size = 0
	}

	string_map StringMapField {
		value = "string"
	}

	string_map StringMapFieldEmbedded {
		value = "EmptyModel"
	}

	int32 Int32Field {
		default = 32
	}

	int32_array Int32ArrayField {
		initial_size = 0
	}

	int32_map Int32MapField {
		value = "int32"
	}

	int32_map Int32MapFieldEmbedded {
		value = "EmptyModel"
	}

	int64 Int64Field {
		default = 64
	}

	int64_array Int64ArrayField {
		initial_size = 0
	}

	int64_map Int64MapField {
		value = "int64"
	}

	int64_map Int64MapFieldEmbedded {
		value = "EmptyModel"
	}

	uint32 Uint32Field {
		default = 32
	}

	uint32_array Uint32ArrayField {
		initial_size = 0
	}

	uint32_map Uint32MapField {
		value = "uint32"
	}

	uint32_map Uint32MapFieldEmbedded {
		value = "EmptyModel"
	}

	uint64 Uint64Field {
		default = 64
	}

	uint64_array Uint64ArrayField {
		initial_size = 0
	}

	uint64_map Uint64MapField {
		value = "uint64"
	}

	uint64_map Uint64MapFieldEmbedded {
		value = "EmptyModel"
	}

	float32 Float32Field {
		default = 32.32
	}

	float32_array Float32ArrayField {
		initial_size = 0
	}

	float32_map Float32MapField {
		value = "float32"
	}

	float32_map Float32MapFieldEmbedded {
		value = "EmptyModel"
	}

	float64 Float64Field {
		default = 64.64
	}

	float64_array Float64ArrayField {
		initial_size = 0
	}

	float64_map Float64MapField {
		value = "float64"
	}

	float64_map Float64MapFieldEmbedded {
		value = "EmptyModel"
	}

	bool BoolField {
		default = true
	}

	bool_array BoolArrayField {
		initial_size = 0
	}

	bytes BytesField {
		initial_size = 512
	}

	bytes_array BytesArrayField {
		initial_size = 0
	}

	enum EnumField {
		reference = "GenericEnum"
		default = "DefaultValue"
	}

	enum_array EnumArrayField {
		reference = "GenericEnum"
		initial_size = 0
	}

	enum_map EnumMapField {
		reference = "GenericEnum"
		value = "string"
	}

	enum_map EnumMapFieldEmbedded {
		reference = "GenericEnum"
		value = "EmptyModel"
	}

	model ModelField {
		reference = "EmptyModel"
	}

	model_array ModelArrayField {
		reference = "EmptyModel"
		initial_size = 0
	}
}
`
