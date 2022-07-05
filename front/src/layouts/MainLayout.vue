<script setup>
/**
 * @file Main layout contains sidebar and header.
 */
import { ref } from 'vue'
import TheDropdownProfile from '../components/molecules/TheDropdownProfile.vue'
import AppItem from '../components/atoms/AppItem.vue'
import { useRestaurantStore } from '../stores/restaurant'

const essentialLinks = [
  {
    title: 'Acceuil',
    icon: 'fa-solid fa-house',
    link: 'RestaurantHome',
  },
  {
    title: 'Votre restaurant',
    icon: 'fa-solid fa-utensils',
    link: 'RestaurantInfo',
  },
  {
    title: 'Modifier votre carte',
    icon: 'fa-regular fa-square-plus',
    link: 'RestaurantMenu',
  },
  {
    title: 'Voir votre carte',
    icon: 'fa-solid fa-book-open-reader',
    link: 'RestaurantShow',
  },
]

const storeRestaurant = useRestaurantStore()

const leftDrawerOpen = ref(false)

const toggleLeftDrawer = () => {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>

<template>
  <q-layout view="hHh lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />
        <q-toolbar-title class="text-white"> Orderizi </q-toolbar-title>

        <TheDropdownProfile :dropdownData="{ name: storeRestaurant.getRestaurantName }" />
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" class="bg-accent" show-if-above :width="265" :breakpoint="850">
      <q-list class="q-mt-xl">
        <div v-for="link in essentialLinks" :key="link.title" v-bind="link">
          <AppItem :to="{ name: link.link }" :iconLeft="link.icon" size="20px" :label="link.title" class="text-white" />
        </div>
      </q-list>
    </q-drawer>

    <q-page-container class="">
      <q-page class="q-pa-lg bg-secondary">
        <router-view />
      </q-page>
    </q-page-container>
  </q-layout>
</template>
