<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" lg="4">
        <TaskSummaryByDay color="info"></TaskSummaryByDay>
      </v-col>
      <v-col cols="12" lg="4">
        <MiningSummaryByDay color="info"></MiningSummaryByDay>
      </v-col>
      <v-col cols="12" lg="4">
        <RunningTaskPieChart type="Pie"></RunningTaskPieChart>
      </v-col>

      <v-col cols="12" lg="3">
        <StatsCard
          color="green"
          icon="mdi-firework"
          title="Running Tasks"
          value="223"
          sub-icon="mdi-calendar"
          sub-text="Last 24 Hours"
        />
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          color="orange"
          icon="mdi-sd"
          title="24H Storage Power"
          value="+654"
          small-value="GB"
          sub-icon="mdi-alert"
          sub-icon-color="error"
          sub-text="Get More Space..."
          sub-text-color="text-primary"
        />
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          color="red"
          icon="mdi-gift"
          title="24H Mining Blocks"
          value="981"
          sub-icon="mdi-tag"
          sub-text="Tracked from Github"
        />
      </v-col>
      <v-col cols="12" lg="3">
        <StatsCard
          color="info"
          icon="mdi-history"
          title="24H Total Tasks"
          value="432321"
          sub-icon="mdi-update"
          sub-text="Just Updated"
        />
      </v-col>
      <v-col cols="12" lg="6">
          <ChainConnectivity title="Chain Connectivity"/>
      </v-col>
      <v-col cols="12" lg="6">
        <RecentTasks></RecentTasks>
      </v-col>
      <v-col cols="12">
        <PipelineSummary title="Pipelines"></PipelineSummary>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts" setup>
import PipelineSummary from "@/widgets/PipelineSummary.vue";
import ChainConnectivity from "@/widgets/ChainConnectivity.vue";
import TaskSummaryByDay from "@/widgets/TaskSummaryByDay.vue";
import MiningSummaryByDay from "@/widgets/MiningSummaryByDay.vue";
import RunningTaskPieChart from "@/widgets/RunningTaskChart.vue";
import RecentTasks from "@/widgets/RecentTasks.vue";
</script>
