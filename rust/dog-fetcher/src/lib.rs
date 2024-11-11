use serde_json::json;
use wasmcloud_component::http::{ErrorCode, IncomingRequest, OutgoingBody, Response, Server};

mod http;
mod objstore;

#[derive(serde::Deserialize)]
struct DogResponse {
    message: String,
}

const CONTAINER_NAME: &str = "doggies";

struct DogFetcher;

impl Server for DogFetcher {
    fn handle(_request: IncomingRequest) -> Result<Response<impl OutgoingBody>, ErrorCode> {
        // Get dog picture URL
        let dog_picture_url = http::request("https://dog.ceo/api/breeds/image/random")?;
        let dog_response: DogResponse = serde_json::from_reader(
            dog_picture_url
                .stream()
                .map_err(|_| internal_error("failed to stream dog API response"))?,
        )
        .map_err(|_| internal_error("failed to deserialize dog API response"))?;

        // remove for PART 2:
        Ok(Response::new(
            json!({"dog_file": &dog_response.message}).to_string(),
        ))

        // PART 2: uncomment
        // // Form the object key using the last 3 segments of the URL
        // let dog_image_name = dog_response
        //     .message
        //     .split('/')
        //     .skip(4)
        //     .collect::<Vec<_>>()
        //     .join("-");
        // let dog_picture = http::request(&dog_response.message)?;
        // objstore::write_object(dog_picture, &CONTAINER_NAME.to_string(), &dog_image_name)?;

        // Ok(Response::new(
        //     json!({"dog_file": format!("/tmp/{CONTAINER_NAME}/{dog_image_name}")}).to_string(),
        // ))
    }
}

/// Helper function to return an internal error
pub(crate) fn internal_error(err: impl ToString) -> ErrorCode {
    ErrorCode::InternalError(Some(err.to_string()))
}

wasmcloud_component::http::export!(DogFetcher);
