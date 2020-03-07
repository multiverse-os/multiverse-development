#![deny(missing_docs)]
//! Utility for configuring the CPUID (CPU identification) for the guest microVM.

extern crate kvm_bindings;
extern crate kvm_ioctls;

use kvm_ioctls::CpuId;

mod common;
use common::*;

/// Contains helper methods for bit operations.
pub mod bit_helper;

mod template;
pub use template::c3;
pub use template::t2;

mod cpu_leaf;

mod transformer;
use transformer::*;
pub use transformer::{Error, VmSpec};

mod brand_string;

/// Sets up the CPUID entries for the given vcpu.
///
/// # Arguments
///
/// * `kvm_cpuid` - KVM related structure holding the relevant CPUID info.
/// * `vm_spec` - The specifications of the VM.
///
/// # Example
/// ```
/// extern crate cpuid;
/// extern crate kvm_ioctls;
///
/// use cpuid::{filter_cpuid, VmSpec};
/// use kvm_ioctls::{CpuId, Kvm, MAX_KVM_CPUID_ENTRIES};
///
/// let kvm = Kvm::new().unwrap();
/// let mut kvm_cpuid: CpuId = kvm.get_supported_cpuid(MAX_KVM_CPUID_ENTRIES).unwrap();
///
/// let vm_spec = VmSpec::new(0, 1, true).unwrap();
///
/// filter_cpuid(&mut kvm_cpuid, &vm_spec).unwrap();
///
/// // Get expected `kvm_cpuid` entries.
/// let entries = kvm_cpuid.mut_entries_slice();
/// ```
#[cfg(any(target_arch = "x86", target_arch = "x86_64"))]
pub fn filter_cpuid(kvm_cpuid: &mut CpuId, vm_spec: &VmSpec) -> Result<(), Error> {
    let maybe_cpuid_transformer: Option<&dyn CpuidTransformer> = match vm_spec.cpu_vendor_id() {
        VENDOR_ID_INTEL => Some(&intel::IntelCpuidTransformer {}),
        VENDOR_ID_AMD => Some(&amd::AmdCpuidTransformer {}),
        _ => None,
    };

    if let Some(cpuid_transformer) = maybe_cpuid_transformer {
        cpuid_transformer.process_cpuid(kvm_cpuid, &vm_spec)?;
    }

    Ok(())
}
