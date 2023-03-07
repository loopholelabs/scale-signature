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
	"fmt"
	"strings"
)

type ModelSchema struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,optional"`

	Models      []*ModelReferenceSchema      `hcl:"model,block"`
	ModelArrays []*ModelReferenceArraySchema `hcl:"modelArray,block"`
	ModelMaps   []*ModelReferenceMapSchema   `hcl:"modelMap,block"`

	Strings      []*StringSchema      `hcl:"string,block"`
	StringArrays []*StringArraySchema `hcl:"stringArray,block"`
	StringMaps   []*StringMapSchema   `hcl:"stringMap,block"`

	Bools      []*BoolSchema      `hcl:"bool,block"`
	BoolArrays []*BoolArraySchema `hcl:"boolArray,block"`

	Bytes       []*BytesSchema      `hcl:"bytes,block"`
	BytesArrays []*BytesArraySchema `hcl:"bytesArray,block"`
	BytesMaps   []*BytesMapSchema   `hcl:"bytesMap,block"`

	Enums      []*EnumSchema      `hcl:"enum,block"`
	EnumArrays []*EnumArraySchema `hcl:"enumArray,block"`
	EnumMaps   []*EnumMapSchema   `hcl:"enumMap,block"`

	Int32s      []*Int32Schema      `hcl:"int32,block"`
	Int32Arrays []*Int32ArraySchema `hcl:"int32Array,block"`
	Int32Maps   []*Int32MapSchema   `hcl:"int32Map,block"`

	Int64s      []*Int64Schema      `hcl:"int64,block"`
	Int64Arrays []*Int64ArraySchema `hcl:"int64Array,block"`
	Int64Maps   []*Int64MapSchema   `hcl:"int64Map,block"`

	Uint32s      []*Uint32Schema      `hcl:"uint32,block"`
	Uint32Arrays []*Uint32ArraySchema `hcl:"uint32Array,block"`
	Uint32Maps   []*Uint32MapSchema   `hcl:"uint32Map,block"`

	Uint64s      []*Uint64Schema      `hcl:"uint64,block"`
	Uint64Arrays []*Uint64ArraySchema `hcl:"uint64Array,block"`
	Uint64Maps   []*Uint64MapSchema   `hcl:"uint64Map,block"`

	Float32s      []*Float32Schema      `hcl:"float32,block"`
	Float32Arrays []*Float32ArraySchema `hcl:"float32Array,block"`
	Float32Maps   []*Float32MapSchema   `hcl:"float32Map,block"`

	Float64s      []*Float64Schema      `hcl:"float64,block"`
	Float64Arrays []*Float64ArraySchema `hcl:"float64Array,block"`
	Float64Maps   []*Float64MapSchema   `hcl:"float64Map,block"`
}

type ModelReferenceSchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}

func (m *ModelReferenceSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(m.Name) {
		return fmt.Errorf("invalid %s.model name: %s", model.Name, m.Name)
	}

	if !ValidLabel.MatchString(m.Reference) {
		return fmt.Errorf("invalid %s.%s.reference: %s", model.Name, m.Name, m.Reference)
	}

	return nil
}

type ModelReferenceArraySchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}

func (m *ModelReferenceArraySchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(m.Name) {
		return fmt.Errorf("invalid %s.model name: %s", model.Name, m.Name)
	}

	if !ValidLabel.MatchString(m.Reference) {
		return fmt.Errorf("invalid %s.%s.reference: %s", model.Name, m.Name, m.Reference)
	}

	return nil
}

type ModelReferenceMapSchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
	Value     string `hcl:"value,attr"`
	Accessor  bool   `hcl:"accessor,optional"`
}

func (m *ModelReferenceMapSchema) Validate(model *ModelSchema) error {
	if !ValidLabel.MatchString(m.Name) {
		return fmt.Errorf("invalid %s.model name: %s", model.Name, m.Name)
	}

	if !ValidLabel.MatchString(m.Reference) {
		return fmt.Errorf("invalid %s.%s.reference: %s", model.Name, m.Name, m.Reference)
	}

	return nil
}

func (m *ModelSchema) Normalize() {
	m.Name = TitleCaser.String(m.Name)
	for _, modelReference := range m.Models {
		modelReference.Name = TitleCaser.String(modelReference.Name)
		modelReference.Reference = TitleCaser.String(modelReference.Reference)
	}

	for _, modelReferenceArray := range m.ModelArrays {
		modelReferenceArray.Name = TitleCaser.String(modelReferenceArray.Name)
		modelReferenceArray.Reference = TitleCaser.String(modelReferenceArray.Reference)
	}

	for _, modelReferenceMap := range m.ModelMaps {
		modelReferenceMap.Name = TitleCaser.String(modelReferenceMap.Name)
		modelReferenceMap.Reference = TitleCaser.String(modelReferenceMap.Reference)

		if !ValidPrimitiveType(strings.ToLower(modelReferenceMap.Value)) {
			modelReferenceMap.Value = TitleCaser.String(modelReferenceMap.Value)
		} else {
			modelReferenceMap.Value = strings.ToLower(modelReferenceMap.Value)
		}
	}

	for _, str := range m.Strings {
		str.Name = TitleCaser.String(str.Name)
	}

	for _, strArray := range m.StringArrays {
		strArray.Name = TitleCaser.String(strArray.Name)
	}

	for _, strMap := range m.StringMaps {
		strMap.Name = TitleCaser.String(strMap.Name)

		if !ValidPrimitiveType(strings.ToLower(strMap.Value)) {
			strMap.Value = TitleCaser.String(strMap.Value)
		} else {
			strMap.Value = strings.ToLower(strMap.Value)
		}
	}

	for _, i32 := range m.Int32s {
		i32.Name = TitleCaser.String(i32.Name)
	}

	for _, i32Array := range m.Int32Arrays {
		i32Array.Name = TitleCaser.String(i32Array.Name)
	}

	for _, i32Map := range m.Int32Maps {
		i32Map.Name = TitleCaser.String(i32Map.Name)

		if !ValidPrimitiveType(strings.ToLower(i32Map.Value)) {
			i32Map.Value = TitleCaser.String(i32Map.Value)
		} else {
			i32Map.Value = strings.ToLower(i32Map.Value)
		}
	}

	for _, i64 := range m.Int64s {
		i64.Name = TitleCaser.String(i64.Name)
	}

	for _, i64Array := range m.Int64Arrays {
		i64Array.Name = TitleCaser.String(i64Array.Name)
	}

	for _, i64Map := range m.Int64Maps {
		i64Map.Name = TitleCaser.String(i64Map.Name)

		if !ValidPrimitiveType(strings.ToLower(i64Map.Value)) {
			i64Map.Value = TitleCaser.String(i64Map.Value)
		} else {
			i64Map.Value = strings.ToLower(i64Map.Value)
		}
	}

	for _, u32 := range m.Uint32s {
		u32.Name = TitleCaser.String(u32.Name)
	}

	for _, u32Array := range m.Uint32Arrays {
		u32Array.Name = TitleCaser.String(u32Array.Name)
	}

	for _, u32Map := range m.Uint32Maps {
		u32Map.Name = TitleCaser.String(u32Map.Name)

		if !ValidPrimitiveType(strings.ToLower(u32Map.Value)) {
			u32Map.Value = TitleCaser.String(u32Map.Value)
		} else {
			u32Map.Value = strings.ToLower(u32Map.Value)
		}
	}

	for _, u64 := range m.Uint64s {
		u64.Name = TitleCaser.String(u64.Name)
	}

	for _, u64Array := range m.Uint64Arrays {
		u64Array.Name = TitleCaser.String(u64Array.Name)
	}

	for _, u64Map := range m.Uint64Maps {
		u64Map.Name = TitleCaser.String(u64Map.Name)

		if !ValidPrimitiveType(strings.ToLower(u64Map.Value)) {
			u64Map.Value = TitleCaser.String(u64Map.Value)
		} else {
			u64Map.Value = strings.ToLower(u64Map.Value)
		}
	}

	for _, f32 := range m.Float32s {
		f32.Name = TitleCaser.String(f32.Name)
	}

	for _, f32Array := range m.Float32Arrays {
		f32Array.Name = TitleCaser.String(f32Array.Name)
	}

	for _, f32Map := range m.Float32Maps {
		f32Map.Name = TitleCaser.String(f32Map.Name)

		if !ValidPrimitiveType(strings.ToLower(f32Map.Value)) {
			f32Map.Value = TitleCaser.String(f32Map.Value)
		} else {
			f32Map.Value = strings.ToLower(f32Map.Value)
		}
	}

	for _, f64 := range m.Float64s {
		f64.Name = TitleCaser.String(f64.Name)
	}

	for _, f64Array := range m.Float64Arrays {
		f64Array.Name = TitleCaser.String(f64Array.Name)
	}

	for _, f64Map := range m.Float64Maps {
		f64Map.Name = TitleCaser.String(f64Map.Name)

		if !ValidPrimitiveType(strings.ToLower(f64Map.Value)) {
			f64Map.Value = TitleCaser.String(f64Map.Value)
		} else {
			f64Map.Value = strings.ToLower(f64Map.Value)
		}
	}

	for _, b := range m.Bools {
		b.Name = TitleCaser.String(b.Name)
	}

	for _, bArray := range m.BoolArrays {
		bArray.Name = TitleCaser.String(bArray.Name)
	}

	for _, enum := range m.Enums {
		enum.Name = TitleCaser.String(enum.Name)
		enum.Default = TitleCaser.String(enum.Default)
		for i := range enum.Values {
			enum.Values[i] = TitleCaser.String(enum.Values[i])
		}
	}

	for _, enumArray := range m.EnumArrays {
		enumArray.Name = TitleCaser.String(enumArray.Name)
		for i := range enumArray.Values {
			enumArray.Values[i] = TitleCaser.String(enumArray.Values[i])
		}
	}

	for _, enumMap := range m.EnumMaps {
		enumMap.Name = TitleCaser.String(enumMap.Name)
		for i := range enumMap.Values {
			enumMap.Values[i] = TitleCaser.String(enumMap.Values[i])
		}

		if !ValidPrimitiveType(strings.ToLower(enumMap.Value)) {
			enumMap.Value = TitleCaser.String(enumMap.Value)
		} else {
			enumMap.Value = strings.ToLower(enumMap.Value)
		}
	}

	for _, b := range m.Bytes {
		b.Name = TitleCaser.String(b.Name)
	}

	for _, bArray := range m.BytesArrays {
		bArray.Name = TitleCaser.String(bArray.Name)
	}

	for _, bMap := range m.BytesMaps {
		bMap.Name = TitleCaser.String(bMap.Name)

		if !ValidPrimitiveType(strings.ToLower(bMap.Value)) {
			bMap.Value = TitleCaser.String(bMap.Value)
		} else {
			bMap.Value = strings.ToLower(bMap.Value)
		}
	}

}

func (m *ModelSchema) Validate(knownModels map[string]struct{}) error {
	if !ValidLabel.MatchString(m.Name) {
		return fmt.Errorf("invalid model name: %s", m.Name)
	}
	if _, ok := knownModels[m.Name]; ok {
		return fmt.Errorf("duplicate model name: %s", m.Name)
	} else {
		knownModels[m.Name] = struct{}{}
	}

	for _, modelReference := range m.Models {
		err := modelReference.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, modelReferenceArray := range m.ModelArrays {
		err := modelReferenceArray.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, modelReferenceMap := range m.ModelMaps {
		err := modelReferenceMap.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, str := range m.Strings {
		err := str.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, strArray := range m.StringArrays {
		err := strArray.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, strMap := range m.StringMaps {
		err := strMap.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i32 := range m.Int32s {
		err := i32.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i32Array := range m.Int32Arrays {
		err := i32Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i32Map := range m.Int32Maps {
		err := i32Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i64 := range m.Int64s {
		err := i64.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i64Array := range m.Int64Arrays {
		err := i64Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, i64Map := range m.Int64Maps {
		err := i64Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u32 := range m.Uint32s {
		err := u32.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u32Array := range m.Uint32Arrays {
		err := u32Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u32Map := range m.Uint32Maps {
		err := u32Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u64 := range m.Uint64s {
		err := u64.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u64Array := range m.Uint64Arrays {
		err := u64Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, u64Map := range m.Uint64Maps {
		err := u64Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f32 := range m.Float32s {
		err := f32.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f32Array := range m.Float32Arrays {
		err := f32Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f32Map := range m.Float32Maps {
		err := f32Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f64 := range m.Float64s {
		err := f64.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f64Array := range m.Float64Arrays {
		err := f64Array.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, f64Map := range m.Float64Maps {
		err := f64Map.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, b := range m.Bools {
		err := b.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, bArray := range m.BoolArrays {
		err := bArray.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, b := range m.Bytes {
		err := b.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, bArray := range m.BytesArrays {
		err := bArray.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, bMap := range m.BytesMaps {
		err := bMap.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, enum := range m.Enums {
		err := enum.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, enumArray := range m.EnumArrays {
		err := enumArray.Validate(m)
		if err != nil {
			return err
		}
	}

	for _, enumMap := range m.EnumMaps {
		err := enumMap.Validate(m)
		if err != nil {
			return err
		}
	}

	return nil
}
