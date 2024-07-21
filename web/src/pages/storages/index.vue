<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {GetStoragePaths} from "@/pages/storages/query";
import {StoragePath, StorageStats} from "@/typed-graph";
import {formatBytes} from "@/utils/formatBytes";
import {format} from "timeago.js";
import {definePage} from "unplugin-vue-router/runtime";
import {ComputedRef} from "vue";
import StorageUsagesChart from "@/widgets/StorageUsagesChart.vue";

definePage({
  meta: {
    title: 'Storages',
  },
})


const { result, loading, refetch, error } = useQuery(GetStoragePaths, null, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[StoragePath]> = computed(() => result.value?.storagePaths || []);
const stats: ComputedRef<[StorageStats]> = computed(() => result.value?.storageStats || []);

const headers = [
  { title: 'Type', key: 'type' },
  { title: 'Storage ID', key: 'storageId' },
  { title: 'URLs', key: 'urls' },
  { title: 'Weight', key: 'weight' },
  { title: 'Capacity', key: 'capacity' },
  { title: 'Available', key: 'available' },
  { title: 'Filesystem Available', key: 'fsAvailable' },
  { title: 'Reserved', key: 'reserved' },
  { title: 'Used', key: 'used' },
  { title: 'Last Heartbeat', key: 'lastHeartbeat' },
];

type SortItem = {
  key: string;
  order?: 'asc' | 'desc';
};
const groupBy: readonly SortItem[] = [{ key: 'type', order: 'asc' }];

function getTypeStats(type: string) {
  const res = stats.value?.find((s: StorageStats) => s.type === type)
  if (res) {
    return res
  } else {
    return {
      type: type,
      totalAvailable: 0,
      totalCapacity: 0,
      totalFsAvailable: 0,
      totalReserved: 0,
      totalUsed: 0
    }
  }
}

function percentUsed(type: string) {
  return (1-getTypeStats(type).totalAvailable/getTypeStats(type).totalCapacity)*100
}
</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center" align="center">
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="Seal Usage"
          :value="formatBytes(getTypeStats('Seal').totalUsed).combined + ' / ' + formatBytes(getTypeStats('Seal').totalCapacity).combined"
          :progress="percentUsed('Seal') || 1"
          icon="mdi-database"
          color="red"
        ></StatsCard>
      </v-col>

      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="Store Usage"
          :value="formatBytes(getTypeStats('Store').totalUsed).combined + ' / ' + formatBytes(getTypeStats('Store').totalCapacity).combined"
          :progress="percentUsed('Store') || 1"
          icon="mdi-nas"
          color="primary"
        ></StatsCard>
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="Hybrid Usage"
          :value="formatBytes(getTypeStats('Hybrid').totalUsed).combined + ' / ' + formatBytes(getTypeStats('Hybrid').totalCapacity).combined"
          :progress="percentUsed('Hybrid') || 1"
          icon="mdi-nas"
          color="secondary"
        ></StatsCard>
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          :loading="loading"
          title="Readonly Usage"
          :value="formatBytes(getTypeStats('Readonly').totalUsed).combined + ' / ' + formatBytes(getTypeStats('Readonly').totalCapacity).combined"
          :progress="percentUsed('Readonly') || 1"
          icon="mdi-tune"
          color="warning"
        ></StatsCard>
      </v-col>
      <v-col cols="12">
        <StorageUsagesChart></StorageUsagesChart>
      </v-col>
      <v-col cols="12">
        <Card title="Storage Paths" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
          </template>
          <v-data-table
            :group-by="groupBy"
            :headers="headers"
            :items="items"
            :loading="loading"
            item-value="storageId">
            <template v-slot:group-header="{ item, toggleGroup, isGroupOpen }">
              <tr>
                <td :colspan="1">
                  <VBtn
                    :icon="isGroupOpen(item) ? '$expand' : '$next'"
                    size="small"
                    variant="text"
                    @click="toggleGroup(item)"
                  ></VBtn>
                  <v-chip color="primary">{{ item.value }}</v-chip>
                </td>
                <td :colspan="4">
                  <v-progress-linear
                    min="0"
                    :max="getTypeStats(item.value).totalCapacity"
                    :model-value="getTypeStats(item.value).totalUsed"
                    :buffer-value="(getTypeStats(item.value).totalCapacity- getTypeStats(item.value).totalAvailable)"
                    :buffer-opacity="getTypeStats(item.value).totalReserved"
                    buffer-color="warning"
                    :color="percentUsed(item.value) > 80 ? 'red' : 'green'"
                    height="25"
                    rounded
                  ><strong>{{ Math.ceil(percentUsed(item.value)) }}%</strong></v-progress-linear>
                </td>
                <td :colspan="1">{{ formatBytes(getTypeStats(item.value).totalCapacity).combined }}</td>
                <td :colspan="1">{{ formatBytes(getTypeStats(item.value).totalAvailable).combined }}</td>
                <td :colspan="1">{{ formatBytes(getTypeStats(item.value).totalFsAvailable).combined }}</td>
                <td :colspan="1">{{ formatBytes(getTypeStats(item.value).totalReserved).combined }}</td>
                <td :colspan="2">{{ formatBytes(getTypeStats(item.value).totalUsed).combined }}</td>
              </tr>
            </template>
            <template v-slot:[`item.type`]="{item}">
              {{item.type}}
            </template>
            <template v-slot:[`item.capacity`]="{item}">
              {{ formatBytes(item.capacity).combined }}
            </template>
            <template v-slot:[`item.available`]="{item}">
              {{ formatBytes(item.available).combined }}
            </template>
            <template v-slot:[`item.reserved`]="{item}">
              {{ formatBytes(item.reserved).combined }}
            </template>
            <template v-slot:[`item.fsAvailable`]="{item}">
              {{ formatBytes(item.fsAvailable).combined }}
            </template>
            <template v-slot:[`item.used`]="{item}">
              {{ formatBytes(item.used).combined }}
            </template>
            <template v-slot:[`item.lastHeartbeat`]="{item}">
              {{format(item.lastHeartbeat)}}
            </template>
          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
