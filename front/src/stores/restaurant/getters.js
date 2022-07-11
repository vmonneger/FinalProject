export const getters = {
  getRestaurantName(state) {
    return state.name
  },

  getRestaurantCategory(state) {
    const categoryParse = state.category.map((category) => ({
      category: category,
    }))
    return categoryParse
  },

  getRestaurantMenu(state) {
    return state.menu?.menu ? state.menu?.menu : []
  },
}
