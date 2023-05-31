//go:build integration && golang

package tests

import (
	"github.com/loopholelabs/polyglot-go"
	"github.com/stretchr/testify/require"
	"os"
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
	modelWithSingleStringField.StringField = "hello world"
	modelWithSingleStringField.Encode(buf)
	err = os.WriteFile("./model_with_single_string_field.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleStringFieldAndDescription := NewModelWithSingleStringFieldAndDescription()
	modelWithSingleStringFieldAndDescription.StringField = "hello world"
	modelWithSingleStringFieldAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_single_string_field_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleInt32Field := NewModelWithSingleInt32Field()
	modelWithSingleInt32Field.Int32Field = 42
	modelWithSingleInt32Field.Encode(buf)
	err = os.WriteFile("./model_with_single_int32_field.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithSingleInt32FieldAndDescription := NewModelWithSingleInt32FieldAndDescription()
	modelWithSingleInt32FieldAndDescription.Int32Field = 42
	modelWithSingleInt32FieldAndDescription.Encode(buf)
	err = os.WriteFile("./model_with_single_int32_field_and_description.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFields := NewModelWithMultipleFields()
	modelWithMultipleFields.StringField = "hello world"
	modelWithMultipleFields.Int32Field = 42
	modelWithMultipleFields.Encode(buf)
	err = os.WriteFile("./model_with_multiple_fields.bin", buf.Bytes(), 0644)
	require.NoError(t, err)
	buf.Reset()

	modelWithMultipleFieldsAndDescription := NewModelWithMultipleFieldsAndDescription()
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

	//modelWithEnumAccessor := NewModelWithEnumAccessor()
	//require.Equal(t, GenericEnumDefaultValue, modelWithEnumAccessor.enumField)
	//modelWithEnumAccessor.SetEnumField(GenericEnumSecondValue)
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
}
