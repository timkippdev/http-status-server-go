# HTTP Status Code Server

This server is meant to be used to test various response scenarios.

## Usage

Request a desired status code response by adding the status code at the end of the url.  

http://some-domain-here.com/xxx <-- replace `xxx` with your desired status code

If you desire a JSON response, you can add the `Accept` header, with the value of `application/json`, to your request.

### Example Request/Response - Normal

```text
curl -i localhost:8080/404
HTTP/1.1 404 Not Found
Content-Type: text/plain
Date: Tue, 03 Dec 2019 01:30:24 GMT
Content-Length: 13

404 Not Found
```

### Example Request/Response - JSON

```text
curl -i -H "Accept:application/json" localhost:8080/404
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Tue, 03 Dec 2019 01:33:52 GMT
Content-Length: 48

{"description":"404 Not Found","statusCode":404}
```

## Running Locally

### Prerequisites
1) Golang (Go) Installed - https://golang.org/doc/install
2) Docker Installed - https://docs.docker.com/v17.09/engine/installation/

### Start Server
Inside your project directory, run the following commands:
```bash
# starting server
docker build -t http-status-server:0.0.1 .
docker run -p 8080:8080 -d --rm --name http-status-server http-status-server:0.0.1

# stopping server
docker container stop http-status-server
```
Your server should now be available at http://localhost:8080