<script setup>
/**
 * @file The dropdown for user profile.
 */
import AppItem from '../atoms/AppItem.vue'
import AppIcon from '../atoms/AppIcon.vue'
import { computed, ref } from 'vue'
import { useRestaurantStore } from 'src/stores/restaurant/index'

const storeRestaurant = useRestaurantStore()

const props = defineProps({
  dropdownData: {
    type: Object,
    default: () => {},
  },
})

const isOpen = ref(false)

const toggle = () => {
  isOpen.value = !isOpen.value
}

const toggleHideMenu = () => {
  isOpen.value = false
}

const logoutRoute = computed(() => ({
  label: 'Se déconnecter',
  icon: 'fas fa-sign-out-alt',
  click: () => {
    storeRestaurant.logoutRestaurant()
  },
}))
</script>

<template>
  <q-btn class="app-dropdown-profile q-px-sm" no-caps flat stretch @click="toggle">
    <div class="items-center no-wrap row q-gutter-sm">
      <div class="gt-xs text-white">{{ props.dropdownData?.name }}</div>
      <AppIcon name="fa-solid fa-angle-down" :class="{ 'is-open': isOpen }" class="text-white" color="white" />
    </div>
    <q-menu @hide="toggleHideMenu">
      <AppItem :iconLeft="logoutRoute.icon" :label="logoutRoute.label" @click="logoutRoute.click" />
    </q-menu>
  </q-btn>
</template>

<style lang="scss">
.app-dropdown-profile {
  .q-icon {
    color: inherit !important;
    margin-top: 5px;
    transition: transform 0.28s;
    &.is-open {
      transform: rotate(180deg);
    }
  }
}
</style>
