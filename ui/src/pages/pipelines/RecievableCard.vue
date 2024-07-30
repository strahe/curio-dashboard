<script setup lang="ts">
import { FileDoneOutlined, InfoCircleOutlined } from '@ant-design/icons-vue';
</script>
<template>
  <v-card variant="flat" class="recievable-card">
    <v-card-item>
      <div class="d-flex align-end">
        <div class="d-flex">
          <v-avatar size="40" color="primary" rounded>
            <FileDoneOutlined :style="{ fontSize: '20px' }" />
          </v-avatar>
          <div class="text-white ml-2">
            <p class="text-h6 mb-0">
              Total Recievables
              <InfoCircleOutlined class="ml-1" />
            </p>
            <div class="d-flex">
              <p class="text-caption mb-0">Current</p>
              <p class="text-h6 ml-2 mb-0">109.1k</p>
            </div>
          </div>
        </div>
        <div class="d-flex text-white ml-auto">
          <p class="text-caption mb-0">Overdue</p>
          <p class="text-h6 ml-2 mb-0">62k</p>
        </div>
      </div>
      <h4 class="text-h4 text-white mt-3 mb-0">$43,078</h4>
      <div class="d-flex align-center">
        <v-progress-linear
          aria-label="progressbar"
          color="warning"
          model-value="90"
          bg-opacity="0.4"
          height="6"
          rounded
        ></v-progress-linear>
        <span class="text-caption text-white ml-2 d-block">90%</span>
      </div>
    </v-card-item>
  </v-card>
</template>
<style lang="scss">
.recievable-card {
  background: var(--v-gradient2);
}
</style>
