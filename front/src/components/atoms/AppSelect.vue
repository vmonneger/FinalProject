<script setup>
/**
 * @file Component for dropdown select.
 */

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false,
  },
  lazyRules: {
    type: Boolean,
    default: false,
  },
  placeholder: {
    type: String,
    default: '',
  },
  useInput: {
    type: Boolean,
    default: false,
  },
  options: {
    type: Array,
    default: () => [],
  },
})
</script>

<template>
  <q-select
    class="app-select"
    :options="props.options"
    :disable="props.disabled"
    :lazy-rules="props.lazyRules"
    :use-input="props.useInput"
    no-error-icon
    outlined
    :popup-content-style="{ width: '100px' }"
  />
</template>

<style lang="scss">
.app-select {
  color: $primary !important;
  .q-field__control {
    background-color: $greyLight;
  }
  &.q-field--outlined {
    .q-field__control::before {
      border-color: $greyLight;
    }
    .q-field__control::after {
      border-color: $greyLight;
    }
  }
  .q-icon {
    color: $primary !important;
  }
}
.q-virtual-scroll__content .q-item.q-item-type {
  background-color: $greyLight;
}

.app-select-label {
  color: $primary !important;
  font-weight: 600;
}
</style>
