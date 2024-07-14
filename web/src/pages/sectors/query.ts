import gql from 'graphql-tag'


export const SectorFragment = {
  meta: gql`
    fragment SectorMetaFragment on SectorMeta {
      id
      spId
      sectorNum
      regSealProof
      ticketEpoch
      ticketValue
      origSealedCid
      origUnsealedCid
      curSealedCid
      curUnsealedCid
      msgCidPrecommit
      msgCidCommit
      msgCidUpdate
      seedEpoch
      seedValue
    }
  `
}

export const GetSectorsMeta = gql`
  query GetSectorsMeta($actor: ActorID, $sectorNumber: Int, $offset: Int!, $limit: Int!) {
    sectors(actor: $actor, sectorNumber: $sectorNumber, offset: $offset, limit: $limit) {
      meta {
        ...SectorMetaFragment
      }
    }
    sectorsCount(actor: $actor)
    actors {
      address
    }
  }
  ${SectorFragment.meta}
`

export const GetSectorMeta = gql`
  query GetSectorsMeta($actor: ActorID, $sectorNumber: Int, $offset: Int!, $limit: Int!) {
    sectors(actor: $actor, sectorNumber: $sectorNumber, offset: $offset, limit: $limit) {
      meta {
        ...SectorMetaFragment
      }
    }
  }
  ${SectorFragment.meta}
`
