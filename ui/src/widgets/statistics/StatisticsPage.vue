<script setup lang="ts">
import { ref, shallowRef } from 'vue';
import BaseBreadcrumb from '@/components/shared/BaseBreadcrumb.vue';

// import
import WidgetFive from './components/WidgetFive.vue';
import TaskCard from './components/TaskCard.vue';
import SkillCard from './components/SkillCard.vue';

// icons
import {
  BarChartOutlined,
  CalendarOutlined,
  FileTextOutlined,
  DownloadOutlined,
  DollarCircleFilled,
  ScheduleFilled,
  ShoppingFilled,
  FacebookOutlined,
  TwitterOutlined,
  LinkedinOutlined,
  YoutubeOutlined,
  EyeOutlined,
  AimOutlined,
  FieldTimeOutlined,
  ContactsOutlined,
  FileProtectOutlined,
  RedditOutlined
} from '@ant-design/icons-vue';

const page = ref({ title: 'Statistics' });
const breadcrumbs = ref([
  {
    title: 'Widget',
    disabled: false,
    href: '#'
  },
  {
    title: 'Statistics',
    disabled: true,
    href: '#'
  }
]);

const cards = shallowRef([
  { earn: '$30200', text: 'All Earnings', icon: BarChartOutlined, color: 'secondary' },
  { earn: '145', text: 'Task', icon: CalendarOutlined, color: 'error' },
  { earn: '290+', text: 'Page Views', icon: FileTextOutlined, color: 'success' },
  { earn: '500', text: 'Downloads', icon: DownloadOutlined, color: 'primary' }
]);

const twocards = shallowRef([
  { title: 'Revenue', earn: '$4,500', text: '$50,032 Last Month', icon: DollarCircleFilled, color: 'primary' },
  { title: 'Orders Received', earn: '486', text: '20% Increase', icon: ScheduleFilled, color: 'warning' },
  { title: 'Total Sales', earn: '1641', text: '$1,055 Revenue Generated', icon: ShoppingFilled, color: 'success' }
]);

const threecards = shallowRef([
  { earn: '1165 +', text: 'Facebook Users', icon: FacebookOutlined, color: 'primary' },
  { earn: '780 +', text: 'Twitter Users', icon: TwitterOutlined, color: 'info' },
  { earn: '998 +', text: 'Linked In Users', icon: LinkedinOutlined, color: 'darkText' },
  { earn: '650 +', text: 'Youtube Videos', icon: YoutubeOutlined, color: 'error' }
]);

const fourcards = shallowRef([
  { earn: '1,563', text: 'Impressions', icon: EyeOutlined, color: 'primary', color2: 'lightprimary', duedate: 'May 23 - June 01 (2018)' },
  { earn: '30,564', text: 'Goal', icon: AimOutlined, color: 'success', color2: 'lightsuccess', duedate: 'May 28 - June 01 (2018)' },
  { earn: '42.6%', text: 'Impact', icon: FieldTimeOutlined, color: 'warning', color2: 'lightwarning', duedate: 'May 30 - June 01 (2018)' }
]);

const sixcards = shallowRef([
  { earn: '1,658', text: 'Daily user', icon: ContactsOutlined, color: 'success' },
  { earn: '1K', text: 'Daily page view', icon: FileProtectOutlined, color: 'primary' },
  { earn: '5,678', text: 'Last month visitor', icon: RedditOutlined, color: 'warning' }
]);
</script>

<template>
  <BaseBreadcrumb :title="page.title" :breadcrumbs="breadcrumbs"></BaseBreadcrumb>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 1 -->
  <!-- --------------------------------------------------------------------- -->
  <v-row class="mb-0">
    <v-col cols="12" sm="6" md="3" v-for="(card1, i) in cards" :key="i" :value="card1">
      <v-card elevation="0">
        <v-card variant="outlined">
          <v-card-text>
            <div class="d-flex align-items-center justify-space-between">
              <div>
                <h4 class="text-h4 mb-2">{{ card1.earn }}</h4>
                <span class="text-medium-emphasis">{{ card1.text }}</span>
              </div>
              <span class="d-flex align-center">
                <component :is="card1.icon" :style="{ fontSize: '30px' }" :class="'text-' + card1.color" />
              </span>
            </div>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 2 -->
  <!-- --------------------------------------------------------------------- -->
  <v-row class="mb-0">
    <v-col cols="12" md="4" v-for="(card2, i) in twocards" :key="i" :value="card2">
      <v-card elevation="0" :class="`bg-${card2.color} widget-gradient overflow-hidden`">
        <v-card variant="outlined">
          <v-card-text class="px-md-9 py-md-8 px-5 py-5">
            <div class="d-flex justify-space-between">
              <span class="d-flex">
                <component :is="card2.icon" :style="{ fontSize: '56px', opacity: 0.23 }" class="text-white opacity-50" />
              </span>
              <div class="text-right">
                <h4 class="text-h4 font-weight-medium mb-2">{{ card2.title }}</h4>
                <h2 class="text-h2 mb-4">{{ card2.earn }}</h2>
                <span class="text-medium-emphasis text-h5 font-weight-regular text-white">{{ card2.text }}</span>
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 3-->
  <!-- --------------------------------------------------------------------- -->
  <v-row class="mb-0">
    <v-col cols="12" sm="6" md="3" v-for="(card3, i) in threecards" :key="i" :value="card3">
      <v-card elevation="0" :class="'hover-icon-card bg-' + card3.color">
        <v-card elevation="0" :class="'bg-' + card3.color">
          <v-card-text>
            <div class="d-flex align-items-center justify-space-between">
              <div>
                <h3 class="text-h3 text-white mb-0">{{ card3.earn }}</h3>
                <span class="text-medium-emphasis text-caption font-weight-medium text-white">{{ card3.text }}</span>
              </div>
              <span class="d-flex align-center">
                <component :is="card3.icon" :style="{ fontSize: '36px' }" class="text-white opacity-50 hover-icon" />
              </span>
            </div>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 4 -->
  <!-- --------------------------------------------------------------------- -->
  <v-row class="mb-0">
    <v-col cols="12" md="4" v-for="(card4, i) in fourcards" :key="i" :value="card4">
      <v-card elevation="0">
        <v-card variant="outlined">
          <v-card-text>
            <div class="d-flex align-items-center justify-space-between">
              <div>
                <h5 class="text-h5">{{ card4.text }}</h5>
                <h3 class="text-h3 my-2">{{ card4.earn }}</h3>
                <h6 class="text-caption font-weight-medium mb-0">{{ card4.duedate }}</h6>
              </div>
              <span class="d-flex align-center">
                <v-btn :class="'text-' + card4.color" :color="card4.color2" rounded="md" icon variant="flat">
                  <component :is="card4.icon" :style="{ fontSize: '20px' }" />
                </v-btn>
              </span>
            </div>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 5 -->
  <!-- --------------------------------------------------------------------- -->
  <WidgetFive />

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 6 -->
  <!-- --------------------------------------------------------------------- -->
  <v-row class="mb-0">
    <v-col cols="12" md="4" v-for="(card6, i) in sixcards" :key="i" :value="card6">
      <v-card elevation="0" :class="'overflow-hidden position-relative bg-' + card6.color">
        <v-card :class="'text-center bg-' + card6.color">
          <component :is="card6.icon" :style="{ fontSize: '110px', opacity: 0.35 }" class="text-white icon-scale" />
          <v-card-text>
            <h3 class="text-h3 my-2">{{ card6.earn }}</h3>
            <span>{{ card6.text }}</span>
          </v-card-text>
        </v-card>
      </v-card>
    </v-col>
  </v-row>

  <!-- --------------------------------------------------------------------- -->
  <!-- Cards 7 -->
  <!-- --------------------------------------------------------------------- -->
  <v-row>
    <v-col cols="12" lg="7">
      <SkillCard />
    </v-col>
    <v-col cols="12" lg="5">
      <TaskCard />
    </v-col>
  </v-row>
</template>
