<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {GetSectorsSdrPipeline} from "@/pages/pipelines/query";
import {Pipeline} from "@/typed-graph";
import {format} from "timeago.js";
import {definePage} from "unplugin-vue-router/runtime";
import PipelineSummary from "@/widgets/PipelineSummary.vue";
import PipelineState from "@/widgets/PipelineState.vue";

definePage({
  meta: {
    title: "Pipelines",
  },
})

const { result, loading, refetch, error } = useQuery(GetSectorsSdrPipeline,  null, ()=>({
  fetchPolicy: 'cache-first',
}))
const items: ComputedRef<[Pipeline]> = computed(() => result.value?.pipelines || []);
const headers = [
  { title: 'Miner', key: 'spId' },
  { title: 'Sector', key: 'sectorNumber' },
  { title: 'Created', key: 'createTime' },
  { title: 'State', key: 'state',maxWidth: '600px',sortable: false },
  { title: '  ', key: 'data-table-expand' },
];
</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="12">
        <PipelineSummary></PipelineSummary>
      </v-col>
      <v-col cols="12">
        <Card title="Pipelines" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
            items-per-page="5"
            show-expand
          >
            <template v-slot:[`item.createTime`]="{value}">
              {{format(value)}}
            </template>
            <template v-slot:[`item.state`]="{item}">
              <PipelineState :data="item"></PipelineState>
            </template>
          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>


<style scoped lang="sass">

</style>
