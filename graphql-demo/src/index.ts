import { ApolloServer, gql } from "apollo-server";

const typeDefs = gql`
  type GreetResponse {
    message: String!
    weather: String!
    time: String!
  }

  type Query {
    greet(name: String!): GreetResponse!
  }
`;

const resolvers = {
  Query: {
    greet: (_: any, args: { name: string }) => {
      const message = `Hello, ${args.name}!`;
      const weather = "Sunny"; // Placeholder, can be dynamic
      const time = new Date().toLocaleString();
      return { message, weather, time };
    },
  },
};

const server = new ApolloServer({ typeDefs, resolvers });

server.listen({ port: 4000 }).then(({ url }) => {
  console.log(`🚀 GraphQL server ready at ${url}`);
});
