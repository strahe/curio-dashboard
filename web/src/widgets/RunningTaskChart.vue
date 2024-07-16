<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import {GetRunningTasks} from "@/query/task";
import {Task} from "@/typed-graph";
import {ComputedRef} from "vue";

const props = defineProps({
  title: {
    type: String,
    default: 'Running Tasks'
  },
  type: {
    type: String,
    default: 'Pie'
  }
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

const chartData = computed(() => ({
  data: {
    labels: Object.keys(nameCounts.value),
    datasets: [{ label: 'value', data: Object.values(nameCounts.value) }],
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'left',
        display: props.type === 'Pie',
      },
      colors: {
        enabled: true,
        forceOverride: true
      }
    },
  }
}));

</script>

<template>
<ChartCard
  :loading="loading"
  :title="title"
  :type="type"
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
