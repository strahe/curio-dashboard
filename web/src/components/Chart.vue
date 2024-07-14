<template>
  <div style="max-height: 300px">
    <component :id="props.name" :is="component" :data="props.data" :options="props.options" />
  </div>
</template>

<script lang="ts" setup>
import { Bar, Pie, Line, Chart, Bubble, Radar, Doughnut, Scatter, PolarArea } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  PointElement,
  LineElement,
  CategoryScale,
  LinearScale, ArcElement,
  Colors
} from 'chart.js';
import { computed} from 'vue';

ChartJS.register(Title, Tooltip, Legend, BarElement, PointElement, LineElement, CategoryScale, LinearScale, ArcElement, Colors);

const props = defineProps({
  name: {
    type: String,
    default: null,
  },
  type: {
    type: String,
    required: true,
  },
  data: {
    type: Object,
    required: true,
    default: () => ({
      labels: [],
      datasets: [],
    }),
  },
  options: {
    type: Object,
    required: true,
    default: () => ({
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false,
        },
      },
    }),
  },
});

const component = computed(() => {
  const found = Object.entries({Bar, Pie, Line, Chart, Bubble, Radar, Doughnut, Scatter, PolarArea}).find(
    ([key]) => key === props.type
  );
  return found ? found[1] : null;
});

</script>

<style lang="scss">
</style>
