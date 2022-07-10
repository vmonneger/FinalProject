import { api } from 'boot/axios'
import { notification } from '../../helpers/notifications'

export const actions = {
  async registerRestaurantUser(data) {
    const { email, password } = data
    try {
      await api.post('/auth/signin', { email, password })
    } catch (e) {
      if (e.response.data?.msg.includes('E11000')) notification('Cet email est déjà enregistré', 'error')
      notification('Erreur veuillez contacter le support', 'error')
      throw new Error(e)
    }
  },

  async loginRestaurantUser(data) {
    const { email, password } = data
    try {
      const response = await api.post('/auth/login', { email, password })
      if (response.status === 201) {
        localStorage.setItem('token', response.data.token)
      }
    } catch (e) {
      notification('Vérifiez votre email ou votre mot de passe', 'error')
      throw new Error(e)
    }
  },

  async logoutRestaurant() {
    localStorage.clear()
    this.$reset()
    this.router.push({ name: 'Login' })
  },

  async queryGetResataurantInfo() {
    try {
      const response = await api.get('/restaurant')
      console.log('get', response)
      this.$patch({
        ...response?.data?.data,
      })
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryPostResataurantInfo(data) {
    const { name, description } = data
    try {
      const response = await api.post('/restaurant', { name, description })
      if (response.status === 201) {
        this.$patch((state) => {
          ;(state.name = response.data.data.name.replace(/\s*$/, '')),
            (state.description = response.data.data.description.replace(/\s*$/, ''))
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryPostResataurantCategory(data) {
    let categories = []
    for (let i = 0; i < data.categoriesMenu.value.length; i++) {
      categories.push(data.categoriesMenu.value[i].category)
    }
    try {
      const response = await api.post('/restaurant/category', { name: categories })
      if (response.status === 201) {
        console.log(response)
        this.$patch((state) => {
          state.category = response.data.data
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryDeleteResataurantCategory(data) {
    try {
      const response = await api.delete('/restaurant/category', { data: { name: [data] } })
      if (response.status === 201) {
        this.$patch((state) => {
          state.category = state.category.filter((category) => category !== data)
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryPostResataurantMenu(data) {
    const { menuItems } = data
    console.log(this.menu)
    let storeMenuFilter = []
    if (this.menu.menu) {
      storeMenuFilter = this.menu.menu.filter((menu) => menu.category !== menuItems.value[0].category)
    }

    const newMenu = [...menuItems.value, ...storeMenuFilter]
    console.log(newMenu)

    try {
      const response = await api.post('/restaurant/menu', { menu: newMenu })
      if (response.status === 201) {
        console.log(response)
        this.$patch((state) => {
          state.menu = response.data.data.menu
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryDeleteResataurantMenu(data) {
    console.log(data)
    try {
      const response = await api.delete('/restaurant/menu', {
        data: { category: data.category, description: data.description, title: data.title },
      })
      if (response.status === 201) {
        console.log(response)
        this.$patch((state) => {
          state.menu = state.menu.filter((category) => category !== data)
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },
}
