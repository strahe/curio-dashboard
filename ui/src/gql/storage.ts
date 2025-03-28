import gql from 'graphql-tag'

export const storagePathFragment = gql`
  fragment StoragePathAll on StoragePath {
    id
    storageId
    urls
    weight
    maxStorage
    canSeal
    canStore
    type
    groups
    allowTo
    allowTypes
    denyTypes
    capacity
    available
    fsAvailable
    reserved
    used
    lastHeartbeat
    heartbeatErr
    allowMiners
    denyMiners
  }`

export const storageLivenessFragment = gql`
  fragment StorageLivenessAll on StorageLiveness {
    storageId
    url
    lastChecked
    lastLive
    lastDead
    lastDeadReason
  }`

export const GetStoragePaths = gql`
  query GetStoragePaths {
    storagePaths {
      ...StoragePathAll
    }
  }
  ${storagePathFragment}
`

export const GetStorageStats = gql`
  query GetStorageStats {
    storageStats {
      type
      totalAvailable
      totalCapacity
      totalFsAvailable
      totalReserved
      totalUsed
    }
  }`

export const GetStorage = gql`
  query GetStorage($id: String!) {
    storage(id: $id) {
      id
      path {
        ...StoragePathAll
      }
      liveness {
        ...StorageLivenessAll
      }
    }
  }
  ${storagePathFragment}
  ${storageLivenessFragment}
`

export const GetMachineStorages = gql`
    query GetMachineStorages($id: Int!) {
      machine(id: $id) {
        id
        storages {
          ...StoragePathAll
        }
      }
    }
    ${storagePathFragment}
`
