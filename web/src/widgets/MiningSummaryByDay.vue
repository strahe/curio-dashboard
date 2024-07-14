<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {MiningSummaryDay} from "@/typed-graph";
import {ChartData} from "chart.js";
import {formatDay} from "@/utils/formatDay";

const { result, loading,refetch } = useQuery(gql`
  query GetMiningSummaryByDay($days: Int!) {
    miningSummaryByDay(lastDays: $days) {
      day
      sp_id
      won
    }
  }
`, {
  days: 7, // todo: make this a prop
}, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[MiningSummaryDay]> = computed(() => result.value?.miningSummaryByDay || []);
const chartData = ref<ChartData<'bar'>>({
  labels: [],
  datasets: [],
})
watch(items, (newItems) => {
  // Step 1: Normalize and sort the days
  const labelsSet = new Set<string>();
  newItems.forEach((item) => {
    labelsSet.add(formatDay(new Date(item.day)));
  });
  const sortedLabels = Array.from(labelsSet).sort();

// Step 2: Initialize datasets for each `sp_id` with pre-filled data arrays of zeros
  const datasetsBySpId: { [sp_id: number]: { label: string; data: number[] } } = {};
  newItems.forEach((item) => {
    if (!datasetsBySpId[item.sp_id]) {
      datasetsBySpId[item.sp_id] = {
        label: item.sp_id,
        data: Array(sortedLabels.length).fill(0),
      };
    }
  });

// Step 3: Populate the datasets
  newItems.forEach((item) => {
    const day = formatDay(new Date(item.day));
    const dayIndex = sortedLabels.indexOf(day);
    if (dayIndex !== -1) {
      datasetsBySpId[item.sp_id].data[dayIndex] += item.won;
    }
  });

  const datasets = Object.values(datasetsBySpId);
  chartData.value = {
    labels: sortedLabels,
    datasets: datasets,
  };
})
const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
  },
}
</script>

<template>
<ChartCard
  v-bind="$attrs"
  :loading="loading"
  name="win-summary-by-day"
  title="Block Wins"
  type='Bar'
  :data="chartData"
  :options="chartOptions"
>
  <template #titleAction>
    <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
  </template>
</ChartCard>
</template>

<style scoped>

</style>
