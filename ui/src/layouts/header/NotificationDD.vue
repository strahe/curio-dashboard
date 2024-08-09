<script setup lang="ts">
import { ref } from 'vue'
// icons
import { CheckCircleOutlined } from '@ant-design/icons-vue'
import { BellIcon } from 'vue-tabler-icons'

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
        :icon="true"
        rounded="sm"
        size="small"
        v-bind="props"
      >
        <v-badge color="primary" :content="isActive ? '2' : '0'" offset-x="-4" offset-y="-5">
          <BellIcon />
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
