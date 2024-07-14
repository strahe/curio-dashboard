<script lang="ts" setup>
import Card from "@/components/Card.vue";

defineProps({
  title: {
    type: String,
    default: undefined,
  },
  color: {
    type: String,
    default: 'orange',
  },
  error: {
    type: Error,
    default: undefined,
  },
  ...Card.props,
})
</script>

<template>
<Card :title="title" :color="color" :error="error">
  <template #titleAction>
    <slot name="titleAction"></slot>
  </template>
  <template #actions v-if="$slots.actions">
    <slot name="actions"></slot>
  </template>
  <Chart v-bind="$attrs"></Chart>
</Card>
</template>
