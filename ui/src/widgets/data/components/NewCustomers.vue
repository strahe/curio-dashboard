<script setup lang="ts">
import { shallowRef } from 'vue';

// assets
import Avatar1 from '@/assets/images/users/avatar-1.png';
import Avatar2 from '@/assets/images/users/avatar-2.png';
import Avatar3 from '@/assets/images/users/avatar-3.png';
import Avatar4 from '@/assets/images/users/avatar-4.png';
import Avatar5 from '@/assets/images/users/avatar-5.png';

const customers = shallowRef([
  {
    avatar: Avatar1,
    name: 'Alex Thompson',
    time: '5 mins ago',
    msg: 'Cheers!',
    status: 'online'
  },
  {
    avatar: Avatar2,
    name: 'John Doue',
    time: 'Today',
    msg: 'stay hungry stay foolish!',
    status: 'online'
  },
  {
    avatar: Avatar3,
    name: 'Alex Thompson',
    time: '30 min ago',
    msg: 'Cheers!',
    status: ''
  },
  {
    avatar: Avatar4,
    name: 'John Doue',
    time: '10 min ago',
    msg: 'stay hungry stay foolish!',
    status: ''
  },
  {
    avatar: Avatar5,
    name: 'Shirley Hoe',
    time: '30 min ago',
    msg: 'Cheers!',
    status: ''
  },
  {
    avatar: Avatar1,
    name: 'Alex Thompson',
    time: '5 mins ago',
    msg: 'Cheers!',
    status: 'online'
  },
  {
    avatar: Avatar2,
    name: 'John Doue',
    time: 'Today',
    msg: 'stay hungry stay foolish!',
    status: 'online'
  },
  {
    avatar: Avatar3,
    name: 'Alex Thompson',
    time: '10 min ago',
    msg: 'Cheers!',
    status: ''
  },
  {
    avatar: Avatar4,
    name: 'John Doue',
    time: '10 min ago',
    msg: 'stay hungry stay foolish!',
    status: ''
  },
  {
    avatar: Avatar5,
    name: 'Shirley Hoe',
    time: '30 min ago',
    msg: 'Cheers!',
    status: ''
  }
]);
</script>

<template>
  <perfect-scrollbar v-bind:style="{ height: '335px' }">
    <v-list class="py-0" lines="two" aria-label="customers" aria-busy="true">
      <v-list-item v-for="(customer, i) in customers" :key="i" :value="customer" rounded="sm" color="secondary" class="no-spacer">
        <template v-slot:prepend>
          <v-avatar size="40" class="mr-3 py-2">
            <img :src="customer.avatar" width="40" alt="Julia" />
          </v-avatar>
        </template>
        <div class="d-inline-flex align-center justify-space-between w-100">
          <div>
            <h6 class="text-subtitle-1 mb-0">{{ customer.name }}</h6>
            <p class="text-caption text-lightText mb-0">{{ customer.msg }}</p>
          </div>
          <v-avatar size="10" color="success" variant="flat" v-if="customer.status == 'online'"></v-avatar>
          <span class="text-subtitle-2 text-lightText" v-else>
            {{ customer.time }}
          </span>
        </div>
      </v-list-item>
    </v-list>
  </perfect-scrollbar>
</template>
