import { ApolloServer, gql } from "apollo-server";

const typeDefs = gql`
  type Query {
    greet(name: String!): String!
  }
`;

const resolvers = {
  Query: {
    greet: (_: any, args: { name: string }) => `Hello, ${args.name}!`,
  },
};

const server = new ApolloServer({ typeDefs, resolvers });

server.listen({ port: 4000 }).then(({ url }) => {
  console.log(`🚀 GraphQL server ready at ${url}`);
});
