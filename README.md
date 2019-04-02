echo-api is a simple console application.

It's a super simple api that saves the last `POST` body and returns that as the response for a `GET` request.
a `DELETE` request removes the state.

The API will respond with a 200 to ANY HttpMethod and supports any uri.
