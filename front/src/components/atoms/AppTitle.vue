<script setup>
/**
 * @file Component to display a title.
 */

const titles = {
  xl: 'h1',
  lg: 'h2',
  md: 'h3',
  sm: 'h4',
  xs: 'h5',
}
const props = defineProps({
  weight: {
    type: String,
    default: 'regular',
  },
  size: {
    type: String,
    default: 'md',
  },
  color: {
    type: String,
    default: 'primary',
  },
})
</script>
<template>
  <component :is="titles[props.size]" class="app-title" :class="`title-${props.size} ${props.color}  ${props.weight} `">
    <slot></slot
  ></component>
</template>
<style lang="scss">
.app-title {
  color: inherit;
  margin: 0;
  &.primary {
    color: $primary !important;
  }
  &.secondary {
    color: $secondary !important;
  }
  &.negative {
    color: $negative !important;
  }
  &.positive {
    color: $positive !important;
  }
  &.warning {
    color: $warning !important;
  }
  &.white {
    color: white !important;
  }
  &.title-xl {
    font-size: 32pt;
    line-height: 40pt;
  }
  &.title-lg {
    font-size: 24pt;
    line-height: 32pt;
  }
  &.title-md {
    font-size: 16pt;
    line-height: 22pt;
  }
  &.title-sm {
    font-size: 12pt;
    line-height: 18pt;
  }
  &.title-xs {
    font-size: 10pt;
    line-height: 16pt;
  }
  &.regular {
    font-weight: 500;
  }
  &.bold {
    font-weight: 700;
  }
}
</style>
