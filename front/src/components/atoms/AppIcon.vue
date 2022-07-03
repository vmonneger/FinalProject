<script setup>
/**
 * @file Component for icon.
 */

const props = defineProps({
  color: {
    type: String,
    default: 'primary',
  },
  size: {
    type: String,
    default: 'xs',
  },
  hoverEffect: {
    type: Boolean,
    default: false,
  },
  hoverText: {
    type: String,
    default: '',
  },
  name: {
    type: String,
  },
})
</script>

<template>
  <q-icon :name="props.name" :color="props.color" :size="props.size" :class="{ 'hover-effect': props.hoverEffect }">
  </q-icon>
</template>
<style lang="scss">
.q-icon {
  &.text-primary {
    color: $primary !important;
  }
  &.text-secondary {
    color: $secondary !important;
  }
  &.text-negative {
    color: $negative !important;
  }
  &.text-positive {
    color: $positive !important;
  }
  &.text-warning {
    color: $warning !important;
  }
  &.text-greyLight {
    color: $greyLight !important;
  }
}

.q-icon {
  &.hover-effect {
    background: transparent;
    cursor: pointer;
    padding: 12px;
    border-radius: 50px;
    transition: 0.25s;
  }
  &.hover-effect:hover {
    background: $greyLight;
  }
}
</style>
