<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";

defineProps({
  title: {
    type: String,
    default: undefined,
  },
  color: {
    type: String,
    default: undefined,
  }
})

const { result, loading, refetch, error } = useQuery(gql`
    query GetPipelineSummary {
        pipelinesSummary {
        id
        sdr
        trees
        precommitMsg
        waitSeed
        porep
        commitMsg
        done
        failed
        }
  }
`,  null, ()=>({
  fetchPolicy: 'cache-first',
}))

const items = computed(() => result.value?.pipelinesSummary || []);
const headers = [
  { title: 'Miner', key: 'id' },
  { title: 'SDR', key: 'sdr' },
  { title: 'Trees', key: 'trees' },
  { title: 'Precommit Msg', key: 'precommitMsg' },
  { title: 'Wait Seed', key: 'waitSeed' },
  { title: 'PoRep', key: 'porep' },
  { title: 'Commit Msg', key: 'commitMsg' },
  { title: 'Done', key: 'done' },
  { title: 'Failed', key: 'failed' },
]
</script>

<template>
  <Card :title="title" :color="color" :error="error as Error">
    <template #titleAction>
      <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
    </template>
    <v-data-table-virtual
      :loading="loading"
      :headers="headers"
      :items="items"
    ></v-data-table-virtual>
  </Card>
</template>

<style scoped lang="sass">

</style>
