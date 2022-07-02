<script setup>
/**
 * @file Component for input.
 */
import { ref, watchEffect, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'text',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  lazyRules: {
    type: Boolean,
    default: false,
  },
  loading: {
    type: Boolean,
    default: false,
  },
  resetText: {
    type: Boolean,
    default: false,
  },
  fullWidth: {
    type: Boolean,
    default: false,
  },
  bgColor: {
    type: String,
    default: 'text',
  },
  label: String,
})
const emit = defineEmits(['update:modelValue'])

const showPassword = ref(false)
const inputType = ref('')

const value = computed({
  /**
   * Getter (to update the value).
   *
   * @returns {void}
   */
  get() {
    return props.modelValue
  },
  /**
   * Setter (to update the value).
   *
   * @param {void} value - Value.
   */
  set(value) {
    emit('update:modelValue', value)
  },
})

watchEffect(() => {
  inputType.value = props.type
  if (props.type === 'password') {
    if (showPassword.value) {
      inputType.value = 'text'
    } else {
      inputType.value = 'password'
    }
  }
})
</script>

<template inheritsAttr="false">
  <label v-show="props.label" class="app-input-label">
    {{ props.label }}
  </label>
  <q-input
    class="app-input"
    :type="inputType"
    :for="props.name"
    :name="props.name"
    :id="props.name"
    :placeholder="props.placeholder"
    :lazy-rules="props.lazyRules"
    :loading="props.loading"
    :rules="props.rules"
    :bg-color="props.bgColor"
    v-model="value"
    v-bind="$attrs"
    outlined
    no-error-icon
    @input="emit('update:modelValue', $event.target.value)"
  >
    <template v-slot:prepend v-if="props.iconLeft">
      <q-icon v-show="props.iconLeft" color="dark" :name="props.iconLeft" size="xs" />
    </template>
    <template v-slot:append v-if="props.iconRight || props.type === 'password'">
      <q-icon
        v-show="props.type === 'password'"
        class="cursor-pointer q-pa-lg"
        color="primary"
        :name="showPassword ? 'fa-regular fa-eye' : 'fa-regular fa-eye-slash'"
        @click="showPassword = !showPassword"
      />
      <q-icon
        v-show="props.resetText"
        class="cursor-pointer q-pa-lg"
        color="primary"
        name="fa fa-times"
        @click="$emit('update:modelValue', '')"
      />
      <q-icon v-show="props.iconRight" color="primary" :name="props.iconRight" size="xs" />
    </template>
  </q-input>
</template>
