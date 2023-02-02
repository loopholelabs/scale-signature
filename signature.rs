pub type NewSignature<T> = fn() -> T;

pub trait Signature {
    fn runtime_context(&mut self) -> &mut dyn RuntimeContext;
}

pub trait Context {
    fn guest_context(&mut self) -> &mut dyn GuestContext;
}

pub trait RuntimeContext {
    fn read(&mut self, b: &mut Vec<u8>) -> Option<Box<dyn std::error::Error>>;
    fn write(&self) -> Vec<u8>;
    fn error(&self, err: Box<dyn std::error::Error>) -> Vec<u8>;
}

pub trait GuestContext {
    unsafe fn from_read_buffer(&mut self) -> Option<Box<dyn std::error::Error>>;
    unsafe fn to_write_buffer(&mut self) -> (u32, u32);
    unsafe fn error_write_buffer(&mut self, err: Box<dyn std::error::Error>) -> (u32, u32);
}