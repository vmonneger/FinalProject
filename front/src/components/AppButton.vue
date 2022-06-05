<script setup>
const props = defineProps({
  size: String,
  label: String,
})
</script>
<template>
  <q-btn color="primary" :label="props.label" :size="props.size" />
</template>
