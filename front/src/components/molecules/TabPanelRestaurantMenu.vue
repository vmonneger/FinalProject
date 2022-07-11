<script setup>
/**
 * @file Component tab panel for restaurant menu.
 */
import { ref, watchEffect } from 'vue'
import FormRestaurantMenu from './FormRestaurantMenu.vue'
import { useRestaurantStore } from '../../stores/restaurant'

const props = defineProps({
  tabName: {
    type: Array,
    default: () => [],
  },
})

const storeRestaurant = useRestaurantStore()
const formData = ref()

const tab = ref()

watchEffect(() => {
  const menu = storeRestaurant.getRestaurantMenu
  formData.value = menu.filter((el) => el.category === tab.value)
})
</script>

<template>
  <q-tabs
    v-model="tab"
    dense
    class="text-grey"
    active-color="primary"
    indicator-color="primary"
    align="justify"
    narrow-indicator
  >
    <template v-for="(item, index) in props.tabName" :key="index">
      <q-tab :name="item" :label="item" />
    </template>
  </q-tabs>

  <q-separator />

  <q-tab-panels v-model="tab" animated class="bg-secondary">
    <template v-for="(item, index) in props.tabName" :key="index">
      <q-tab-panel :name="item">
        <FormRestaurantMenu :dataMenu="formData" :category="item" />
      </q-tab-panel>
    </template>
  </q-tab-panels>
</template>
