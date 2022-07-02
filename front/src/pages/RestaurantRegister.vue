<script setup>
import FormRegister from '../components/molecules/FormRegister.vue'
import AppCard from '../components/atoms/AppCard.vue'
import AppTitle from '../components/atoms/AppTitle.vue'
</script>

<template>
  <div class="row col-6 justify-center items-center window-height">
    <div class="q-pa-md">
      <AppCard class="">
        <AppTitle class="text-center q-pb-lg">Vous êtes restaurateur ?</AppTitle>
        <p class="text-center text-secondary text-bold">
          Créez votre compte pour pouvoir bénéficier de notre solution de commande
        </p>
        <FormRegister label="S'enregistrer" />
      </AppCard>
      <div class="row justify-center">
        <router-link style="text-decoration: none" :to="{ name: 'Login' }">
          <p class="text-white text-center q-mt-md">Vous avez déjà un compte ?</p>
        </router-link>
      </div>
    </div>
  </div>
</template>
