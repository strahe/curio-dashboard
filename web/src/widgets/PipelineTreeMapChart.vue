<script setup lang="ts">

let series = [
  {x: 'SDR', y: 30},
  {x: 'TreeD', y: 10},
  {x: 'TreeC', y: 5},
  {x: 'TreeR', y: 5},
  {x: 'PreCommit', y: 5},
  {x: 'PreMsg', y: 5},
  {x: 'PoRep', y: 20},
  {x: 'Finalize', y: 5},
  {x: 'MoveStorage', y: 10},
  {x: 'Commit', y: 20},
  {x: 'CommitMsg', y: 5},

]
const chartData=  {

    series: [
      {
        data: series
      },
      {
        data: [
          {x: 'Failed', y: 10},
        ]
      }
    ],
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
    colors: [
      '#3B93A5',
      '#F7B844',
      '#ADD8C7',
      '#EC3C65',
      '#CDD7B6',
      '#C1F666',
      '#D43F97',
      '#1E5D8C',
      '#421243',
      '#7F94B0',
      '#EF6537',
      '#C0ADDB'
    ],
    plotOptions: {
      treemap:{
        enableShades: false
      }
    },
      dataLabels: {
        enabled: true
      },
    },
  }

</script>

<template>
  <apexchart :options="chartData.options" :series="chartData.series" type="treemap" height="150"></apexchart>
</template>

<style scoped>

</style>
