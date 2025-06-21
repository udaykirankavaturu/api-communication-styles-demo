# GraphQL API Demo (TypeScript + Apollo Server)

This demo shows a simple GraphQL API built with TypeScript and Apollo Server, containerized with Docker.

## Prerequisites

- Docker Engine running

## Build the Docker image

```sh
docker build -t graphql-demo .
```

## Run the container

```sh
docker run -d -p 4000:4000 --name graphql-demo-test graphql-demo
```

## Test the API

Open the GraphQL Playground in your browser:

http://localhost:4000/

Try the following query:

```
query {
  greet(name: "World")
}
```

Expected response:

```
{
  "data": {
    "greet": "Hello, World!"
  }
}
```

## Stop and remove the container

```sh
docker stop graphql-demo-test && docker rm graphql-demo-test
```
