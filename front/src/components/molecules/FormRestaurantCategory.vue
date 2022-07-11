<script setup>
/**
 * @file Form for restaurant category.
 */
import { ref } from 'vue'
import { ruleInputRequired } from '../../helpers/rules'
import AppIcon from '../atoms/AppIcon.vue'
import AppInput from '../atoms/AppInput.vue'
import AppButton from '../atoms/AppButton.vue'
import { useRestaurantStore } from '../../stores/restaurant'
import { notificationSaved } from '../../helpers/notifications'

const storeRestaurant = useRestaurantStore()
const loading = ref(false)
const categoriesMenu = ref(storeRestaurant.getRestaurantCategory)

const removeCategoryMenu = async (index, category) => {
  await storeRestaurant.queryDeleteResataurantCategory(category)
  categoriesMenu.value.splice(index, 1)
  notificationSaved()
}

const addCategoryMenu = () => {
  categoriesMenu.value = [
    ...(categoriesMenu.value || []),
    {
      category: '',
    },
  ]
}

const onSubmit = async () => {
  loading.value = true
  try {
    await storeRestaurant.queryPostResataurantCategory({
      categoriesMenu,
    })
  } catch (e) {
    loading.value = false
    throw new Error(e)
  }
  notificationSaved()
  loading.value = false
}
</script>

<template>
  <div class="row col-12 q-col-gutter-md">
    <q-form class="col-12" @submit.prevent.stop @validation-success="onSubmit">
      <div class="row col-12 items-center q-col-gutter-lg">
        <div class="row col-12 items-center" v-for="(categoryMenu, index) in categoriesMenu" :key="index">
          <div class="col-grow">
            <AppInput
              v-model="categoryMenu.category"
              name="socialNetworkUrl"
              dense
              placeholder="Ex: Boissons, Desserts, Entrées, Apéritifs..."
              :rules="ruleInputRequired"
              class="no-padding"
              lazy-rules
            />
          </div>
          <div class="col-auto">
            <AppIcon
              name="fa-regular fa-trash-can"
              color="negative"
              hover-effect
              @click="removeCategoryMenu(index, categoryMenu.category)"
            />
          </div>
        </div>
      </div>
      <div class="row justify-between col-12 q-gutter-y-sm q-mt-md">
        <AppButton color="primary" @click="addCategoryMenu" label="Ajouter une catégorie" />
        <AppButton color="accent" type="submit" label="Enregister" />
      </div>
    </q-form>
  </div>
</template>

<style lang="scss">
.app-modal {
  .bg-primary {
    background-color: $primary !important;
  }
}
</style>
