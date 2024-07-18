<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import {GetRunningTasks} from "@/query/task";
import {Task} from "@/typed-graph";
import {ComputedRef} from "vue";

defineProps({
  title: {
    type: String,
    default: null
  },
})

const { result, loading, refetch, error } = useQuery(GetRunningTasks,  null, ()=>({
  fetchPolicy: 'cache-first',
}))
const items: ComputedRef<[Task]> = computed(() => result.value?.tasks || []);
const nameCounts = computed(() => {
  return items.value.reduce((acc: Record<string, number>, item) => {
    acc[item.name] = (acc[item.name] || 0) + 1;
    return acc;
  }, {});
});

const chartData = computed(() => {
  return {
    series: [{
      name: 'Running',
      data: Object.values(nameCounts.value)
    }],
    options: {
      chart: {
        type: 'bar',
        toolbar: { show: false }
      },
      plotOptions: {
        bar: {
          borderRadius: 4,
          borderRadiusApplication: 'end',
          horizontal: false,
        }
      },
      dataLabels: {
        enabled: false
      },
      xaxis: {
        categories: Object.keys(nameCounts.value),
      }
    },
  }
})
</script>

<template>
<ChartCard
  v-bind="$attrs"
  :loading="loading"
  :title="title"
  :series="chartData.series"
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
