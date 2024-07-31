<script setup lang="ts">
import { ref } from 'vue';
// assets
import imageReader from '@/assets/images/analytics/reader.svg';
// icons
import { PlusOutlined } from '@ant-design/icons-vue';

const p1 = ref(35);
</script>
<template>
  <v-card elevation="0" variant="outlined" class="overflow-hidden bg-surface">
    <v-card-text class="pa-0">
      <v-row class="ma-0">
        <v-col cols="12" sm="7" class="pa-0">
          <div class="bg-primary pa-5 position-relative h-100 overflow-hidden">
            <h5 class="text-h5">What would you want to learn today</h5>
            <p class="text-caption" style="max-width: 55%">Your learning capacity is 80% as daily analytics</p>
            <h4 class="text-h4 pt-sm-13 pt-8">35% Completed</h4>
            <div style="max-width: 60%" class="pb-7">
              <v-progress-linear
                color="success"
                height="6"
                aria-label="progressbar"
                rounded
                v-model="p1"
                bg-color="lightsuccess"
                bg-opacity="1"
              ></v-progress-linear>
            </div>
            <v-img :src="imageReader" class="readMedia" alt="reader" cover />
          </div>
        </v-col>
        <v-col cols="12" sm="5" class="pa-0">
          <div class="pa-5 h-100">
            <h6 class="text-h6">Get started with new basic skills</h6>
            <p class="text-lightText text-h6 mb-0">Last Date 5th Nov 2020</p>
            <v-divider class="my-6"></v-divider>
            <div class="d-flex justify-space-between">
              <div class="d-flex align-center justify-center flex-wrap flex-row-reverse">
                <v-avatar size="35" variant="flat" color="secondary" class="ml-n2 cursor-pointer">
                  +2
                  <v-tooltip activator="parent" location="top" aria-label="tooltip">
                    <div class="d-flex align-center justify-center flex-wrap flex-row-reverse">
                      <v-avatar size="40" variant="outlined" color="surface" class="ml-n2">
                        <img src="@/assets/images/users/avatar-6.png" width="40" alt="vector" />
                      </v-avatar>
                      <v-avatar size="40" variant="outlined" color="surface" class="ml-n2">
                        <img src="@/assets/images/users/avatar-5.png" width="40" alt="vector" />
                      </v-avatar>
                      <v-avatar size="40" variant="outlined" color="surface">
                        <img src="@/assets/images/users/avatar-4.png" width="40" alt="vector" />
                      </v-avatar>
                    </div>
                  </v-tooltip>
                </v-avatar>
                <v-avatar size="35" variant="outlined" color="surface" class="ml-n2 cursor-pointer">
                  <img src="@/assets/images/users/avatar-3.png" width="35" alt="vector" />
                </v-avatar>
                <v-avatar size="35" variant="outlined" color="surface" class="ml-n2 cursor-pointer">
                  <img src="@/assets/images/users/avatar-2.png" width="35" alt="vector" />
                </v-avatar>
                <v-avatar size="35" variant="outlined" color="surface" class="cursor-pointer">
                  <img src="@/assets/images/users/avatar-1.png" width="35" alt="vector" />
                </v-avatar>
              </div>
              <v-btn icon rounded color="primary" variant="flat" size="small">
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
