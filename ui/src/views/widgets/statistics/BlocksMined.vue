<script setup lang="ts">

import { computed, ComputedRef } from 'vue'
import { useQuery } from '@vue/apollo-composable'
import { GetMiningBlockCount } from '@/views/query/mining'
import { MiningCount } from '@/typed-graph'
import { IconBox } from '@tabler/icons-vue'

const props = defineProps({
  sp: String,
  lastHours: {
    type: Number,
    default: 24,
  },
})

const currentEnd = new Date()
const currentStart = new Date(currentEnd.getTime() - props.lastHours * 60 * 60 * 1000)

const { result } = useQuery(GetMiningBlockCount, {
  sp: props.sp,
  end: currentEnd,
  start: currentStart,
})

const count: ComputedRef<MiningCount> = computed(() => result.value?.miningCount || {
  include: 0,
  exclude: 0,
})

</script>

<template>
  <v-card elevation="0">
    <v-card variant="outlined">
      <v-card-text>
        <div class="d-flex align-items-center justify-space-between">
          <div>
            <h5 class="text-h5">{{ $t('fields.Blocks Mined') }}</h5>
            <h3 class="text-h3 my-2">{{ count.include }}</h3>
            <h6 class="text-caption font-weight-medium mb-0">
              {{ $d(currentStart, 'short') }} - {{ $d(currentEnd, 'short') }}
            </h6>
          </div>
          <span class="d-flex align-center">
            <v-btn
              icon="true"
              rounded="md"
              variant="flat"
            >
              <IconBox :size="20" />
            </v-btn>
          </span>
        </div>
      </v-card-text>
    </v-card>
  </v-card>
</template>

<style scoped lang="scss">

</style>
