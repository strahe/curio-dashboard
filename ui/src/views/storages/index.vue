<script setup lang="ts">
import { ref } from 'vue'
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue'
import StoragePathList from '@/views/storages/StoragePathList.vue'
import StorageUsageChart from '@/views/storages/StorageUsageChart.vue'

const page = ref({ title: 'Storages' })
const breadcrumbs = ref([
  {
    title: 'Storages',
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
      <StorageUsageChart />
    </v-col>
    <v-col cols="12">
      <StoragePathList />
    </v-col>
  </v-row>
</template>

<style scoped>

</style>
