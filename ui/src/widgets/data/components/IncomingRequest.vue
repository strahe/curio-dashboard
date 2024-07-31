<script setup lang="ts">
import { shallowRef } from 'vue';

const requests = shallowRef([
  {
    color: 'success',
    name: 'Incoming requests'
  },
  {
    color: 'error',
    name: 'You have 2 pending requests..'
  },
  {
    color: 'warning',
    name: 'You have 3 pending tasks'
  },
  {
    color: 'primary',
    name: 'New order received'
  },
  {
    color: 'success',
    name: 'Incoming requests'
  },
  {
    color: 'primary',
    name: 'You have 2 pending requests..'
  },
  {
    color: 'warning',
    name: 'You have 3 pending tasks'
  },
  {
    color: 'error',
    name: 'New order received'
  }
]);
</script>

<template>
  <perfect-scrollbar v-bind:style="{ height: '335px' }">
    <v-list class="py-0" aria-label="request list" aria-busy="true" border>
      <v-list-item class="d-flex align-center no-spacer" v-for="(request, i) in requests" :key="i" :value="request">
        <template v-slot:prepend>
          <v-avatar size="12" :color="request.color" variant="flat" class="mr-3"> </v-avatar>
        </template>
        <div class="d-inline-flex align-center justify-space-between w-100">
          <h6 class="text-h6 mb-0">{{ request.name }}</h6>
        </div>
      </v-list-item>
    </v-list>
  </perfect-scrollbar>
</template>
