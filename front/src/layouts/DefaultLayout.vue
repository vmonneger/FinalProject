<script setup>
/**
 * @file Default layout.
 */
</script>
<template>
  <q-layout class="h-100 bg-secondary items-center">
    <q-scroll-area style="height: 100vh; max-width: 100%">
      <router-view />
    </q-scroll-area>
  </q-layout>
</template>
