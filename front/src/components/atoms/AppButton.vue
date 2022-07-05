<script setup>
/**
 * @file Component for button.
 */
const props = defineProps({
  color: {
    type: String,
    default: 'accent',
  },
  size: {
    type: String,
    default: 'md',
  },
  label: String,
})
</script>

<template>
  <q-btn :color="props.color" :label="props.label" :size="props.size" no-caps unelevated>
    <template v-slot:loading>
      <q-spinner-tail />
    </template>
  </q-btn>
</template>
