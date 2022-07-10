<script setup>
/**
 * @file Component for Social Network Link.
 */
import { ref } from 'vue'
import AppIcon from '../atoms/AppIcon.vue'
import AppInput from '../atoms/AppInput.vue'
import AppButton from '../atoms/AppButton.vue'
import { useRestaurantStore } from '../../stores/restaurant'
import { notificationSaved } from '../../helpers/notifications'
import { ruleInputRequired, ruleSpecialCharactersLight } from '../../helpers/rules'

const props = defineProps({
  dataMenu: {
    type: Array,
    default: () => [],
  },
  category: {
    type: String,
    default: '',
  },
})

const storeRestaurant = useRestaurantStore()

const loading = ref(false)

const menuItems = ref(props.dataMenu)

const removeMenuItem = async (index, title, description, category) => {
  const menuItem = {
    title,
    description,
    category,
  }
  await storeRestaurant.queryDeleteResataurantMenu(menuItem)
  menuItems.value.splice(index, 1)
}

const addMenuItem = () => {
  menuItems.value = [...(menuItems.value || []), { category: props.category }]
}

const onSubmit = async () => {
  loading.value = true
  console.log(menuItems.value)
  try {
    await storeRestaurant.queryPostResataurantMenu({
      menuItems,
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
        <div class="row justify-center col-12" v-for="(menuItem, index) in menuItems" :key="index">
          <div class="col-8">
            <AppInput
              v-model="menuItem.title"
              name="title"
              placeholder="Titre"
              :rules="ruleInputRequired"
              dense
              lazy-rules
            />
            <AppInput
              v-model="menuItem.description"
              :rules="[...ruleSpecialCharactersLight]"
              lazy-rules
              name="description"
              dense
              type="textarea"
              placeholder="Description"
            />
          </div>
          <div class="col-auto">
            <AppIcon
              name="fa-regular fa-trash-can"
              color="negative"
              hover-effect
              @click="removeMenuItem(index, menuItem.title, menuItem.description, menuItem.category)"
            />
          </div>
        </div>
      </div>
      <div class="row justify-between col-12 q-gutter-y-sm q-mt-md">
        <AppButton color="primary" @click="addMenuItem" label="Ajouter à la carte" />
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
