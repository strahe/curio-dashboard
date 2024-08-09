<script setup lang="ts">
import { shallowRef } from 'vue'

// import icons
import { FileTextFilled, MessageFilled, ShoppingFilled } from '@ant-design/icons-vue'

const feeds = shallowRef([
  {
    color: 'primary',
    icon: MessageFilled,
    name: 'You have 3 pending tasks.',
    time: '5 mins ago',
  },
  {
    color: 'error',
    icon: ShoppingFilled,
    name: 'New order received',
    time: '1 day ago',
  },
  {
    color: 'success',
    icon: FileTextFilled,
    name: 'You have 3 pending tasks.',
    time: '3 week ago',
  },
  {
    color: 'primary',
    icon: MessageFilled,
    name: 'New order received',
    time: 'around month',
  },
  {
    color: 'warning',
    icon: ShoppingFilled,
    name: 'Order cancelled',
    time: '2 month ago',
  },
])
</script>

<template>
  <v-list
    aria-busy="true"
    aria-label="feeds"
    class="py-0"
    density="compact"
    lines="two"
  >
    <v-list-item v-for="(feed, i) in feeds" :key="i" class="no-spacer">
      <template #prepend>
        <v-avatar class="mr-3 py-2" :color="feed.color" size="32" variant="flat">
          <component :is="feed.icon" class="text-white" />
        </v-avatar>
      </template>
      <div class="d-inline-flex align-center justify-space-between w-100">
        <p class="text-caption mb-0">{{ feed.name }}</p>
        <p class="text-caption text-lightText ml-auto mb-0">{{ feed.time }}</p>
      </div>
    </v-list-item>
  </v-list>
</template>
