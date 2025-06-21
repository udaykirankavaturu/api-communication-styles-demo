# REST API Demo (Node.js + Express)

This demo shows a simple REST API built with Node.js and Express, containerized with Docker.

## Prerequisites

- Docker Engine running

## Build the Docker image

```sh
docker build -t rest-demo .
```

## Run the container

```sh
docker run -d -p 3000:3000 --name rest-demo-test rest-demo
```

## Test the API

Get a greeting:

```sh
curl http://localhost:3000/greet?name=World
```

Expected response:

```json
{ "message": "Hello, World!" }
```

## Stop and remove the container

```sh
docker stop rest-demo-test && docker rm rest-demo-test
```
