/**
 * @file Define the user store.
 */

import { defineStore } from 'pinia'
import { actions } from './actions'
import { getters } from './getters'
import { state } from './state'

export const useUserStore = defineStore('User', {
  state: state,
  getters: getters,
  actions: actions,
})
