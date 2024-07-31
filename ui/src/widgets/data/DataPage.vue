<script setup lang="ts">
import { ref } from 'vue';
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue';
import CardHeader from '@/components/shared/CardHeader.vue';
import WidgetCard from '@/components/shared/WidgetCard.vue';

// icons
import { PlusCircleOutlined } from '@ant-design/icons-vue';

// imports widgets
import TodoList from './components/TodoList.vue';
import TrafficSources from './components/TrafficSources.vue';
import TeamMembers from './components/TeamMembers.vue';
import UserActivity from './components/UserActivity.vue';
import LatestPosts from './components/LatestPosts.vue';
import Messages from './components/MessageList.vue';
import Projects from './components/ProjectList.vue';
import ProductSales from './components/ProductSales.vue';
import Tasks from './components/TaskList.vue';
import ApplicationSales from './components/ApplicationSales.vue';
import ActiveTickets from './components/ActiveTickets.vue';
import Feeds from './components/FeedList.vue';
import LatestCustomer from './components/LatestCustomer.vue';
import LatestOrder from './components/LatestOrder.vue';
import IncomingRequest from './components/IncomingRequest.vue';
import TotalRevenue from './components/TotalRevenue.vue';
import NewCustomers from './components/NewCustomers.vue';
import RecentTickets from './components/RecentTickets.vue';

const page = ref({ title: 'Data' });
const breadcrumbs = ref([
  {
    title: 'Widget',
    disabled: false,
    href: '#'
  },
  {
    title: 'Data',
    disabled: true,
    href: '#'
  }
]);
</script>

<template>
  <BaseBreadcrumb :title="page.title" :breadcrumbs="breadcrumbs"></BaseBreadcrumb>

  <v-row>
    <!-- ---------------------------------------------------------------- -->
    <!-- Todo List -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader title="To Do List" class="headerWithBtn">
        <template v-slot:header>
          <v-btn color="primary" variant="text" size="small" rounded icon>
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
      <WidgetCard title="Traffic Sources" subtitle="You’re getting more and more sources, keep it up!">
        <TrafficSources />
      </WidgetCard>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Team Members -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader title="Team Members" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <TeamMembers />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Messages -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="6">
      <CardHeader title="Latest Messages">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <Messages />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- User Activity -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="6">
      <CardHeader title="User Activity">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <UserActivity />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Projects -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="6">
      <CardHeader title="Projects" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <Projects />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Product Sales -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="6">
      <CardHeader title="Product Sales" class="overflow-hidden">
        <ProductSales />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Tasks -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="4" md="5">
      <CardHeader title="Tasks">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <Tasks />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Application Sales -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8" md="7">
      <CardHeader title="Application Sales" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <ApplicationSales />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Active Tickets -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="8">
      <CardHeader title="Active Tickets">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <ActiveTickets />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Posts -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader title="Latest Posts" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <LatestPosts />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Feeds -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="4" md="5">
      <CardHeader title="Feeds" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <Feeds />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Customers -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8" md="7">
      <CardHeader title="Latest Customers" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <LatestCustomer />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Latest Order -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12">
      <CardHeader title="Latest Order" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <LatestOrder />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Incoming Requests -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader title="Incoming Requests" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <IncomingRequest />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Total Revenue -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4" sm="6">
      <CardHeader title="Total Revenue" class="overflow-hidden">
        <TotalRevenue />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- New Customers -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" md="4">
      <CardHeader title="New Customers" class="overflow-hidden">
        <NewCustomers />
      </CardHeader>
    </v-col>

    <!-- ---------------------------------------------------------------- -->
    <!-- Recent Tickets -->
    <!-- ---------------------------------------------------------------- -->
    <v-col cols="12" lg="8">
      <CardHeader title="Recent Tickets" class="overflow-hidden">
        <template v-slot:header>
          <router-link to="/widget/data" class="text-h6 text-primary link-hover">View all</router-link>
        </template>
        <RecentTickets />
      </CardHeader>
    </v-col>
  </v-row>
</template>
