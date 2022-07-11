<script setup>
/**
 * @file Page for restaurant info.
 */
import FormRestaurantInfo from '../components/molecules/FormRestaurantInfo.vue'
import AppCard from '../components/atoms/AppCard.vue'
import AppTitle from '../components/atoms/AppTitle.vue'
</script>

<template>
  <AppCard class="bg-secondary">
    <AppTitle class="q-pb-lg"> Renseigner les informations de votre restaurant </AppTitle>
    <FormRestaurantInfo />
  </AppCard>
</template>
