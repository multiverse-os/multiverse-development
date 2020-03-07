// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

use std::result;

use futures::sync::oneshot;
use hyper::Method;
use serde_json::Value;

use request::{IntoParsedRequest, ParsedRequest};
use vmm::VmmAction;

// The names of the members from this enum must precisely correspond (as a string) to the possible
// values of "action_type" from the json request body. This is useful to get a strongly typed
// struct from the Serde deserialization process.
#[derive(Clone, Debug, Deserialize, PartialEq, Serialize)]
enum ActionType {
    BlockDeviceRescan,
    FlushMetrics,
    InstanceStart,
    SendCtrlAltDel,
}

// The model of the json body from a sync request. We use Serde to transform each associated
// json body into this.
#[derive(Clone, Debug, Deserialize, PartialEq, Serialize)]
#[serde(deny_unknown_fields)]
pub struct ActionBody {
    action_type: ActionType,
    #[serde(skip_serializing_if = "Option::is_none")]
    payload: Option<Value>,
}

fn validate_payload(action_body: &ActionBody) -> Result<(), String> {
    match action_body.action_type {
        ActionType::BlockDeviceRescan => {
            match action_body.payload {
                Some(ref payload) => {
                    // Expecting to have drive_id as a String in the payload.
                    if !payload.is_string() {
                        return Err(
                            "Invalid payload type. Expected a string representing the drive_id"
                                .to_string(),
                        );
                    }
                    Ok(())
                }
                None => Err("Payload is required for block device rescan.".to_string()),
            }
        }
        ActionType::FlushMetrics | ActionType::InstanceStart | ActionType::SendCtrlAltDel => {
            // Neither FlushMetrics nor InstanceStart should have a payload.
            if action_body.payload.is_some() {
                return Err(format!(
                    "{:?} does not support a payload.",
                    action_body.action_type
                ));
            }
            Ok(())
        }
    }
}

impl IntoParsedRequest for ActionBody {
    fn into_parsed_request(
        self,
        _: Option<String>,
        _: Method,
    ) -> result::Result<ParsedRequest, String> {
        validate_payload(&self)?;
        match self.action_type {
            ActionType::BlockDeviceRescan => {
                // Safe to unwrap because we validated the payload in the validate_payload func.
                let block_device_id = self.payload.unwrap().as_str().unwrap().to_string();
                let (sync_sender, sync_receiver) = oneshot::channel();
                Ok(ParsedRequest::Sync(
                    Box::new(VmmAction::RescanBlockDevice(block_device_id, sync_sender)),
                    sync_receiver,
                ))
            }
            ActionType::FlushMetrics => {
                let (sync_sender, sync_receiver) = oneshot::channel();
                Ok(ParsedRequest::Sync(
                    Box::new(VmmAction::FlushMetrics(sync_sender)),
                    sync_receiver,
                ))
            }
            ActionType::InstanceStart => {
                let (sync_sender, sync_receiver) = oneshot::channel();
                Ok(ParsedRequest::Sync(
                    Box::new(VmmAction::StartMicroVm(sync_sender)),
                    sync_receiver,
                ))
            }
            ActionType::SendCtrlAltDel => {
                let (sync_sender, sync_receiver) = oneshot::channel();
                Ok(ParsedRequest::Sync(
                    Box::new(VmmAction::SendCtrlAltDel(sync_sender)),
                    sync_receiver,
                ))
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json;

    #[test]
    fn test_validate_payload() {
        // Test InstanceStart.
        let action_body = ActionBody {
            action_type: ActionType::InstanceStart,
            payload: None,
        };
        assert!(validate_payload(&action_body).is_ok());
        // Error case: InstanceStart with payload.
        let action_body = ActionBody {
            action_type: ActionType::InstanceStart,
            payload: Some(Value::String("dummy-payload".to_string())),
        };
        assert!(validate_payload(&action_body).is_err());

        // Test BlockDeviceRescan
        let action_body = ActionBody {
            action_type: ActionType::BlockDeviceRescan,
            payload: Some(Value::String(String::from("dummy_id"))),
        };
        assert!(validate_payload(&action_body).is_ok());
        // Error case: no payload.
        let action_body = ActionBody {
            action_type: ActionType::BlockDeviceRescan,
            payload: None,
        };
        assert!(validate_payload(&action_body).is_err());
        // Error case: payload is not String.
        let action_body = ActionBody {
            action_type: ActionType::BlockDeviceRescan,
            payload: Some(Value::Bool(false)),
        };
        assert!(validate_payload(&action_body).is_err());

        // Test FlushMetrics.
        let action_body = ActionBody {
            action_type: ActionType::FlushMetrics,
            payload: None,
        };

        assert!(validate_payload(&action_body).is_ok());
        // Error case: FlushMetrics with payload.
        let action_body = ActionBody {
            action_type: ActionType::FlushMetrics,
            payload: Some(Value::String("metrics-payload".to_string())),
        };
        let res = validate_payload(&action_body);
        assert!(res.is_err());
        assert_eq!(res.unwrap_err(), "FlushMetrics does not support a payload.");

        // Test SendCtrlAltDel.
        let action_body = ActionBody {
            action_type: ActionType::SendCtrlAltDel,
            payload: None,
        };
        assert!(validate_payload(&action_body).is_ok());
        // Error case: SendCtrlAltDel with payload.
        let action_body = ActionBody {
            action_type: ActionType::SendCtrlAltDel,
            payload: Some(Value::String("dummy-payload".to_string())),
        };
        assert!(validate_payload(&action_body).is_err());
    }

    #[test]
    fn test_into_parsed_request() {
        {
            let json = r#"{
                "action_type": "BlockDeviceRescan",
                "payload": "dummy_id"
              }"#;
            let (sender, receiver) = oneshot::channel();
            let req = ParsedRequest::Sync(
                Box::new(VmmAction::RescanBlockDevice("dummy_id".to_string(), sender)),
                receiver,
            );

            let result: Result<ActionBody, serde_json::Error> = serde_json::from_str(json);
            assert!(result.is_ok());
            assert!(result
                .unwrap()
                .into_parsed_request(None, Method::Put)
                .unwrap()
                .eq(&req));
        }

        {
            let json = r#"{
                "action_type": "InstanceStart"
            }"#;

            let (sender, receiver) = oneshot::channel();
            let req: ParsedRequest =
                ParsedRequest::Sync(Box::new(VmmAction::StartMicroVm(sender)), receiver);
            let result: Result<ActionBody, serde_json::Error> = serde_json::from_str(json);
            assert!(result.is_ok());
            assert!(result
                .unwrap()
                .into_parsed_request(None, Method::Put)
                .unwrap()
                .eq(&req));
        }

        {
            let json = r#"{
                "action_type": "SendCtrlAltDel"
            }"#;

            let (sender, receiver) = oneshot::channel();
            let req: ParsedRequest =
                ParsedRequest::Sync(Box::new(VmmAction::SendCtrlAltDel(sender)), receiver);
            let result: Result<ActionBody, serde_json::Error> = serde_json::from_str(json);
            assert!(result.is_ok());
            assert!(result
                .unwrap()
                .into_parsed_request(None, Method::Put)
                .unwrap()
                .eq(&req));
        }

        {
            let json = r#"{
                "action_type": "FlushMetrics"
            }"#;

            let (sender, receiver) = oneshot::channel();
            let req: ParsedRequest =
                ParsedRequest::Sync(Box::new(VmmAction::FlushMetrics(sender)), receiver);
            let result: Result<ActionBody, serde_json::Error> = serde_json::from_str(json);
            assert!(result.is_ok());
            assert!(result
                .unwrap()
                .into_parsed_request(None, Method::Put)
                .unwrap()
                .eq(&req));

            let json = r#"{
                "action_type": "FlushMetrics",
                "payload": "metrics-payload"
            }"#;

            let result: Result<ActionBody, serde_json::Error> = serde_json::from_str(json);
            assert!(result.is_ok());
            let res = result.unwrap().into_parsed_request(None, Method::Put);
            assert!(res.is_err());
            assert!(res == Err("FlushMetrics does not support a payload.".to_string()));
        }
    }
}
