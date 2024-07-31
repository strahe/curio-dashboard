<script setup lang="ts">
import { shallowRef } from 'vue';

// import icons
import { MessageFilled, ShoppingFilled, FileTextFilled } from '@ant-design/icons-vue';

const feeds = shallowRef([
  {
    color: 'primary',
    icon: MessageFilled,
    name: 'You have 3 pending tasks.',
    time: '5 mins ago'
  },
  {
    color: 'error',
    icon: ShoppingFilled,
    name: 'New order received',
    time: '1 day ago'
  },
  {
    color: 'success',
    icon: FileTextFilled,
    name: 'You have 3 pending tasks.',
    time: '3 week ago'
  },
  {
    color: 'primary',
    icon: MessageFilled,
    name: 'New order received',
    time: 'around month'
  },
  {
    color: 'warning',
    icon: ShoppingFilled,
    name: 'Order cancelled',
    time: '2 month ago'
  }
]);
</script>

<template>
  <v-list class="py-0" lines="two" density="compact" aria-label="feeds" aria-busy="true">
    <v-list-item v-for="(feed, i) in feeds" :key="i" class="no-spacer">
      <template v-slot:prepend>
        <v-avatar size="32" :color="feed.color" variant="flat" class="mr-3 py-2">
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
