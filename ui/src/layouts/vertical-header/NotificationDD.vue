<script setup lang="ts">
import { ref } from 'vue'
// icons
import { BellOutlined, CheckCircleOutlined, GiftOutlined, MessageOutlined, SettingOutlined } from '@ant-design/icons-vue'

const isActive = ref(true)

function deactivateItem () {
    isActive.value = false
}
</script>

<template>
  <!-- ---------------------------------------------- -->
  <!-- notifications DD -->
  <!-- ---------------------------------------------- -->
  <v-menu :close-on-content-click="false" offset="6, 0">
    <template #activator="{ props }">
      <v-btn
        class="text-secondary ml-sm-2 ml-1"
        color="darkText"
        icon
        rounded="sm"
        size="small"
        v-bind="props"
      >
        <v-badge color="primary" :content="isActive ? '2' : '0'" offset-x="-4" offset-y="-5">
          <BellOutlined :style="{ fontSize: '16px' }" />
        </v-badge>
      </v-btn>
    </template>
    <v-sheet class="notification-dropdown" rounded="md" width="387">
      <div class="pa-4">
        <div class="d-flex align-center justify-space-between">
          <h6 class="text-subtitle-1 mb-0">Notifications</h6>
          <v-btn
            :class="isActive ? 'd-block' : 'd-none'"
            color="success"
            icon
            rounded
            size="small"
            variant="text"
            @click="deactivateItem()"
          >
            <CheckCircleOutlined :style="{ fontSize: '16px' }" />
            <v-tooltip activator="parent" aria-label="tooltip" :content-class="isActive ? 'custom-tooltip' : 'd-none'" location="bottom">
              <span class="text-caption">Mark as all read</span>
            </v-tooltip>
          </v-btn>
        </div>
      </div>
      <v-divider />
      <perfect-scrollbar style="height: calc(100vh - 300px); max-height: 265px">
        <v-list aria-busy="true" aria-label="notification list" class="py-0" lines="two">
          <v-list-item :active="isActive" class="no-spacer py-1" color="secondary" value="1">
            <template #prepend>
              <v-avatar class="mr-3 py-2 text-success" color="lightsuccess" size="36" variant="flat">
                <GiftOutlined />
              </v-avatar>
            </template>
            <div class="d-inline-flex justify-space-between w-100">
              <h6 class="text-subtitle-1 font-weight-regular mb-0">
                It's <span style="font-weight: 600">Cristina danny's</span> birthday today.
              </h6>
              <span class="text-caption">3:00 AM</span>
            </div>

            <p class="text-caption text-medium-emphasis my-0">2 min ago</p>
          </v-list-item>
          <v-divider />
          <v-list-item class="no-spacer" color="secondary" value="2">
            <template #prepend>
              <v-avatar class="mr-3 py-2 text-primary" color="lightprimary" size="36" variant="flat">
                <MessageOutlined />
              </v-avatar>
            </template>
            <div class="d-inline-flex justify-space-between w-100">
              <h6 class="text-subtitle-1 font-weight-regular mb-0"><span style="font-weight: 600">Aida Burg</span> commented your post.</h6>
              <span class="text-caption">6:00 PM</span>
            </div>

            <p class="text-caption text-medium-emphasis my-0">5 August</p>
          </v-list-item>
          <v-divider />
          <v-list-item :active="isActive" class="no-spacer" color="secondary" value="3">
            <template #prepend>
              <v-avatar class="mr-3 py-2 text-error" color="lighterror" size="36" variant="flat">
                <SettingOutlined />
              </v-avatar>
            </template>
            <div class="d-inline-flex justify-space-between w-100">
              <h6 class="text-subtitle-1 font-weight-regular mb-0">Your Profile is Complete <span style="font-weight: 600">60%</span></h6>
              <span class="text-caption">2:45 PM</span>
            </div>

            <p class="text-caption text-medium-emphasis my-0">7 hours ago</p>
          </v-list-item>
          <v-divider />
          <v-list-item class="no-spacer" color="secondary" value="4">
            <template #prepend>
              <v-avatar class="mr-3 py-2 text-primary" color="lightprimary" size="36" variant="flat"> C </v-avatar>
            </template>
            <div class="d-inline-flex justify-space-between w-100">
              <h6 class="text-subtitle-1 font-weight-regular mb-0">
                <span style="font-weight: 600">Cristina Danny</span> invited to join <span style="font-weight: 600">Metting.</span>
              </h6>
              <span class="text-caption">9:10 PM</span>
            </div>

            <p class="text-caption text-medium-emphasis my-0">Daily scrum meeting time</p>
          </v-list-item>
        </v-list>
      </perfect-scrollbar>
      <v-divider />
      <div class="pa-2 text-center">
        <v-btn color="primary" variant="text">View All</v-btn>
      </div>
    </v-sheet>
  </v-menu>
</template>

<style lang="scss">
.v-tooltip {
  > .v-overlay__content.custom-tooltip {
    padding: 2px 6px;
  }
}
</style>
