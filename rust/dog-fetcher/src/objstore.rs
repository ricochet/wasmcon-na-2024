//! Object storage related functionality and helper methods

use wasmcloud_component::trace;
use wasmcloud_component::wasi::blobstore::blobstore::{
    container_exists, create_container, get_container,
};
use wasmcloud_component::wasi::blobstore::container::Container;
use wasmcloud_component::wasi::blobstore::types::OutgoingValue;
use wasmcloud_component::wasi::http::outgoing_handler::ErrorCode;
use wasmcloud_component::wasi::http::types::IncomingBody;
use wasmcloud_component::wasi::io::streams::StreamError;

use crate::internal_error;

/// A helper that will automatically create a container if it doesn't exist and returns an owned copy of the name for immediate use
pub(crate) fn ensure_container(name: &String) -> Result<Container, ErrorCode> {
    if container_exists(name).map_err(internal_error)? {
        return get_container(name).map_err(internal_error);
    }
    create_container(name).map_err(internal_error)
}

/// Stream a binary blob from an HTTP [`IncomingBody`] to object storage
pub(crate) fn write_object(
    object_body: IncomingBody,
    bucket: &String,
    key: &String,
) -> Result<(), ErrorCode> {
    let container = ensure_container(&String::from(bucket))?;

    let data = OutgoingValue::new_outgoing_value();
    let data_body = data
        .outgoing_value_write_body()
        .map_err(|()| internal_error("failed to get data output stream"))?;
    container.write_data(key, &data).map_err(internal_error)?;
    let object_stream = object_body
        .stream()
        .map_err(|()| internal_error("failed to stream object"))?;
    loop {
        match data_body.blocking_splice(&object_stream, u64::MAX) {
            Ok(0) | Err(StreamError::Closed) => break,
            Ok(n) => {
                trace!("wrote {n} bytes to object storage");
                continue;
            }
            Err(e) => return Err(internal_error(e)),
        }
    }
    drop(data_body);
    drop(object_stream);
    OutgoingValue::finish(data).map_err(internal_error)?;
    let _trailers = IncomingBody::finish(object_body);

    Ok(())
}
