<script setup lang="ts">
import { ref, watch } from 'vue'
import { useSubscription } from '@vue/apollo-composable'
import { SubscribeCompletedTask } from '@/gql/task'
import { TaskHistory } from '@/typed-graph'
import { formatDuration } from '@/utils/helpers/formatDuration'
import { IconPlayerPause, IconPlayerPlay } from '@tabler/icons-vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  maxMessages: {
    type: Number,
    default: 20,
  },
  host: {
    type: String,
    default: null,
  },
})

const { d, t } = useI18n()

const isStop = ref(false)
const { result, loading, stop, start } = useSubscription(SubscribeCompletedTask, {
  last: props.maxMessages,
  host: props.host,
})

const tasks = ref<TaskHistory[]>([])

watch(
  result,
  data => {
    if (tasks.value.length >= props.maxMessages) {
      tasks.value.pop()
    }
    tasks.value.unshift(data.completedTask)
  },
)

const headers = [
  { title: 'Name', key: 'name' },
  { title: 'ID', key: 'taskId' },
  { title: 'Machine', key: 'completedByHostAndPort' },
  { title: 'Posted', key: 'posted' },
  { title: 'Start', key: 'workStart' },
  { title: 'Queued', key: 'queued' },
  { title: 'Took', key: 'took' },
  { title: 'Result', key: 'result' },
  { title: 'Error', key: 'err' },
]

</script>

<template>
  <UiWidgetCard
    class-name="px-0 pb-0 rounded-md"
    :title="t('fields.Recently Finished Tasks')"
  >
    <template #append>
      <v-btn
        :icon="isStop ? IconPlayerPlay : IconPlayerPause"
        rounded
        variant="text"
        @click="isStop = !isStop; isStop ? stop() : start()"
      />
    </template>
    <template #subtitle>
      <router-link :to="{name: 'TaskHistory'}">
        {{ t('fields.View All') }}
      </router-link>
    </template>
    <v-data-table-virtual
      fixed-header
      :headers="headers"
      height="500"
      hover
      :items="tasks"
      :loading="loading"
    >
      <template #item.name="{ value }">
        <router-link :to="{ name: 'TaskHistory', query: { name: value } }">
          {{ value }}
        </router-link>
      </template>
      <template #item.posted="{ value }">
        {{ d(value, 'short') }}
      </template>
      <template #item.workStart="{ value }">
        {{ d(value, 'short') }}
      </template>
      <template #item.completedByHostAndPort="{ item }">
        <RouterLink
          v-if="item.completedBy"
          :to="{ name: 'MachineInfo', params: { id: item.completedBy.id } }"
        >
          {{ item.completedBy.hostAndPort }}
        </RouterLink>
        <span v-else>{{ item.completedByHostAndPort }}</span>
      </template>
      <template #item.queued="{ item }">
        {{ formatDuration(new Date(item.workStart).getTime() - new Date(item.posted).getTime()) }}
      </template>
      <template #item.took="{ item }">
        {{ formatDuration(new Date(item.workEnd).getTime() - new Date(item.workStart).getTime()) }}
      </template>
      <template #item.result="{ item }">
        <v-chip
          class="px-0"
          size="small"
          variant="text"
        >
          <v-avatar
            class="mr-2"
            :color="item.result ? 'success' : 'error'"
            size="8"
            variant="flat"
          />
          <p class="text-h6 mb-0">
            {{ item.result ? 'success' : 'failure' }}
          </p>
        </v-chip>
      </template>
    </v-data-table-virtual>
  </UiWidgetCard>
</template>
