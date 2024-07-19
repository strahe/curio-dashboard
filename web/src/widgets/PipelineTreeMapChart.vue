<script setup lang="ts">

import {Pipeline} from "@/typed-graph";


const props = defineProps({
  data: {
    type: Object as () => Pipeline,
    required: true,
  },
  height: {
    type: Number,
    default: 150
  }
})

const items = [
  {x: 'SDR', y: 30, key: 'afterSdr'},
  {x: 'TreeD', y: 10, key: 'afterTreeD'},
  {x: 'TreeC', y: 5, key: 'afterTreeC'},
  {x: 'TreeR', y: 5, key: 'afterTreeR'},
  {x: 'PreCommit', y: 5, key: 'afterPrecommitMsg'},
  {x: 'PreMsg', y: 5, key: 'afterPrecommitMsgSuccess'},
  {x: 'PoRep', y: 20, key: 'afterPorep'},
  {x: 'Finalize', y: 5, key: 'afterFinalize'},
  {x: 'MoveStorage', y: 10, key: 'afterMoveStorage'},
  {x: 'Commit', y: 20, key: 'afterCommitMsg'},
  {x: 'CommitMsg', y: 5, key: 'afterCommitMsgSuccess'},
]


function generateColors(data: Pipeline): string[] {
  let colors = items.map(() => '#797a78');
  items.forEach((item, index) => {
    if (data[item.key as keyof typeof data]) {
      colors[index] = '#1ef109'; // Assuming green color is '#1ef109'
    }
  });

  return colors;
}

const chartData = computed(() => {
  return {
    series: items.map((item) => {
      return {
        data: [{
          x: item.x,
          y: item.y
        }]
      }
    }),
    options: {
      legend: {
        show: false
      },
      chart: {
        height: 100,
        type: 'treemap',
        toolbar: {
          show: false
        }
      },
      colors: generateColors(props.data),
      plotOptions: {
        treemap:{
          enableShades: false
        }
      },
      dataLabels: {
        enabled: true
      },
      tooltip: {
        enabled: false
      },
      theme: {
        // mode:  theme.global.current.value?.dark ? 'dark' : 'light'
      }
    },
  }
})
</script>

<template>
  <apexchart :options="chartData.options" :series="chartData.series" type="treemap" :height="height"></apexchart>
</template>

<style scoped>

</style>
