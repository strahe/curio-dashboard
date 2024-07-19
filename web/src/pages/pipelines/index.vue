<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {GetSectorsSdrPipeline} from "@/pages/pipelines/query";
import {Pipeline} from "@/typed-graph";
import {format} from "timeago.js";
import {definePage} from "unplugin-vue-router/runtime";
import PipelineSummary from "@/widgets/PipelineSummary.vue";
import {ComputedRef} from "vue";
import PipelineTreeMapChart from "@/widgets/PipelineTreeMapChart.vue";

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
  { title: 'ID', key: 'id'},
  { title: 'Miner', key: 'spId' },
  { title: 'Sector', key: 'sectorNumber' },
  { title: 'Created', key: 'createTime' },
  { title: 'Status', key: 'status' },
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
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
            items-per-page="15"
            show-expand
          >
            <template v-slot:[`item.createTime`]="{value}">
              {{format(value)}}
            </template>
            <template v-slot:[`item.status`]="{item}">
              <v-chip :color="item.status === 'Failed' ? 'red' :'green'">{{item.status}}</v-chip>
            </template>
            <template v-slot:expanded-row="{ columns, item }">
              <tr>
                <td :colspan="columns.length">
                  <PipelineTreeMapChart :data="item"></PipelineTreeMapChart>
                </td>
              </tr>
            </template>

          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>


<style scoped lang="sass">

</style>
