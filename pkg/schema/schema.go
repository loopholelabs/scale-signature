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
	Name   string         `hcl:"name,attr"`
	Tag    string         `hcl:"tag,attr"`
	Models []*ModelSchema `hcl:"model,block"`
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

	// Validate all models
	knownModels := make(map[string]struct{})
	for _, model := range s.Models {
		err := model.Validate(knownModels)
		if err != nil {
			return err
		}
	}

	// Ensure all model references are valid
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

		for _, modelReferenceMap := range model.ModelMaps {
			if _, ok := knownModels[modelReferenceMap.Reference]; !ok {
				return fmt.Errorf("unknown %s.%s.reference: %s", model.Name, modelReferenceMap.Name, modelReferenceMap.Reference)
			}

			if !ValidPrimitiveType(modelReferenceMap.Value) {
				if _, ok := knownModels[modelReferenceMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, modelReferenceMap.Name, modelReferenceMap.Value)
				}
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

		for _, bMap := range model.BytesMaps {
			if !ValidPrimitiveType(bMap.Value) {
				if _, ok := knownModels[bMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, bMap.Name, bMap.Value)
				}
			}
		}

		for _, enumMap := range model.EnumMaps {
			if !ValidPrimitiveType(enumMap.Value) {
				if _, ok := knownModels[enumMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, enumMap.Name, enumMap.Value)
				}
			}
		}
	}

	return nil
}

func ValidPrimitiveType(t string) bool {
	switch t {
	case "string", "int32", "int64", "uint32", "uint64", "float32", "float64", "bool", "bytes":
		return true
	default:
		return false
	}
}

func PolyglotPrimitive(t string) string {
	switch t {
	case "string":
		return "polyglot.StringKind"
	case "int32":
		return "polyglot.Int32Kind"
	case "int64":
		return "polyglot.Int64Kind"
	case "uint32":
		return "polyglot.Uint32Kind"
	case "uint64":
		return "polyglot.Uint64Kind"
	case "float32":
		return "polyglot.Float32Kind"
	case "float64":
		return "polyglot.Float64Kind"
	case "bool":
		return "polyglot.BoolKind"
	case "bytes":
		return "polyglot.BytesKind"
	default:
		return "polyglot.AnyKind"
	}
}

func PolyglotPrimitiveEncode(t string) string {
	switch t {
	case "string":
		return "String"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint32":
		return "Uint32"
	case "uint64":
		return "Uint64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "bool":
		return "Bool"
	case "bytes":
		return "Bytes"
	default:
		return ""
	}
}
