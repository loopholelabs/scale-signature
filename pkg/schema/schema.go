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
	ValidLabel    = regexp.MustCompile(`^[A-Za-z0-9]*$`)
	InvalidString = regexp.MustCompile(`[^A-Za-z0-9-.]`)
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

	if InvalidString.MatchString(s.Tag) {
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

		for _, i64 := range model.Int64s {
			err := i64.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, i64Array := range model.Int64Arrays {
			err := i64Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, i64Map := range model.Int64Maps {
			err := i64Map.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u32 := range model.Uint32s {
			err := u32.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u32Array := range model.Uint32Arrays {
			err := u32Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u32Map := range model.Uint32Maps {
			err := u32Map.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u64 := range model.Uint64s {
			err := u64.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u64Array := range model.Uint64Arrays {
			err := u64Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, u64Map := range model.Uint64Maps {
			err := u64Map.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f32 := range model.Float32s {
			err := f32.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f32Array := range model.Float32Arrays {
			err := f32Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f32Map := range model.Float32Maps {
			err := f32Map.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f64 := range model.Float64s {
			err := f64.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f64Array := range model.Float64Arrays {
			err := f64Array.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, f64Map := range model.Float64Maps {
			err := f64Map.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, b := range model.Bools {
			err := b.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, bArray := range model.BoolArrays {
			err := bArray.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, b := range model.Bytes {
			err := b.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, bArray := range model.BytesArrays {
			err := bArray.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, bMap := range model.BytesMaps {
			err := bMap.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, enum := range model.Enums {
			err := enum.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, enumArray := range model.EnumArrays {
			err := enumArray.Validate(model)
			if err != nil {
				return err
			}
		}

		for _, enumMap := range model.EnumMaps {
			err := enumMap.Validate(model)
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

		for _, i64Map := range model.Int64Maps {
			if !validPrimitiveType(i64Map.Value) {
				if _, ok := knownModels[i64Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, i64Map.Name, i64Map.Value)
				}
			}
		}

		for _, u32Map := range model.Uint32Maps {
			if !validPrimitiveType(u32Map.Value) {
				if _, ok := knownModels[u32Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, u32Map.Name, u32Map.Value)
				}
			}
		}

		for _, u64Map := range model.Uint64Maps {
			if !validPrimitiveType(u64Map.Value) {
				if _, ok := knownModels[u64Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, u64Map.Name, u64Map.Value)
				}
			}
		}

		for _, f32Map := range model.Float32Maps {
			if !validPrimitiveType(f32Map.Value) {
				if _, ok := knownModels[f32Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, f32Map.Name, f32Map.Value)
				}
			}
		}

		for _, f64Map := range model.Float64Maps {
			if !validPrimitiveType(f64Map.Value) {
				if _, ok := knownModels[f64Map.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, f64Map.Name, f64Map.Value)
				}
			}
		}

		for _, bMap := range model.BytesMaps {
			if !validPrimitiveType(bMap.Value) {
				if _, ok := knownModels[bMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, bMap.Name, bMap.Value)
				}
			}
		}

		for _, enumMap := range model.EnumMaps {
			if !validPrimitiveType(enumMap.Value) {
				if _, ok := knownModels[enumMap.Value]; !ok {
					return fmt.Errorf("unknown %s.%s.value: %s", model.Name, enumMap.Name, enumMap.Value)
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
