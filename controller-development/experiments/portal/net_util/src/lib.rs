#![deny(missing_docs)]
//! # Network-related utilities
//!
//! Provides tools for representing and handling network related concepts like MAC addresses and
//! network interfaces.

extern crate libc;
extern crate serde;

extern crate net_gen;
#[macro_use]
extern crate sys_util;

mod mac;
mod tap;

pub use mac::{MacAddr, MAC_ADDR_LEN};
pub use tap::{Error as TapError, Tap};
