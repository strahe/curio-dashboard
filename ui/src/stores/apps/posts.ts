import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';
import type { Reply } from '@/types/user/post';

export const usePostsStore = defineStore({
  id: 'post',
  state: () => ({
    posts: []
  }),
  getters: {},
  actions: {
    // Fetch Posts from action
    async fetchPosts() {
      try {
        const response = await axios.get('/api/posts/list');
        this.posts = response.data;
      } catch (error) {
        alert(error);
      }
    },
    // like post
    async likePost(postId: string) {
      try {
        const response = await axios.post('/api/posts/list/like', { postId });
        this.posts = response.data.posts;
      } catch (error) {
        alert(error);
      }
    },
    // add Comment
    async addComment(postId: string, comment: Reply) {
      try {
        const response = await axios.post('/api/comments/add', { postId, comment });
        this.posts = response.data.posts;
      } catch (error) {
        alert(error);
      }
    },
    // add Comment
    async addReply(postId: string, commentId: string, reply: Reply) {
      try {
        const response = await axios.post('/api/replies/add', { postId, commentId, reply });
        this.posts = response.data.posts;
      } catch (error) {
        alert(error);
      }
    }
  }
});
