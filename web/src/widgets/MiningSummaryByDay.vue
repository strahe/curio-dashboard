<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {MiningSummaryDay} from "@/typed-graph";
import {formatDay} from "@/utils/formatDay";
import {ComputedRef} from "vue";

const { result, loading,refetch, error } = useQuery(gql`
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

const chartData = computed(() => {
  const data = prepareData(items.value);
  return {
    data: {
      labels: data.labels,
      datasets: data.datasets,
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
    },
  };
});

function prepareData(newItems) {
  const labelsSet = new Set<string>();
  newItems.forEach((item) => {
    labelsSet.add(formatDay(new Date(item.day)));
  });
  const sortedLabels = Array.from(labelsSet).sort();

  const datasetsBySpId: { [sp_id: number]: { label: string; data: number[] } } = {};
  newItems.forEach((item) => {
    if (!datasetsBySpId[item.sp_id]) {
      datasetsBySpId[item.sp_id] = {
        label: item.sp_id,
        data: Array(sortedLabels.length).fill(0),
      };
    }
  });

  newItems.forEach((item) => {
    const day = formatDay(new Date(item.day));
    const dayIndex = sortedLabels.indexOf(day);
    if (dayIndex !== -1) {
      datasetsBySpId[item.sp_id].data[dayIndex] += item.won;
    }
  });

  const datasets = Object.values(datasetsBySpId);
  return  {
    labels: sortedLabels,
    datasets: datasets,
  };
}

</script>

<template>
<ChartCard
  v-bind="$attrs"
  :loading="loading"
  name="win-summary-by-day"
  title="Block Wins"
  type='Bar'
  :data="chartData.data"
  :options="chartData.options"
  :error="error as Error"
>
  <template #titleAction>
    <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
  </template>
</ChartCard>
</template>

<style scoped>

</style>
