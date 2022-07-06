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

  async queryPostResataurantMenu(data) {
    let categories = []
    for (let i = 0; i < data.categoriesMenu.value.length; i++) {
      categories.push(data.categoriesMenu.value[i].category)
    }
    try {
      const response = await api.post('/restaurant/category', { category: categories })
      if (response.status === 201) {
        this.$patch((state) => {
          state.category = response.data.data.category
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },
}
