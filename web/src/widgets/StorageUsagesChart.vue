<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {ComputedRef} from "vue";
import {StorageUsage} from "@/typed-graph";
import {useTheme} from "vuetify";
import {formatBytes} from "@/utils/formatBytes";

const theme = useTheme();

const props = defineProps({
  title: {
    type: String,
    default: undefined,
  },
})

const { result, loading, refetch, error } = useQuery(gql`
    query GetStorageUsages($storageID: String, $lastDays: Int!) {
        storageUsages(storageID: $storageID, lastDays: $lastDays) {
            time
            available
            fsAvailable
            reserved
            used
        }
    }
`,  {
  lastDays: 7
}, ()=>({
  fetchPolicy: 'cache-first',
}))

const items: ComputedRef<[StorageUsage]> = computed(() => result.value?.storageUsages || []);

const series = computed(() => {
  return [{
    name: 'Available',
    data: items.value.map((item: any) => item.available)
  },{
    name: 'FS Available',
    data: items.value.map((item: any) => item.fsAvailable)
  },{
    name: 'Reserved',
    data: items.value.map((item: any) => item.reserved)
  },{
    name: 'Used',
    data: items.value.map((item: any) => item.used)
  }]
});

const firstRender = ref(true);
function updated(chartContext) {
  if (firstRender.value) {
    firstRender.value = false;
    chartContext.hideSeries("FS Available");
    chartContext.hideSeries("Reserved");
  }
}

const chartData = computed(() => {
  return {
    series: series.value,
    options: {
      chart: {
        id: 'storage-usage',
      },
      xaxis: {
        type: 'datetime',
        categories: items.value.map((item: any) => item.time)
      },
      title: {
        text: props.title,
      },
      stroke: {
        curve: 'smooth'
      },
      dataLabels: {
        enabled: false
      },
      theme: {
        mode: theme.current.value.dark ? 'dark' : 'light'
      },
      fill: {
        type: 'gradient',
        gradient: {
          shadeIntensity: 1,
          inverseColors: false,
          opacityFrom: 0.5,
          opacityTo: 0,
          stops: [0, 90, 100]
        },
      },
      tooltip:{
        y: {
          formatter: function(value: number) {
            return formatBytes(value).combined;
          }
        }
      },
      yaxis: {
        labels: {
          formatter: function(value: number) {
            return formatBytes(value).combined;
          }
        }
      }
    }
  }
})

</script>

<template>
  <Card :loading="loading" :error="error as Error">
    <template #titleAction>
      <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
    </template>
    <apexchart
      :options="chartData.options"
      :series="chartData.series"
      type="area"
      height="250"
      @updated="updated"
    ></apexchart>
  </Card>
</template>

<style scoped>

</style>
