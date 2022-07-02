import { api } from 'boot/axios'

export const actions = {
  async createRestaurantUser(data) {
    console.log(data)
    const { email, password } = data
    try {
      const response = await api.post('/auth/signin', { email, password })
      if (response.statusCode === 200) {
        console.log(response)
      }
    } catch (e) {
      console.log(e)
    }
  },
}
