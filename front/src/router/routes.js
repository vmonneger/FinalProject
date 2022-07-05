const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '/home/restaurant',
        name: 'RestaurantHome',
        component: () => import('pages/RestaurantHome.vue'),
      },
      {
        path: '/restaurant/info',
        name: 'RestaurantInfo',
        component: () => import('pages/RestaurantInfo.vue'),
      },
      {
        path: '/restaurant/menu',
        name: 'RestaurantMenu',
        component: () => import('pages/RestaurantMenu.vue'),
      },
      {
        path: '/restaurant/show',
        name: 'RestaurantShow',
        component: () => import('pages/RestaurantShow.vue'),
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
