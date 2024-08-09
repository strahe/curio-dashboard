<script setup lang="ts">
import { shallowRef } from 'vue'

// icons
import { BugFilled, ClockCircleFilled, MobileFilled, TwitterCircleFilled, WarningFilled } from '@ant-design/icons-vue'

const messages = shallowRef([
  {
    color: 'success',
    icon: TwitterCircleFilled,
    title: 'You’re getting more and more followers, keep it up!',
    time: '8:50',
  },
  {
    color: 'primary',
    icon: ClockCircleFilled,
    title: 'Design mobile Application',
    time: 'Sat, 5 Mar',
  },
  {
    color: 'error',
    icon: BugFilled,
    name: 'Jenny',
    title: ' assign you a task ',
    link: 'Mockup Design.',
    time: 'Sun, 17 Feb',
  },
  {
    color: 'warning',
    icon: WarningFilled,
    title: 'Design logo',
    time: 'Sat, 18 Mar ',
  },
  {
    color: 'success',
    icon: MobileFilled,
    title: 'Design mobile Application',
    time: 'Sat, 22 Mar',
  },
])
</script>

<template>
  <v-timeline class="my-1 justify-start oppositeRemove px-5" line-color="borderLight" side="end">
    <v-timeline-item
      v-for="(message, i) in messages"
      :key="i"
      :dot-color="message.color"
      fill-dot
      size="small"
    >
      <template #icon>
        <component :is="message.icon" class="text-white" />
      </template>
      <v-card elevation="0" rounded="0">
        <span class="text-caption text-lightText">{{ message.time }}</span>
        <p class="text-caption text-wrap mb-0">
          <router-link v-if="message.name" class="text-primary link-hover" to="/widget/data">{{ message.name }}</router-link>
          {{ message.title }}
          <router-link v-if="message.link" class="text-primary link-hover" to="/widget/data">{{ message.link }}</router-link>
        </p>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>
<style lang="scss">
.oppositeRemove {
  .v-timeline-item__opposite {
    display: none;
  }
  &.v-timeline--vertical.v-timeline.v-timeline--side-end {
    .v-timeline-item {
      .v-timeline-item__body {
        padding-inline-start: 16px;
      }
    }
  }
}
</style>
