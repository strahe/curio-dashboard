<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {Machine, MachineSummary} from "@/typed-graph";
import {GetMachines} from "@/pages/machines/query";
import {formatBytes} from "@/utils/formatBytes";
import {format} from "timeago.js";
import {definePage} from "unplugin-vue-router/runtime";

definePage({
  meta: {
    title: 'Machines',
  },
})

const { result, loading, refetch, error } = useQuery(GetMachines, null, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[Machine]> = computed(() => result.value?.machines || []);
const stats: ComputedRef<MachineSummary> = computed(() => result.value?.machineSummary || {});

const headers = ref([
  { title: 'ID', key: 'id' },
  { title: 'Name', key: 'detail.machineName' },
  { title: 'Last Contact', key: 'lastContact' },
  { title: 'Startup', key: 'detail.startupTime' },
  { title: 'CPU', key: 'cpu' },
  { title: 'GPU', key: 'gpu' },
  { title: 'RAM', key: 'ram' },
  { title: 'Host and Port', key: 'hostAndPort' },
  { title: '  ', key: 'data-table-expand' },
]);
//
// let expanded = ref([]);

</script>
<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="TOTAL CPU"
          :value="String(stats.totalCpu || 0)"
          icon="mdi-cpu-64-bit"
          color="secondary"
        ></StatsCard>
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="TOTAL GPU"
          :value="String(stats.totalGpu || 0)"
          icon="mdi-expansion-card"
          color="primary"
        ></StatsCard>
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="TOTAL RAM"
          :value="formatBytes(stats.totalRam || 0).combined"
          icon="mdi-speedometer"
          color="warning"
        ></StatsCard>
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="UP / DOWN"
          :value="`${stats.totalUp || 0} / ${stats.totalDown || 0}`"
          icon="mdi-check-circle-outline"
          color="green"
        ></StatsCard>

      </v-col>
      <v-col cols="12">
        <Card title="Machines" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
            show-expand
          >
            <template v-slot:[`item.detail.machineName`]="{item}">
              {{ item.detail?.machineName? item.detail.machineName : item.hostAndPort }}
            </template>
            <template v-slot:[`item.ram`]="{value}">
              {{ formatBytes(value).combined }}
            </template>
            <template v-slot:[`item.lastContact`]="{value}">
              <div :title="value">{{format(value)}}</div>
            </template>
            <template v-slot:[`item.detail.startupTime`]="{value}">
              <div :title="value">{{format(value)}}</div>
            </template>
          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped lang="sass">

</style>
