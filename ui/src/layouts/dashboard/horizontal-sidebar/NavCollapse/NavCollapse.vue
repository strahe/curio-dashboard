<script setup>
import NavItem from '../NavItem/NavItem.vue';
// icons
import { RightOutlined } from '@ant-design/icons-vue';

const props = defineProps({ item: Object, level: Number });
</script>

<template>
  <!---Dropdown  -->
  <a class="navItemLink rounded-md cursor-pointer">
    <!---Icon  -->
    <i class="navIcon">
      <component :is="props.item.icon" :level="level" :style="{ fontSize: '16px' }"></component>
    </i>
    <!---Title  -->
    <span class="mr-auto">{{ item.title }}</span>
    <!---If Caption-->
    <small v-if="item.subCaption" class="text-caption mt-n1 hide-menu">
      {{ item.subCaption }}
    </small>
    <i class="ddIcon ml-2"><RightOutlined :style="{ fontSize: '15px' }" /></i>
  </a>
  <!---Sub Item-->
  <ul :class="`ddMenu px-0 ddLevel-${level + 1}`">
    <li v-for="(subitem, i) in item.children" :key="i" class="navItem rounded-0">
      <NavCollapse :item="subitem" v-if="subitem.children" :level="props.level + 1" />
      <NavItem :item="subitem" :level="props.level + 1" v-else></NavItem>
    </li>
  </ul>
  <!---End Item Sub Header -->
</template>
