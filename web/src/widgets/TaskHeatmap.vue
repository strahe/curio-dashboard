<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {ComputedRef} from "vue";
import {TaskAggregate} from "@/typed-graph";
import {useTheme} from "vuetify";

const theme = useTheme();

const props = defineProps({
  title: {
    type: String,
    default: undefined,
  },
  lastDays: {
    type: Number,
    default: 7,
  },
})

const { result, loading, refetch, error } = useQuery(gql`
  query GetTaskAggregates($lastHours: Int!) {
    taskAggregatesByHour(lastHours: $lastHours) {
    time
    success
    }
  }
`, {
  lastHours: props.lastDays * 24,
}, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[TaskAggregate]> = computed(() => result.value?.taskAggregatesByHour || []);

function getHourlyData(day: string) {
  return Array.from({ length: 24 }, (_, hour) => {
    const pd = new Date(`${day} ${hour}:00:00`);
    const item = items.value.find(item => areDatesEqualAtHourLevel(pd, new Date(item.time)));
    return item ? { x: hour, y: item.success}: {x: hour, y: 0};
  });
}

function areDatesEqualAtHourLevel(date1: Date, date2: Date): boolean {
  return  date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate() &&
    date1.getHours() === date2.getHours();
}

function generatePastDates(lastDays: number): string[] {
  const dates: string[] = [];
  const currentDate = new Date();

  for (let i = 0; i < lastDays; i++) {
    const pastDate = new Date(currentDate);
    pastDate.setDate(currentDate.getDate() - i);
    const formattedDate = `${pastDate.getFullYear()}-${(pastDate.getMonth() + 1).toString().padStart(2, '0')}-${pastDate.getDate().toString().padStart(2, '0')}`;
    dates.unshift(formattedDate); // Add to the beginning to have dates in ascending order
  }

  return dates;
}

const series = computed(() => {
  const lastDays = generatePastDates(7);
  return lastDays.reverse().map(day => ({
    name: day.slice(5),
    data: getHourlyData(day)
  }));
});

const chartData = computed(() => {
  return {
    series: series.value,
    options: {
      chart: {
        id: 'task-heatmap',
        toolbar: {
          show: true,
          tools: {
            download: true,
            selection: false,
            zoom: false,
            zoomin: false,
            zoomout: false,
            pan: false,
            reset: false
          },
        },
      },
      dataLabels: {
        enabled: false
      },
      title: {
        text: props.title,
      },
      colors: [theme.current.value.colors.success],
      theme: {
        mode: theme.current.value.dark ? 'dark' : 'light'
      },
      plotOptions: {
        heatmap: {
          radius: 2,
          enableShades: true,
        }}
      }
  }
})
</script>

<template>
  <Card :loading="loading" :error="error as Error">
    <template #titleAction>
      <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
    </template>
    <apexchart :options="chartData.options" :series="chartData.series" type="heatmap" height="250"></apexchart>
  </Card>
</template>

<style scoped>

</style>
