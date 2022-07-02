/**
 * @file This file contains quasar related rules for inputs.
 * @see https://quasar.dev/vue-components/input
 */

import { regexSpecialCharacters, regexSpecialCharactersLight } from './regex'

/**
 * Rule to use in an input check if the value is set and if its length is not 0.
 *
 * @param {string} val - Input value.
 * @returns {(number|boolean)} - True if the rule is ok or string error message.
 */
export const ruleInputRequired = [(val) => (val && val.length > 0) || 'Ce champs est requis']

/**
 * Rule to use in an input.
 * Check if input contains special characters (refer to regexSpecialCharacters).
 *
 * @param {string} val - Input value.
 * @returns {boolean} True or false.
 * Rule that checks if input contains special characters.
 */
export const ruleSpecialCharacters = [
  (val) => !regexSpecialCharacters.test(val) || 'Les charactères spéciaux ne sont pas autorisés',
]

/**
 * Rule that checks if input contains a subset of special characters.
 *
 * @param {string} val - Input value.
 * @returns {(number|boolean)} - True if the rule is ok or string error message.
 */
export const ruleSpecialCharactersLight = [
  (val) => !regexSpecialCharactersLight.test(val) || 'Les charactères spéciaux ne sont pas autorisés',
]

/**
 * Rule that checks if if input is type of email.
 *
 * @param {string} val - Input value.
 * @returns {(number|boolean)} - True if the rule is ok or string error message.
 */
export const ruleEmail = [(val) => (val.includes('@') && val.includes('.')) || "L'email n'est pas valide"]

/**
 * Rule that checks if input is equal to the password confirmation.
 *
 * @param {string} password - Password Input value.
 * @returns {Array} - Used by quasar for rules.
 */
export const ruleVerifyPassword = (password) => [(val) => val === password || 'Le mot de passe est différent']
