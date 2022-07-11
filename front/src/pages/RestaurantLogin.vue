<script setup>
/**
 * @file Page for restaurant login.
 */
import FormLogin from '../components/molecules/FormLogin.vue'
import AppCard from '../components/atoms/AppCard.vue'
import AppTitle from '../components/atoms/AppTitle.vue'
</script>

<template>
  <div class="row col-6 justify-center items-center window-height">
    <div class="q-pa-md">
      <AppCard class="">
        <AppTitle class="text-center q-pb-lg">Connectez vous à votre compte</AppTitle>
        <FormLogin />
      </AppCard>
      <div class="row justify-center">
        <router-link style="text-decoration: none" :to="{ name: 'Register' }">
          <p class="text-white text-center q-mt-md">Vous n'avez pas de compte ?</p>
        </router-link>
      </div>
    </div>
  </div>
</template>
