import gql from "graphql-tag";

export const GetPendingDeals = gql`
  query GetPendingDeals{
    dealsPending {
      spID
      sectorNumber
      pieceIndex
      pieceCID
      pieceSize
      dataURL
      dataHeaders
      dataRawSize
      dataDeleteOnFinalize
      f05PublishCID
      f05DealID
      f05DealProposal
      f05DealStartEpoch
      f05DealEndEpoch
      directStartEpoch
      directEndEpoch
      directPieceActivationManifest
      createdAt
      isSnap
    }
  }
`
