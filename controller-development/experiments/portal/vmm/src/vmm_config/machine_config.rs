use serde::{de, Deserialize};
use std::fmt::{Display, Formatter, Result};

/// Portal aims to support small scale workloads only, so limit the maximum
/// vCPUs supported.
pub const MAX_SUPPORTED_VCPUS: u8 = 32;

/// Errors associated with configuring the microVM.
#[derive(Debug, PartialEq)]
pub enum VmConfigError {
    /// The vcpu count is invalid. When hyperthreading is enabled, the `cpu_count` must be either
    /// 1 or an even number.
    InvalidVcpuCount,
    /// The memory size is invalid. The memory can only be an unsigned integer.
    InvalidMemorySize,
    /// Cannot update the configuration of the microvm post boot.
    UpdateNotAllowedPostBoot,
}

impl Display for VmConfigError {
    fn fmt(&self, f: &mut Formatter) -> Result {
        use self::VmConfigError::*;
        match *self {
            InvalidVcpuCount => write!(
                f,
                "The vCPU number is invalid! The vCPU number can only \
                 be 1 or an even number when hyperthreading is enabled.",
            ),
            InvalidMemorySize => write!(f, "The memory size (MiB) is invalid.",),
            UpdateNotAllowedPostBoot => {
                write!(f, "The update operation is not allowed after boot.")
            }
        }
    }
}

/// Strongly typed structure that represents the configuration of the
/// microvm.
#[derive(Clone, Debug, Deserialize, PartialEq, Serialize)]
#[serde(deny_unknown_fields)]
pub struct VmConfig {
    /// Number of vcpu to start.
    #[serde(
        default,
        skip_serializing_if = "Option::is_none",
        deserialize_with = "validate_vcpu_num"
    )]
    pub vcpu_count: Option<u8>,
    /// The memory size in MiB.
    #[serde(skip_serializing_if = "Option::is_none")]
    pub mem_size_mib: Option<usize>,
    /// Enables or disabled hyperthreading.
    #[serde(skip_serializing_if = "Option::is_none")]
    pub ht_enabled: Option<bool>,
    /// A CPU template that it is used to filter the CPU features exposed to the guest.
    #[serde(skip_serializing_if = "Option::is_none")]
    pub cpu_template: Option<CpuFeaturesTemplate>,
}

impl Default for VmConfig {
    fn default() -> Self {
        VmConfig {
            vcpu_count: Some(4),
            mem_size_mib: Some(4192),
            ht_enabled: Some(false),
            cpu_template: None,
        }
    }
}

fn validate_vcpu_num<'de, D>(d: D) -> std::result::Result<Option<u8>, D::Error>
where
    D: de::Deserializer<'de>,
{
    let val = Option::<u8>::deserialize(d)?;
    if let Some(ref value) = val {
        if *value > MAX_SUPPORTED_VCPUS {
            return Err(de::Error::invalid_value(
                de::Unexpected::Unsigned(u64::from(*value)),
                &"number of vCPUs exceeds the maximum limitation",
            ));
        }
    }
    Ok(val)
}

/// Template types available for configuring the CPU features that map
/// to EC2 instances.
#[derive(Clone, Copy, Debug, Deserialize, PartialEq, Serialize)]
pub enum CpuFeaturesTemplate {
    /// C3 Template.
    C3,
    /// T2 Template.
    T2,
}

impl Display for CpuFeaturesTemplate {
    fn fmt(&self, f: &mut Formatter) -> Result {
        match self {
            CpuFeaturesTemplate::C3 => write!(f, "C3"),
            CpuFeaturesTemplate::T2 => write!(f, "T2"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_display_cpu_features_template() {
        assert_eq!(CpuFeaturesTemplate::C3.to_string(), "C3".to_string());
        assert_eq!(CpuFeaturesTemplate::T2.to_string(), "T2".to_string());
    }

    #[test]
    fn test_display_vm_config_error() {
        let expected_str = "The vCPU number is invalid! The vCPU number can only \
                            be 1 or an even number when hyperthreading is enabled.";
        assert_eq!(VmConfigError::InvalidVcpuCount.to_string(), expected_str);

        let expected_str = "The memory size (MiB) is invalid.";
        assert_eq!(VmConfigError::InvalidMemorySize.to_string(), expected_str);

        let expected_str = "The update operation is not allowed after boot.";
        assert_eq!(
            VmConfigError::UpdateNotAllowedPostBoot.to_string(),
            expected_str
        );
    }
}
