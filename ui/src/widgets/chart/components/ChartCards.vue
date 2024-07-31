<script setup lang="ts">
import { computed } from 'vue';
import { useTheme } from 'vuetify';

const theme = useTheme();
import { getPrimary } from '../../../forms/charts/apex-chart/UpdateColors';
const warningColor = theme.current.value.colors.warning;

// icons
import { RiseOutlined, FallOutlined } from '@ant-design/icons-vue';

const chartOptions1 = computed(() => {
  return {
    chart: {
      type: 'bar',
      height: 100,
      fontFamily: `inherit`,
      sparkline: {
        enabled: true
      }
    },
    dataLabels: {
      enabled: false
    },
    colors: getPrimary.value,
    stroke: {
      curve: 'smooth',
      width: 0
    },
    tooltip: {
      fixed: {
        enabled: false
      },
      x: {
        show: false
      },
      y: {
        title: {
          formatter: () => 'Users : $'
        }
      }
    }
  };
});

// chart 2
const chartOptions2 = computed(() => {
  return {
    chart: {
      type: 'area',
      height: 100,
      offsetY: 10,
      fontFamily: `inherit`,
      sparkline: {
        enabled: true
      }
    },
    dataLabels: {
      enabled: false
    },
    colors: ['#ff4d4f'],
    fill: {
      type: 'gradient',
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.6,
        opacityTo: 0.7,
        stops: [0, 100]
      }
    },
    stroke: {
      curve: 'smooth',
      width: 3
    },
    tooltip: {
      fixed: {
        enabled: false
      },
      x: {
        show: false
      },
      y: {
        title: {
          formatter: () => 'Orders :'
        }
      }
    }
  };
});

// chart 3
const chartOptions3 = computed(() => {
  return {
    chart: {
      type: 'bar',
      height: 100,
      fontFamily: `inherit`,
      sparkline: {
        enabled: true
      }
    },
    dataLabels: {
      enabled: false
    },
    colors: warningColor,
    stroke: {
      curve: 'smooth',
      width: 0
    },
    tooltip: {
      fixed: {
        enabled: false
      },
      x: {
        show: false
      },
      y: {
        title: {
          formatter: () => 'Sales : $'
        }
      }
    }
  };
});

// chart 4
const chartOptions4 = computed(() => {
  return {
    chart: {
      type: 'area',
      height: 100,
      fontFamily: `inherit`,
      sparkline: {
        enabled: true
      }
    },
    dataLabels: {
      enabled: false
    },
    colors: [getPrimary.value],
    fill: {
      type: 'gradient',
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.7,
        opacityTo: 0.9,
        stops: [0, 100]
      }
    },
    stroke: {
      curve: 'straight',
      width: 1.5
    },
    tooltip: {
      fixed: {
        enabled: false
      },
      x: {
        show: false
      },
      y: {
        title: {
          formatter: () => 'Marketing : $'
        }
      }
    }
  };
});

// chart 1
const barChart1 = {
  series: [
    {
      name: 'Users',
      data: [
        220, 230, 240, 220, 225, 215, 205, 195, 185, 150, 185, 195, 80, 205, 215, 225, 240, 225, 215, 205, 80, 215, 225, 240, 215, 210, 190
      ]
    }
  ]
};

// chart 2
const areaChart2 = {
  series: [
    {
      name: 'Orders',
      data: [1800, 1500, 1800, 1700, 1400, 1200, 1000, 800, 600, 500, 600, 800, 500, 700, 400, 600, 500, 600]
    }
  ]
};

// chart 3
const barChart2 = {
  series: [
    {
      name: 'Sales',
      data: [
        220, 230, 240, 220, 225, 215, 205, 195, 185, 150, 185, 195, 80, 205, 215, 225, 240, 225, 215, 205, 80, 215, 225, 240, 215, 210, 180
      ]
    }
  ]
};

// chart 4
const areaChart4 = {
  series: [
    {
      name: 'Marketing',
      data: [100, 140, 100, 240, 115, 125, 90, 100, 80, 150, 160, 150, 170, 155, 150, 160, 145, 200, 140, 160]
    }
  ]
};
</script>

<template>
  <v-row class="mb-0">
    <!-- chart 1 -->
    <v-col cols="12" sm="6" md="3">
      <v-card variant="outlined" elevation="0" class="bg-surface">
        <v-card-text>
          <div class="d-flex align-start">
            <div>
              <h6 class="text-h6 text-lightText">Total Users</h6>
              <h4 class="text-h4 mb-0 d-flex align-center flex-wrap gap-1">
                78,250
                <v-chip color="primary" border="primary solid thin opacity-50" class="ml-2" size="small" label>
                  <template v-slot:prepend>
                    <component :is="RiseOutlined" :style="{ fontSize: '12px' }" :class="'mr-1 text-primary'" />
                  </template>
                  70.5%
                </v-chip>
              </h4>
            </div>
          </div>
        </v-card-text>
        <apexchart type="bar" height="100" :options="chartOptions1" :series="barChart1.series"> </apexchart>
      </v-card>
    </v-col>

    <!-- chart 2 -->
    <v-col cols="12" sm="6" md="3">
      <v-card variant="outlined" elevation="0" class="bg-surface overflow-hidden">
        <v-card-text>
          <div class="d-flex align-start">
            <div>
              <h6 class="text-h6 text-lightText">Total Order</h6>
              <h4 class="text-h4 mb-0 d-flex align-center flex-wrap gap-1">
                18,800
                <v-chip color="warning" border="warning solid thin opacity-50" class="ml-2" size="small" label>
                  <template v-slot:prepend>
                    <component :is="FallOutlined" :style="{ fontSize: '12px' }" :class="'mr-1 text-warning'" />
                  </template>
                  27.4%
                </v-chip>
              </h4>
            </div>
          </div>
        </v-card-text>
        <apexchart type="area" height="100" :options="chartOptions2" :series="areaChart2.series"> </apexchart>
      </v-card>
    </v-col>

    <!-- chart 3 -->
    <v-col cols="12" sm="6" md="3">
      <v-card variant="outlined" elevation="0" class="bg-surface">
        <v-card-text>
          <div class="d-flex align-start">
            <div>
              <h6 class="text-h6 text-lightText">Total Sales</h6>
              <h4 class="text-h4 mb-0 d-flex align-center flex-wrap gap-1">
                $35,078
                <v-chip color="warning" border="warning solid thin opacity-50" class="ml-2" size="small" label>
                  <template v-slot:prepend>
                    <component :is="FallOutlined" :style="{ fontSize: '12px' }" :class="'mr-1 text-warning'" />
                  </template>
                  27.4%
                </v-chip>
              </h4>
            </div>
          </div>
        </v-card-text>
        <apexchart type="bar" height="100" :options="chartOptions3" :series="barChart2.series"> </apexchart>
      </v-card>
    </v-col>

    <!-- chart 4 -->
    <v-col cols="12" sm="6" md="3">
      <v-card variant="outlined" elevation="0" class="bg-surface overflow-hidden">
        <v-card-text>
          <div class="d-flex align-start">
            <div>
              <h6 class="text-h6 text-lightText">Total Marketing</h6>
              <h4 class="text-h4 mb-0 d-flex align-center flex-wrap gap-1">
                $1,12,083
                <v-chip color="primary" border="primary solid thin opacity-50" class="ml-2" size="small" label>
                  <template v-slot:prepend>
                    <component :is="RiseOutlined" :style="{ fontSize: '12px' }" :class="'mr-1 text-primary'" />
                  </template>
                  27.4%
                </v-chip>
              </h4>
            </div>
          </div>
        </v-card-text>
        <apexchart type="area" height="100" :options="chartOptions4" :series="areaChart4.series"> </apexchart>
      </v-card>
    </v-col>
  </v-row>
</template>
