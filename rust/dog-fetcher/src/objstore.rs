//! Object storage related functionality and helper methods

use anyhow::{anyhow, bail, Result};
use bytes::Bytes;

use crate::bindings::wasi::blobstore::blobstore::{container_exists, create_container, get_container};
use crate::bindings::wasi::blobstore::container::Container;

use crate::bindings::wasi::blobstore::types::{IncomingValue, OutgoingValue};

/// Maximum bytes to read at a time from the incoming request body
/// this value is chosen somewhat arbitrarily, and is not a limit for bytes read,
/// but is instead the amount of bytes to be read *at once*
const MAX_READ_BYTES: u32 = 2048;

/// Maximum bytes to write at a time, due to the limitations on wasi-io's blocking_write_and_flush()
const MAX_WRITE_BYTES: usize = 4096;

/// A helper that will automatically create a container if it doesn't exist and returns an owned copy of the name for immediate use
pub(crate) fn ensure_container(name: &String) -> Result<Container> {
    if container_exists(name)
        .map_err(|e| anyhow!("error checking for container: {e}"))?
    {
        return get_container(name).map_err(|e| anyhow!("failed to get container: {e}"));
    }
    create_container(name).map_err(|e| anyhow!("failed to create container: {e}"))
}

/// Write a binary blob to object storage
pub(crate) fn write_object(object_bytes: Bytes, bucket: &str, key: &String) -> Result<()> {
    let container = ensure_container(&String::from(bucket))?;

    let data = OutgoingValue::new_outgoing_value();
    let data_body = data
        .outgoing_value_write_body()
        .map_err(|()| anyhow!("failed to get data output stream"))?;
    if let Err(e) = container.write_data(key, &data) {
        bail!("failed to write data: {e}");
    };
    for chunk in object_bytes.chunks(MAX_WRITE_BYTES) {
        data_body
            .blocking_write_and_flush(chunk)
            .map_err(|e| anyhow!("failed to write chunk: {e}"))?;
    }
    drop(data_body);
    OutgoingValue::finish(data).map_err(|e| anyhow!("failed to write data: {e}"))?;

    Ok(())
}

/// Read a binary blob from object storage
pub(crate) fn read_object(bucket: &str, key: &str) -> Result<Bytes> {
    let key = &String::from(key);
    let container = ensure_container(&String::from(bucket))?;
    let metadata = container
        .object_info(key)
        .map_err(|e| anyhow!("failed to get object metadata: {e}"))?;
    let incoming = container
        .get_data(key, 0, metadata.size)
        .map_err(|e| anyhow!("failed to get data: {e}"))?;
    let body = IncomingValue::incoming_value_consume_sync(incoming)
        .map_err(|e| anyhow!("failed to consume incoming value: {e}"))?;

    Ok(Bytes::from(body))
}
