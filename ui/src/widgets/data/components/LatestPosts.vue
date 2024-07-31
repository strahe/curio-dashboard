<script setup lang="ts">
import { shallowRef } from 'vue';

import Post1 from '@/assets/images/background/post1.jpg';
import Post3 from '@/assets/images/background/post3.jpg';

const posts = shallowRef([
  {
    avatar: Post1,
    name: 'Up unpacked friendly',
    time: '14 minutes ago'
  },
  {
    video: '668nUCeBHyY',
    name: 'Up unpacked friendly',
    time: '14 minutes ago'
  },
  {
    avatar: Post3,
    name: 'Up unpacked friendly',
    time: '14 minutes ago'
  }
]);
</script>

<template>
  <v-list class="py-1" lines="two" aria-label="post list" aria-busy="true">
    <v-list-item v-for="(post, i) in posts" :key="i">
      <template v-slot:prepend>
        <div class="mr-4">
          <v-img v-if="post.avatar" :src="post.avatar" width="90" height="80" class="rounded-md" cover :alt="post.avatar" />
          <!--If Video-->
          <div v-if="post.video" class="rounded-md overflow-hidden">
            <iframe
              :src="`https://www.youtube.com/embed/${post?.video}`"
              aria-label="video"
              frameborder="0"
              width="90"
              height="80"
            ></iframe>
          </div>
        </div>
      </template>
      <h6 class="text-subtitle-1 mb-1">{{ post.name }}</h6>
      <p class="text-caption text-lightText mb-0">Video | {{ post.time }}</p>
    </v-list-item>
  </v-list>
</template>
