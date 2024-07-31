<script setup lang="ts">
import { shallowRef } from 'vue';

import Avatar1 from '@/assets/images/users/avatar-1.png';
import Avatar2 from '@/assets/images/users/avatar-2.png';
import Avatar3 from '@/assets/images/users/avatar-3.png';
import Avatar4 from '@/assets/images/users/avatar-4.png';

const products = shallowRef([
  {
    avatar: Avatar1,
    name: 'John Deo',
    time: 12,
    title: '[#1183] Workaround for OS X selects printing bug',
    subtext: 'Chrome fixed the bug several versions ago, thus rendering this...'
  },
  {
    avatar: Avatar2,
    name: 'Jems Win',
    time: 16,
    title: '[#1249] Vertically center carousel controls',
    subtext: 'Try any carousel control and reduce the screen width below...'
  },
  {
    avatar: Avatar3,
    name: 'Jeny Wiliiam',
    time: 40,
    title: '[#1254] Inaccurate small pagination height',
    subtext: 'The height of pagination elements is not consistent with...'
  },
  {
    avatar: Avatar4,
    name: 'Jems Win',
    time: 12,
    title: '[#1249] Vertically center carousel controls',
    subtext: 'Try any carousel control and reduce the screen width below...'
  }
]);
</script>

<template>
  <v-table class="bordered-table" hover density="comfortable">
    <thead class="bg-containerBg">
      <tr>
        <th class="text-center text-caption font-weight-bold text-uppercase">Due</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Name</th>
        <th class="text-left text-caption font-weight-bold text-uppercase" style="min-width: 356px">Position</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in products" :key="item.time">
        <td class="py-3 text-center">
          <h6 class="text-subtitle-1 mb-0">{{ item.time }}</h6>
          <span class="text-lightText text-subtitle-2">hours</span>
        </td>
        <td class="py-3">
          <div class="d-flex align-center">
            <img :src="item.avatar" width="40" alt="Julia" class="rounded-circle" />
            <h5 class="text-subtitle-1 ml-4 mb-0">{{ item.name }}</h5>
          </div>
        </td>
        <td class="py-3" style="min-width: 356px">
          <h5 class="text-subtitle-1 mb-0">{{ item.title }}</h5>
          <span class="text-lightText text-caption">{{ item.subtext }}</span>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>
