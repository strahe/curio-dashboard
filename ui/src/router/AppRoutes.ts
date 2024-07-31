const AppRoutes = {
  path: '/app',
  meta: {
    requiresAuth: true,
  },
  component: () => import('@/layouts/DashboardLayout.vue'),
  children: [
    {
      name: 'Overview',
      path: '/app/overview',
      component: () => import('@/pages/overview/index.vue'),
    },
    {
      name: 'Analytics',
      path: '/app/analytics',
      component: () => import('@/pages/analytics/index.vue'),
    },

    {
      name: 'Machines',
      path: '/app/machines',
      component: () => import('@/pages/machines/index.vue'),
    },
    {
      name: 'Storages',
      path: '/app/storages',
      component: () => import('@/pages/storages/index.vue'),
    },
    {
      name: 'Pipelines',
      path: '/app/pipelines',
      component: () => import('@/pages/pipelines/index.vue'),
    },
    {
      name: 'Configurations',
      path: '/app/configurations',
      component: () => import('@/pages/configurations/index.vue'),
    },
  ],
}

export default AppRoutes
