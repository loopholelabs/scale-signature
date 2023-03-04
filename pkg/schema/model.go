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

type ModelSchema struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,optional"`

	Models      []ModelReferenceSchema      `hcl:"model,block"`
	ModelArrays []ModelReferenceArraySchema `hcl:"modelArray,block"`
	ModelMaps   []ModelReferenceMapSchema   `hcl:"modelMap,block"`

	Strings      []StringSchema      `hcl:"string,block"`
	StringArrays []StringArraySchema `hcl:"stringArray,block"`
	StringMaps   []StringMapSchema   `hcl:"stringMap,block"`

	Bools      []BoolSchema      `hcl:"bool,block"`
	BoolArrays []BoolArraySchema `hcl:"boolArray,block"`

	Bytes       []BytesSchema      `hcl:"bytes,block"`
	BytesArrays []BytesArraySchema `hcl:"bytesArray,block"`
	BytesMaps   []BytesMapSchema   `hcl:"bytesMap,block"`

	Enums      []EnumSchema      `hcl:"enum,block"`
	EnumArrays []EnumArraySchema `hcl:"enumArray,block"`
	EnumMaps   []EnumMapSchema   `hcl:"enumMap,block"`

	Int32s     []Int32Schema      `hcl:"int32,block"`
	Int32Array []Int32ArraySchema `hcl:"int32Array,block"`
	Int32Maps  []Int32MapSchema   `hcl:"int32Map,block"`

	Int64s      []Int64Schema      `hcl:"int64,block"`
	Int64Arrays []Int64ArraySchema `hcl:"int64Array,block"`
	Int64Maps   []Int64MapSchema   `hcl:"int64Map,block"`

	Uint32s      []Uint32Schema      `hcl:"uint32,block"`
	Uint32Arrays []Uint32ArraySchema `hcl:"uint32Array,block"`
	Uint32Maps   []Uint32MapSchema   `hcl:"uint32Map,block"`

	Uint64s      []Uint64Schema      `hcl:"uint64,block"`
	Uint64Arrays []Uint64ArraySchema `hcl:"uint64Array,block"`
	Uint64Maps   []Uint64MapSchema   `hcl:"uint64Map,block"`

	Float32s      []Float32Schema      `hcl:"float32,block"`
	Float32Arrays []Float32ArraySchema `hcl:"float32Array,block"`
	Float32Maps   []Float32MapSchema   `hcl:"float32Map,block"`

	Float64s      []Float64Schema      `hcl:"float64,block"`
	Float64Arrays []Float64ArraySchema `hcl:"float64Array,block"`
	Float64Maps   []Float64MapSchema   `hcl:"float64Map,block"`
}

type ModelReferenceSchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}

type ModelReferenceArraySchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}

type ModelReferenceMapSchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
	Value     string `hcl:"value,attr"`
	Accessor  bool   `hcl:"accessor,optional"`
}
