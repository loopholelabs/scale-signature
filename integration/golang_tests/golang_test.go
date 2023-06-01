//go:build integration && golang

package golang_tests

import (
	"github.com/loopholelabs/polyglot-go"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	buf := polyglot.NewBuffer()

	emptyModel := NewEmptyModel()
	emptyModel.Encode(buf)
	err := os.WriteFile("./empty_model.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	emptyModelWithDescription := NewEmptyModelWithDescription()
	emptyModelWithDescription.Encode(buf)
	err = os.WriteFile("./empty_model_with_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleStringField := NewModelWithSingleStringField()
	require.Equal(t, "DefaultValue", modelWithSingleStringField.StringField)
	modelWithSingleStringField.StringField = "hello world"
	modelWithSingleStringField.Encode(buf)
	err = os.WriteFile("./model_with_single_string_field.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleStringFieldAndDescription := NewModelWithSingleStringFieldAndDescription()
	require.Equal(t, "DefaultValue", modelWithSingleStringFieldAndDescription.StringField)
	modelWithSingleStringFieldAndDescription.StringField = "hello world"
	modelWithSingleStringFieldAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_single_string_field_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleInt32Field := NewModelWithSingleInt32Field()
	require.Equal(t, int32(32), modelWithSingleInt32Field.Int32Field)
	modelWithSingleInt32Field.Int32Field = 42
	modelWithSingleInt32Field.Encode(buf)
	err = os.WriteFile("./model_with_single_int32_field.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleInt32FieldAndDescription := NewModelWithSingleInt32FieldAndDescription()
	require.Equal(t, int32(32), modelWithSingleInt32FieldAndDescription.Int32Field)
	modelWithSingleInt32FieldAndDescription.Int32Field = 42
	modelWithSingleInt32FieldAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_single_int32_field_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFields := NewModelWithMultipleFields()
	require.Equal(t, "DefaultValue", modelWithMultipleFields.StringField)
	require.Equal(t, int32(32), modelWithMultipleFields.Int32Field)
	modelWithMultipleFields.StringField = "hello world"
	modelWithMultipleFields.Int32Field = 42
	modelWithMultipleFields.Encode(buf)
	err = os.WriteFile("./model_with_multiple_fields.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFieldsAndDescription := NewModelWithMultipleFieldsAndDescription()
	require.Equal(t, "DefaultValue", modelWithMultipleFieldsAndDescription.StringField)
	require.Equal(t, int32(32), modelWithMultipleFieldsAndDescription.Int32Field)
	modelWithMultipleFieldsAndDescription.StringField = "hello world"
	modelWithMultipleFieldsAndDescription.Int32Field = 42
	modelWithMultipleFieldsAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_multiple_fields_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEnum := NewModelWithEnum()
	require.Equal(t, GenericEnumDefaultValue, modelWithEnum.EnumField)
	modelWithEnum.EnumField = GenericEnumSecondValue
	modelWithEnum.Encode(buf)
	err = os.WriteFile("./model_with_enum.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEnumAndDescription := NewModelWithEnumAndDescription()
	require.Equal(t, GenericEnumDefaultValue, modelWithEnumAndDescription.EnumField)
	modelWithEnumAndDescription.EnumField = GenericEnumSecondValue
	modelWithEnumAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_enum_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEnumAccessor := NewModelWithEnumAccessor()
	defaultEnumValue, err := modelWithEnumAccessor.GetEnumField()
	require.NoError(t, err)
	require.Equal(t, GenericEnumDefaultValue, defaultEnumValue)
	err = modelWithEnumAccessor.SetEnumField(GenericEnumSecondValue)
	require.NoError(t, err)
	modelWithEnumAccessor.Encode(buf)
	err = os.WriteFile("./model_with_enum_accessor.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEnumAccessorAndDescription := NewModelWithEnumAccessorAndDescription()
	defaultEnumValue, err = modelWithEnumAccessorAndDescription.GetEnumField()
	require.NoError(t, err)
	require.Equal(t, GenericEnumDefaultValue, defaultEnumValue)
	err = modelWithEnumAccessorAndDescription.SetEnumField(GenericEnumSecondValue)
	require.NoError(t, err)
	modelWithEnumAccessorAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_enum_accessor_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFieldsAccessor := NewModelWithMultipleFieldsAccessor()
	stringValue, err := modelWithMultipleFieldsAccessor.GetStringField()
	require.NoError(t, err)
	require.Equal(t, "DefaultValue", stringValue)
	err = modelWithMultipleFieldsAccessor.SetStringField("hello world")
	require.ErrorContains(t, err, "value must match ^[a-zA-Z0-9]*$")
	err = modelWithMultipleFieldsAccessor.SetStringField("")
	require.ErrorContains(t, err, "length must be between 1 and 20")
	err = modelWithMultipleFieldsAccessor.SetStringField("hello")
	require.NoError(t, err)
	stringValue, err = modelWithMultipleFieldsAccessor.GetStringField()
	require.NoError(t, err)
	require.Equal(t, strings.ToUpper("hello"), stringValue)
	int32Value, err := modelWithMultipleFieldsAccessor.GetInt32Field()
	require.NoError(t, err)
	require.Equal(t, int32(32), int32Value)
	err = modelWithMultipleFieldsAccessor.SetInt32Field(-1)
	require.ErrorContains(t, err, "value must be between 0 and 100")
	err = modelWithMultipleFieldsAccessor.SetInt32Field(101)
	require.ErrorContains(t, err, "value must be between 0 and 100")
	err = modelWithMultipleFieldsAccessor.SetInt32Field(42)
	require.NoError(t, err)
	modelWithMultipleFieldsAccessor.Encode(buf)
	err = os.WriteFile("./model_with_multiple_fields_accessor.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFieldsAccessorAndDescription := NewModelWithMultipleFieldsAccessorAndDescription()
	stringValue, err = modelWithMultipleFieldsAccessorAndDescription.GetStringField()
	require.NoError(t, err)
	require.Equal(t, "DefaultValue", stringValue)
	err = modelWithMultipleFieldsAccessorAndDescription.SetStringField("hello world")
	require.NoError(t, err)
	int32Value, err = modelWithMultipleFieldsAccessorAndDescription.GetInt32Field()
	require.NoError(t, err)
	require.Equal(t, int32(32), int32Value)
	err = modelWithMultipleFieldsAccessorAndDescription.SetInt32Field(42)
	require.NoError(t, err)
	modelWithMultipleFieldsAccessorAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_multiple_fields_accessor_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEmbeddedModels := NewModelWithEmbeddedModels()
	require.NotNil(t, modelWithEmbeddedModels.EmbeddedEmptyModel)
	require.NotNil(t, modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor)
	require.Equal(t, 64, cap(modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor))
	require.Equal(t, 0, len(modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor))
	require.IsType(t, []*ModelWithMultipleFieldsAccessor{}, modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor)
	modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor = append(modelWithEmbeddedModels.EmbeddedModelArrayWithMultipleFieldsAccessor, modelWithMultipleFieldsAccessor)
	modelWithEmbeddedModels.Encode(buf)
	err = os.WriteFile("./model_with_embedded_models.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEmbeddedModelsAndDescription := NewModelWithEmbeddedModelsAndDescription()
	require.NotNil(t, modelWithEmbeddedModelsAndDescription.EmbeddedEmptyModel)
	require.NotNil(t, modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor)
	require.Equal(t, 0, cap(modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor))
	require.Equal(t, 0, len(modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor))
	require.IsType(t, []*ModelWithMultipleFieldsAccessor{}, modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor)
	modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor = append(modelWithEmbeddedModelsAndDescription.EmbeddedModelArrayWithMultipleFieldsAccessor, modelWithMultipleFieldsAccessor)
	modelWithEmbeddedModelsAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_embedded_models_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEmbeddedModelsAccessor := NewModelWithEmbeddedModelsAccessor()
	embeddedModel, err := modelWithEmbeddedModelsAccessor.GetEmbeddedEmptyModel()
	require.NoError(t, err)
	require.NotNil(t, embeddedModel)
	embeddedModelArray, err := modelWithEmbeddedModelsAccessor.GetEmbeddedModelArrayWithMultipleFieldsAccessor()
	require.NoError(t, err)
	require.NotNil(t, embeddedModelArray)
	require.Equal(t, 0, cap(embeddedModelArray))
	require.Equal(t, 0, len(embeddedModelArray))
	require.IsType(t, []*ModelWithMultipleFieldsAccessor{}, embeddedModelArray)
	err = modelWithEmbeddedModelsAccessor.SetEmbeddedModelArrayWithMultipleFieldsAccessor([]*ModelWithMultipleFieldsAccessor{modelWithMultipleFieldsAccessor})
	require.NoError(t, err)
	modelWithEmbeddedModelsAccessor.Encode(buf)
	err = os.WriteFile("./model_with_embedded_models_accessor.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithEmbeddedModelsAccessorAndDescription := NewModelWithEmbeddedModelsAccessorAndDescription()
	embeddedModel, err = modelWithEmbeddedModelsAccessorAndDescription.GetEmbeddedEmptyModel()
	require.NoError(t, err)
	require.NotNil(t, embeddedModel)
	embeddedModelArray, err = modelWithEmbeddedModelsAccessorAndDescription.GetEmbeddedModelArrayWithMultipleFieldsAccessor()
	require.NoError(t, err)
	require.NotNil(t, embeddedModelArray)
	require.Equal(t, 0, cap(embeddedModelArray))
	require.Equal(t, 0, len(embeddedModelArray))
	require.IsType(t, []*ModelWithMultipleFieldsAccessor{}, embeddedModelArray)
	err = modelWithEmbeddedModelsAccessorAndDescription.SetEmbeddedModelArrayWithMultipleFieldsAccessor([]*ModelWithMultipleFieldsAccessor{modelWithMultipleFieldsAccessor})
	require.NoError(t, err)
	modelWithEmbeddedModelsAccessorAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_embedded_models_accessor_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithAllFieldTypes := NewModelWithAllFieldTypes()

	require.Equal(t, "DefaultValue", modelWithAllFieldTypes.StringField)
	modelWithAllFieldTypes.StringField = "hello world"
	require.Equal(t, 0, cap(modelWithAllFieldTypes.StringArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.StringArrayField))
	require.IsType(t, []string{}, modelWithAllFieldTypes.StringArrayField)
	modelWithAllFieldTypes.StringArrayField = append(modelWithAllFieldTypes.StringArrayField, "hello", "world")
	require.Equal(t, 0, len(modelWithAllFieldTypes.StringMapField))
	require.IsType(t, map[string]string{}, modelWithAllFieldTypes.StringMapField)
	modelWithAllFieldTypes.StringMapField["hello"] = "world"
	require.Equal(t, 0, len(modelWithAllFieldTypes.StringMapFieldEmbedded))
	require.IsType(t, map[string]*EmptyModel{}, modelWithAllFieldTypes.StringMapFieldEmbedded)
	modelWithAllFieldTypes.StringMapFieldEmbedded["hello"] = emptyModel

	require.Equal(t, int32(32), modelWithAllFieldTypes.Int32Field)
	modelWithAllFieldTypes.Int32Field = 42
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Int32ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int32ArrayField))
	require.IsType(t, []int32{}, modelWithAllFieldTypes.Int32ArrayField)
	modelWithAllFieldTypes.Int32ArrayField = append(modelWithAllFieldTypes.Int32ArrayField, 42, 84)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int32MapField))
	require.IsType(t, map[int32]int32{}, modelWithAllFieldTypes.Int32MapField)
	modelWithAllFieldTypes.Int32MapField[42] = 84
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int32MapFieldEmbedded))
	require.IsType(t, map[int32]*EmptyModel{}, modelWithAllFieldTypes.Int32MapFieldEmbedded)
	modelWithAllFieldTypes.Int32MapFieldEmbedded[42] = emptyModel

	require.Equal(t, int64(64), modelWithAllFieldTypes.Int64Field)
	modelWithAllFieldTypes.Int64Field = 100
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Int64ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int64ArrayField))
	require.IsType(t, []int64{}, modelWithAllFieldTypes.Int64ArrayField)
	modelWithAllFieldTypes.Int64ArrayField = append(modelWithAllFieldTypes.Int64ArrayField, 100, 200)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int64MapField))
	require.IsType(t, map[int64]int64{}, modelWithAllFieldTypes.Int64MapField)
	modelWithAllFieldTypes.Int64MapField[100] = 200
	require.Equal(t, 0, len(modelWithAllFieldTypes.Int64MapFieldEmbedded))
	require.IsType(t, map[int64]*EmptyModel{}, modelWithAllFieldTypes.Int64MapFieldEmbedded)
	modelWithAllFieldTypes.Int64MapFieldEmbedded[100] = emptyModel

	require.Equal(t, uint32(32), modelWithAllFieldTypes.Uint32Field)
	modelWithAllFieldTypes.Uint32Field = 42
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Uint32ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint32ArrayField))
	require.IsType(t, []uint32{}, modelWithAllFieldTypes.Uint32ArrayField)
	modelWithAllFieldTypes.Uint32ArrayField = append(modelWithAllFieldTypes.Uint32ArrayField, 42, 84)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint32MapField))
	require.IsType(t, map[uint32]uint32{}, modelWithAllFieldTypes.Uint32MapField)
	modelWithAllFieldTypes.Uint32MapField[42] = 84
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint32MapFieldEmbedded))
	require.IsType(t, map[uint32]*EmptyModel{}, modelWithAllFieldTypes.Uint32MapFieldEmbedded)
	modelWithAllFieldTypes.Uint32MapFieldEmbedded[42] = emptyModel

	require.Equal(t, uint64(64), modelWithAllFieldTypes.Uint64Field)
	modelWithAllFieldTypes.Uint64Field = 100
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Uint64ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint64ArrayField))
	require.IsType(t, []uint64{}, modelWithAllFieldTypes.Uint64ArrayField)
	modelWithAllFieldTypes.Uint64ArrayField = append(modelWithAllFieldTypes.Uint64ArrayField, 100, 200)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint64MapField))
	require.IsType(t, map[uint64]uint64{}, modelWithAllFieldTypes.Uint64MapField)
	modelWithAllFieldTypes.Uint64MapField[100] = 200
	require.Equal(t, 0, len(modelWithAllFieldTypes.Uint64MapFieldEmbedded))
	require.IsType(t, map[uint64]*EmptyModel{}, modelWithAllFieldTypes.Uint64MapFieldEmbedded)
	modelWithAllFieldTypes.Uint64MapFieldEmbedded[100] = emptyModel

	require.Equal(t, float32(32.32), modelWithAllFieldTypes.Float32Field)
	modelWithAllFieldTypes.Float32Field = 42.0
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Float32ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float32ArrayField))
	require.IsType(t, []float32{}, modelWithAllFieldTypes.Float32ArrayField)
	modelWithAllFieldTypes.Float32ArrayField = append(modelWithAllFieldTypes.Float32ArrayField, 42.0, 84.0)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float32MapField))
	require.IsType(t, map[float32]float32{}, modelWithAllFieldTypes.Float32MapField)
	modelWithAllFieldTypes.Float32MapField[42.0] = 84.0
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float32MapFieldEmbedded))
	require.IsType(t, map[float32]*EmptyModel{}, modelWithAllFieldTypes.Float32MapFieldEmbedded)
	modelWithAllFieldTypes.Float32MapFieldEmbedded[42.0] = emptyModel

	require.Equal(t, float64(64.64), modelWithAllFieldTypes.Float64Field)
	modelWithAllFieldTypes.Float64Field = 100.0
	require.Equal(t, 0, cap(modelWithAllFieldTypes.Float64ArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float64ArrayField))
	require.IsType(t, []float64{}, modelWithAllFieldTypes.Float64ArrayField)
	modelWithAllFieldTypes.Float64ArrayField = append(modelWithAllFieldTypes.Float64ArrayField, 100.0, 200.0)
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float64MapField))
	require.IsType(t, map[float64]float64{}, modelWithAllFieldTypes.Float64MapField)
	modelWithAllFieldTypes.Float64MapField[100.0] = 200.0
	require.Equal(t, 0, len(modelWithAllFieldTypes.Float64MapFieldEmbedded))
	require.IsType(t, map[float64]*EmptyModel{}, modelWithAllFieldTypes.Float64MapFieldEmbedded)
	modelWithAllFieldTypes.Float64MapFieldEmbedded[100.0] = emptyModel

	require.Equal(t, true, modelWithAllFieldTypes.BoolField)
	modelWithAllFieldTypes.BoolField = false
	require.Equal(t, 0, cap(modelWithAllFieldTypes.BoolArrayField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.BoolArrayField))
	require.IsType(t, []bool{}, modelWithAllFieldTypes.BoolArrayField)
	modelWithAllFieldTypes.BoolArrayField = append(modelWithAllFieldTypes.BoolArrayField, true, false)

	require.Equal(t, 512, cap(modelWithAllFieldTypes.BytesField))
	require.Equal(t, 0, len(modelWithAllFieldTypes.BytesField))
	require.IsType(t, []byte{}, modelWithAllFieldTypes.BytesField)
	modelWithAllFieldTypes.BytesField = append(modelWithAllFieldTypes.BytesField, []byte{42, 84}...)
	require.Equal(t, 0, len(modelWithAllFieldTypes.BytesArrayField))
	require.IsType(t, [][]byte{}, modelWithAllFieldTypes.BytesArrayField)
	modelWithAllFieldTypes.BytesArrayField = append(modelWithAllFieldTypes.BytesArrayField, []byte{42, 84}, []byte{84, 42})
}

func TestInput(t *testing.T) {
	emptyModel := NewEmptyModel()
	emptyModelData, err := os.ReadFile("./empty_model.bin")
	require.NoError(t, err)
	err = emptyModel.Decode(emptyModelData)
	require.NoError(t, err)

	emptyModelWithDescription := NewEmptyModelWithDescription()
	emptyModelWithDescriptionData, err := os.ReadFile("./empty_model_with_description.bin")
	require.NoError(t, err)
	err = emptyModelWithDescription.Decode(emptyModelWithDescriptionData)
	require.NoError(t, err)

	modelWithSingleStringField := NewModelWithSingleStringField()
	modelWithSingleStringFieldData, err := os.ReadFile("./model_with_single_string_field.bin")
	require.NoError(t, err)
	err = modelWithSingleStringField.Decode(modelWithSingleStringFieldData)
	require.NoError(t, err)
	require.Equal(t, "hello world", modelWithSingleStringField.StringField)

	modelWithSingleStringFieldAndDescription := NewModelWithSingleStringFieldAndDescription()
	modelWithSingleStringFieldAndDescriptionData, err := os.ReadFile("./model_with_single_string_field_and_description.bin")
	require.NoError(t, err)
	err = modelWithSingleStringFieldAndDescription.Decode(modelWithSingleStringFieldAndDescriptionData)
	require.NoError(t, err)

	modelWithSingleInt32Field := NewModelWithSingleInt32Field()
	modelWithSingleInt32FieldData, err := os.ReadFile("./model_with_single_int32_field.bin")
	require.NoError(t, err)
	err = modelWithSingleInt32Field.Decode(modelWithSingleInt32FieldData)
	require.NoError(t, err)
	require.Equal(t, int32(42), modelWithSingleInt32Field.Int32Field)

	modelWithSingleInt32FieldAndDescription := NewModelWithSingleInt32FieldAndDescription()
	modelWithSingleInt32FieldAndDescriptionData, err := os.ReadFile("./model_with_single_int32_field_and_description.bin")
	require.NoError(t, err)
	err = modelWithSingleInt32FieldAndDescription.Decode(modelWithSingleInt32FieldAndDescriptionData)
	require.NoError(t, err)
	require.Equal(t, int32(42), modelWithSingleInt32FieldAndDescription.Int32Field)

	modelWithMultipleFields := NewModelWithMultipleFields()
	modelWithMultipleFieldsData, err := os.ReadFile("./model_with_multiple_fields.bin")
	require.NoError(t, err)
	err = modelWithMultipleFields.Decode(modelWithMultipleFieldsData)
	require.NoError(t, err)
	require.Equal(t, "hello world", modelWithMultipleFields.StringField)
	require.Equal(t, int32(42), modelWithMultipleFields.Int32Field)

	modelWithMultipleFieldsAndDescription := NewModelWithMultipleFieldsAndDescription()
	modelWithMultipleFieldsAndDescriptionData, err := os.ReadFile("./model_with_multiple_fields_and_description.bin")
	require.NoError(t, err)
	err = modelWithMultipleFieldsAndDescription.Decode(modelWithMultipleFieldsAndDescriptionData)
	require.NoError(t, err)
	require.Equal(t, "hello world", modelWithMultipleFieldsAndDescription.StringField)
	require.Equal(t, int32(42), modelWithMultipleFieldsAndDescription.Int32Field)

	modelWithEnum := NewModelWithEnum()
	modelWithEnumData, err := os.ReadFile("./model_with_enum.bin")
	require.NoError(t, err)
	err = modelWithEnum.Decode(modelWithEnumData)
	require.NoError(t, err)
	require.Equal(t, GenericEnumSecondValue, modelWithEnum.EnumField)

	modelWithEnumAndDescription := NewModelWithEnumAndDescription()
	modelWithEnumAndDescriptionData, err := os.ReadFile("./model_with_enum_and_description.bin")
	require.NoError(t, err)
	err = modelWithEnumAndDescription.Decode(modelWithEnumAndDescriptionData)
	require.NoError(t, err)
	require.Equal(t, GenericEnumSecondValue, modelWithEnumAndDescription.EnumField)

	modelWithEnumAccessor := NewModelWithEnumAccessor()
	modelWithEnumAccessorData, err := os.ReadFile("./model_with_enum_accessor.bin")
	require.NoError(t, err)
	err = modelWithEnumAccessor.Decode(modelWithEnumAccessorData)
	require.NoError(t, err)
	enumValue, err := modelWithEnumAccessor.GetEnumField()
	require.NoError(t, err)
	require.Equal(t, GenericEnumSecondValue, enumValue)
}
