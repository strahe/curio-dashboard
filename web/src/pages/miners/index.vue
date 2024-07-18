<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import {GetActors} from "@/pages/miners/query";
import {Actor} from "@/typed-graph";
import {ComputedRef} from "vue";
import {formatBytes} from "@/utils/formatBytes";
import {formatFIL} from "@/utils/formatFIL";
import {definePage} from "unplugin-vue-router/runtime";

definePage({
  meta: {
    title: 'Miners',
  },
})


const { result, loading, refetch, error } = useQuery(GetActors, null, () => ({
  fetchPolicy: 'cache-first',
}));
const items: ComputedRef<[Actor]> = computed(() => result.value?.actors || []);
const headers = ref([
  { title: 'Address', key: 'address' },
  {title: 'Layers', key: 'layers'},
  { title: 'Balance', key: 'actorBalance'},
  { title: 'Available Balance', key: 'actorAvailableBalance' },
  { title: 'Worker Balance', key: 'workerBalance' },
  { title: 'QA Power', key: 'qualityAdjustedPower' },
  { title: 'RAW Power', key: 'rawBytePower' },
  { title: '  ', key: 'data-table-expand' },
]);
const totalQAPower = computed(() => items.value.reduce((acc, item) => acc + item.qualityAdjustedPower, 0));
const totalRawPower = computed(() => items.value.reduce((acc, item) => acc + item.rawBytePower, 0));
const totalBalance = computed(() => items.value.reduce((acc, item) => acc + item.actorBalance, 0));
const totalAvailableBalance = computed(() => items.value.reduce((acc, item) => acc + item.actorAvailableBalance, 0));
const totalWorkerBalance = computed(() => items.value.reduce((acc, item) => acc + item.workerBalance, 0));
</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total Miner"
          :value="String(items.length)"
          icon="mdi-account-hard-hat"
          color="green"
        ></StatsCard>
      </v-col>
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total QA Power"
          :value="formatBytes(totalQAPower).combined"
          icon="mdi-relative-scale"
          color="secondary"
        ></StatsCard>
      </v-col>
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total Raw Bytes Power"
          :value="formatBytes(totalRawPower).combined"
          icon="mdi-nas"
          color="primary"
        ></StatsCard>
      </v-col>
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total Balance"
          :value="formatFIL(totalBalance)"
          icon="mdi-bank"
          color="warning"
        ></StatsCard>
      </v-col>
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total Available Balance"
          :value="formatFIL(totalAvailableBalance)"
          icon="mdi-wallet"
          color="success"
        ></StatsCard>
      </v-col>
      <v-col cols="6" lg="2">
        <StatsCard
          :loading="loading"
          title="Total Worker Balance"
          :value="formatFIL(totalWorkerBalance)"
          icon="mdi-wallet-outline"
          color="error"
        ></StatsCard>
      </v-col>
      <v-col cols="12">
        <Card title="Miners" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
            item-value="address"
            show-expand
          >
            <template v-slot:[`item.layers`]="{value}">
              <div v-for="(layer, index) in value" :key="index">
                <v-chip>{{ layer }}</v-chip>
              </div>
            </template>
            <template v-slot:[`item.actorBalance`]="{value}">
              {{ formatFIL(value) }}
            </template>
            <template v-slot:[`item.actorAvailableBalance`]="{value}">
              {{ formatFIL(value) }}
            </template>
            <template v-slot:[`item.workerBalance`]="{value}">
              {{ formatFIL(value) }}
            </template>
            <template v-slot:[`item.qualityAdjustedPower`]="{value}">
              {{ formatBytes(value).combined }}
            </template>
            <template v-slot:[`item.rawBytePower`]="{value}">
              {{ formatBytes(value).combined }}
            </template>
          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>

</template>

<style scoped>

</style>
