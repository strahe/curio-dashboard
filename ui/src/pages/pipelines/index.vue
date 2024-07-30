<script setup lang="ts">
import { ref } from 'vue'
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue'
import WidgetCharts from '@/pages/pipelines/WidgetCharts.vue'
import RecievableCard from '@/pages/pipelines/RecievableCard.vue'
import OrderTab from '@/pages/pipelines/OrderTab.vue'
const page = ref({ title: 'Pipelines' })
const breadcrumbs = ref([
    {
        title: 'Invoice',
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
    <v-col cols="12" lg="8">
      <WidgetCharts />
    </v-col>
    <v-col cols="12" lg="4">
      <RecievableCard />
    </v-col>
    <v-col cols="12">
      <OrderTab />
    </v-col>
  </v-row>
</template>

<style scoped>

</style>
