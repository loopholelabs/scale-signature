pub mod generated;

#[cfg(test)]
mod tests {
    use std::fs;
    use std::error::Error;
    use std::io::Cursor;
    use crate::generated::Decode;

    use super::*;

    #[test]
    fn test_input() -> Result<(), Box<dyn Error>> {
        let mut empty_model_data = fs::read("../binaries/empty_model.bin")?;
        let _empty_model = generated::EmptyModel::decode(&mut Cursor::new(&mut empty_model_data))?;

        let mut empty_model_with_description_data = fs::read("../binaries/empty_model_with_description.bin")?;
        let _empty_model_with_description = generated::EmptyModelWithDescription::decode(&mut Cursor::new(&mut empty_model_with_description_data))?;

        let mut model_with_single_string_field_data = fs::read("../binaries/model_with_single_string_field.bin")?;
        let model_with_single_string_field = generated::ModelWithSingleStringField::decode(&mut Cursor::new(&mut model_with_single_string_field_data))?;
        assert_eq!(model_with_single_string_field.string_field, "hello world");

        let mut model_with_single_string_field_and_description_data = fs::read("../binaries/model_with_single_string_field_and_description.bin")?;
        let model_with_single_string_field_and_description = generated::ModelWithSingleStringFieldAndDescription::decode(&mut Cursor::new(&mut model_with_single_string_field_and_description_data))?;
        assert_eq!(model_with_single_string_field_and_description.string_field, "hello world");

        let mut model_with_single_int32_field_data = fs::read("../binaries/model_with_single_int32_field.bin")?;
        let model_with_single_int32_field = generated::ModelWithSingleInt32Field::decode(&mut Cursor::new(&mut model_with_single_int32_field_data))?;
        assert_eq!(model_with_single_int32_field.int32_field, 42);

        let mut model_with_single_int32_field_and_description_data = fs::read("../binaries/model_with_single_int32_field_and_description.bin")?;
        let model_with_single_int32_field_and_description = generated::ModelWithSingleInt32FieldAndDescription::decode(&mut Cursor::new(&mut model_with_single_int32_field_and_description_data))?;
        assert_eq!(model_with_single_int32_field_and_description.int32_field, 42);

        let mut model_with_multiple_fields_data = fs::read("../binaries/model_with_multiple_fields.bin")?;
        let model_with_multiple_fields = generated::ModelWithMultipleFields::decode(&mut Cursor::new(&mut model_with_multiple_fields_data))?;
        assert_eq!(model_with_multiple_fields.string_field, "hello world");
        assert_eq!(model_with_multiple_fields.int32_field, 42);

        let mut model_with_multiple_fields_and_description_data = fs::read("../binaries/model_with_multiple_fields_and_description.bin")?;
        let model_with_multiple_fields_and_description = generated::ModelWithMultipleFieldsAndDescription::decode(&mut Cursor::new(&mut model_with_multiple_fields_and_description_data))?;
        assert_eq!(model_with_multiple_fields_and_description.string_field, "hello world");
        assert_eq!(model_with_multiple_fields_and_description.int32_field, 42);

        let mut model_with_enum_data = fs::read("../binaries/model_with_enum.bin")?;
        let model_with_enum = generated::ModelWithEnum::decode(&mut Cursor::new(&mut model_with_enum_data))?;
        assert_eq!(model_with_enum.enum_field, generated::GenericEnum::SecondValue);

        let mut model_with_enum_and_description_data = fs::read("../binaries/model_with_enum_and_description.bin")?;
        let model_with_enum_and_description = generated::ModelWithEnumAndDescription::decode(&mut Cursor::new(&mut model_with_enum_and_description_data))?;
        assert_eq!(model_with_enum_and_description.enum_field, generated::GenericEnum::SecondValue);

        let mut model_with_enum_accessor_data = fs::read("../binaries/model_with_enum_accessor.bin")?;
        let model_with_enum_accessor = generated::ModelWithEnumAccessor::decode(&mut Cursor::new(&mut model_with_enum_accessor_data))?;
        let enum_value = model_with_enum_accessor.get_enum_field();
        assert_eq!(*enum_value, generated::GenericEnum::SecondValue);

        let mut model_with_enum_accessor_and_description_data = fs::read("../binaries/model_with_enum_accessor_and_description.bin")?;
        let model_with_enum_accessor_and_description = generated::ModelWithEnumAccessorAndDescription::decode(&mut Cursor::new(&mut model_with_enum_accessor_and_description_data))?;
        let enum_value = model_with_enum_accessor_and_description.get_enum_field();
        assert_eq!(*enum_value, generated::GenericEnum::SecondValue);

        let mut model_with_multiple_fields_accessor_data = fs::read("../binaries/model_with_multiple_fields_accessor.bin")?;
        let model_with_multiple_fields_accessor = generated::ModelWithMultipleFieldsAccessor::decode(&mut Cursor::new(&mut model_with_multiple_fields_accessor_data))?;
        let string_field_value = model_with_multiple_fields_accessor.get_string_field();
        assert_eq!(string_field_value, "HELLO");
        let int32_field_value = model_with_multiple_fields_accessor.get_int32_field();
        assert_eq!(int32_field_value, 42);

        let mut model_with_multiple_fields_accessor_and_description_data = fs::read("../binaries/model_with_multiple_fields_accessor_and_description.bin")?;
        let model_with_multiple_fields_accessor_and_description = generated::ModelWithMultipleFieldsAccessorAndDescription::decode(&mut Cursor::new(&mut model_with_multiple_fields_accessor_and_description_data))?;
        let string_field_value = model_with_multiple_fields_accessor_and_description.get_string_field();
        assert_eq!(string_field_value, "hello world");
        let int32_field_value = model_with_multiple_fields_accessor_and_description.get_int32_field();
        assert_eq!(int32_field_value, 42);

        let mut model_with_embedded_models_data = fs::read("../binaries/model_with_embedded_models.bin")?;
        let model_with_embedded_models = generated::ModelWithEmbeddedModels::decode(&mut Cursor::new(&mut model_with_embedded_models_data))?;
        assert!(model_with_embedded_models.embedded_empty_model.is_some());
        assert_eq!(model_with_embedded_models.embedded_model_array_with_multiple_fields_accessor.len(), 1);
        assert_eq!(model_with_embedded_models.embedded_model_array_with_multiple_fields_accessor[0].get_int32_field(), 42);
        assert_eq!(model_with_embedded_models.embedded_model_array_with_multiple_fields_accessor[0].get_string_field(), "HELLO");

        let mut model_with_embedded_models_and_description_data = fs::read("../binaries/model_with_embedded_models_and_description.bin")?;
        let model_with_embedded_models_and_description = generated::ModelWithEmbeddedModelsAndDescription::decode(&mut Cursor::new(&mut model_with_embedded_models_and_description_data))?;
        assert!(model_with_embedded_models_and_description.embedded_empty_model.is_some());
        assert_eq!(model_with_embedded_models_and_description.embedded_model_array_with_multiple_fields_accessor.len(), 1);
        assert_eq!(model_with_embedded_models_and_description.embedded_model_array_with_multiple_fields_accessor[0].get_int32_field(), 42);
        assert_eq!(model_with_embedded_models_and_description.embedded_model_array_with_multiple_fields_accessor[0].get_string_field(), "HELLO");

        let mut model_with_embedded_models_accessor_data = fs::read("../binaries/model_with_embedded_models_accessor.bin")?;
        let model_with_embedded_models_accessor = generated::ModelWithEmbeddedModelsAccessor::decode(&mut Cursor::new(&mut model_with_embedded_models_accessor_data))?;
        let embedded_empty_model = model_with_embedded_models_accessor.get_embedded_empty_model();
        assert!(embedded_empty_model.is_some());
        let embedded_model_array_with_multiple_fields_accessor = model_with_embedded_models_accessor.get_embedded_model_array_with_multiple_fields_accessor().unwrap();
        assert_eq!(embedded_model_array_with_multiple_fields_accessor.len(), 1);
        assert_eq!(embedded_model_array_with_multiple_fields_accessor[0].get_int32_field(), 42);
        assert_eq!(embedded_model_array_with_multiple_fields_accessor[0].get_string_field(), "HELLO");

        let mut model_with_embedded_models_accessor_and_description_data = fs::read("../binaries/model_with_embedded_models_accessor_and_description.bin")?;
        let model_with_embedded_models_accessor_and_description = generated::ModelWithEmbeddedModelsAccessorAndDescription::decode(&mut Cursor::new(&mut model_with_embedded_models_accessor_and_description_data))?;
        let embedded_empty_model = model_with_embedded_models_accessor_and_description.get_embedded_empty_model();
        assert!(embedded_empty_model.is_some());
        let embedded_model_array_with_multiple_fields_accessor = model_with_embedded_models_accessor_and_description.get_embedded_model_array_with_multiple_fields_accessor().unwrap();
        assert_eq!(embedded_model_array_with_multiple_fields_accessor[0].get_int32_field(), 42);
        assert_eq!(embedded_model_array_with_multiple_fields_accessor[0].get_string_field(), "HELLO");

        let mut model_with_all_field_types_data = fs::read("../binaries/model_with_all_field_types.bin")?;
        let model_with_all_field_types = generated::ModelWithAllFieldTypes::decode(&mut Cursor::new(&mut model_with_all_field_types_data))?;
        assert_eq!(model_with_all_field_types.string_field, "hello world");
        assert_eq!(model_with_all_field_types.string_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.string_array_field[0], "hello");
        assert_eq!(model_with_all_field_types.string_array_field[1], "world");
        assert_eq!(model_with_all_field_types.string_map_field.get("hello"), Some(&"world".to_string()));
        assert!(model_with_all_field_types.string_map_field_embedded.get("hello").is_some());

        assert_eq!(model_with_all_field_types.int32_field, 42);
        assert_eq!(model_with_all_field_types.int32_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.int32_array_field[0], 42);
        assert_eq!(model_with_all_field_types.int32_array_field[1], 84);
        assert_eq!(model_with_all_field_types.int32_map_field.get(&42), Some(&84));
        assert!(model_with_all_field_types.int32_map_field_embedded.get(&42).is_some());

        assert_eq!(model_with_all_field_types.int64_field, 100);
        assert_eq!(model_with_all_field_types.int64_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.int64_array_field[0], 100);
        assert_eq!(model_with_all_field_types.int64_array_field[1], 200);
        assert_eq!(model_with_all_field_types.int64_map_field.get(&100), Some(&200));
        assert!(model_with_all_field_types.int64_map_field_embedded.get(&100).is_some());

        assert_eq!(model_with_all_field_types.uint32_field, 42);
        assert_eq!(model_with_all_field_types.uint32_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.uint32_array_field[0], 42);
        assert_eq!(model_with_all_field_types.uint32_array_field[1], 84);
        assert_eq!(model_with_all_field_types.uint32_map_field.get(&42), Some(&84));
        assert!(model_with_all_field_types.uint32_map_field_embedded.get(&42).is_some());

        assert_eq!(model_with_all_field_types.uint64_field, 100);
        assert_eq!(model_with_all_field_types.uint64_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.uint64_array_field[0], 100);
        assert_eq!(model_with_all_field_types.uint64_array_field[1], 200);
        assert_eq!(model_with_all_field_types.uint64_map_field.get(&100), Some(&200));
        assert!(model_with_all_field_types.uint64_map_field_embedded.get(&100).is_some());

        assert_eq!(model_with_all_field_types.float32_field, 42.0);
        assert_eq!(model_with_all_field_types.float32_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.float32_array_field[0], 42.0);
        assert_eq!(model_with_all_field_types.float32_array_field[1], 84.0);

        assert_eq!(model_with_all_field_types.float64_field, 100.0);
        assert_eq!(model_with_all_field_types.float64_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.float64_array_field[0], 100.0);
        assert_eq!(model_with_all_field_types.float64_array_field[1], 200.0);

        assert_eq!(model_with_all_field_types.bool_field, false);
        assert_eq!(model_with_all_field_types.bool_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.bool_array_field[0], true);
        assert_eq!(model_with_all_field_types.bool_array_field[1], false);

        assert_eq!(model_with_all_field_types.bytes_field, &[42, 84]);
        assert_eq!(model_with_all_field_types.bytes_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.bytes_array_field[0], &[42, 84]);
        assert_eq!(model_with_all_field_types.bytes_array_field[1], &[84, 42]);

        assert_eq!(model_with_all_field_types.enum_field, generated::GenericEnum::SecondValue);
        assert_eq!(model_with_all_field_types.enum_array_field.len(), 2);
        assert_eq!(model_with_all_field_types.enum_array_field[0], generated::GenericEnum::FirstValue);
        assert_eq!(model_with_all_field_types.enum_array_field[1], generated::GenericEnum::SecondValue);
        assert_eq!(model_with_all_field_types.enum_map_field.get(&generated::GenericEnum::FirstValue), Some(&"hello world".to_string()));
        assert!(model_with_all_field_types.enum_map_field_embedded.get(&generated::GenericEnum::FirstValue).is_some());

        assert_eq!(model_with_all_field_types.model_array_field.len(), 2);

        Ok(())
    }
}