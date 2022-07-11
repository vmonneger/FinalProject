import { route } from 'quasar/wrappers'
import { createRouter, createMemoryHistory, createWebHistory, createWebHashHistory } from 'vue-router'
import routes from './routes'
import { notification } from '../helpers/notifications'
import { useRestaurantStore } from '../stores/restaurant/index'

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === 'history'
    ? createWebHistory
    : createWebHashHistory

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  })
  Router.beforeEach(async (to) => {
    const storeRestaurant = useRestaurantStore()

    const publicPages = ['/login', '/register']
    const authRequired = !publicPages.includes(to.path)
    const isLogin = localStorage.getItem('token')

    if (authRequired && !isLogin && to.name !== 'Login' && to.name !== 'Register') {
      notification("Vous n'êtes pas connecté", 'warning')
      return { name: 'Login' }
    }

    if (isLogin && (to.name === 'Login' || to.name === 'Register')) {
      return { name: 'RestaurantHome' }
    }

    if (isLogin) {
      if (!storeRestaurant.id) {
        await storeRestaurant.queryGetRestaurantInfo()
      }
    }
  })
  return Router
})
