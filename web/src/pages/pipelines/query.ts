import gql from "graphql-tag";

export const GetSectorsSdrPipeline = gql`
  query GetSectorsSdrPipeline {
    pipelines {
      id
      spId
      sectorNumber
      createTime
      regSealProof
      status
      failed
      failedAt
      failedReason
      failedReasonMsg
    }
  }`
