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
        api.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`
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

  async queryGetRestaurantInfo() {
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

  async queryPostResataurantCategory(data) {
    let categories = []
    for (let i = 0; i < data.categoriesMenu.value.length; i++) {
      categories.push(data.categoriesMenu.value[i].category)
    }
    try {
      const response = await api.post('/restaurant/category', { name: categories })
      if (response.status === 201) {
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
    let storeMenuFilter = []
    if (this.menu?.menu) {
      storeMenuFilter = [...this.menu.menu]
    }

    // A comparer used to determine if two entries are equal.
    const isSameMenuItem = (a, b) => a.title === b.title && a.category === b.category

    // Get items that only occur in the left array,
    // using the compareFunction to determine equality.
    const onlyInLeft = (left, right, compareFunction) =>
      left.filter((leftValue) => !right.some((rightValue) => compareFunction(leftValue, rightValue)))

    const onlyInA = onlyInLeft(menuItems.value, storeMenuFilter, isSameMenuItem)

    const result = [...onlyInA]

    try {
      const response = await api.post('/restaurant/menu', { menu: result })
      if (response.status === 201) {
        this.$patch((state) => {
          state.menu.menu = [...response.data.data]
        })
      }
    } catch (e) {
      throw new Error(e)
    }
  },

  async queryDeleteResataurantMenu(data) {
    try {
      const response = await api.delete('/restaurant/menu', {
        data: { category: data.category, description: data.description, title: data.title },
      })
      if (response.status === 201) {
        if (response.data.ModifiedCount > 0) {
          this.$patch((state) => {
            state.menu.menu = state.menu.menu.filter(
              (menuItem) =>
                menuItem.title !== data.title &&
                menuItem.description !== data.description &&
                menuItem.category !== data.category
            )
          })
        }
      }
    } catch (e) {
      throw new Error(e)
    }
  },
}
