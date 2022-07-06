<script setup>
import FormRestaurantCategory from '../components/molecules/FormRestaurantCategory.vue'
import AppCard from '../components/atoms/AppCard.vue'
import AppTitle from '../components/atoms/AppTitle.vue'
</script>

<template>
  <div class="row">
    <AppCard class="col-xs-8 col-md-8 col-lg-6 col-12 bg-secondary">
      <AppTitle class="q-pb-lg"> Renseigner les catégories de votre menu </AppTitle>
      <p>Entrez une catégorie par champs</p>
      <FormRestaurantCategory />
    </AppCard>
  </div>
</template>
