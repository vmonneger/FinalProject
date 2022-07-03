<script setup>
/**
 * @file Item clickable.
 */

const props = defineProps({
  size: {
    type: String,
    default: '13px',
  },
  disable: {
    type: Boolean,
    default: false,
  },
  label: {
    type: String,
  },
  iconLeft: {
    type: String,
  },
  iconRight: {
    type: String,
  },
})
</script>

<template>
  <q-item
    clickable
    class="text-primary items-center"
    active-class="text-primary q-item-drawer text-bold"
    :disable="props.disable"
  >
    <q-item-section avatar class="q-pr-xs" v-if="props.iconLeft">
      <q-icon :name="props.iconLeft" :size="props.size" align="left" />
    </q-item-section>
    <q-item-section v-if="props.label">
      {{ props.label }}
    </q-item-section>
    <q-item-section side v-if="props.iconRight">
      <q-icon :name="props.iconRight" :size="props.size" />
    </q-item-section>
  </q-item>
</template>
