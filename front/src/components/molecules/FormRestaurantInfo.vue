<script setup>
import AppButton from '../atoms/AppButton.vue'
import AppInput from '../atoms/AppInput.vue'
import { ref, reactive } from 'vue'
import { notificationSaved } from '../../helpers/notifications'
import { useRestaurantStore } from 'src/stores/restaurant'
import { ruleInputRequired, ruleSpecialCharactersLight, ruleSpecialCharacters } from '../../helpers/rules'

const restaurantStore = useRestaurantStore()

const form = reactive({
  name: restaurantStore.name,
  description: restaurantStore.description,
})

const loading = ref(false)

const onSubmit = async () => {
  loading.value = true
  try {
    await restaurantStore.queryPostResataurantInfo({
      name: form.name,
      description: form.description,
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
  <q-form @submit.prevent.stop @validation-success="onSubmit">
    <div class="q-gutter-y-md">
      <div>
        <AppInput
          v-model="form.name"
          :rules="[...ruleInputRequired, ...ruleSpecialCharacters]"
          lazy-rules
          name="name"
          dense
          label="Nom du restaurant"
        />
      </div>
      <div>
        <AppInput
          v-model="form.description"
          :rules="[...ruleInputRequired, ...ruleSpecialCharactersLight]"
          lazy-rules
          name="description"
          dense
          type="textarea"
          label="Description de votre restaurant"
        />
      </div>
      <div class="fit text-right">
        <AppButton label="Enregistrer" type="submit" no-caps :loading="loading" />
      </div>
    </div>
  </q-form>
</template>
