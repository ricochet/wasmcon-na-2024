// PART 2: This can be uncommented
// use std::{
//     io::{self, BufRead, BufReader},
//     sync::LazyLock,
// };

// use crate::bindings::wasi::blobstore::{blobstore, types::IncomingValue};
// use crate::bindings::wasi::io0_2_1::streams::StreamError;

// static PASSWORD_LIST_FILE: LazyLock<String> =
//     LazyLock::new(|| "500-worst-passwords.txt".to_string());
// static PASSWORD_BUCKET_NAME: LazyLock<String> = LazyLock::new(|| "passwords".to_string());

// pub fn get_password_list() -> anyhow::Result<Vec<String>> {
//     let container = blobstore::get_container(&PASSWORD_BUCKET_NAME)
//         .map_err(|e| anyhow::anyhow!("failed to get container: {e}"))?;
//     container
//         .has_object(&PASSWORD_LIST_FILE)
//         .map_err(|_| anyhow::anyhow!("failed to get bad passwords file"))?;
//     let size = container
//         .object_info(&PASSWORD_LIST_FILE)
//         .map_err(|_| anyhow::anyhow!("failed to get bad passwords file"))?
//         .size;
//     let stream = container
//         .get_data(&PASSWORD_LIST_FILE, 0, size)
//         .map_err(|_| anyhow::anyhow!("failed to get bad passwords file"))?;
//     let reader = IncomingValue::incoming_value_consume_async(stream)
//         .map_err(|_| anyhow::anyhow!("failed to get bad passwords file"))?;

//     BufReader::new(reader)
//         .lines()
//         .collect::<Result<Vec<String>, _>>()
//         .map_err(|_| anyhow::anyhow!("failed to get bad passwords file"))
// }

// impl io::Read for crate::bindings::wasi::io0_2_1::streams::InputStream {
//     fn read(&mut self, buf: &mut [u8]) -> io::Result<usize> {
//         let n = buf
//             .len()
//             .try_into()
//             .map_err(|e| io::Error::new(io::ErrorKind::Other, e))?;
//         match self.blocking_read(n) {
//             Ok(chunk) => {
//                 let n = chunk.len();
//                 if n > buf.len() {
//                     return Err(io::Error::new(
//                         io::ErrorKind::Other,
//                         "more bytes read than requested",
//                     ));
//                 }
//                 buf[..n].copy_from_slice(&chunk);
//                 Ok(n)
//             }
//             Err(StreamError::Closed) => Ok(0),
//             Err(StreamError::LastOperationFailed(e)) => {
//                 Err(io::Error::new(io::ErrorKind::Other, e.to_debug_string()))
//             }
//         }
//     }
// }
