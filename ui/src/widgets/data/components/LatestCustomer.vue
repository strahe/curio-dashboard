<script setup lang="ts">
import { shallowRef } from 'vue';

import Flag1 from '@/assets/images/flags/1.jpg';
import Flag2 from '@/assets/images/flags/2.jpg';
import Flag3 from '@/assets/images/flags/3.jpg';
import Flag4 from '@/assets/images/flags/4.jpg';
import Flag5 from '@/assets/images/flags/5.jpg';

const products = shallowRef([
  {
    flag: Flag1,
    country: 'Germany',
    name: 'Angelina Jolly',
    average: '56.23'
  },
  {
    flag: Flag2,
    country: 'USA',
    name: 'John Deo',
    average: '25.23'
  },
  {
    flag: Flag3,
    country: 'Australia',
    name: 'Jenifer Vintage',
    average: '12.45'
  },
  {
    flag: Flag4,
    country: 'United Kingdom',
    name: 'Lori Moore',
    average: '8.65'
  },
  {
    flag: Flag5,
    country: 'Brazil',
    name: 'Allianz Dacron',
    average: '3.56'
  },
  {
    flag: Flag1,
    country: 'Australia',
    name: 'Jenifer Vintage',
    average: '12.45'
  },
  {
    flag: Flag3,
    country: 'USA',
    name: 'John Deo',
    average: '25.23'
  },
  {
    flag: Flag5,
    country: 'Australia',
    name: 'Jenifer Vintage',
    average: '12.45'
  },
  {
    flag: Flag2,
    country: 'United Kingdom',
    name: 'Lori Moore',
    average: '8.65'
  }
]);
</script>

<template>
  <perfect-scrollbar v-bind:style="{ height: '267px' }">
    <v-table class="text-no-wrap bordered-table" hover density="comfortable">
      <thead class="bg-containerBg">
        <tr>
          <th class="text-left text-caption font-weight-bold text-uppercase">#</th>
          <th class="text-left text-caption font-weight-bold text-uppercase">Country</th>
          <th class="text-left text-caption font-weight-bold text-uppercase">Name</th>
          <th class="text-right text-caption font-weight-bold text-uppercase">Average</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in products" :key="item.name">
          <td>
            <img :src="item.flag" width="30" :alt="item.flag" />
          </td>
          <td>{{ item.country }}</td>
          <td>{{ item.name }}</td>
          <td class="text-right">{{ item.average }}%</td>
        </tr>
      </tbody>
    </v-table>
  </perfect-scrollbar>
</template>
