//! Small system utility modules for usage by other modules.

extern crate libc;
use libc::c_int;

#[macro_use]
pub mod ioctl;

mod eventfd;
mod signal;
mod struct_util;
mod terminal;

pub use eventfd::*;
pub use ioctl::*;
pub use signal::*;
pub use struct_util::{read_struct, read_struct_slice};
pub use terminal::*;

/// Wrapper to interpret syscall exit codes and provide a rustacean `io::Result`
pub struct SyscallReturnCode(pub c_int);

impl SyscallReturnCode {
    /// Returns the last OS error if value is -1 or Ok(value) otherwise.
    pub fn into_result(self) -> std::io::Result<c_int> {
        if self.0 == -1 {
            Err(std::io::Error::last_os_error())
        } else {
            Ok(self.0)
        }
    }

    /// Returns the last OS error if value is -1 or Ok(()) otherwise.
    pub fn into_empty_result(self) -> std::io::Result<()> {
        self.into_result().map(|_| ())
    }
}
