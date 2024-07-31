<script setup lang="ts">
import { shallowRef } from 'vue';

const projects = shallowRef([
  {
    subject: 'Website down for one week',
    dept: 'Support',
    date: 'Today 2:00',
    status: 'Open'
  },
  {
    subject: 'Loosing control on server',
    dept: 'Support',
    date: 'Yesterday',
    status: 'Progress'
  },
  {
    subject: 'Authorizations keys',
    dept: 'Support',
    date: '27, Aug',
    status: 'Closed'
  },
  {
    subject: 'Restoring default settings',
    dept: 'Support',
    date: 'Today 9:00 ',
    status: 'Open'
  },
  {
    subject: 'Loosing control on server',
    dept: 'Support',
    date: 'Yesterday',
    status: 'Progress'
  },
  {
    subject: 'Authorizations keys',
    dept: 'Support',
    date: '27, Aug',
    status: 'Closed'
  },
  {
    subject: 'Restoring default settings',
    dept: 'Support',
    date: 'Today 9:00',
    status: 'Open'
  },
  {
    subject: 'Authorizations keys',
    dept: 'Support',
    date: '27, Aug',
    status: 'Closed'
  }
]);
</script>

<template>
  <v-table class="bordered-table" hover density="comfortable">
    <thead class="bg-containerBg">
      <tr>
        <th class="text-left text-caption font-weight-bold text-uppercase">Subject</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Department</th>
        <th class="text-left text-caption font-weight-bold text-uppercase">Date</th>
        <th class="text-right text-caption font-weight-bold text-uppercase">Status</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in projects" :key="item.subject">
        <td class="py-3">
          {{ item.subject }}
        </td>
        <td class="py-3">{{ item.dept }}</td>
        <td class="py-3">{{ item.date }}</td>
        <td class="py-3 text-right">
          <v-chip variant="outlined" label color="secondary" size="small">{{ item.status }}</v-chip>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>
