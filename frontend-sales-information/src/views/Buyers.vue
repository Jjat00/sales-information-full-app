<template>
  <v-container fluid>
    <div class="buyertitle">
      <h1 class="mb-5">Here you can get the data of your buyers</h1>
    </div>
    <div class="buyers">
      <ListBuyers/>
    </div>

  </v-container>
</template>

<script>
import ListBuyers from '@/components/ListBuyers.vue'

import { mapMutations } from 'vuex'

export default {
  components: {
    ListBuyers
  },
  methods:{
    
  }
}
</script>

<style scoped>
  .buyertitle{
    text-align: center;
  }
  .buyers{
    height: 215px;
    box-shadow: 0 0 20px 0 rgba(53, 62, 139, 0.3);
  }

</style>