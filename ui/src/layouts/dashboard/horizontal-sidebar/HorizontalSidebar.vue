<script setup lang="ts">
import { shallowRef } from 'vue';
import { useDisplay } from 'vuetify';
import { useCustomizerStore } from '@/stores/customizer';
import HorizontalItems from './horizontalItems';
import NavItem from './NavItem/NavItem.vue';
import NavCollapse from './NavCollapse/NavCollapse.vue';
import VerticalSidebar from '../vertical-sidebar/VerticalSidebar.vue';

const customizer = useCustomizerStore();
const sidebarMenu = shallowRef(HorizontalItems);
const { mdAndUp } = useDisplay();
</script>

<template>
  <template v-if="mdAndUp">
    <div class="horizontalMenu">
      <v-container fluid class="py-0">
        <ul class="gap-1 horizontal-navbar px-0" :class="customizer.boxed ? 'maxWidth' : ''">
          <!---Menu Loop -->
          <li v-for="(item, i) in sidebarMenu" :key="i" class="navItem">
            <!---If Has Child -->
            <NavCollapse :item="item" :level="0" v-if="item.children" />
            <!---Single Item-->
            <NavItem :item="item" v-else />
            <!---End Single Item-->
          </li>
        </ul>
      </v-container>
    </div>
  </template>
  <div v-else class="mobile-menu">
    <VerticalSidebar />
  </div>
</template>
<style lang="scss"></style>
