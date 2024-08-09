<script setup lang="ts">
import { shallowRef } from 'vue'

// icons
import { CheckOutlined, ShoppingOutlined, TwitterOutlined, UserOutlined } from '@ant-design/icons-vue'

const messages = shallowRef([
  {
    color: 'info',
    icon: TwitterOutlined,
    time: '2 hrs ago',
    title: '+ 1652 Followers',
    subtext: 'You’re getting more and more followers, keep it up!',
  },
  {
    color: 'error',
    icon: ShoppingOutlined,
    time: '4 hrs ago',
    title: '+ 5 New Products were added!',
    subtext: 'Congratulations!',
  },
  {
    color: 'success',
    icon: CheckOutlined,
    time: '1 day ago',
    title: 'Database backup completed!',
    subtext: 'Download the',
    link: 'latest backup.',
  },
  {
    color: 'primary',
    icon: UserOutlined,
    time: '2 day ago',
    title: '+2 Friend Requests',
    subtext: 'This is great, keep it up!',
  },
])
</script>

<template>
  <v-timeline class="my-1 custom-timeline timeline-icon-circle justify-start px-5" line-color="borderLight" side="end">
    <v-timeline-item v-for="(message, i) in messages" :key="i" :dot-color="message.color" fill-dot>
      <template #icon>
        <component :is="message.icon" :style="{ fontSize: '16px' }" />
      </template>
      <template #opposite>
        <span class="text-subtitle-2 text-medium-emphasis">{{ message.time }}</span>
      </template>
      <v-card elevation="0">
        <h6 class="text-subtitle-1 mb-0">
          {{ message.title }}
        </h6>
        <span class="text-caption text-lightText">
          {{ message.subtext }}
          <router-link v-if="message.link" class="text-primary link-hover" to="/widget/data">{{ message.link }}</router-link>
        </span>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>
