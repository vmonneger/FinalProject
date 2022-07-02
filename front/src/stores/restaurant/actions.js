import { api } from 'boot/axios'

export const actions = {
  async registerRestaurantUser(data) {
    console.log(data)
    const { email, password } = data
    try {
      const response = await api.post('/auth/signin', { email, password })
      if (response.status === 201) {
        console.log(response)
      }
    } catch (e) {
      console.log(e)
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
