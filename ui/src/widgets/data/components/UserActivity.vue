<script setup lang="ts">
import { shallowRef } from 'vue';
// assets
import Avatar1 from '@/assets/images/users/avatar-5.png';
import Avatar2 from '@/assets/images/users/avatar-6.png';
import Avatar3 from '@/assets/images/users/avatar-7.png';

// icons
import { ClockCircleOutlined } from '@ant-design/icons-vue';

const activities = shallowRef([
  {
    avatar: Avatar1,
    color: 'error',
    name: 'John Deo',
    time: 'now',
    subtext: 'Lorem Ipsum is simply dummy text.'
  },
  {
    avatar: Avatar2,
    color: 'success',
    name: 'John Deo',
    time: '2 min ago',
    subtext: 'Lorem Ipsum is simply dummy text.'
  },
  {
    avatar: Avatar3,
    color: 'primary',
    name: 'John Deo',
    time: '1 day ago',
    subtext: 'Lorem Ipsum is simply dummy text.'
  },
  {
    avatar: Avatar1,
    color: 'warning',
    name: 'John Deo',
    time: '3 week ago',
    subtext: 'Lorem Ipsum is simply dummy text.'
  }
]);
</script>

<template>
  <v-list class="py-2" lines="two" aria-label="user list" aria-busy="true">
    <v-list-item v-for="(activity, i) in activities" :key="i" rounded="sm">
      <template v-slot:prepend>
        <div class="mr-3">
          <v-badge dot :color="activity.color" location="bottom end">
            <v-avatar size="40" class="py-2">
              <img :src="activity.avatar" width="40" alt="Julia" />
            </v-avatar>
          </v-badge>
        </div>
      </template>
      <div class="d-inline-flex align-center justify-space-between w-100">
        <h6 class="text-subtitle-1 mb-1">{{ activity.name }}</h6>
        <span class="text-caption d-flex align-center">
          <span class="text-lightText">{{ activity.time }}</span>
          <ClockCircleOutlined :style="{ fontSize: '11px' }" class="ml-1" />
        </span>
      </div>

      <p class="text-caption text-lightText mb-0">{{ activity.subtext }}</p>
    </v-list-item>
  </v-list>
</template>
