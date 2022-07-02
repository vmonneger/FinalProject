/**
 * @file Helper for notification manipulation.
 * Notifications are small messages that pops on the UI and disapear after several seconds
 * to inform the user that something happened.
 */

import { Notify } from 'quasar'

/**
 * Display a notification.
 *
 * @param {string} message - Message of the notification.
 * @param {string} type - Type of notification: 'error' | 'warning' | 'info' | 'primary'.
 * @returns {void}.
 */
export const notification = (message, type) => {
  if (type === 'error') {
    return Notify.create({
      color: 'negative',
      textColor: 'white',
      icon: 'fa-solid fa-triangle-exclamation',
      message: message,
    })
  }
  if (type === 'warning') {
    return Notify.create({
      color: 'warning',
      textColor: 'white',
      icon: 'fa-solid fa-triangle-exclamation',
      message: message,
    })
  }
  if (type === 'info') {
    return Notify.create({
      color: 'info',
      textColor: 'white',
      message: message,
    })
  }
  if (type === 'primary')
    return Notify.create({
      color: 'primary',
      textColor: 'white',
      icon: 'fa-regular fa-circle-check',
      message: message,
    })
}

/**
 * Display a notification with the "saved" message.
 *
 * @returns {void}
 */
export const notificationSaved = () => {
  return Notify.create({
    color: 'primary',
    textColor: 'white',
    icon: 'fa-regular fa-circle-check',
    message: 'Enregistré',
  })
}
