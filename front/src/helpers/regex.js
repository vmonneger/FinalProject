/**
 * @file Regex that can be reused by all the app.
 */

/**
 * Check special character `*+?|{[()^$.,#`.
 */
export const regexSpecialCharacters = /[*+?|{[()^$.,#]/g

/**
 * Check special character `*+|{[^`.
 */
export const regexSpecialCharactersLight = /[|{[^]/g
