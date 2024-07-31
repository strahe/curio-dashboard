<script setup lang="ts">
import { computed } from 'vue';
import { useTheme } from 'vuetify';
import { getLightBorder } from '../../../forms/charts/apex-chart/UpdateColors';
import UiTitleCard from '@/components/shared/UiTitleCard.vue';

const theme = useTheme();
const WarningColor = theme.current.value.colors.warning;

const chartOptions1 = computed(() => {
  return {
    chart: {
      type: 'line',
      height: 340,
      fontFamily: `inherit`,
      foreColor: '#a1aab2',
      toolbar: {
        show: false
      }
    },
    colors: [WarningColor],
    dataLabels: {
      enabled: false
    },
    labels: ['Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec', 'Jan'],
    xaxis: {
      crosshairs: {
        width: 1
      },
      labels: {
        offsetX: 8
      },
      axisTicks: {
        show: false
      },
      axisBorder: {
        show: false
      }
    },
    yaxis: {
      show: false
    },
    stroke: {
      curve: 'smooth',
      width: 1.5
    },
    grid: {
      strokeDashArray: 4,
      borderColor: getLightBorder.value
    },
    tooltip: {
      fixed: {
        enabled: false
      },
      x: {
        show: false
      },
      marker: {
        show: false
      }
    }
  };
});

// chart 1
const lineChart1 = {
  series: [
    {
      name: 'Income',
      data: [58, 90, 38, 83, 63, 75, 35, 55]
    }
  ]
};
</script>

<template>
  <UiTitleCard title="Analytics Report" class-name="px-0 rounded-md overflow-hidden pb-0">
    <v-list class="py-0" aria-label="report list" aria-busy="true" border>
      <v-list-item value="Company Finance Growth">
        <div class="d-inline-flex align-center justify-space-between w-100 ga-2">
          <h6 class="text-h6 mb-0">Company Finance Growth</h6>
          <h5 class="ml-auto text-h5 mb-0">+45.14%</h5>
        </div>
      </v-list-item>
      <v-list-item value="Company Expenses Ratio">
        <div class="d-inline-flex align-center justify-space-between w-100 ga-2">
          <h6 class="text-h6 mb-0">Company Expenses Ratio</h6>
          <h5 class="ml-auto text-h5 mb-0">0.58%</h5>
        </div>
      </v-list-item>
    </v-list>
    <v-divider></v-divider>
    <apexchart type="line" height="340" :options="chartOptions1" :series="lineChart1.series"> </apexchart>
  </UiTitleCard>
</template>
