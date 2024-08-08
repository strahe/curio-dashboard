<script setup lang="ts">
import { shallowRef } from 'vue'

// import icons
import { CaretDownFilled, CaretUpFilled } from '@ant-design/icons-vue'

const revenues = shallowRef([
  {
    name: 'Bitcoin',
    price: 145.58,
  },
  {
    name: 'Ethereum',
    price: 6.368,
  },
  {
    name: 'Ripple',
    price: 458.63,
  },
  {
    name: 'Neo',
    price: 5.631,
  },
  {
    name: 'Ethereum',
    price: 6.368,
  },
  {
    name: 'Ripple',
    price: 458.63,
  },
  {
    name: 'Neo',
    price: 5.631,
  },
  {
    name: 'Ethereum',
    price: 6.368,
  },
  {
    name: 'Ripple',
    price: 458.63,
  },
  {
    name: 'Neo',
    price: 5.631,
  },
])
</script>

<template>
  <perfect-scrollbar :style="{ height: '335px' }">
    <v-list aria-busy="true" aria-label="revenue list" border class="py-0">
      <v-list-item v-for="(revenue, i) in revenues" :key="i" color="primary" :value="revenue">
        <template #prepend>
          <CaretUpFilled v-if="revenue.price > 145" class="mr-3 text-success" />
          <CaretDownFilled v-else class="mr-3 text-error" />
        </template>
        <div class="d-inline-flex align-center justify-space-between w-100">
          <h6 class="text-h6 mb-0">{{ revenue.name }}</h6>
          <div v-if="revenue.price > 145" class="ml-auto text-h6 text-success">+ ${{ revenue.price }}</div>
          <div v-else class="ml-auto text-h6 text-error">- ${{ revenue.price }}</div>
        </div>
      </v-list-item>
    </v-list>
  </perfect-scrollbar>
</template>
