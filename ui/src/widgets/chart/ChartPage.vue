<script setup lang="ts">
import { ref } from 'vue';
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue';

// import components
import ChartCards from './components/ChartCards.vue';
import UniqueVisitor from './components/UniqueVisitor.vue';
import IncomeOverview from './components/IncomeOverview.vue';
import IncomeCompare from './components/IncomeCompare.vue';
import AnalyticsReport from './components/AnalyticsReport.vue';
import SalesReport from './components/SalesReport.vue';
import AcquisitionChannels from './components/AcquisitionChannels.vue';

const page = ref({ title: 'Charts' });
const breadcrumbs = ref([
  {
    title: 'Widget',
    disabled: false,
    href: '#'
  },
  {
    title: 'Charts',
    disabled: true,
    href: '#'
  }
]);
</script>

<template>
  <BaseBreadcrumb :title="page.title" :breadcrumbs="breadcrumbs"></BaseBreadcrumb>

  <ChartCards />
  <v-row class="mb-0">
    <v-col cols="12" md="8">
      <UniqueVisitor />
    </v-col>

    <v-col cols="12" md="4">
      <IncomeOverview />
    </v-col>
  </v-row>

  <v-row class="mb-0">
    <v-col cols="12" md="4">
      <AnalyticsReport />
    </v-col>
    <v-col cols="12" md="8">
      <IncomeCompare />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="12" md="8">
      <SalesReport />
    </v-col>
    <v-col cols="12" md="4">
      <AcquisitionChannels />
    </v-col>
  </v-row>
</template>
