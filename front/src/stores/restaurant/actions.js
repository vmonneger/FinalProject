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
      console.log(response)
      if (response.status === 201) {
        localStorage.setItem('token', response.data.data.data.token)
      }
    } catch (e) {
      console.log(e)
    }
  },
}
