use std::io::Error;
use std::io::Cursor;

pub trait RuntimeContext {
    fn read(&mut self) -> Error;
    fn write(&self) -> Vec<u8>;
    fn error(&self, err: Error) -> Vec<u8>;
}

pub trait GuestContext {
    fn to_write_buffer(self) -> (u32, u32);
    fn error_write_buffer(self, error: &str) -> (u32, u32);
    fn from_read_buffer(self, read_buff: &mut Cursor<&mut Vec<u8>>) -> Error;
}
