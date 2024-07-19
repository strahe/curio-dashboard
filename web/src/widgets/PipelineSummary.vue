<script setup lang="ts">

import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {PipelineSummary} from "@/typed-graph";
import {useTheme} from "vuetify";

const theme = useTheme();

const props = defineProps({
  title: {
    type: String,
    default: undefined,
  },
  height: {
    type: Number,
    default: 250
  }
})

const { result, loading, error } = useQuery(gql`
    query GetPipelineSummary {
        pipelinesSummary {
        id
        sdr
        trees
        precommitMsg
        waitSeed
        porep
        commitMsg
        done
        failed
        }
  }
`,  null, ()=>({
  fetchPolicy: 'cache-first',
}))

const items = computed(() => result.value?.pipelinesSummary || []);
const categories = computed(() => items.value.map((item: PipelineSummary) => item.id))
const series = computed(() => {
  return [
    {
      name: 'SDR',
      data: items.value.map((item: any) => item.sdr)
    },{
      name: 'Trees',
      data: items.value.map((item: any) => item.trees)
    },{
      name: 'Precommit',
      data: items.value.map((item: any) => item.precommitMsg)
    },{
      name: 'Wait Seed',
      data: items.value.map((item: any) => item.waitSeed)
    },{
      name: 'PoRep',
      data: items.value.map((item: any) => item.porep)
    },{
      name: 'Commit',
      data: items.value.map((item: any) => item.commitMsg)
    },{
      name: 'Done',
      data: items.value.map((item: any) => item.done)
    },{
      name: 'Failed',
      data: items.value.map((item: any) => item.failed)
    }
  ]
})

const chartData = computed(() => {
  return {
    series: series.value,
    options: {
      chart: {
        id: 'pipeline-summary',
        stacked: true,
      },
      plotOptions: {
        bar: {
          horizontal: true,
          dataLabels: {
            total: {
              enabled: true,
              offsetX: 0,
              style: {
                fontSize: '13px',
                fontWeight: 900
              }
            }
          }
        },
      },
      title: {
        text: props.title,
      },
      stroke: {
        width: 1,
        colors: ['#fff']
      },
      xaxis: {
        categories: categories.value,
        labels: {
          show: false
        }
      },
      yaxis: {
        title: {
          text: undefined
        },
      },
      fill: {
        opacity: 1
      },
      legend: {
        position: 'top',
        horizontalAlign: 'left',
        offsetX: 40
      },
      theme: {
        mode: theme.current.value.dark ? 'dark' : 'light'
      }
    }
  }
})

</script>

<template>
  <Card :loading="loading" :error="error as Error">
    <apexchart :options="chartData.options" :series="chartData.series" type="bar" :height="height"></apexchart>
  </Card>
</template>

<style scoped lang="sass">

</style>
