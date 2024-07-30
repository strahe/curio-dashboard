import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

export const useFrinedsStore = defineStore({
  id: 'Frineds',
  state: () => ({
    friends: []
  }),
  getters: {},
  actions: {
    // Fetch followers from action
    async fetchFrineds() {
      try {
        const response = await axios.get('/api/friends/list');
        this.friends = response.data.friends;
      } catch (error) {
        alert(error);
      }
    }
  }
});
