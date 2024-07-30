const AppRoutes = {
  path: '/app',
  meta: {
    requiresAuth: true,
  },
  component: () => import('@/layouts/dashboard/DashboardLayout.vue'),
  children: [
    {
      name: 'Machines',
      path: '/app/machines',
      component: () => import('@/pages/machines/index.vue'),
    },
    {
      name: 'Pipelines',
      path: '/app/pipelines',
      component: () => import('@/pages/pipelines/index.vue'),
    },
  ],
}

export default AppRoutes
