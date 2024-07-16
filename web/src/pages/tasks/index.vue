<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {Task} from "@/typed-graph";
import {format} from "timeago.js";
import TaskDetails from "@/widgets/TaskDetails.vue";
import {definePage} from "unplugin-vue-router/runtime";
import {GetRunningTasks} from "@/query/task";
import RunningTaskPieChart from "@/widgets/RunningTaskChart.vue";

definePage({
  meta: {
    title: 'Tasks',
  },
})

const { result, loading,refetch, error } = useQuery(GetRunningTasks,  null, ()=>({
  fetchPolicy: 'cache-first',
}))

const items: ComputedRef<[Task]> = computed(() => result.value?.tasks || []);
const headers = ref([
  { title: 'Task', key: 'name', width: '200px'},
  {title: 'ID', key: 'id'},
  { title: 'Posted', key: 'postedTime' },
  { title: 'Updated', key: 'updateTime' },
  { title: '  ', key: 'data-table-expand' },
]);

const selectedNames = ref([]);
const nameCounts = computed(() => {
  return items.value.reduce((acc: Record<string, number>, item) => {
    acc[item.name] = (acc[item.name] || 0) + 1;
    return acc;
  }, {});
});

const chartData = computed(() => ({
  labels: Object.keys(nameCounts.value),
  datasets: [{ data: Object.values(nameCounts.value) }],
}));

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
    colors: {
      enabled: true,
      forceOverride: true
    }
  },
}

</script>

<template>
<v-container fill-height fluid grid-list-xl>
  <v-row justify="center">
    <v-col cols="12">
      <RunningTaskPieChart type="Bar"></RunningTaskPieChart>
    </v-col>
    <v-col cols="12">
      <v-card>
        <v-alert v-if="error" type="error">
          {{ error.message }}
        </v-alert>
        <v-card-title class="d-flex justify-space-between">
          <div>Running Tasks</div>
          <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
        </v-card-title>
        <v-card-text>
          <v-data-table
            :items="items"
            :headers="headers"
            :loading="loading"
            show-expand
          >
            <template v-slot:[`item.postedTime`]="{ item }">
              <v-chip>
                {{ format(item.postedTime, 'en_US') }}
              </v-chip>
            </template>
            <template v-slot:[`item.updateTime`]="{ item }">
              <v-chip>
                {{ format(item.updateTime, 'en_US') }}
              </v-chip>
            </template>
            <template v-slot:expanded-row="{ columns, item }">
              <tr>
                <td :colspan="columns.length">
                  <v-row class="mt-2">
                    <TaskDetails :id="item.id"></TaskDetails>
                  </v-row>
                </td>
              </tr>
            </template>
            <template v-slot:[`header.name`]="{ column }">
              <v-select
                variant="underlined"
                v-model="selectedNames"
                :label="column.title"
                :items="nameCounts ? Object.keys(nameCounts) : []"
                multiple
              >
                <template v-slot:selection="{ item, index }">
                  <v-chip v-if="index < 1">
                    <span>{{ item.title }}</span>
                  </v-chip>
                  <span
                    v-if="index === 1"
                    class="text-grey text-caption align-self-center"
                  >
        (+{{ selectedNames.length - 1 }} others)
      </span>
                </template>
              </v-select>
            </template>
          </v-data-table>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</v-container>
</template>

<style scoped>

</style>
