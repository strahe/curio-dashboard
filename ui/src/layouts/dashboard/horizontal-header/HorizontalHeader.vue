<script setup lang="ts">
import { ref, watch } from 'vue';
import { useCustomizerStore } from '../../../stores/customizer';
// icons
import { MenuFoldOutlined, WindowsOutlined, TranslationOutlined, SettingOutlined, SearchOutlined } from '@ant-design/icons-vue';
import Logo from '../logo/LogoMain.vue';
// dropdown imports
import LanguageDD from '../vertical-header/LanguageDD.vue';
import NotificationDD from '../vertical-header/NotificationDD.vue';
import ProfileDD from '../vertical-header/ProfileDD.vue';
import MegaMenuDD from '../vertical-header/MegaMenuDD.vue';
import Searchbar from '../vertical-header/SearchBarPanel.vue';
import MessageDD from '../vertical-header/MessageDD.vue';
import FullScreen from '../vertical-header/FullScreen.vue';

const customizer = useCustomizerStore();
const priority = ref(customizer.setHorizontalLayout ? 0 : 0);
watch(priority, (newPriority) => {
  // yes, console.log() is a side effect
  priority.value = newPriority;
});
</script>

<template>
  <v-app-bar elevation="0" :priority="priority" height="60">
    <div class="pa-5 hidden-md-and-down">
      <Logo />
    </div>
    <v-btn
      class="hidden-lg-and-up text-secondary ms-3"
      color="darkText"
      icon
      rounded="sm"
      variant="text"
      @click.stop="customizer.SET_SIDEBAR_DRAWER"
      size="small"
    >
      <MenuFoldOutlined :style="{ fontSize: '16px' }" />
    </v-btn>

    <!-- search mobile -->
    <v-menu :close-on-content-click="false" class="hidden-lg-and-up" offset="10, 0">
      <template v-slot:activator="{ props }">
        <v-btn
          class="hidden-lg-and-up text-secondary ml-1"
          color="lightsecondary"
          icon
          rounded="sm"
          variant="flat"
          size="small"
          v-bind="props"
        >
          <SearchOutlined :style="{ fontSize: '17px' }" />
        </v-btn>
      </template>
      <v-sheet class="search-sheet v-col-12 pa-0" width="320">
        <v-text-field persistent-placeholder placeholder="Search here.." color="primary" variant="solo" hide-details>
          <template v-slot:prepend-inner>
            <SearchOutlined :style="{ fontSize: '17px' }" />
          </template>
        </v-text-field>
      </v-sheet>
    </v-menu>

    <!-- ---------------------------------------------- -->
    <!-- Search part -->
    <!-- ---------------------------------------------- -->
    <v-sheet class="d-none d-lg-block" width="250">
      <Searchbar />
    </v-sheet>

    <!---/Search part -->

    <v-spacer />
    <!-- ---------------------------------------------- -->
    <!---right part -->
    <!-- ---------------------------------------------- -->

    <!-- ---------------------------------------------- -->
    <!-- Messages -->
    <!-- ---------------------------------------------- -->
    <v-menu :close-on-content-click="false" offset="10, 320">
      <template v-slot:activator="{ props }">
        <v-btn
          icon
          class="text-secondary hidden-sm-and-down d-lg-block d-none"
          color="darkText"
          rounded="sm"
          size="small"
          variant="text"
          v-bind="props"
        >
          <WindowsOutlined :style="{ fontSize: '16px' }" />
        </v-btn>
      </template>
      <v-sheet width="1024" height="325" rounded="md" class="d-lg-block d-none">
        <MegaMenuDD />
      </v-sheet>
    </v-menu>
    <!-- ---------------------------------------------- -->
    <!-- translate -->
    <!-- ---------------------------------------------- -->
    <v-menu :close-on-content-click="false" location="bottom" offset="6, 80">
      <template v-slot:activator="{ props }">
        <v-btn icon class="ml-sm-2 ml-1" color="darkText" rounded="sm" size="small" v-bind="props">
          <TranslationOutlined :style="{ fontSize: '16px' }" />
        </v-btn>
      </template>
      <v-sheet rounded="md" width="200">
        <LanguageDD />
      </v-sheet>
    </v-menu>

    <!-- ---------------------------------------------- -->
    <!-- Notification -->
    <!-- ---------------------------------------------- -->
    <NotificationDD />

    <!-- ---------------------------------------------- -->
    <!-- Message -->
    <!-- ---------------------------------------------- -->
    <MessageDD />

    <!-- ---------------------------------------------- -->
    <!-- Fullscreen -->
    <!-- ---------------------------------------------- -->
    <FullScreen />

    <!-- ---------------------------------------------- -->
    <!-- Customizer -->
    <!-- ---------------------------------------------- -->
    <v-btn
      class="customizer-btn ml-sm-2 ml-1"
      icon
      color="darkText"
      rounded="sm"
      size="small"
      variant="text"
      @click.stop="customizer.SET_CUSTOMIZER_DRAWER(!customizer.Customizer_drawer)"
    >
      <SettingOutlined class="icon" :style="{ fontSize: '16px' }" />
    </v-btn>

    <!-- ---------------------------------------------- -->
    <!-- User Profile -->
    <!-- ---------------------------------------------- -->
    <v-menu :close-on-content-click="false" offset="8, 0">
      <template v-slot:activator="{ props }">
        <v-btn class="profileBtn" variant="text" rounded="sm" v-bind="props">
          <div class="d-flex align-center">
            <v-avatar class="mr-sm-2 mr-0 py-2">
              <img src="@/assets/images/users/avatar-1.png" alt="Julia" />
            </v-avatar>
            <h6 class="text-subtitle-1 mb-0 d-sm-block d-none">JWT User</h6>
          </div>
        </v-btn>
      </template>
      <v-sheet rounded="md" width="290">
        <ProfileDD />
      </v-sheet>
    </v-menu>
  </v-app-bar>
</template>
