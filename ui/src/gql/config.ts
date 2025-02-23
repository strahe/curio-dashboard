import gql from 'graphql-tag'

export const GetConfigs = gql`
  query GetConfigs {
    configs {
      id
      title
      config
      usedBy {
        id
        machineId
        machineName
        minersArray
        miners
        startupTime
      }
    }
  }
`

export const GetConfig = gql`
    query GetConfig($layer: String!) {
      config(layer: $layer) {
        id
        title
        config
        usedBy {
          id
          machineId
          machineName
          minersArray
          miners
          startupTime
        }
      }
    }`

export const UpdateConfig = gql`
  mutation UpdateConfig($title: String!, $config: String!) {
    updateConfig(title: $title, config: $config) {
      id
      title
      config
    }
  }
`

export const CreateConfig = gql`
  mutation CreateConfig($title: String!, $config: String!) {
    createConfig(title: $title, config: $config) {
      id
      title
      config
    }
  }
`

export const RemoveConfig = gql`
  mutation RemoveConfig($title: String!) {
    removeConfig(title: $title) {
      id
      title
      config
    }
  }
`
