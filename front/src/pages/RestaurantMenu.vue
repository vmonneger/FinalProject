<script setup>
/**
 * @file Page for restaurant menu.
 */
import FormRestaurantCategory from '../components/molecules/FormRestaurantCategory.vue'
import AppCard from '../components/atoms/AppCard.vue'
import AppTitle from '../components/atoms/AppTitle.vue'
import TabPanelRestaurantMenu from '../components/molecules/TabPanelRestaurantMenu.vue'
import { useRestaurantStore } from '../stores/restaurant'

const storeRestaurant = useRestaurantStore()
</script>

<template>
  <div class="row q-col-gutter-md">
    <div class="col-xs-8 col-md-8 col-lg-6 col-12">
      <AppCard class="bg-secondary">
        <AppTitle class="q-pb-lg"> Renseigner les catégories de votre carte </AppTitle>
        <p>Entrez une catégorie par champs</p>
        <FormRestaurantCategory />
      </AppCard>
    </div>
    <div v-if="storeRestaurant.category.length > 0" class="col-xs-12 col-lg-6">
      <AppCard class="bg-secondary">
        <AppTitle class="q-pb-lg"> Compléter votre carte </AppTitle>
        <TabPanelRestaurantMenu :tabName="storeRestaurant.category"> </TabPanelRestaurantMenu>
      </AppCard>
    </div>
  </div>
</template>
