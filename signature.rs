use std::io::Error;

type NewSignature<T> = fn() -> T;

pub trait Signature {
    fn runtime_context(&self) -> dyn RuntimeContext;
}

pub trait Context {
    fn guest_context(&self) -> dyn GuestContext;
}

pub trait RuntimeContext {
    fn read(&mut self) -> Error;
    fn write(&self) -> Vec<u8>;
    fn error(&self, err: Error) -> Vec<u8>;
}

pub trait GuestContext {
    fn to_write_buffer(self) -> (u32, u32);
    fn error_write_buffer(self, error: &str) -> (u32, u32);
    unsafe fn from_read_buffer(self) -> Option<Error>;
}