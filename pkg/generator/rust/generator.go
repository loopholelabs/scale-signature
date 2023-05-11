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

package rust

import (
	"bytes"
	"errors"
	"github.com/loopholelabs/polyglot-rs/pkg/utils"
	"github.com/loopholelabs/scale-signature/pkg/generator/templates"
	"github.com/loopholelabs/scale-signature/pkg/schema"
	"os/exec"
	"text/template"
)

const (
	defaultPackageName = "types"
)

type Generator struct {
	templ *template.Template
}

func New() (*Generator, error) {
	templ, err := template.New("").Funcs(templateFunctions()).ParseFS(templates.FS, "*rs.templ")
	if err != nil {
		return nil, err
	}

	return &Generator{
		templ: templ,
	}, nil
}

func (g *Generator) Generate(schema *schema.Schema, packageName string, version string) ([]byte, error) {
	if packageName == "" {
		packageName = defaultPackageName
	}

	buf := new(bytes.Buffer)
	err := g.templ.ExecuteTemplate(buf, "types.rs.templ", map[string]any{
		"schema":  schema,
		"version": version,
		"package": packageName,
	})
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("rustfmt")
	cmd.Stdin = bytes.NewReader(buf.Bytes())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"Primitive":               Primitive,
		"IsPrimitive":             schema.ValidPrimitiveType,
		"PolyglotPrimitive":       PolyglotPrimitive,
		"PolyglotPrimitiveEncode": PolyglotPrimitiveEncode,
		"PolyglotPrimitiveDecode": PolyglotPrimitiveDecode,
		"Deref":                   func(i *bool) bool { return *i },
		"LowerFirst":              func(s string) string { return string(s[0]+32) + s[1:] },
		"SnakeCase":               utils.SnakeCase,
		"Params":                  Params,
	}
}

func Primitive(t string) string {
	switch t {
	case "string":
		return "String"
	case "int32":
		return "i32"
	case "int64":
		return "i64"
	case "uint32":
		return "u32"
	case "uint64":
		return "u64"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "bool":
		return "bool"
	case "bytes":
		return "Vec<u8>"
	default:
		return t
	}
}

func PolyglotPrimitive(t string) string {
	switch t {
	case "string":
		return "Kind::String"
	case "int32":
		return "Kind::I32"
	case "int64":
		return "Kind::I64"
	case "uint32":
		return "Kind::U32"
	case "uint64":
		return "Kind::U64"
	case "float32":
		return "Kind::F32"
	case "float64":
		return "Kind::F64"
	case "bool":
		return "Kind::Bool"
	case "bytes":
		return "Kind::Bytes"
	default:
		return "Kind::Any"
	}
}

func PolyglotPrimitiveEncode(t string) string {
	switch t {
	case "string":
		return "encode_string"
	case "int32":
		return "encode_i32"
	case "int64":
		return "encode_i64"
	case "uint32":
		return "encode_u32"
	case "uint64":
		return "encode_u64"
	case "float32":
		return "encode_f32"
	case "float64":
		return "encode_f64"
	case "bool":
		return "encode_bool"
	case "bytes":
		return "encode_bytes"
	default:
		return t
	}
}

func PolyglotPrimitiveDecode(t string) string {
	switch t {
	case "string":
		return "decode_string"
	case "int32":
		return "decode_i32"
	case "int64":
		return "decode_i64"
	case "uint32":
		return "decode_u32"
	case "uint64":
		return "decode_u64"
	case "float32":
		return "decode_f32"
	case "float64":
		return "decode_f64"
	case "bool":
		return "decode_bool"
	case "bytes":
		return "decode_bytes"
	default:
		return ""
	}
}

func Params(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("parameters must be a list of key/value pairs")
	}
	params := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("keys must be strings")
		}
		params[key] = values[i+1]
	}
	return params, nil
}
