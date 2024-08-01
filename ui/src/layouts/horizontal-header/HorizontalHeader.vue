<script setup lang="ts">
import { ref, watch } from 'vue'
// icons
import { MenuFoldOutlined, SearchOutlined, SettingOutlined, TranslationOutlined, WindowsOutlined } from '@ant-design/icons-vue'
import Logo from '../logo/LogoMain.vue'
// dropdown imports
import LanguageDD from '../vertical-header/LanguageDD.vue'
import NotificationDD from '../vertical-header/NotificationDD.vue'
import ProfileDD from '../vertical-header/ProfileDD.vue'
import MegaMenuDD from '../vertical-header/MegaMenuDD.vue'
import Searchbar from '../vertical-header/SearchBarPanel.vue'
import MessageDD from '../vertical-header/MessageDD.vue'
import FullScreen from '../vertical-header/FullScreen.vue'
import { useCustomizerStore } from '@/stores/customizer'

const customizer = useCustomizerStore()
const priority = ref(customizer.setHorizontalLayout ? 0 : 0)
watch(priority, newPriority => {
  // yes, console.log() is a side effect
  priority.value = newPriority
})
</script>

<template>
  <v-app-bar elevation="0" height="60" :priority="priority">
    <div class="pa-5 hidden-md-and-down">
      <Logo />
    </div>
    <v-btn
      class="hidden-lg-and-up text-secondary ms-3"
      color="darkText"
      icon
      rounded="sm"
      size="small"
      variant="text"
      @click.stop="customizer.SET_SIDEBAR_DRAWER"
    >
      <MenuFoldOutlined :style="{ fontSize: '16px' }" />
    </v-btn>

    <!-- search mobile -->
    <v-menu class="hidden-lg-and-up" :close-on-content-click="false" offset="10, 0">
      <template #activator="{ props }">
        <v-btn
          class="hidden-lg-and-up text-secondary ml-1"
          color="lightsecondary"
          icon
          rounded="sm"
          size="small"
          variant="flat"
          v-bind="props"
        >
          <SearchOutlined :style="{ fontSize: '17px' }" />
        </v-btn>
      </template>
      <v-sheet class="search-sheet v-col-12 pa-0" width="320">
        <v-text-field
          color="primary"
          hide-details
          persistent-placeholder
          placeholder="Search here.."
          variant="solo"
        >
          <template #prepend-inner>
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
      <template #activator="{ props }">
        <v-btn
          class="text-secondary hidden-sm-and-down d-lg-block d-none"
          color="darkText"
          icon
          rounded="sm"
          size="small"
          variant="text"
          v-bind="props"
        >
          <WindowsOutlined :style="{ fontSize: '16px' }" />
        </v-btn>
      </template>
      <v-sheet class="d-lg-block d-none" height="325" rounded="md" width="1024">
        <MegaMenuDD />
      </v-sheet>
    </v-menu>
    <!-- ---------------------------------------------- -->
    <!-- translate -->
    <!-- ---------------------------------------------- -->
    <v-menu :close-on-content-click="false" location="bottom" offset="6, 80">
      <template #activator="{ props }">
        <v-btn
          class="ml-sm-2 ml-1"
          color="darkText"
          icon
          rounded="sm"
          size="small"
          v-bind="props"
        >
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
      color="darkText"
      icon
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
      <template #activator="{ props }">
        <v-btn class="profileBtn" rounded="sm" variant="text" v-bind="props">
          <div class="d-flex align-center">
            <v-avatar class="mr-sm-2 mr-0 py-2">
              <img alt="Julia" src="@/assets/images/users/avatar-1.png">
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
