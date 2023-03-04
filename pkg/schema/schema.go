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
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"os"
	"regexp"
)

var (
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidTag  = errors.New("invalid tag")
)

var (
	ValidLabel = regexp.MustCompile(`^[A-Za-z0-9]*$`)
)

type Schema struct {
	Name   string        `hcl:"name,attr"`
	Tag    string        `hcl:"tag,attr"`
	Models []ModelSchema `hcl:"model,block"`
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

	if !ValidLabel.MatchString(s.Tag) {
		return ErrInvalidTag
	}

	knownModels := make(map[string]struct{})

	for _, model := range s.Models {
		if !ValidLabel.MatchString(model.Name) {
			return fmt.Errorf("invalid model name: %s", model.Name)
		}
		knownModels[model.Name] = struct{}{}

		for _, modelReference := range model.Models {
			err := modelReference.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, modelReferenceArray := range model.ModelArrays {
			err := modelReferenceArray.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, modelReferenceMap := range model.ModelMaps {
			err := modelReferenceMap.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, str := range model.Strings {
			err := str.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, strArray := range model.StringArrays {
			err := strArray.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, strMap := range model.StringMaps {
			err := strMap.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, i32 := range model.Int32s {
			err := i32.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, i32Array := range model.Int32Arrays {
			err := i32Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, i32Map := range model.Int32Maps {
			err := i32Map.Validate(model)
			if err != nil {
				return err
			}
		}
	}

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

			if !validPrimitiveType(modelReferenceMap.Value) {
				if _, ok := knownModels[modelReferenceMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, modelReferenceMap.Name, modelReferenceMap.Value)
				}
			}
		}

		for _, strMap := range model.StringMaps {
			if !validPrimitiveType(strMap.Value) {
				if _, ok := knownModels[strMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, strMap.Name, strMap.Value)
				}
			}
		}

		for _, i32Map := range model.Int32Maps {
			if !validPrimitiveType(i32Map.Value) {
				if _, ok := knownModels[i32Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, i32Map.Name, i32Map.Value)
				}
			}
		}
	}

	return nil
}

func validPrimitiveType(t string) bool {
	switch t {
	case "string", "int32", "int64", "uint32", "uint64", "float32", "float64", "bool", "bytes":
		return true
	default:
		return false
	}
}
