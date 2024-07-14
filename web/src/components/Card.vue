<script setup lang="ts">
defineProps({
  title: {
    type: String,
    default: null,
  },
  color: {
    type: String,
    default: 'orange',
  },
  error: {
    type: Error,
    default: null,
  },
})
</script>

<template>
<v-card>
  <v-card-title v-if="title" :style="color ? `color: ${color};` : ''" class="d-flex justify-space-between">
    <div>{{ title }}</div>
    <slot name="titleAction" />
  </v-card-title>
  <v-card-text>
    <v-alert v-if="error" type="error">
      {{ error.message }}
    </v-alert>
    <slot></slot>
  </v-card-text>
  <v-card-actions v-if="$slots.actions">
    <slot name="actions" />
  </v-card-actions>
</v-card>
</template>

<style scoped>

</style>
