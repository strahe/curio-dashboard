<script setup lang="ts">
import { computed } from 'vue';
import { getPrimary, getLightBorder } from '../../../forms/charts/apex-chart/UpdateColors';
import UiTitleCard from '@/components/shared/UiTitleCard.vue';

// icons
import { CaretDownFilled, DownloadOutlined } from '@ant-design/icons-vue';

const chartOptions1 = computed(() => {
  return {
    chart: {
      type: 'area',
      height: 355,
      fontFamily: `inherit`,
      foreColor: '#a1aab2',
      toolbar: {
        show: false
      }
    },
    colors: [getPrimary.value],
    labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
    stroke: {
      curve: 'straight',
      width: 1
    },
    fill: {
      type: 'gradient',
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.4,
        opacityTo: 0.2,
        stops: [0, 100]
      }
    },
    grid: {
      borderColor: getLightBorder.value,
      xaxis: {
        lines: {
          show: true
        }
      }
    },
    xaxis: {
      axisBorder: {
        show: true,
        color: getLightBorder.value
      },
      axisTicks: {
        color: getLightBorder.value
      }
    },
    dataLabels: {
      enabled: false
    }
  };
});

// chart 1
const areaChart1 = {
  series: [
    {
      name: 'Income',
      data: [100, 20, 60, 20, 20, 80, 20]
    }
  ]
};
</script>

<template>
  <UiTitleCard title="Income Overview" class-name="px-0 rounded-md overflow-hidden pb-0">
    <v-row class="justify-sm-space-between justify-center py-5 px-4">
      <v-col cols="12" sm="6">
        <div class="d-flex align-center text-error justify-sm-start justify-center">
          <CaretDownFilled />
          <h6 class="text-h6 mb-0 ml-1">$1,12,900 (45.67%)</h6>
        </div>
        <p class="text-h6 text-lightText mb-0 text-sm-left text-center">Compare to : 01 Dec 2021-08 Jan 2022</p>
      </v-col>
      <v-col cols="12" sm="6">
        <div class="d-flex align-center flex-wrap ga-2 justify-sm-end justify-center">
          <v-btn-group divided variant="outlined" color="secondary" density="compact">
            <v-btn>Week</v-btn>
            <v-btn>Month</v-btn>
          </v-btn-group>
          <v-btn icon rounded variant="outlined" color="secondary" size="small" disabled>
            <DownloadOutlined :style="{ fontSize: '14px' }" />
          </v-btn>
        </div>
      </v-col>
    </v-row>
    <apexchart type="area" height="355" :options="chartOptions1" :series="areaChart1.series"> </apexchart>
  </UiTitleCard>
</template>
