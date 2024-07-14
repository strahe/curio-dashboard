import gql from "graphql-tag";

export const GetConfigs = gql`
    query GetConfigs {
        configs {
          id
          title
          config
          usedBy {
            machineId
            machineName
          }
        }
    }
`
