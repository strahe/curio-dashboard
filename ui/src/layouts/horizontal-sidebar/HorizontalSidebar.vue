<script setup lang="ts">
import { shallowRef } from 'vue'
import { useDisplay } from 'vuetify'
import { useCustomizerStore } from '@/stores/customizer'
import HorizontalItems from './horizontalItems'
import NavItem from './NavItem/NavItem.vue'
import NavCollapse from './NavCollapse/NavCollapse.vue'
import VerticalSidebar from '../vertical-sidebar/VerticalSidebar.vue'

const customizer = useCustomizerStore()
const sidebarMenu = shallowRef(HorizontalItems)
const { mdAndUp } = useDisplay()
</script>

<template>
  <template v-if="mdAndUp">
    <div class="horizontalMenu">
      <v-container class="py-0" fluid>
        <ul class="gap-1 horizontal-navbar px-0" :class="customizer.boxed ? 'maxWidth' : ''">
          <!---Menu Loop -->
          <li v-for="(item, i) in sidebarMenu" :key="i" class="navItem">
            <!---If Has Child -->
            <NavCollapse v-if="item.children" :item="item" :level="0" />
            <!---Single Item-->
            <NavItem v-else :item="item" />
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
