<script lang="ts" setup>
import Card from "@/components/Card.vue";
import VueApexCharts from "vue3-apexcharts";

defineProps({
  type: {
    type: String
  },
  title: {
    type: String,
    default: undefined,
  },
  error: {
    type: Error,
    default: undefined,
  },
  options: {
    type: Object
  },
  series: {
    type: Array,
    required: true
  },
  width: {
    default: "100%"
  },
  height: {
    default: "auto"
  }
})
</script>

<template>
<Card :title="title" :error="error">
  <template #titleAction>
    <slot name="titleAction"></slot>
  </template>
  <template #actions v-if="$slots.actions">
    <slot name="actions"></slot>
  </template>
  <VueApexCharts :type="type" :options="options" :series="series" :width="width" :height="height"></VueApexCharts>
</Card>
</template>
