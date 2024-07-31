const AppRoutes = {
  path: '/app',
  meta: {
    requiresAuth: true,
  },
  component: () => import('@/layouts/DashboardLayout.vue'),
  children: [
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
