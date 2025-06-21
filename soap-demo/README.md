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

## Files

- `Dockerfile`: Container setup
- `app.py`: SOAP service implementation
