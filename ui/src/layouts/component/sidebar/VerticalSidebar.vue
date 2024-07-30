<script setup lang="ts">
import { ref, shallowRef, computed } from 'vue';
import sidebarItems from './sidebarItem';

// icons
import { SearchOutlined } from '@ant-design/icons-vue';

import NavTitle from './NavTitle.vue';
import NavItem from './NavItem.vue';

const originalSidebarMenu = shallowRef(sidebarItems);
const searchValue = ref('');

const sidebarMenu = computed(() => {
  return originalSidebarMenu.value.filter((item) => {
    // Check if item and item.title are defined
    if (item && item.title) {
      return item.title.toLowerCase().includes(searchValue.value.toLowerCase());
    } else if (item && item.header) {
      return item.header.toLowerCase().includes(searchValue.value.toLowerCase());
    }
    // If item or item.title is undefined, exclude it from the filtered list
    return false;
  });
});
</script>

<template>
  <v-card variant="outlined" class="bg-surface br-0 position-sticky componentSidebar">
    <v-card-text class="pt-5 px-0">
      <div class="px-5">
        <v-text-field single-line variant="outlined" placeholder="Search Components" v-model="searchValue">
          <template v-slot:prepend-inner>
            <SearchOutlined :style="{ fontSize: '14px' }" />
          </template>
        </v-text-field>
      </div>
      <!-- ---------------------------------------------- -->
      <!---Navigation -->
      <!-- ---------------------------------------------- -->
      <perfect-scrollbar class="scrollnavbar">
        <v-list aria-busy="true" aria-label="menu list">
          <!---Menu Loop -->
          <template v-for="(item, i) in sidebarMenu" :key="i">
            <NavTitle :item="item" v-if="item && item.header" :key="item.title" />
            <NavItem :item="item" v-else />
          </template>
        </v-list>
      </perfect-scrollbar>
    </v-card-text>
  </v-card>
</template>
<style lang="scss">
.componentSidebar {
  @media (min-width: 1280px) {
    top: 100px;
  }
}
</style>
