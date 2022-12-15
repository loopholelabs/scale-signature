use std::io::Error;
use std::io::Cursor;

use scale_signature_http::http_signature::{HttpContext};

pub trait RuntimeContext {
    fn read(&mut self) -> HttpContext;
    fn write(&self) -> Vec<u8>;
    fn error(&self, err: std::io::Error) -> Vec<u8>;
    fn generated(&self) -> &HttpContext;
    fn new(self) -> Self;
}

pub trait GuestContext {
    fn from_read_buffer(self, read_buff: &mut Cursor<&mut Vec<u8>>) -> Result<HttpContext, Error> ;
    fn to_write_buffer(self) -> Result<(u32, u32), Error>;
    fn error_write_buffer(self, error: &str) -> (u32, u32);
    fn next(self) -> Self;
    fn new() -> Self;
}
