import * as generated from "./generated";
import * as polyglot from "@loopholelabs/polyglot-ts";
import * as fs from 'fs';

test('test-output', () => {
    const emptyModelEncoder = new polyglot.Encoder();
    const emptyModel = new generated.EmptyModel();
    emptyModel.encode(emptyModelEncoder);
    fs.writeFileSync('../binaries/empty_model.bin', emptyModelEncoder.bytes, 'binary');

    const emptyModelWithDescriptionEncoder = new polyglot.Encoder();
    const emptyModelWithDescription = new generated.EmptyModelWithDescription();
    emptyModelWithDescription.encode(emptyModelWithDescriptionEncoder);
    fs.writeFileSync('../binaries/empty_model_with_description.bin', emptyModelWithDescriptionEncoder.bytes, 'binary');

    const modelWithSingleStringFieldEncoder = new polyglot.Encoder();
    const modelWithSingleStringField = new generated.ModelWithSingleStringField();
    expect(modelWithSingleStringField.stringField).toEqual('DefaultValue');
    modelWithSingleStringField.stringField = 'hello world';
    modelWithSingleStringField.encode(modelWithSingleStringFieldEncoder);
    fs.writeFileSync('../binaries/model_with_single_string_field.bin', modelWithSingleStringFieldEncoder.bytes, 'binary');

    const modelWithSingleStringFieldAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithSingleStringFieldAndDescription = new generated.ModelWithSingleStringFieldAndDescription();
    expect(modelWithSingleStringFieldAndDescription.stringField).toEqual('DefaultValue');
    modelWithSingleStringFieldAndDescription.stringField = 'hello world';
    modelWithSingleStringFieldAndDescription.encode(modelWithSingleStringFieldAndDescriptionEncoder);
    fs.writeFileSync(
        '../binaries/model_with_single_string_field_and_description.bin',
        modelWithSingleStringFieldAndDescriptionEncoder.bytes,
        'binary'
    );

    const modelWithSingleInt32FieldEncoder = new polyglot.Encoder();
    const modelWithSingleInt32Field = new generated.ModelWithSingleInt32Field();
    expect(modelWithSingleInt32Field.int32Field).toEqual(32);
    modelWithSingleInt32Field.int32Field = 42;
    modelWithSingleInt32Field.encode(modelWithSingleInt32FieldEncoder);
    fs.writeFileSync('../binaries/model_with_single_int32_field.bin', modelWithSingleInt32FieldEncoder.bytes, 'binary');

    const modelWithSingleInt32FieldAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithSingleInt32FieldAndDescription = new generated.ModelWithSingleInt32FieldAndDescription();
    expect(modelWithSingleInt32FieldAndDescription.int32Field).toEqual(32);
    modelWithSingleInt32FieldAndDescription.int32Field = 42;
    modelWithSingleInt32FieldAndDescription.encode(modelWithSingleInt32FieldAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_single_int32_field_and_description.bin', modelWithSingleInt32FieldAndDescriptionEncoder.bytes, 'binary');

    const modelWithMultipleFieldsEncoder = new polyglot.Encoder();
    const modelWithMultipleFields = new generated.ModelWithMultipleFields();
    expect(modelWithMultipleFields.stringField).toEqual('DefaultValue');
    expect(modelWithMultipleFields.int32Field).toEqual(32);
    modelWithMultipleFields.stringField = 'hello world';
    modelWithMultipleFields.int32Field = 42;
    modelWithMultipleFields.encode(modelWithMultipleFieldsEncoder);
    fs.writeFileSync('../binaries/model_with_multiple_fields.bin', modelWithMultipleFieldsEncoder.bytes, 'binary');

    const modelWithMultipleFieldsAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithMultipleFieldsAndDescription = new generated.ModelWithMultipleFieldsAndDescription();
    expect(modelWithMultipleFieldsAndDescription.stringField).toEqual('DefaultValue');
    expect(modelWithMultipleFieldsAndDescription.int32Field).toEqual(32);
    modelWithMultipleFieldsAndDescription.stringField = 'hello world';
    modelWithMultipleFieldsAndDescription.int32Field = 42;
    modelWithMultipleFieldsAndDescription.encode(modelWithMultipleFieldsAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_multiple_fields_and_description.bin', modelWithMultipleFieldsAndDescriptionEncoder.bytes, 'binary');

    const modelWithEnumEncoder = new polyglot.Encoder();
    const modelWithEnum = new generated.ModelWithEnum();
    expect(modelWithEnum.enumField).toEqual(generated.GenericEnum.DefaultValue);
    modelWithEnum.enumField = generated.GenericEnum.SecondValue;
    modelWithEnum.encode(modelWithEnumEncoder);
    fs.writeFileSync('../binaries/model_with_enum.bin', modelWithEnumEncoder.bytes, 'binary');

    const modelWithEnumAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithEnumAndDescription = new generated.ModelWithEnumAndDescription();
    expect(modelWithEnumAndDescription.enumField).toEqual(generated.GenericEnum.DefaultValue);
    modelWithEnumAndDescription.enumField = generated.GenericEnum.SecondValue;
    modelWithEnumAndDescription.encode(modelWithEnumAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_enum_and_description.bin', modelWithEnumAndDescriptionEncoder.bytes, 'binary');

    const modelWithEnumAccessorEncoder = new polyglot.Encoder();
    const modelWithEnumAccessor = new generated.ModelWithEnumAccessor();
    let enumValue = modelWithEnumAccessor.enumField
    expect(enumValue).toEqual(generated.GenericEnum.DefaultValue);
    modelWithEnumAccessor.enumField = generated.GenericEnum.SecondValue;
    modelWithEnumAccessor.encode(modelWithEnumAccessorEncoder);
    fs.writeFileSync('../binaries/model_with_enum_accessor.bin', modelWithEnumAccessorEncoder.bytes, 'binary');

    const modelWithEnumAccessorAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithEnumAccessorAndDescription = new generated.ModelWithEnumAccessorAndDescription();
    enumValue = modelWithEnumAccessorAndDescription.enumField
    expect(enumValue).toEqual(generated.GenericEnum.DefaultValue);
    modelWithEnumAccessorAndDescription.enumField = generated.GenericEnum.SecondValue;
    modelWithEnumAccessorAndDescription.encode(modelWithEnumAccessorAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_enum_accessor_and_description.bin', modelWithEnumAccessorAndDescriptionEncoder.bytes, 'binary');

    const modelWithMultipleFieldsAccessorEncoder = new polyglot.Encoder();
    const modelWithMultipleFieldsAccessor = new generated.ModelWithMultipleFieldsAccessor();
    let stringFieldValue = modelWithMultipleFieldsAccessor.stringField;
    expect(stringFieldValue).toEqual('DefaultValue');
    try {
        modelWithMultipleFieldsAccessor.stringField = 'hello world';
        fail('Expected error to be thrown');
    } catch (e) {
        // @ts-ignore
        expect(e.message).toEqual('value must match ^[a-zA-Z0-9]*$');
    }
    try {
        modelWithMultipleFieldsAccessor.stringField = "";
        fail('Expected error to be thrown');
    } catch (e) {
        // @ts-ignore
        expect(e.message).toEqual('length must be between 1 and 20');
    }
    modelWithMultipleFieldsAccessor.stringField = 'hello';
    stringFieldValue = modelWithMultipleFieldsAccessor.stringField;
    expect(stringFieldValue).toEqual('HELLO');
    let int32FieldValue = modelWithMultipleFieldsAccessor.int32Field;
    expect(int32FieldValue).toEqual(32);
    try {
        modelWithMultipleFieldsAccessor.int32Field = -1;
        fail('Expected error to be thrown');
    } catch (e) {
        // @ts-ignore
        expect(e.message).toEqual('value must be between 0 and 100');
    }
    try {
        modelWithMultipleFieldsAccessor.int32Field = 101;
        fail('Expected error to be thrown');
    } catch (e) {
        // @ts-ignore
        expect(e.message).toEqual('value must be between 0 and 100');
    }
    modelWithMultipleFieldsAccessor.int32Field = 42;
    modelWithMultipleFieldsAccessor.encode(modelWithMultipleFieldsAccessorEncoder);
    fs.writeFileSync('../binaries/model_with_multiple_fields_accessor.bin', modelWithMultipleFieldsAccessorEncoder.bytes, 'binary');

    const modelWithMultipleFieldsAccessorAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithMultipleFieldsAccessorAndDescription = new generated.ModelWithMultipleFieldsAccessorAndDescription();
    stringFieldValue = modelWithMultipleFieldsAccessorAndDescription.stringField;
    expect(stringFieldValue).toEqual('DefaultValue');
    modelWithMultipleFieldsAccessorAndDescription.stringField = 'hello world';
    int32FieldValue = modelWithMultipleFieldsAccessorAndDescription.int32Field;
    expect(int32FieldValue).toEqual(32);
    modelWithMultipleFieldsAccessorAndDescription.int32Field = 42;
    modelWithMultipleFieldsAccessorAndDescription.encode(modelWithMultipleFieldsAccessorAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_multiple_fields_accessor_and_description.bin', modelWithMultipleFieldsAccessorAndDescriptionEncoder.bytes, 'binary');

    const modelWithEmbeddedModelsEncoder = new polyglot.Encoder();
    const modelWithEmbeddedModels = new generated.ModelWithEmbeddedModels();
    expect(modelWithEmbeddedModels.embeddedEmptyModel).not.toBeNull();
    expect(modelWithEmbeddedModels.embeddedEmptyModel).not.toBeUndefined();
    expect(modelWithEmbeddedModels.embeddedModelArrayWithMultipleFieldsAccessor).not.toBeNull();
    expect(modelWithEmbeddedModels.embeddedModelArrayWithMultipleFieldsAccessor).not.toBeUndefined();
    expect(modelWithEmbeddedModels.embeddedModelArrayWithMultipleFieldsAccessor.length).toEqual(0);
    modelWithEmbeddedModels.embeddedModelArrayWithMultipleFieldsAccessor.push(modelWithMultipleFieldsAccessor);
    modelWithEmbeddedModels.encode(modelWithEmbeddedModelsEncoder);
    fs.writeFileSync('../binaries/model_with_embedded_models.bin', modelWithEmbeddedModelsEncoder.bytes, 'binary');

    const modelWithEmbeddedModelsAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithEmbeddedModelsAndDescription = new generated.ModelWithEmbeddedModelsAndDescription();
    expect(modelWithEmbeddedModelsAndDescription.embeddedEmptyModel).not.toBeNull();
    expect(modelWithEmbeddedModelsAndDescription.embeddedEmptyModel).not.toBeUndefined();
    expect(modelWithEmbeddedModelsAndDescription.embeddedModelArrayWithMultipleFieldsAccessor).not.toBeNull();
    expect(modelWithEmbeddedModelsAndDescription.embeddedModelArrayWithMultipleFieldsAccessor).not.toBeUndefined();
    expect(modelWithEmbeddedModelsAndDescription.embeddedModelArrayWithMultipleFieldsAccessor.length).toEqual(0);
    modelWithEmbeddedModelsAndDescription.embeddedModelArrayWithMultipleFieldsAccessor.push(modelWithMultipleFieldsAccessor);
    modelWithEmbeddedModelsAndDescription.encode(modelWithEmbeddedModelsAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_embedded_models_and_description.bin', modelWithEmbeddedModelsAndDescriptionEncoder.bytes, 'binary');

    const modelWithEmbeddedModelsAccessorEncoder = new polyglot.Encoder();
    const modelWithEmbeddedModelsAccessor = new generated.ModelWithEmbeddedModelsAccessor();
    let embeddedModel = modelWithEmbeddedModelsAccessor.embeddedEmptyModel;
    expect(embeddedModel).not.toBeNull();
    expect(embeddedModel).not.toBeUndefined();
    let embeddedModelArray = modelWithEmbeddedModelsAccessor.embeddedModelArrayWithMultipleFieldsAccessor;
    expect(embeddedModelArray).not.toBeNull();
    expect(embeddedModelArray).not.toBeUndefined();
    expect(embeddedModelArray.length).toEqual(0);
    modelWithEmbeddedModelsAccessor.embeddedModelArrayWithMultipleFieldsAccessor.push(modelWithMultipleFieldsAccessor);
    modelWithEmbeddedModelsAccessor.encode(modelWithEmbeddedModelsAccessorEncoder);
    fs.writeFileSync('../binaries/model_with_embedded_models_accessor.bin', modelWithEmbeddedModelsAccessorEncoder.bytes, 'binary');

    const modelWithEmbeddedModelsAccessorAndDescriptionEncoder = new polyglot.Encoder();
    const modelWithEmbeddedModelsAccessorAndDescription = new generated.ModelWithEmbeddedModelsAccessorAndDescription();
    embeddedModel = modelWithEmbeddedModelsAccessorAndDescription.embeddedEmptyModel;
    expect(embeddedModel).not.toBeNull();
    expect(embeddedModel).not.toBeUndefined();
    embeddedModelArray = modelWithEmbeddedModelsAccessorAndDescription.embeddedModelArrayWithMultipleFieldsAccessor;
    expect(embeddedModelArray).not.toBeNull();
    expect(embeddedModelArray).not.toBeUndefined();
    expect(embeddedModelArray.length).toEqual(0);
    modelWithEmbeddedModelsAccessorAndDescription.embeddedModelArrayWithMultipleFieldsAccessor.push(modelWithMultipleFieldsAccessor);
    modelWithEmbeddedModelsAccessorAndDescription.encode(modelWithEmbeddedModelsAccessorAndDescriptionEncoder);
    fs.writeFileSync('../binaries/model_with_embedded_models_accessor_and_description.bin', modelWithEmbeddedModelsAccessorAndDescriptionEncoder.bytes, 'binary');

    const modelWithAllFieldTypesEncoder = new polyglot.Encoder();
    const modelWithAllFieldTypes = new generated.ModelWithAllFieldTypes();
    expect(modelWithAllFieldTypes.stringField).toEqual('DefaultValue');
    modelWithAllFieldTypes.stringField = 'hello world';
    expect(modelWithAllFieldTypes.stringArrayField.length).toEqual(0);
    expect(modelWithAllFieldTypes.stringMapField).toEqual(new Map<string, string>());
    expect(modelWithAllFieldTypes.stringMapFieldEmbedded).toEqual(new Map<string, generated.EmptyModel>());
    modelWithAllFieldTypes.stringArrayField.push('hello', 'world');
    modelWithAllFieldTypes.stringMapField.set('hello', 'world');
    modelWithAllFieldTypes.stringMapFieldEmbedded.set('hello', emptyModel);

    expect(modelWithAllFieldTypes.int32Field).toEqual(32);
    modelWithAllFieldTypes.int32Field = 42;
    expect(modelWithAllFieldTypes.int32ArrayField.length).toEqual(0);
    expect(modelWithAllFieldTypes.int32MapField).toEqual(new Map<number, number>());
    expect(modelWithAllFieldTypes.int32MapFieldEmbedded).toEqual(new Map<number, generated.EmptyModel>());
    modelWithAllFieldTypes.int32ArrayField.push(42, 84);
    modelWithAllFieldTypes.int32MapField.set(42, 84);
    modelWithAllFieldTypes.int32MapFieldEmbedded.set(42, emptyModel);

    expect(modelWithAllFieldTypes.int64Field).toEqual(BigInt(64));
    modelWithAllFieldTypes.int64Field = BigInt(100);
    expect(modelWithAllFieldTypes.int64ArrayField.length).toEqual(0);
    expect(modelWithAllFieldTypes.int64MapField).toEqual(new Map<bigint, bigint>());
    expect(modelWithAllFieldTypes.int64MapFieldEmbedded).toEqual(new Map<bigint, generated.EmptyModel>());
    modelWithAllFieldTypes.int64ArrayField.push(BigInt(100), BigInt(200));
    modelWithAllFieldTypes.int64MapField.set(BigInt(100), BigInt(200));
    modelWithAllFieldTypes.int64MapFieldEmbedded.set(BigInt(100), emptyModel);
});

test('test-input', () => {
    const emptyModelData = fs.readFileSync("../binaries/empty_model.bin")
    const emptyModel = generated.EmptyModel.decode(new polyglot.Decoder(emptyModelData));
});