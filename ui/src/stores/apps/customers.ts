import { defineStore } from 'pinia'
// project imports
import axios from '@/utils/axios'
// types
import type { CustomerStateProps } from '@/types/customers/index'

export const useCustomers = defineStore({
  id: 'customers',
  state: (): CustomerStateProps => ({
    customers: [],
    orders: [],
    products: [],
    productreviews: [],
  }),
  getters: {
    // Get Customers from Getters
    getCustomers (state) {
      return state.customers
    },
    // Get orders from Getters
    getOrders (state) {
      return state.orders
    },
    // Get orders from Getters
    getProducts (state) {
      return state.products
    },
    // Get orders from Getters
    getProductsreviews (state) {
      return state.productreviews
    },
  },
  actions: {
    // Fetch Customers from action
    async fetchCustomers () {
      try {
        const data = await axios.get('/api/data/customers')
        this.customers = data.data
      } catch (error) {
        alert(error)
      }
    },

    // Fetch Orders from action
    async fetchOrders () {
      try {
        const data = await axios.get('/api/data/orders')
        this.orders = data.data
      } catch (error) {
        alert(error)
      }
    },

    // Fetch products from action
    async fetchProducts () {
      try {
        const data = await axios.get('/api/data/products')
        this.products = data.data
      } catch (error) {
        alert(error)
      }
    },

    // Fetch products from action
    async fetchReviews () {
      try {
        const data = await axios.get('/api/data/productreviews')
        this.productreviews = data.data
      } catch (error) {
        alert(error)
      }
    },

    // Delete Customer
    deleteCustomer (itemId: string) {
      this.customers = this.customers.filter(object => {
        return object.name !== itemId
      })
    },

    // Delete Orders
    deleteOrder (itemId: string) {
      this.orders = this.orders.filter(object => {
        return object.id !== itemId
      })
    },
  },
})
