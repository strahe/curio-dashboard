<script setup lang="ts">
defineProps({
  title: {
    type: String,
    required: true
  },
  value: {
    type: String,
    required: true
  },
  progress: {
    type: Number || undefined,
    default: undefined,
    validator: (value: number) => value >= 0 && value <= 100
  },
  loading: {
    type: Boolean,
    default: false
  },
  color: {
    type: String,
    default: 'red'
  },
  icon: {
    type: String,
    default: 'mdi-help'
  },
})
</script>

<template>
<v-card :loading="loading">
  <v-card-subtitle class="mt-3">{{title}}</v-card-subtitle>

  <v-card-text>
    <v-row>
      <v-col cols="8" class="value-font-size">
        {{value}}
      </v-col>
      <v-col cols="4">
        <v-icon size="40" :color="color">
          {{ icon }}
        </v-icon>
      </v-col>
    </v-row>
  </v-card-text>
  <v-card-actions v-if="progress">
    <v-progress-linear :model-value="progress" :color="color"></v-progress-linear>
  </v-card-actions>
</v-card>
</template>

<style scoped>
.value-font-size {
  font-size: 20px; /* Adjust the font size as needed */
}
</style>
