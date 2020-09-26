<template>
    <v-container>
      <!-- Header -->
        <v-app-bar app color="primary" dark >
          <v-app-bar-nav-icon 
            @click="drawer = !drawer">
          </v-app-bar-nav-icon>
          <v-toolbar-title>Buyers Data</v-toolbar-title>
          <v-spacer/>
            <v-btn class="mx-2" icon :to="{name:'Home'}">
                  <v-icon>home</v-icon>
            </v-btn>
            <v-btn class="mx-2" icon :to="{name:'Upload'}">
                  <v-icon>cloud_upload</v-icon>
            </v-btn>            
            <v-btn class="mx-2" icon :to="{name:'Buyers'}">
                  <v-icon>people</v-icon >
            </v-btn>
        </v-app-bar>
      <!-- Navigations -->
      <v-navigation-drawer app v-model="drawer" 
        absolute temporary dark color="secondary">
      <v-list-item>
        <v-list-item-avatar>
          <v-icon>person</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>Buyers Data</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <!-- Navigations actions -->
      <v-divider></v-divider>
      <v-row mt=4>
        <v-col>
          <v-btn block class="my-1" 
            color="primary" 
            large :to="{name:'Home'}">Home
          </v-btn>
          <v-btn block class="my-1" 
            color="primary" 
            large :to="{name:'Upload'}">Upload
          </v-btn>          
          <v-btn  block class="my-1" 
            color="primary" 
            large :to="{name:'Buyers'}">Buyers
          </v-btn>
        </v-col>
      </v-row>
    </v-navigation-drawer>
    </v-container>
</template>

<script>
export default {
    name: "Header",
    data(){
      return{
        drawer: true
      }
    }
}
</script>