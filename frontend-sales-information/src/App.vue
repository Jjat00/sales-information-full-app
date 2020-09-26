<template>
  <v-app>
    <v-container>
      <!-- Header nav menu-->
      <Header></Header>
      <v-main>
      <!-- router -->
        <router-view/>
      <!-- loading dialog-->
        <v-dialog v-model="loading.state" hide-overlay persistent width="300" height="100">
          <v-card color="primary" dark>
            <v-card-title>
              loading...
            </v-card-title>
            <v-card-text>
              {{loading.title}}
              <v-progress-linear indeterminate color="white" class="mb-0"></v-progress-linear>
            </v-card-text>
          </v-card>
        </v-dialog>
      </v-main>
      <!-- Footer -->
      <Footer></Footer>
    </v-container>
  </v-app>
</template>

<script>
import Header from "./components/Header.vue"
import Footer from "./components/Footer.vue"
import { mapState } from 'vuex';
export default {
  name: 'App',
  components: {
      Header,
      Footer
  },
  computed:{
    ...mapState(['loading'])
  },
};
</script>


<style scoped>
  .nav{
    align-items: center;
    text-align: center;
  }
  .nav .items {
    padding: 5px;
  }
</style>