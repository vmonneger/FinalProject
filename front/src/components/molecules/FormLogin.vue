<script setup>
/**
 * @file Form for restaurant register (authentification).
 */
import { ref, reactive } from 'vue'
import { useRestaurantStore } from 'src/stores/restaurant'
import { useRouter } from 'vue-router'
import AppButton from '../atoms/AppButton.vue'
import AppInput from '../atoms/AppInput.vue'
import { ruleInputRequired, ruleEmail } from '../../helpers/rules'

const router = useRouter()
const restaurantStore = useRestaurantStore()

const form = reactive({})

const loading = ref(false)

/**
 * The submit of the form.
 * Dispatch an action from firebase to create a user in firebase and create a new user in bakend.
 */
const onSubmit = async () => {
  loading.value = true
  try {
    await restaurantStore.loginRestaurantUser({
      email: form.email,
      password: form.password,
    })
    loading.value = false
  } catch (e) {
    loading.value = false
    throw e
  }
  loading.value = false
  router.push({ name: 'Login' })
}
</script>

<template>
  <q-form @submit.prevent.stop @validation-success="onSubmit">
    <div class="row q-col-gutter-sm">
      <AppInput
        v-model="form.email"
        class="col-12"
        name="email"
        placeholder="Email"
        dense
        :rules="ruleEmail"
        lazy-rules
      />
      <AppInput
        v-model="form.password"
        class="col-12"
        name="password"
        placeholder="Mot de passe"
        dense
        :rules="ruleInputRequired"
        lazy-rules
        type="password"
      />
    </div>
    <div class="row q-mb-md q-mt-lg">
      <AppButton label="Se connecter" type="submit" class="fit" no-caps :loading="loading" />
    </div>
  </q-form>
</template>
