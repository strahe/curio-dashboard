<template>
  <v-navigation-drawer id="app-drawer" v-model="drawer" width="260" elevation="5" floating>
    <v-list-item class="text-center" two-line>
      <v-avatar>
        <v-img
          src="@/assets/logo.png"
          height="34"
          contain
        />
      </v-avatar>

      <v-list-item-title class="title">Curio Dashboard</v-list-item-title>
    </v-list-item>

    <v-divider class="mx-3 mb-3" />

    <v-list density="compact" nav>
      <v-list-item v-for="(link, i) in links" :key="i" :to="link.to" active-class="primary white--text">
        <template #prepend>
          <v-icon>{{ link.icon }}</v-icon>
        </template>
        <v-list-item-title>{{ link.text }}</v-list-item-title>
      </v-list-item>
    </v-list>

    <template #append>
      <v-list density="compact" nav>
        <v-list-item>
          <template #prepend>
            <v-icon>mdi-package-up</v-icon>
          </template>
          <v-list-item-title class="font-weight-light">Upgrade</v-list-item-title>
        </v-list-item>
      </v-list>
    </template>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>
import {useAppStore} from "@/stores/app";
const links = [
  {
    to: '/',
    icon: 'mdi-view-dashboard',
    text: 'Dashboard',
  },
  {
    to: '/machines',
    icon: 'mdi-server',
    text: 'Machines',
  },
  {
    to: '/pipelines',
    icon: 'mdi-pipe',
    text: 'Pipelines',
  },
  {
    to: '/tasks',
    icon: 'mdi-hexagon-multiple-outline',
    text: 'Tasks',
  },
  {
    to: '/miners',
    icon: 'mdi-account-hard-hat',
    text: 'Miners',
  },
  {
    to: '/sectors',
    icon: 'mdi-oci',
    text: 'Sectors',
  },
  {
    to: '/storages',
    icon: 'mdi-nas',
    text: 'Storages',
  },
  {
    to: '/configs',
    icon: 'mdi-cog',
    text: 'Configs',
  }
]

const drawer = computed({
  get: () => useAppStore().drawer,
  set: (val) => useAppStore().setDrawer(val),
});
</script>
