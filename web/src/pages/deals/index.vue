<script setup lang="ts">
import {definePage} from "unplugin-vue-router/runtime";
import {useQuery} from "@vue/apollo-composable";
import {GetPendingDeals} from "@/pages/deals/graphql";
import {OpenSectorPiece} from "@/typed-graph";
import {ComputedRef} from "vue";
import {format} from "timeago.js";
import {formatBytes} from "@/utils/formatBytes";
import {VDataTable} from "vuetify/components";

definePage({
  meta: {
    title: 'Pending Deals',
  },
})

const { result, loading, refetch, error } = useQuery(GetPendingDeals, null, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[OpenSectorPiece]> = computed(() => result.value?.dealsPending || []);

const headers = ref([
  { title: 'Miner', key: 'spID' },
  { title: 'Sector', key: 'sectorNumber' },
  { title: 'Piece Index', key: 'pieceIndex' },
  { title: 'Piece CID', key: 'pieceCID', sortable: false },
  { title: 'Piece Size', key: 'pieceSize' },
  { title: 'Data Raw Size', key: 'dataRawSize' },
  { title: 'Delete On Finalize', key: 'dataDeleteOnFinalize' },
  { title: 'SnapDeals', key: 'isSnap' },
  { title: 'Created', key: 'createdAt' },
  { title: 'Action', key: 'action', sortable: false },
]);

const sortBy = [
  { key: 'sectorNumber', order: 'asc' },
  { key: 'pieceIndex', order: 'asc' }];

function sealNow(item: OpenSectorPiece) {
  console.log('seal now ', item.sectorNumber)
  // todo: call the api to seal the sector
}

</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="12">
        <Card
        title="Pending Deals"
        :error="error as Error"
        >
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
            :sort-by="sortBy"
            items-per-page="15"
          >
            <template v-slot:[`item.createdAt`]="{value}">
              {{format(value)}}
            </template>
            <template v-slot:[`item.dataDeleteOnFinalize`]="{value}">
              <v-chip :color="value ? 'red' : 'green'">{{value ? 'Yes' : 'No'}}</v-chip>
            </template>
            <template v-slot:[`item.pieceSize`]="{value}">
              {{ formatBytes(value).combined }}
            </template>
            <template v-slot:[`item.dataRawSize`]="{value}">
              {{ formatBytes(value).combined }}
            </template>
            <template v-slot:[`item.action`]="{item}">
              <v-btn
                @click="sealNow(item)"
                size="small"
                color="primary"
                prepend-icon="mdi-play">Seal Now</v-btn>
            </template>
          </v-data-table>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
