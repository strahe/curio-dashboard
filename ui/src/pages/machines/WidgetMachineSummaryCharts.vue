<script setup lang="ts">
  import { computed, ComputedRef, ref } from 'vue'
  import { useQuery } from '@vue/apollo-composable'
  import { GetMachinesSummary } from '@/pages/machines/graphql'
  import { MachineSummary } from '@/typed-graph'
  import { formatBytes } from '@/utils/helpers/formatBytes'
  import moment from 'moment'
  import { mdiCpu64Bit } from '@mdi/js'

  const { result } = useQuery(GetMachinesSummary, null, () => ({
    fetchPolicy: 'cache-first',
    pollInterval: 600000,
  }))

  const stats: ComputedRef<MachineSummary> = computed(() => result.value?.machineSummary || {})

  const cards = ref([
    { value: stats.value.total, text: 'Machine', icon: mdiCpu64Bit, color: 'primary', color2: 'lightprimary', duedate: moment().calendar() },
    { value: stats.value.totalCpu, text: 'CPU', icon: mdiCpu64Bit, color: 'primary', color2: 'lightprimary', duedate: moment().calendar() },
    { value: stats.value.totalGpu, text: 'GPU', icon: mdiCpu64Bit, color: 'success', color2: 'lightsuccess', duedate: moment().calendar() },
    { value: formatBytes(stats.value.totalRam).combined, text: 'RAM', icon: mdiCpu64Bit, color: 'warning', color2: 'lightwarning', duedate: moment().calendar() },
  ])

</script>
<template>
  <v-row>
    <v-col
      v-for="(card4, i) in cards"
      :key="i"
      cols="12"
      md="3"
      :value="card4"
    >
      <v-card elevation="0">
        <v-card variant="outlined">
          <v-card-text>
            <div class="d-flex align-items-center justify-space-between">
              <div>
                <h5 class="text-h5">{{ card4.text }}</h5>
                <h3 class="text-h3 my-2">{{ card4.value }}</h3>
                <h6 class="text-caption font-weight-medium mb-0">{{ card4.duedate }}</h6>
              </div>
              <span class="d-flex align-center">
                <v-btn
                  :class="'text-' + card4.color"
                  :color="card4.color2"
                  rounded="md"
                  variant="flat"
                >
                  <v-icon :icon="`mdiSvg:${card4.icon}`" :style="{ fontSize: '20px' }" />
                </v-btn>
              </span>
            </div>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>
</template>
