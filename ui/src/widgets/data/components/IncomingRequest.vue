<script setup lang="ts">
import { shallowRef } from 'vue'

const requests = shallowRef([
  {
    color: 'success',
    name: 'Incoming requests',
  },
  {
    color: 'error',
    name: 'You have 2 pending requests..',
  },
  {
    color: 'warning',
    name: 'You have 3 pending tasks',
  },
  {
    color: 'primary',
    name: 'New order received',
  },
  {
    color: 'success',
    name: 'Incoming requests',
  },
  {
    color: 'primary',
    name: 'You have 2 pending requests..',
  },
  {
    color: 'warning',
    name: 'You have 3 pending tasks',
  },
  {
    color: 'error',
    name: 'New order received',
  },
])
</script>

<template>
  <perfect-scrollbar :style="{ height: '335px' }">
    <v-list aria-busy="true" aria-label="request list" border class="py-0">
      <v-list-item v-for="(request, i) in requests" :key="i" class="d-flex align-center no-spacer" :value="request">
        <template #prepend>
          <v-avatar class="mr-3" :color="request.color" size="12" variant="flat" />
        </template>
        <div class="d-inline-flex align-center justify-space-between w-100">
          <h6 class="text-h6 mb-0">{{ request.name }}</h6>
        </div>
      </v-list-item>
    </v-list>
  </perfect-scrollbar>
</template>
