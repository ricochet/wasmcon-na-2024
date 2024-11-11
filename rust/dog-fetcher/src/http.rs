use url::Url;
use wasmcloud_component::{
    http::ErrorCode,
    wasi::{
        self,
        http::types::{Fields, Scheme},
    },
};

use crate::internal_error;

/// Helper function to make an outgoing HTTP request and return the incoming body from the response
pub(crate) fn request(url: &str) -> Result<wasi::http::types::IncomingBody, ErrorCode> {
    let req = wasi::http::outgoing_handler::OutgoingRequest::new(Fields::new());
    let url = Url::parse(url).map_err(|_| internal_error("failed to parse URL".to_string()))?;

    if url.scheme() == "https" {
        req.set_scheme(Some(&Scheme::Https))
            .map_err(|_| internal_error("failed to set HTTPS scheme".to_string()))?;
    } else if url.scheme() == "http" {
        req.set_scheme(Some(&Scheme::Http))
            .map_err(|_| internal_error("failed to set HTTP scheme".to_string()))?;
    } else {
        req.set_scheme(Some(&Scheme::Other(url.scheme().to_string())))
            .map_err(|_| internal_error("failed to set custom scheme".to_string()))?;
    }

    req.set_authority(Some(url.authority()))
        .map_err(|_| internal_error("failed to set URL authority".to_string()))?;

    req.set_path_with_query(Some(url.path()))
        .map_err(|_| internal_error("failed to set URL path".to_string()))?;

    match wasi::http::outgoing_handler::handle(req, None) {
        Ok(resp) => {
            resp.subscribe().block();
            let response = resp
                .get()
                .ok_or(())
                .map_err(|_| internal_error("HTTP request response missing".to_string()))?
                .map_err(|_| {
                    internal_error("HTTP request response requested more than once".to_string())
                })?
                .map_err(|_| internal_error("HTTP request failed".to_string()))?;

            response
                .consume()
                .map_err(|_| internal_error("failed to consume response".to_string()))
        }
        Err(e) => Err(e),
    }
}
