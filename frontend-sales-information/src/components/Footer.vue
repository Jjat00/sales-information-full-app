<template>
  <v-footer app dark padless>
    <v-card width="100%" class="secondary white--text text-center">
      <v-card-text class="white--text">
        {{ new Date().getFullYear() }} — <strong>Buyers</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
  export default {
    name: 'Footer'
  }
</script>


