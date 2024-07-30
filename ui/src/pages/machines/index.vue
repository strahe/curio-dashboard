<script setup lang="ts">
  import { ref } from 'vue'
  import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue'
  import RecievableCard from '@/pages/pipelines/RecievableCard.vue'
  import MachineList from '@/pages/machines/MachineList.vue'
  import WidgetMachineSummaryCharts from '@/pages/machines/WidgetMachineSummaryCharts.vue'

  const page = ref({ title: 'Pipelines' })
  const breadcrumbs = ref([
    {
      title: 'Pipeline',
      disabled: false,
      href: '#',
    },
    {
      title: 'List',
      disabled: true,
      href: '#',
    },
  ])

</script>

<template>
  <BaseBreadcrumb :breadcrumbs="breadcrumbs" :title="page.title" />
  <v-row>
    <v-col cols="12" lg="12">
      <WidgetMachineSummaryCharts />
    </v-col>
    <v-col cols="12">
      <MachineList />
    </v-col>
  </v-row>
</template>

<style scoped>

</style>
