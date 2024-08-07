<script setup lang="ts">
import { ref, watch } from 'vue'
// icons
import { MenuFoldOutlined, SearchOutlined } from '@ant-design/icons-vue'
import { AdjustmentsHorizontalIcon, LanguageIcon } from 'vue-tabler-icons'

// dropdown imports
import LanguageDD from './LanguageDD.vue'
import NotificationDD from './NotificationDD.vue'
import ProfileDD from './ProfileDD.vue'
import Searchbar from './SearchBarPanel.vue'
import FullScreen from './FullScreen.vue'
import { useCustomizerStore } from '@/stores/customizer'

const customizer = useCustomizerStore()
const priority = ref(customizer.horizontalLayout ? 0 : 0)
watch(priority, newPriority => {
  priority.value = newPriority
})
</script>

<template>
  <v-app-bar elevation="0" height="60" :priority="priority">
    <v-btn
      class="hidden-md-and-down text-secondary mr-3"
      color="darkText"
      icon
      rounded="sm"
      size="small"
      variant="text"
      @click.stop="customizer.setMiniSidebar(!customizer.miniSidebar)"
    >
      <MenuFoldOutlined v-if="!customizer.horizontalLayout" :style="{ fontSize: '16px' }" />
    </v-btn>
    <v-btn
      class="hidden-lg-and-up text-secondary ms-3"
      color="darkText"
      icon
      rounded="sm"
      size="small"
      variant="text"
      @click.stop="customizer.setSidebarDrawer(!customizer.sidebarDrawer)"
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

    <FullScreen />

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
          <LanguageIcon />
        </v-btn>
      </template>
      <v-sheet rounded="md" width="200">
        <LanguageDD />
      </v-sheet>
    </v-menu>

    <!-- ---------------------------------------------- -->
    <!-- Customizer -->
    <!-- ---------------------------------------------- -->
    <v-menu
      :close-on-content-click="false"
      location="bottom"
      offset="6, 80"
    >
      <template #activator="{ props }">
        <v-btn
          icon
          v-bind="props"
        >
          <v-icon :icon="AdjustmentsHorizontalIcon" />
        </v-btn>
      </template>
      <v-card>
        <v-list>
          <v-list-item>
            <v-switch
              color="purple"
              hide-details
              label="Dark Mode"
              :model-value="customizer.dark"
              @click.stop="customizer.setDark(!customizer.dark)"
            />
          </v-list-item>

          <v-list-item>
            <v-switch
              color="purple"
              hide-details
              label="Horizontal Layout"
              :model-value="customizer.horizontalLayout"
              @click.stop="customizer.setHorizontalLayout(!customizer.horizontalLayout)"
            />
          </v-list-item>
        </v-list>
      </v-card>
    </v-menu>
    <NotificationDD />

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
