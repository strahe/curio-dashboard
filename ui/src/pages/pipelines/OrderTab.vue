<script setup lang="ts">
import { ref } from 'vue'
import type { Header, Item } from 'vue3-easy-data-table'
import 'vue3-easy-data-table/dist/style.css'

import user1 from '@/assets/images/users/avatar-5.png'
import user2 from '@/assets/images/users/avatar-4.png'
import user3 from '@/assets/images/users/avatar-7.png'
import user4 from '@/assets/images/users/avatar-1.png'
import user5 from '@/assets/images/users/avatar-2.png'

// icons
import { DeleteOutlined, DownloadOutlined, EditOutlined, EyeOutlined, PlusOutlined, SearchOutlined } from '@ant-design/icons-vue'

const searchField = ref('name')
const searchValue = ref('')

const headers: Header[] = [
    { text: 'Invoice ID', value: 'id', sortable: true },
    { text: 'User info', value: 'name', sortable: true },
    { text: 'Create date', value: 'date', sortable: true },
    { text: 'Due date', value: 'due', sortable: true },
    { text: 'Quantity', value: 'qty', sortable: true },
    { text: 'Status', value: 'status', sortable: true },
    { text: 'Action', value: 'operation' },
]
const items = ref([
    {
        id: 1,
        name: 'Tessi Eneas',
        image: user1,
        mail: 'tass23@gmail.com',
        date: '05/01/2023',
        due: '06/02/2023',
        qty: '1000',
        status: 1,
    },
    {
        id: 2,
        name: 'Abey Boseley',
        image: user2,
        mail: 'aabsl32@gmail.com',
        date: '7/15/2023',
        due: '2/15/2023',
        qty: '2030',
        status: 2,
    },
    {
        id: 3,
        name: 'Shelba Thews',
        image: user3,
        mail: 'slbt37@gmail.com',
        date: '7/6/2023',
        due: '7/8/2023',
        qty: '3000',
        status: 3,
    },
    {
        id: 4,
        name: 'Salvatore Boncore',
        image: user4,
        mail: 'sabf231@gmail.com',
        date: '2/8/2023',
        due: '3/30/2023',
        qty: '2000',
        status: 2,
    },
    {
        id: 5,
        name: 'Mickie Melmoth',
        image: user5,
        mail: 'mmsht23@gmail.com',
        date: '5/5/2023',
        due: '7/11/2023',
        qty: '3000',
        status: 1,
    },
])
const items1 = ref([
    {
        id: 1,
        name: 'Tessi Eneas',
        image: user1,
        mail: 'tass23@gmail.com',
        date: '05/01/2023',
        due: '06/02/2023',
        qty: '1000',
        status: 1,
    },
    {
        id: 5,
        name: 'Mickie Melmoth',
        image: user5,
        mail: 'mmsht23@gmail.com',
        date: '5/5/2023',
        due: '7/11/2023',
        qty: '3000',
        status: 1,
    },
])
const items2 = ref([
    {
        id: 2,
        name: 'Abey Boseley',
        image: user2,
        mail: 'aabsl32@gmail.com',
        date: '7/15/2023',
        due: '2/15/2023',
        qty: '2030',
        status: 2,
    },
    {
        id: 4,
        name: 'Salvatore Boncore',
        image: user4,
        mail: 'sabf231@gmail.com',
        date: '2/8/2023',
        due: '3/30/2023',
        qty: '2000',
        status: 2,
    },
])
const items3 = ref([
    {
        id: 3,
        name: 'Shelba Thews',
        image: user3,
        mail: 'slbt37@gmail.com',
        date: '7/6/2023',
        due: '7/8/2023',
        qty: '3000',
        status: 3,
    },
])
const themeColor = ref('rgb(var(--v-theme-primary))')
const itemsSelected = ref<Item[]>([])

// tabs data
const tab = ref(0)

function deleteRow (index: number) {
    items.value.splice(index, 1)
}
function deleteRow1 (index: number) {
    items1.value.splice(index, 1)
}
function deleteRow2 (index: number) {
    items2.value.splice(index, 1)
}
function deleteRow3 (index: number) {
    items3.value.splice(index, 1)
}
</script>
<template>
  <v-card class="bg-surface" variant="outlined">
    <v-card-item>
      <v-tabs v-model="tab" color="primary">
        <v-tab class="font-weight-medium" value="1">
          All
          <v-chip class="ml-2 font-weight-medium" color="primary" label size="small"> 5 </v-chip>
        </v-tab>
        <v-tab class="font-weight-medium" value="2">
          Paid
          <v-chip class="ml-2 font-weight-medium" color="success" label size="small"> 2 </v-chip>
        </v-tab>
        <v-tab class="font-weight-medium" value="3">
          Unpaid
          <v-chip class="ml-2 font-weight-medium" color="warning" label size="small"> 2 </v-chip>
        </v-tab>
        <v-tab class="font-weight-medium" value="4">
          Cancelled
          <v-chip class="ml-2 font-weight-medium" color="error" label size="small"> 1 </v-chip>
        </v-tab>
      </v-tabs>
      <v-divider />
      <v-window v-model="tab" class="mt-5">
        <v-window-item value="1">
          <v-row class="align-center mb-2" justify="space-between">
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchValue"
                hide-details
                persistent-placeholder
                placeholder="Search 5 records..."
                type="text"
                variant="outlined"
              >
                <template #prepend-inner>
                  <SearchOutlined :style="{ fontSize: '14px' }" />
                </template>
              </v-text-field>
            </v-col>
            <v-col cols="12" md="3">
              <div class="d-flex ga-2 justify-end">
                <v-btn color="primary" to="/app/invoice/create" variant="flat">
                  <template #prepend>
                    <PlusOutlined />
                  </template>
                  Add invoice
                </v-btn>
                <v-btn icon rounded size="small" variant="text">
                  <DownloadOutlined class="text-lightText" :style="{ fontSize: '24px' }" />
                </v-btn>
              </div>
            </v-col>
          </v-row>
          <div class="invoice-table">
            <EasyDataTable
              v-model:items-selected="itemsSelected"
              :headers="headers"
              :items="items"
              :rows-per-page="5"
              :search-field="searchField"
              :search-value="searchValue"
              table-class-name="customize-table"
              :theme-color="themeColor"
            >
              <template #item-name="{ name, image, mail }">
                <div class="player-wrapper d-flex align-center">
                  <v-avatar size="32">
                    <img alt="profile" :src="image" width="32">
                  </v-avatar>
                  <div class="ml-2">
                    <h5 class="text-subtitle-1 mb-0">{{ name }}</h5>
                    <p class="text-h6 text-lightText mb-0">{{ mail }}</p>
                  </div>
                </div>
              </template>
              <template #item-status="{ status }">
                <v-chip v-if="status === 1" color="success" label size="small"> Paid </v-chip>
                <v-chip v-if="status === 2" color="info" label size="small"> Unpaid </v-chip>
                <v-chip v-if="status === 3" color="error" label size="small"> Cancelled </v-chip>
              </template>
              <template #item-operation>
                <div class="operation-wrapper">
                  <v-btn
                    color="secondary"
                    icon
                    rounded
                    to="/app/invoice/details"
                    variant="text"
                  >
                    <EyeOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">View</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn color="primary" icon rounded variant="text">
                    <EditOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Edit</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn
                    color="error"
                    icon
                    rounded
                    variant="text"
                    @click="deleteRow"
                  >
                    <DeleteOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Delete</span>
                    </v-tooltip>
                  </v-btn>
                </div>
              </template>
            </EasyDataTable>
          </div>
        </v-window-item>

        <v-window-item value="2">
          <v-row class="align-center mb-2" justify="space-between">
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchValue"
                hide-details
                persistent-placeholder
                placeholder="Search 5 records..."
                type="text"
                variant="outlined"
              >
                <template #prepend-inner>
                  <SearchOutlined :style="{ fontSize: '14px' }" />
                </template>
              </v-text-field>
            </v-col>
            <v-col cols="12" md="3">
              <div class="d-flex ga-2 justify-end">
                <v-btn color="primary" to="/app/invoice/create" variant="flat">
                  <template #prepend>
                    <PlusOutlined />
                  </template>
                  Add invoice
                </v-btn>
                <v-btn icon rounded size="small" variant="text">
                  <DownloadOutlined class="text-lightText" :style="{ fontSize: '24px' }" />
                </v-btn>
              </div>
            </v-col>
          </v-row>
          <div class="invoice-table">
            <EasyDataTable
              v-model:items-selected="itemsSelected"
              :headers="headers"
              :items="items1"
              :rows-per-page="5"
              :search-field="searchField"
              :search-value="searchValue"
              table-class-name="customize-table"
              :theme-color="themeColor"
            >
              <template #item-name="{ name, image, mail }">
                <div class="player-wrapper d-flex align-center">
                  <v-avatar size="32">
                    <img alt="profile" :src="image" width="32">
                  </v-avatar>
                  <div class="ml-2">
                    <h5 class="text-subtitle-1 mb-0">{{ name }}</h5>
                    <p class="text-h6 text-lightText mb-0">{{ mail }}</p>
                  </div>
                </div>
              </template>
              <template #item-status="{ status }">
                <v-chip v-if="status === 1" color="success" label size="small"> Paid </v-chip>
                <v-chip v-if="status === 2" color="info" label size="small"> Unpaid </v-chip>
                <v-chip v-if="status === 3" color="error" label size="small"> Cancelled </v-chip>
              </template>
              <template #item-operation>
                <div class="operation-wrapper">
                  <v-btn
                    color="secondary"
                    icon
                    rounded
                    to="/app/invoice/details"
                    variant="text"
                  >
                    <EyeOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">View</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn color="primary" icon rounded variant="text">
                    <EditOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Edit</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn
                    color="error"
                    icon
                    rounded
                    variant="text"
                    @click="deleteRow1"
                  >
                    <DeleteOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Delete</span>
                    </v-tooltip>
                  </v-btn>
                </div>
              </template>
            </EasyDataTable>
          </div>
        </v-window-item>

        <v-window-item value="3">
          <v-row class="align-center mb-2" justify="space-between">
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchValue"
                hide-details
                persistent-placeholder
                placeholder="Search 5 records..."
                type="text"
                variant="outlined"
              >
                <template #prepend-inner>
                  <SearchOutlined :style="{ fontSize: '14px' }" />
                </template>
              </v-text-field>
            </v-col>
            <v-col cols="12" md="3">
              <div class="d-flex ga-2 justify-end">
                <v-btn color="primary" to="/app/invoice/create" variant="flat">
                  <template #prepend>
                    <PlusOutlined />
                  </template>
                  Add invoice
                </v-btn>
                <v-btn icon rounded size="small" variant="text">
                  <DownloadOutlined class="text-lightText" :style="{ fontSize: '24px' }" />
                </v-btn>
              </div>
            </v-col>
          </v-row>
          <div class="invoice-table">
            <EasyDataTable
              v-model:items-selected="itemsSelected"
              :headers="headers"
              :items="items2"
              :rows-per-page="5"
              :search-field="searchField"
              :search-value="searchValue"
              table-class-name="customize-table"
              :theme-color="themeColor"
            >
              <template #item-name="{ name, image, mail }">
                <div class="player-wrapper d-flex align-center">
                  <v-avatar size="32">
                    <img alt="profile" :src="image" width="32">
                  </v-avatar>
                  <div class="ml-2">
                    <h5 class="text-subtitle-1 mb-0">{{ name }}</h5>
                    <p class="text-h6 text-lightText mb-0">{{ mail }}</p>
                  </div>
                </div>
              </template>
              <template #item-status="{ status }">
                <v-chip v-if="status === 1" color="success" label size="small"> Paid </v-chip>
                <v-chip v-if="status === 2" color="info" label size="small"> Unpaid </v-chip>
                <v-chip v-if="status === 3" color="error" label size="small"> Cancelled </v-chip>
              </template>
              <template #item-operation>
                <div class="operation-wrapper">
                  <v-btn
                    color="secondary"
                    icon
                    rounded
                    to="/app/invoice/details"
                    variant="text"
                  >
                    <EyeOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">View</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn color="primary" icon rounded variant="text">
                    <EditOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Edit</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn
                    color="error"
                    icon
                    rounded
                    variant="text"
                    @click="deleteRow2"
                  >
                    <DeleteOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Delete</span>
                    </v-tooltip>
                  </v-btn>
                </div>
              </template>
            </EasyDataTable>
          </div>
        </v-window-item>

        <v-window-item value="4">
          <v-row class="align-center mb-2" justify="space-between">
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchValue"
                hide-details
                persistent-placeholder
                placeholder="Search 5 records..."
                type="text"
                variant="outlined"
              >
                <template #prepend-inner>
                  <SearchOutlined :style="{ fontSize: '14px' }" />
                </template>
              </v-text-field>
            </v-col>
            <v-col cols="12" md="3">
              <div class="d-flex ga-2 justify-end">
                <v-btn color="primary" to="/app/invoice/create" variant="flat">
                  <template #prepend>
                    <PlusOutlined />
                  </template>
                  Add invoice
                </v-btn>
                <v-btn icon rounded size="small" variant="text">
                  <DownloadOutlined class="text-lightText" :style="{ fontSize: '24px' }" />
                </v-btn>
              </div>
            </v-col>
          </v-row>
          <div class="invoice-table">
            <EasyDataTable
              v-model:items-selected="itemsSelected"
              :headers="headers"
              :items="items3"
              :rows-per-page="5"
              :search-field="searchField"
              :search-value="searchValue"
              table-class-name="customize-table"
              :theme-color="themeColor"
            >
              <template #item-name="{ name, image, mail }">
                <div class="player-wrapper d-flex align-center">
                  <v-avatar size="32">
                    <img alt="profile" :src="image" width="32">
                  </v-avatar>
                  <div class="ml-2">
                    <h5 class="text-subtitle-1 mb-0">{{ name }}</h5>
                    <p class="text-h6 text-lightText mb-0">{{ mail }}</p>
                  </div>
                </div>
              </template>
              <template #item-status="{ status }">
                <v-chip v-if="status === 1" color="success" label size="small"> Paid </v-chip>
                <v-chip v-if="status === 2" color="info" label size="small"> Unpaid </v-chip>
                <v-chip v-if="status === 3" color="error" label size="small"> Cancelled </v-chip>
              </template>
              <template #item-operation>
                <div class="operation-wrapper">
                  <v-btn
                    color="secondary"
                    icon
                    rounded
                    to="/app/invoice/details"
                    variant="text"
                  >
                    <EyeOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">View</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn color="primary" icon rounded variant="text">
                    <EditOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Edit</span>
                    </v-tooltip>
                  </v-btn>
                  <v-btn
                    color="error"
                    icon
                    rounded
                    variant="text"
                    @click="deleteRow3"
                  >
                    <DeleteOutlined />
                    <v-tooltip activator="parent" aria-label="tooltip" content-class="smallTooltip" location="bottom">
                      <span class="text-caption">Delete</span>
                    </v-tooltip>
                  </v-btn>
                </div>
              </template>
            </EasyDataTable>
          </div>
        </v-window-item>
      </v-window>
    </v-card-item>
  </v-card>
</template>
<style lang="scss">
.v-tooltip {
  > .v-overlay__content.smallTooltip {
    padding: 2px 10px;
  }
}
</style>
