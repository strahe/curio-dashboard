<script setup lang="ts">
import { shallowRef } from 'vue';

// import icons
import { CaretUpFilled, CaretDownFilled } from '@ant-design/icons-vue';

const revenues = shallowRef([
  {
    name: 'Bitcoin',
    price: 145.58
  },
  {
    name: 'Ethereum',
    price: 6.368
  },
  {
    name: 'Ripple',
    price: 458.63
  },
  {
    name: 'Neo',
    price: 5.631
  },
  {
    name: 'Ethereum',
    price: 6.368
  },
  {
    name: 'Ripple',
    price: 458.63
  },
  {
    name: 'Neo',
    price: 5.631
  },
  {
    name: 'Ethereum',
    price: 6.368
  },
  {
    name: 'Ripple',
    price: 458.63
  },
  {
    name: 'Neo',
    price: 5.631
  }
]);
</script>

<template>
  <perfect-scrollbar v-bind:style="{ height: '335px' }">
    <v-list class="py-0" aria-label="revenue list" aria-busy="true" border>
      <v-list-item v-for="(revenue, i) in revenues" :key="i" :value="revenue" color="primary">
        <template v-slot:prepend>
          <CaretUpFilled class="mr-3 text-success" v-if="revenue.price > 145" />
          <CaretDownFilled class="mr-3 text-error" v-else />
        </template>
        <div class="d-inline-flex align-center justify-space-between w-100">
          <h6 class="text-h6 mb-0">{{ revenue.name }}</h6>
          <div class="ml-auto text-h6 text-success" v-if="revenue.price > 145">+ ${{ revenue.price }}</div>
          <div class="ml-auto text-h6 text-error" v-else>- ${{ revenue.price }}</div>
        </div>
      </v-list-item>
    </v-list>
  </perfect-scrollbar>
</template>
