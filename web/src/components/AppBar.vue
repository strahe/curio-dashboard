<template>
  <v-app-bar id="core-app-bar" absolute color="transparent" flat height="88">
    <v-toolbar-title class="font-weight-light align-self-center">
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-btn v-show="!responsive" icon="mdi-view-list" @click.stop="onClick">
      </v-btn>
      {{ title }}
    </v-toolbar-title>
    <v-spacer />
    <v-toolbar-items class="flex-fill">
      <v-row align="center" class="mx-0">
        <v-text-field class="mr-4 purple-input" color="purple" label="Search..." hide-details />

        <v-btn height="48" to="/" icon="mdi-view-dashboard" color="tertiary">
        </v-btn>
        <v-btn height="48" icon="mdi-theme-light-dark" @click="toggleTheme" color="tertiary">
        </v-btn>
      </v-row>
    </v-toolbar-items>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useDisplay, useTheme } from 'vuetify';
import {useAppStore} from "@/stores/app";
import {useCookieTheme} from "@/composables/useCookieTheme";
import { storeToRefs } from 'pinia'

const theme = useTheme();
const route = useRoute();
const cookieTheme = useCookieTheme();
const appStore = useAppStore()
const { drawer } = storeToRefs(appStore)

const toggleTheme = () => {
  const themeValue = theme.global.current.value.dark ? 'light' : 'dark';
  theme.global.name.value = themeValue;
  cookieTheme.set('theme', themeValue);
};

const title = ref("Title");

const responsive = computed(() => {
  const display = useDisplay();
  return display.lgAndUp.value;
});

watch(route, (val) => {
  title.value = val.meta.title as string || val.name || "Title";
});

onMounted(() => {
  appStore.setDrawer(responsive.value);
  title.value = route.name;
});

const onClick = () => {
  appStore.setDrawer(!appStore.drawer);
};
</script>
<style>
/* Fix coming in v2.0.8 */
#core-app-bar {
  width: auto;
}

#core-app-bar a {
  text-decoration: none;
}
</style>
