<script setup lang="ts">
import { ref } from 'vue'
// icons
import {
    CommentOutlined,
    EditOutlined,
    LockOutlined,
    LogoutOutlined,
    ProfileOutlined,
    QuestionCircleOutlined,
    SettingOutlined,
    UnorderedListOutlined,
    UserOutlined,
    WalletOutlined,
} from '@ant-design/icons-vue'
import { useAuthStore } from '@/stores/auth'

const tab = ref(null)
const authStore = useAuthStore()
</script>

<template>
  <!-- ---------------------------------------------- -->
  <!-- profile DD -->
  <!-- ---------------------------------------------- -->
  <div>
    <div class="d-flex align-center pa-5">
      <v-avatar class="mr-2" size="32">
        <img alt="Julia" src="@/assets/images/users/avatar-1.png" width="32">
      </v-avatar>
      <div>
        <h6 class="text-h6 mb-0">JWT User</h6>
        <p class="text-caption mb-0">UI/UX Designer</p>
      </div>
      <div class="ml-auto">
        <v-btn
          color="primary"
          icon
          rounded="sm"
          size="large"
          variant="text"
          @click="authStore.logout()"
        >
          <LogoutOutlined :style="{ fontSize: '20px' }" />
        </v-btn>
      </div>
    </div>
    <v-tabs v-model="tab" color="primary" grow>
      <v-tab value="111"> <UserOutlined class="v-icon--start" /> Profile </v-tab>
      <v-tab value="222"> <SettingOutlined class="v-icon--start" /> Setting </v-tab>
    </v-tabs>
    <perfect-scrollbar style="height: calc(100vh - 300px); max-height: 240px">
      <v-window v-model="tab">
        <v-window-item value="111">
          <v-list aria-busy="true" aria-label="profile list" class="py-0">
            <v-list-item color="primary" rounded="0" value="Edit profile">
              <template #prepend>
                <EditOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Edit Profile</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="View Profile">
              <template #prepend>
                <UserOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> View Profile</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="Social Profile">
              <template #prepend>
                <ProfileOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Social Profile</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="Billing">
              <template #prepend>
                <WalletOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Billing</v-list-item-title>
            </v-list-item>

            <v-list-item color="secondary" rounded="0" @click="authStore.logout()">
              <template #prepend>
                <LogoutOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-subtitle-2"> Logout</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-window-item>
        <v-window-item value="222">
          <v-list aria-busy="true" aria-label="profile list" class="py-0">
            <v-list-item color="primary" rounded="0" value="Support">
              <template #prepend>
                <QuestionCircleOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Support</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="Account">
              <template #prepend>
                <UserOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Account settings</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="Privacy">
              <template #prepend>
                <LockOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Privacy center</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="Feedback">
              <template #prepend>
                <CommentOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> Feedback</v-list-item-title>
            </v-list-item>

            <v-list-item color="primary" rounded="0" value="History">
              <template #prepend>
                <UnorderedListOutlined class="mr-4" :style="{ fontSize: '14px' }" />
              </template>

              <v-list-item-title class="text-h6"> History</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-window-item>
      </v-window>
    </perfect-scrollbar>
  </div>
</template>
