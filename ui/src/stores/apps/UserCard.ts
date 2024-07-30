import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

export const useUserCardStore = defineStore({
  id: 'userCard',
  state: () => ({
    cards: [],
    list: [],
    list2: []
  }),
  getters: {},
  actions: {
    // Fetch cards from action
    async fetchCards() {
      try {
        const response = await axios.get('/api/details-card/list');
        this.cards = response.data;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch filtered mail from action
    async filterCards(filter: string) {
      try {
        const response = await axios.post('/api/details-card/filter', { filter });
        this.cards = response.data.results;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch cards from action
    async fetchlistCards() {
      try {
        const response = await axios.get('/api/avatar-list/s1/list');
        this.list = response.data;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch cards from action
    async fetchlist2Cards() {
      try {
        const response = await axios.get('/api/avatar-list/s2/list');
        this.list2 = response.data;
      } catch (error) {
        alert(error);
      }
    }
  }
});
