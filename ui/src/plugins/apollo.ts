import { ApolloClient, createHttpLink, InMemoryCache, split } from '@apollo/client/core'
import { GraphQLWsLink } from '@apollo/client/link/subscriptions'
import { createClient } from 'graphql-ws'
import { getMainDefinition } from '@apollo/client/utilities'

// HTTP connection to the API
const httpLink = createHttpLink({
  // You should use an absolute URL here
  uri: 'http://localhost:9091/graphql',
})

const wsLink = new GraphQLWsLink(
  createClient({
    url: 'ws://localhost:9091/graphql',
  })
)

const link = split(
  // split based on operation type
  ({ query }) => {
    const definition = getMainDefinition(query)
    return (
      definition.kind === 'OperationDefinition' &&
            definition.operation === 'subscription'
    )
  },
  wsLink,
  httpLink
)

// Cache implementation
const cache = new InMemoryCache({
})

// Create the apollo client
export const apolloClient = new ApolloClient({
  link,
  cache,
})
