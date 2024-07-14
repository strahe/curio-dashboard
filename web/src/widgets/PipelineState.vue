<script setup lang="ts">
import {Pipeline} from "@/typed-graph";

defineProps({
  data: {
    type: Object as () => Pipeline,
    required: true,
  }
})

const items = [
  {
    cols: 1,
    items: [
      { title: 'SDR', key: 'afterSdr' },
      { title: 'TreeD', key: 'afterTreeD' }
    ],
  },
  {
    cols:1,
    items: [
      { title: 'TreeR', key: 'afterTreeR' },
      { title: 'TreeC', key: 'afterTreeC' }
    ],
  },
  {
    cols: 2,
    items: [
      { title: 'Precommit', key: 'afterPrecommitMsg' },
      { title: 'Precommit Wait', key: 'afterPrecommitMsgSuccess' }
    ],
  },
  {
    cols: 2,
    items: [
      { title: 'Wait Seed', key: 'seedEpoch' },
      { title: 'PoRep', key: 'afterPorep' }
    ],
  },
  {
    cols: 2,
    items: [
      { title: 'Commit', key: 'afterCommitMsg' },
      { title: 'Commit Wait', key: 'afterCommitMsgSuccess' }
    ],
  },
  {
    cols: 2,
    items: [
      {title: 'Finalize', key: 'afterFinalize'},
      {title: 'Move Storage', key: 'afterMoveStorage'}
    ],
  },
  {
    cols: 1,
    items: [
      {title: 'On Chain', key: 'afterFinalize'},
      {title: 'Active', key: 'afterFinalize'},
    ]
  },
  {
    cols: 1,
    items: [
      {title: 'State', key: 'afterFinalize'},
    ]
  }
]

</script>

<template>
<v-container fluid>
  <v-row align="center" justify="center" no-gutters>
    <template v-for="item in items" :key="item">
      <v-col :cols="item.cols" >
        <v-row no-gutters>
        <template v-for="subItem in item.items" :key="subItem.title">
          <v-col cols="12">
            <v-sheet :height="60 / item.items.length" border rounded class="d-flex align-center justify-center flex-wrap text-center mx-auto px-4" :color="data[subItem.key as keyof typeof data]? 'green': ''">
              {{ subItem.title }}
            </v-sheet>
          </v-col>
        </template>
        </v-row>
      </v-col>
    </template>
  </v-row>
</v-container>
</template>

<style scoped>

</style>
