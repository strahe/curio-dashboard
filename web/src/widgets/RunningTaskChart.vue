<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import {GetRunningTasks} from "@/query/task";
import {Task} from "@/typed-graph";
import {ComputedRef} from "vue";

const props = defineProps({
  title: {
    type: String,
    default: undefined
  },
  height: {
    type: Number,
    default: 250
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

const chartData = computed(() => {
  return {
    series: [{
      name: 'Running',
      data: Object.values(nameCounts.value)
    }],
    options: {
      chart: {
        type: 'bar',
      },
      plotOptions: {
        bar: {
          borderRadius: 4,
          borderRadiusApplication: 'end',
          horizontal: false,
        }
      },
      title: {
        text: props.title,
      },
      dataLabels: {
        enabled: true
      },
      xaxis: {
        categories: Object.keys(nameCounts.value),
      }
    },
  }
})
</script>

<template>
  <Card :loading="loading" :error="error as Error">
    <template #titleAction>
      <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
    </template>
    <apexchart :options="chartData.options" :series="chartData.series" type="bar" :height="height"></apexchart>
  </Card>
</template>

<style scoped>

</style>
