<script setup lang="ts">
import { shallowRef } from 'vue';

// assets
import Product1 from '@/assets/images/products/1.jpg';
import Product2 from '@/assets/images/products/2.jpg';
import Product3 from '@/assets/images/products/3.jpg';
import Product4 from '@/assets/images/products/4.jpg';

// import icons
import { EditOutlined, DeleteOutlined } from '@ant-design/icons-vue';

const products = shallowRef([
  {
    orderid: 81412314,
    customer: 'John Deo',
    avatar: Product1,
    product: 'Moto G5',
    total: '10',
    date: '17-2-2017',
    status: 'pending'
  },
  {
    orderid: 68457898,
    customer: 'Jenny William',
    avatar: Product2,
    product: 'iPhone 8',
    total: '16',
    date: '20-2-2017',
    status: 'paid'
  },
  {
    orderid: 45457898,
    customer: 'Lori Moore',
    avatar: Product3,
    product: 'Redmi 4',
    total: '20',
    date: '17-2-2017',
    status: 'success'
  },
  {
    orderid: 62446232,
    customer: 'Austin Pena',
    avatar: Product4,
    product: 'Jio',
    total: '15',
    date: '125-4-2017',
    status: 'failed'
  }
]);
</script>

<template>
  <v-table class="bordered-table" hover density="comfortable">
    <thead class="bg-containerBg">
      <tr>
        <th class="text-left text-caption font-weight-bold text-uppercase">Customer</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Order Id</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Photo</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Product</th>

        <th class="text-left text-caption font-weight-bold text-uppercase">Total</th>
        <th class="text-left text-caption font-weight-bold text-uppercase" style="min-width: 110px">Date</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Status</th>
        <th class="text-left text-caption font-weight-bold text-uppercase pl-9">Action</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in products" :key="item.orderid">
        <td class="py-3">{{ item.customer }}</td>
        <td class="py-3">#{{ item.orderid }}</td>
        <td class="py-3">
          <img :src="item.avatar" width="20" height="40" :alt="item.avatar" />
        </td>
        <td class="py-3">{{ item.product }}</td>
        <td class="py-3">{{ item.total }}</td>
        <td class="py-3" style="min-width: 110px">{{ item.date }}</td>
        <td class="py-3">
          <v-chip color="warning" variant="flat" label size="small" v-if="item.status === 'pending'"> Pending</v-chip>
          <v-chip color="primary" variant="flat" label size="small" v-else-if="item.status === 'paid'"> Paid</v-chip>
          <v-chip color="success" variant="flat" label size="small" v-else-if="item.status === 'success'"> Success</v-chip>
          <v-chip color="error" variant="flat" label size="small" v-else>Failed</v-chip>
        </td>
        <td class="py-3">
          <div class="d-flex align-center">
            <v-btn icon color="primary" variant="text" rounded>
              <EditOutlined :style="{ fontSize: '20px' }" />
            </v-btn>
            <v-btn icon color="inherit" variant="text" rounded>
              <DeleteOutlined :style="{ fontSize: '20px' }" />
            </v-btn>
          </div>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>
