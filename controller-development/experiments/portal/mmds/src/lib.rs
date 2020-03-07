#[macro_use]
extern crate lazy_static;
extern crate serde_json;

extern crate micro_http;

pub mod data_store;

use serde_json::{Map, Value};
use std::sync::{Arc, Mutex};

use data_store::{Error as MmdsError, Mmds};
use micro_http::{Body, Request, RequestError, Response, StatusCode, Version};

lazy_static! {
    // A static reference to a global Mmds instance. We currently use this for ease of access during
    // prototyping. We'll consider something like passing Arc<Mutex<Mmds>> references to the
    // appropriate threads in the future.
    pub static ref MMDS: Arc<Mutex<Mmds>> = Arc::new(Mutex::new(Mmds::default()));
}

/// Patch provided JSON document (given as `serde_json::Value`) in-place with JSON Merge Patch
/// [RFC 7396](https://tools.ietf.org/html/rfc7396).
pub fn json_patch(target: &mut Value, patch: &Value) {
    if patch.is_object() {
        if !target.is_object() {
            // Replace target with a serde_json object so we can recursively copy patch values.
            *target = Value::Object(Map::new());
        }

        // This is safe since we make sure patch and target are objects beforehand.
        let doc = target.as_object_mut().unwrap();
        for (key, value) in patch.as_object().unwrap() {
            if value.is_null() {
                // If the value in the patch is null we remove the entry.
                doc.remove(key.as_str());
            } else {
                // Recursive call to update target document.
                // If `key` is not in the target document (it's a new field defined in `patch`)
                // insert a null placeholder and pass it as the new target
                // so we can insert new values recursively.
                json_patch(doc.entry(key.as_str()).or_insert(Value::Null), value);
            }
        }
    } else {
        *target = patch.clone();
    }
}

fn build_response(http_version: Version, status_code: StatusCode, body: Body) -> Response {
    let mut response = Response::new(http_version, status_code);
    response.set_body(body);
    response
}

pub fn parse_request(request_bytes: &[u8]) -> Response {
    let request = Request::try_from(request_bytes);
    match request {
        Ok(request) => {
            let uri = request.uri().get_abs_path();
            if uri.is_empty() {
                return build_response(
                    request.http_version(),
                    StatusCode::BadRequest,
                    Body::new("Invalid URI.".to_string()),
                );
            }

            // The lock can be held by one thread only, so it is safe to unwrap.
            // If another thread poisoned the lock, we abort the execution.
            let response = MMDS
                .lock()
                .expect("Failed to build MMDS response due to poisoned lock")
                .get_value(uri.to_string());
            match response {
                Ok(response) => {
                    let response_body = response.join("\n");
                    build_response(
                        request.http_version(),
                        StatusCode::OK,
                        Body::new(response_body),
                    )
                }
                Err(e) => {
                    match e {
                        MmdsError::NotFound => {
                            // NotFound
                            let error_msg = format!("Resource not found: {}.", uri);
                            build_response(
                                request.http_version(),
                                StatusCode::NotFound,
                                Body::new(error_msg),
                            )
                        }
                        MmdsError::UnsupportedValueType => {
                            // InternalServerError
                            let error_msg =
                                format!("The resource {} has an invalid format.", uri.to_string());
                            build_response(
                                request.http_version(),
                                StatusCode::InternalServerError,
                                Body::new(error_msg),
                            )
                        }
                    }
                }
            }
        }
        Err(e) => match e {
            RequestError::InvalidHttpVersion(err_msg) => build_response(
                Version::default(),
                StatusCode::NotImplemented,
                Body::new(err_msg.to_string()),
            ),
            RequestError::InvalidUri(err_msg) | RequestError::InvalidHttpMethod(err_msg) => {
                build_response(
                    Version::default(),
                    StatusCode::BadRequest,
                    Body::new(err_msg.to_string()),
                )
            }
            RequestError::InvalidRequest => build_response(
                Version::default(),
                StatusCode::BadRequest,
                Body::new("Invalid request.".to_string()),
            ),
            RequestError::InvalidHeader => build_response(
                Version::default(),
                StatusCode::BadRequest,
                Body::new("Invalid headers.".to_string()),
            ),
            RequestError::UnsupportedHeader => unreachable!(),
        },
    }
}

#[cfg(test)]
mod tests {
    extern crate serde_json;
    use super::*;

    #[test]
    fn test_parse_request() {
        let data = r#"{
            "name": {
                "first": "John",
                "second": "Doe"
            },
            "age": "43",
            "phones": {
                "home": {
                    "RO": "+40 1234567",
                    "UK": "+44 1234567"
                },
                "mobile": "+44 2345678"
            }
        }"#;
        MMDS.lock()
            .unwrap()
            .put_data(serde_json::from_str(data).unwrap())
            .unwrap();

        // Test invalid request.
        let request = b"HTTP/1.1";
        let dummy_response = Response::new(Version::Http11, StatusCode::BadRequest);
        assert!(parse_request(request).status() == dummy_response.status());

        // Test unsupported HTTP version.
        let request = b"GET http://169.254.169.255/ HTTP/2.0\r\n\r\n";
        let mut expected_response = Response::new(Version::Http11, StatusCode::NotImplemented);
        expected_response.set_body(Body::new("Unsupported HTTP version.".to_string()));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        // Test invalid HTTP Method.
        let request = b"POST http://169.254.169.255/ HTTP/1.0\r\n\r\n";
        let mut expected_response = Response::new(Version::Http11, StatusCode::BadRequest);
        expected_response.set_body(Body::new("Unsupported HTTP method.".to_string()));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        // Test invalid (empty absolute path) URI.
        let request = b"GET http:// HTTP/1.0\r\n\r\n";
        let mut expected_response = Response::new(Version::Http10, StatusCode::BadRequest);
        expected_response.set_body(Body::new("Invalid URI.".to_string()));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        // Test invalid HTTP format.
        let request = b"GET / HTTP/1.1\r\n";
        let mut expected_response = Response::new(Version::Http11, StatusCode::BadRequest);
        expected_response.set_body(Body::new("Invalid request.".to_string()));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        // Test resource not found.
        let request = b"GET http://169.254.169.254/invalid HTTP/1.0\r\n\r\n";
        let mut expected_response = Response::new(Version::Http10, StatusCode::NotFound);
        expected_response.set_body(Body::new("Resource not found: /invalid.".to_string()));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        // Test Ok path.
        let request = b"GET http://169.254.169.254/ HTTP/1.0\r\n\r\n";
        let mut expected_response = Response::new(Version::Http10, StatusCode::OK);
        let body = "age\nname/\nphones/".to_string();
        expected_response.set_body(Body::new(body));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        let request = b"GET /age HTTP/1.1\r\n\r\n";
        let mut expected_response = Response::new(Version::Http11, StatusCode::OK);
        let body = "43".to_string();
        expected_response.set_body(Body::new(body));
        let actual_response = parse_request(request);

        assert!(expected_response.status() == actual_response.status());
        assert!(expected_response.body().unwrap() == actual_response.body().unwrap());
        assert!(expected_response.http_version() == actual_response.http_version());

        let data = r#"{
            "name": {
                "first": "John",
                "second": "Doe"
            },
            "age": 43
        }"#;
        assert_eq!(
            MMDS.lock()
                .unwrap()
                .put_data(serde_json::from_str(data).unwrap()),
            Err(MmdsError::UnsupportedValueType)
        );
    }

    #[test]
    fn test_json_patch() {
        let mut data = serde_json::json!({
            "name": {
                "first": "John",
                "second": "Doe"
            },
            "age": "43",
            "phones": {
                "home": {
                    "RO": "+40 1234567",
                    "UK": "+44 1234567"
                },
                "mobile": "+44 2345678"
            }
        });

        let patch = serde_json::json!({
            "name": {
                "second": null,
                "last": "Kennedy"
            },
            "age": "44",
            "phones": {
                "home": "+44 1234567",
                "mobile": {
                    "RO": "+40 2345678",
                    "UK": "+44 2345678"
                }
            }
        });
        json_patch(&mut data, &patch);

        // Test value replacement in target document.
        assert_eq!(data["age"], patch["age"]);

        // Test null value removal from target document.
        assert_eq!(data["name"]["second"], Value::Null);

        // Test add value to target document.
        assert_eq!(data["name"]["last"], patch["name"]["last"]);
        assert!(!data["phones"]["home"].is_object());
        assert_eq!(data["phones"]["home"], patch["phones"]["home"]);
        assert!(data["phones"]["mobile"].is_object());
        assert_eq!(
            data["phones"]["mobile"]["RO"],
            patch["phones"]["mobile"]["RO"]
        );
        assert_eq!(
            data["phones"]["mobile"]["UK"],
            patch["phones"]["mobile"]["UK"]
        );
    }
}
