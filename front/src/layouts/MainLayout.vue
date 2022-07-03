<script setup>
/**
 * @file Main layout contains sidebar and header.
 */
import { ref } from 'vue'
import TheDropdownProfile from '../components/molecules/TheDropdownProfile.vue'
import AppItem from '../components/atoms/AppItem.vue'

const essentialLinks = [
  {
    title: 'Modifier votre carte',
    icon: 'fa-regular fa-square-plus',
  },
  {
    title: 'Voir votre carte',
    icon: 'fa-solid fa-book-open-reader',
  },
]

const dropdownData = { name: 'tessst' }

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

        <TheDropdownProfile :dropdownData="dropdownData" />
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" class="bg-accent" show-if-above :width="265" :breakpoint="850">
      <q-list class="q-mt-xl">
        <div v-for="link in essentialLinks" :key="link.title" v-bind="link">
          <AppItem :iconLeft="link.icon" size="20px" :label="link.title" class="text-white" />
        </div>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>
