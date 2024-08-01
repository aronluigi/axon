
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "./../pkg/gql/schema.graphql",
  documents: "src/**/*.graphql",
  generates: {
    "src/generated/graphql.ts": {
      plugins: [
        "typescript",
        "typescript-operations",
        "graphql-codegen-svelte-apollo"
      ],
      config: {
        clientPath: '../apollo.client',
        asyncQuery: true
      }
    },
  },
};

export default config;
