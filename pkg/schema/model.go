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
	ModelsArray []ModelReferenceArraySchema `hcl:"modelArray,block"`

	Strings      []StringSchema      `hcl:"string,block"`
	StringsArray []StringArraySchema `hcl:"stringArray,block"`

	Bools      []BoolSchema      `hcl:"bool,block"`
	BoolsArray []BoolArraySchema `hcl:"boolArray,block"`

	Bytes      []BytesSchema      `hcl:"bytes,block"`
	BytesArray []BytesArraySchema `hcl:"bytesArray,block"`

	Enums      []EnumSchema      `hcl:"enum,block"`
	EnumsArray []EnumArraySchema `hcl:"enumArray,block"`

	Int32s     []Int32Schema      `hcl:"int32,block"`
	Int32Array []Int32ArraySchema `hcl:"int32Array,block"`

	Int64s      []Int64Schema      `hcl:"int64,block"`
	Int64sArray []Int64ArraySchema `hcl:"int64Array,block"`

	Uint32s      []Uint32Schema      `hcl:"uint32,block"`
	Uint32sArray []Uint32ArraySchema `hcl:"uint32Array,block"`

	Uint64s      []Uint64Schema      `hcl:"uint64,block"`
	Uint64sArray []Uint64ArraySchema `hcl:"uint64Array,block"`

	Float32s      []Float32Schema      `hcl:"float32,block"`
	Float32sArray []Float32ArraySchema `hcl:"float32Array,block"`

	Float64s      []Float64Schema      `hcl:"float64,block"`
	Float64sArray []Float64ArraySchema `hcl:"float64Array,block"`
}

type ModelReferenceSchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}

type ModelReferenceArraySchema struct {
	Name      string `hcl:"name,label"`
	Reference string `hcl:"reference,attr"`
}
