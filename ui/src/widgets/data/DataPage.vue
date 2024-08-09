<script setup lang="ts">
import { ref } from 'vue'
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue'
import CardHeader from '@/components/shared/CardHeader.vue'
import WidgetCard from '@/components/shared/WidgetCard.vue'

// icons
import { PlusCircleOutlined } from '@ant-design/icons-vue'

// imports widgets
import TodoList from './components/TodoList.vue'
import TrafficSources from './components/TrafficSources.vue'
import TeamMembers from './components/TeamMembers.vue'
import UserActivity from './components/UserActivity.vue'
import LatestPosts from './components/LatestPosts.vue'
import Messages from './components/MessageList.vue'
import Projects from './components/ProjectList.vue'
import ProductSales from './components/ProductSales.vue'
import Tasks from './components/TaskList.vue'
import ApplicationSales from './components/ApplicationSales.vue'
import ActiveTickets from './components/ActiveTickets.vue'
import Feeds from './components/FeedList.vue'
import LatestCustomer from './components/LatestCustomer.vue'
import LatestOrder from './components/LatestOrder.vue'
import IncomingRequest from './components/IncomingRequest.vue'
import TotalRevenue from './components/TotalRevenue.vue'
import NewCustomers from './components/NewCustomers.vue'
import RecentTickets from './components/RecentTickets.vue'

const page = ref({ title: 'Data' })
const breadcrumbs = ref([
  {
    title: 'Widget',
    disabled: false,
    href: '#',
  },
  {
    title: 'Data',
    disabled: true,
    href: '#',
  },
])
</script>

<template>
  <BaseBreadcrumb :breadcrumbs="breadcrumbs" :title="page.title" />

  <v-row>
    <!-- ---------------------------------------------------------------- -->
    <!-- Todo List -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader class="headerWithBtn" title="To Do List">
        <template #header>
          <v-btn
            color="primary"
            icon
            rounded
            size="small"
            variant="text"
          >
            <PlusCircleOutlined :style="{ fontSize: '16px' }" />
          </v-btn>
        </template>
        <TodoList />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Traffic Sources -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <WidgetCard subtitle="You’re getting more and more sources, keep it up!" title="Traffic Sources">
        <TrafficSources />
      </WidgetCard>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Team Members -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader class="overflow-hidden" title="Team Members">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <TeamMembers />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Messages -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="6">
      <CardHeader title="Latest Messages">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <Messages />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- User Activity -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="6">
      <CardHeader title="User Activity">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <UserActivity />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Projects -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="6">
      <CardHeader class="overflow-hidden" title="Projects">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <Projects />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Product Sales -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="6">
      <CardHeader class="overflow-hidden" title="Product Sales">
        <ProductSales />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Tasks -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="4" md="5">
      <CardHeader title="Tasks">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <Tasks />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Application Sales -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8" md="7">
      <CardHeader class="overflow-hidden" title="Application Sales">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <ApplicationSales />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Active Tickets -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="8">
      <CardHeader title="Active Tickets">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <ActiveTickets />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Posts -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader class="overflow-hidden" title="Latest Posts">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <LatestPosts />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Feeds -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="4" md="5">
      <CardHeader class="overflow-hidden" title="Feeds">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <Feeds />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Customers -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8" md="7">
      <CardHeader class="overflow-hidden" title="Latest Customers">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <LatestCustomer />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Order -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12">
      <CardHeader class="overflow-hidden" title="Latest Order">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <LatestOrder />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Incoming Requests -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader class="overflow-hidden" title="Incoming Requests">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <IncomingRequest />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Total Revenue -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader class="overflow-hidden" title="Total Revenue">
        <TotalRevenue />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- New Customers -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader class="overflow-hidden" title="New Customers">
        <NewCustomers />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Recent Tickets -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8">
      <CardHeader class="overflow-hidden" title="Recent Tickets">
        <template #header>
          <router-link class="text-h6 text-primary link-hover" to="/widget/data">View all</router-link>
        </template>
        <RecentTickets />
      </CardHeader>
    </v-col>
  </v-row>
</template>
