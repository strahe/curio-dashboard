<script setup lang="ts">
import { shallowRef } from 'vue'

import Avatar1 from '@/assets/images/users/avatar-1.png'
import Avatar2 from '@/assets/images/users/avatar-2.png'
import Avatar3 from '@/assets/images/users/avatar-3.png'
import Avatar4 from '@/assets/images/users/avatar-4.png'

const teams = shallowRef([
  {
    avatar: Avatar1,
    name: 'David Jones',
    time: '5 min ago',
    role: 'Project Leader',
  },
  {
    avatar: Avatar2,
    name: 'David Jones',
    time: '1 hour ago',
    role: 'HR Manager',
  },
  {
    avatar: Avatar3,
    name: 'David Jones',
    time: 'Yesterday',
    role: 'Developer',
  },
  {
    avatar: Avatar4,
    name: 'David Jones',
    time: '02-05-2021',
    role: 'UI/UX Designer',
  },
])
</script>

<template>
  <v-list aria-busy="true" aria-label="member list" class="py-3" lines="two">
    <v-list-item v-for="(team, i) in teams" :key="i" class="no-spacer" rounded="sm">
      <template #prepend>
        <v-avatar class="mr-3 py-2" size="40">
          <img alt="Julia" :src="team.avatar" width="40">
        </v-avatar>
      </template>
      <div class="d-inline-flex align-center justify-space-between w-100">
        <div>
          <h6 class="text-subtitle-1 mb-0">{{ team.name }}</h6>
          <p class="text-caption text-lightText mb-0">{{ team.role }}</p>
        </div>
        <span class="text-caption">{{ team.time }}</span>
      </div>
    </v-list-item>
  </v-list>
</template>
