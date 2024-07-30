import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

export const useFollowersStore = defineStore({
  id: 'followers',
  state: () => ({
    followers: []
  }),
  getters: {},
  actions: {
    // Fetch followers from action
    async fetchFollowers() {
      try {
        const response = await axios.get('/api/followers/list');
        this.followers = response.data.followers;
      } catch (error) {
        alert(error);
      }
    }
  }
});
