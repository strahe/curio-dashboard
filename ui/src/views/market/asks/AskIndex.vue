<script setup lang="ts">
import { ref } from "vue"
import BaseBreadcrumb from "@/components/shared/BaseBreadcrumb.vue"
import Asklist from "@/views/market/asks/AskList.vue"

const breadcrumbs = ref([
  {
    title: 'Market',
    disabled: true,
    to: '#',
  },
  {
    title: 'Storage Asks',
    disabled: true,
    to: '#',
  },
])
</script>

<template>
  <BaseBreadcrumb :breadcrumbs="breadcrumbs" />
  <v-row>
    <v-col
      cols="12"
    >
      <Asklist />
    </v-col>
  </v-row>
</template>
