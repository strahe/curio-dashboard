import gql from "graphql-tag";

export const GetRunningTasks = gql`
  query GetRunningTasks {
    tasks {
      id
      name
      postedTime
      updateTime
    }
  }`
