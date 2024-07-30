import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

interface Task {
  id: number;
  description: string;
  userStory: string;
  dueDate: Date;
  priority: string;
  title: string;
}

interface Column {
  id: number;
  title: string;
  tasks: Task[];
}

export const useKanbanStore = defineStore({
  id: 'kanban',
  state: () => ({
    kanbanItems: [] as Column[]
  }),
  getters: {},
  actions: {
    // Fetch Chat from action
    async fetchkanbanItems() {
      try {
        const response = await axios.get('/api/data/kanban');
        this.kanbanItems = response.data;
      } catch (error) {
        alert(error);
      }
    },
    addTask(itemID: number, newTask: Task) {
      const column = this.kanbanItems.find((col) => col.id === itemID);
      if (column) {
        column.tasks.push(newTask);
      }
    }
  }
});
