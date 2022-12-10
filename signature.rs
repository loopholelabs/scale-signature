TODO // document like go Interfaces, rm/make private all but read/write/error
pub trait RuntimeContext {
    fn new(self) -> Self;
    fn read(&mut self) -> HttpContext;
    fn write(&self) -> Vec<u8>;
    fn error(&self, err: std::io::Error) -> Vec<u8>;
    fn generated(&self) -> &HttpContext;
}

pub trait GuestContext {
    fn from_read_buffer(self, read_buff: &mut Cursor<&mut Vec<u8>>) -> Result<HttpContext, Error> ;
    fn to_write_buffer(self) -> Result<(u32, u32), Error>;
    fn error_write_buffer(self, error: &str) -> (u32, u32);

TODO // document like go Interfaces, rm/make private all below
    fn new() -> Self;
    fn next(self) -> Self;
    fn request(&mut self) -> &mut HttpRequest;
    fn response(&mut self) -> &mut HttpResponse;
}
