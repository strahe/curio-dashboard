<script setup lang="ts">

import UiChildCard from '@/components/shared/UiChildCard.vue'
import { useQuery } from '@vue/apollo-composable'
import { computed, ComputedRef } from 'vue'
import { StoragePath } from '@/typed-graph'
import { GetMachineStorages } from '@/gql/storage'
import { IconReload } from '@tabler/icons-vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({
  id: {
    type: Number,
    required: true,
  },
})

const { result, loading, refetch, error } = useQuery(GetMachineStorages, {
  id: props.id,
}, () => ({
  fetchPolicy: 'cache-first',
}))

const paths: ComputedRef<[StoragePath]> = computed(() => result.value?.machine.storages || [])

</script>

<template>
  <UiChildCard :title="t('fields.Storages')">
    <template #action>
      <v-btn
        :icon="IconReload"
        rounded
        variant="text"
        @click="refetch"
      />
    </template>
    <StoragePathTable
      :error="error"
      :items="paths"
      :loading="loading"
    />
  </UiChildCard>
</template>

<style scoped lang="scss">

</style>
