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
        let empty_model = generated::EmptyModel::new();
        let mut empty_model_data = fs::read("../binaries/empty_model.bin")?;
        generated::EmptyModel::decode(&mut Cursor::new(&mut empty_model_data))?;

        // let empty_model_with_description = generated::EmptyModelWithDescription::new();
        // let mut empty_model_with_description_data = fs::read("../binaries/empty_model_with_description.bin")?;
        // empty_model_with_description.decode(&empty_model_with_description_data)?;
        //
        // // The process continues for all models...
        //
        // let model_with_all_field_types = generated::ModelWithAllFieldTypes::new();
        // let mut model_with_all_field_types_data = fs::read("../binaries/model_with_all_field_types.bin")?;
        // model_with_all_field_types.decode(&model_with_all_field_types_data)?;
        //
        // assert_eq!("hello world", model_with_all_field_types.string_field);
        // assert_eq!(2, model_with_all_field_types.string_array_field.len());
        // assert_eq!("hello", model_with_all_field_types.string_array_field[0]);
        // assert_eq!("world", model_with_all_field_types.string_array_field[1]);
        // assert_eq!("world", model_with_all_field_types.string_map_field.get("hello").unwrap());
        // assert_eq!(empty_model, model_with_all_field_types.string_map_field_embedded.get("hello").unwrap());
        //
        // // The process continues for all field types...
        //
        Ok(())
    }
}