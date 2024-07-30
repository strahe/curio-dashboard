/**
 * .eslint.js
 *
 * ESLint configuration file.
 */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'eslint:recommended',
    'plugin:vue/vue3-essential',
    'vuetify',
    '@vue/eslint-config-typescript/recommended',
  ],
  rules: {
    'vue/multi-word-component-names': 'off',
    'vue/script-indent': ['error', 2, { baseIndent: 0 }],
  },
}
