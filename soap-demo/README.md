# SOAP Demo

This demo shows a simple SOAP web service using Python and the `spyne` library.

## Prerequisites

- Docker (Engine running)

## How to Run

1. Build the Docker image:
   ```sh
   docker build -t soap-demo .
   ```
2. Run the container:
   ```sh
   docker run -p 8000:8000 soap-demo
   ```
3. The SOAP service will be available at `http://localhost:8000/`.
4. You can test the service using a SOAP client (e.g., SoapUI) or by sending a SOAP request to the endpoint.
5. Example CURL to run in postman

```sh
curl --location 'http://localhost:8000/' \
--header 'Content-Type: text/xml' \
--data-raw '<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:spy="spyne.examples.hello.soap">
   <soapenv:Header/>
   <soapenv:Body>
      <spy:say_hello>
         <spy:name>YourName</spy:name>
         <spy:times>2</spy:times>
      </spy:say_hello>
   </soapenv:Body>
</soapenv:Envelope>'
```

## Files

- `Dockerfile`: Container setup
- `app.py`: SOAP service implementation
