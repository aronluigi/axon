import { ApolloClient, InMemoryCache } from '@apollo/client/core'

export default new ApolloClient({
  uri: 'http://localhost:8888/query',
  cache: new InMemoryCache()
})
