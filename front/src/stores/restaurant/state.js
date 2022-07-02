/**
 * @file State of the user store.
 */

/**
 * Initialize the state of the restaurant store.
 *
 * @returns {RestaurantState} - The restaurant state.
 */
export const state = () => ({
  uuid: '',
  defaultLanguage: '',
  identity: {},
  managements: [],
  files: [],
  avatarFile: {
    uuid: '',
    path: '',
    src: '',
    loading: false,
  },
  loadingFile: {
    status: false,
    fileType: '',
  },
})
