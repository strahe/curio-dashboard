<script setup lang="ts">
import { ref } from 'vue'
// assets
import imageReader from '@/assets/images/analytics/reader.svg'
// icons
import { PlusOutlined } from '@ant-design/icons-vue'

const p1 = ref(35)
</script>
<template>
  <v-card class="overflow-hidden bg-surface" elevation="0" variant="outlined">
    <v-card-text class="pa-0">
      <v-row class="ma-0">
        <v-col class="pa-0" cols="12" sm="7">
          <div class="bg-primary pa-5 position-relative h-100 overflow-hidden">
            <h5 class="text-h5">What would you want to learn today</h5>
            <p class="text-caption" style="max-width: 55%">Your learning capacity is 80% as daily analytics</p>
            <h4 class="text-h4 pt-sm-13 pt-8">35% Completed</h4>
            <div class="pb-7" style="max-width: 60%">
              <v-progress-linear
                v-model="p1"
                aria-label="progressbar"
                bg-color="lightsuccess"
                bg-opacity="1"
                color="success"
                height="6"
                rounded
              />
            </div>
            <v-img alt="reader" class="readMedia" cover :src="imageReader" />
          </div>
        </v-col>
        <v-col class="pa-0" cols="12" sm="5">
          <div class="pa-5 h-100">
            <h6 class="text-h6">Get started with new basic skills</h6>
            <p class="text-lightText text-h6 mb-0">Last Date 5th Nov 2020</p>
            <v-divider class="my-6" />
            <div class="d-flex justify-space-between">
              <div class="d-flex align-center justify-center flex-wrap flex-row-reverse">
                <v-avatar class="ml-n2 cursor-pointer" color="secondary" size="35" variant="flat">
                  +2
                  <v-tooltip activator="parent" aria-label="tooltip" location="top">
                    <div class="d-flex align-center justify-center flex-wrap flex-row-reverse">
                      <v-avatar class="ml-n2" color="surface" size="40" variant="outlined">
                        <img alt="vector" src="@/assets/images/users/avatar-6.png" width="40">
                      </v-avatar>
                      <v-avatar class="ml-n2" color="surface" size="40" variant="outlined">
                        <img alt="vector" src="@/assets/images/users/avatar-5.png" width="40">
                      </v-avatar>
                      <v-avatar color="surface" size="40" variant="outlined">
                        <img alt="vector" src="@/assets/images/users/avatar-4.png" width="40">
                      </v-avatar>
                    </div>
                  </v-tooltip>
                </v-avatar>
                <v-avatar class="ml-n2 cursor-pointer" color="surface" size="35" variant="outlined">
                  <img alt="vector" src="@/assets/images/users/avatar-3.png" width="35">
                </v-avatar>
                <v-avatar class="ml-n2 cursor-pointer" color="surface" size="35" variant="outlined">
                  <img alt="vector" src="@/assets/images/users/avatar-2.png" width="35">
                </v-avatar>
                <v-avatar class="cursor-pointer" color="surface" size="35" variant="outlined">
                  <img alt="vector" src="@/assets/images/users/avatar-1.png" width="35">
                </v-avatar>
              </div>
              <v-btn
                color="primary"
                icon
                rounded
                size="small"
                variant="flat"
              >
                <PlusOutlined />
              </v-btn>
            </div>
            <p class="text-caption mb-0 text-lightText mt-7">Chrome fixed the bug several versions ago, thus rendering this...</p>
          </div>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>
