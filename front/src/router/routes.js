const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '/home/restaurant',
        name: 'HomeRestaurant',
        component: () => import('pages/RestaurantHome.vue'),
        meta: {
          requiresRestaurantAuth: true,
        },
      },
    ],
  },
  {
    path: '/',
    component: () => import('layouts/DefaultLayout.vue'),
    children: [
      {
        path: '/auth/register',
        name: 'Register',
        component: () => import('pages/RestaurantRegister.vue'),
      },
      {
        path: '/auth/login',
        name: 'Login',
        component: () => import('pages/RestaurantLogin.vue'),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
]

export default routes
