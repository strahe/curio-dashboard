<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import { TaskSummaryDay} from "@/typed-graph";
import {formatDay} from "@/utils/formatDay";
import {ComputedRef} from "vue";

const { result, loading, refetch, error } = useQuery(gql`
  query GetTaskSummaryByDay($days: Int!) {
    taskSummaryByDay(lastDays: $days) {
      day
      falseCount
      trueCount
      totalCount
    }
  }
`, {
  days: 7,
}, () => ({
  fetchPolicy: 'cache-first',
}));
const items: ComputedRef<[TaskSummaryDay]> = computed(() => result.value?.taskSummaryByDay || []);
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
      plugins: {
        legend: {
          display: false,
        },
      },
    },
  };
});

function prepareData(newItems) {
  const sortedItems = [...newItems].sort((a, b) => new Date(a.day).getTime() - new Date(b.day).getTime());
  return {
    labels: sortedItems.map((item) => {
      return formatDay(new Date(item.day));
    }),
    datasets: [
      {
        label: 'Success',
        data: sortedItems.map((item) => item.trueCount),
      },
      {
        label: 'Failed',
        data: sortedItems.map((item) => item.falseCount),
      }
    ]
  }
}

</script>

<template>
<ChartCard
  v-bind="$attrs"
  :loading="loading"
  name="task-summary-by-day"
  title="Completed Tasks"
  type='Line'
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
