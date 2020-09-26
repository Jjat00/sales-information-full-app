<template>
    <v-container>
      <!-- Buttons -->
        <div class="my-3 text-center">
          <v-row>
            <v-col cols="12" sm="6">
              <v-btn 
              @click="getBuyers" 
              dark class="mx-1 my-1"  
              width="150" large color="secondary">
                Get all Buyers
              </v-btn>
              <v-btn 
                @click="deleteBuyers" 
                dark class="mx-1 my-1"  
                width="150" large color="secondary">
                  Clear
              </v-btn>
          </v-col>
          <v-col cols="12" sm="3">
            <v-text-field label="buyerId"
              v-model="buyerID"
              @keyup.enter="getBuyer"/>
          </v-col>
          <v-col cols="12" sm="3">
            <v-btn dark class="mx-1 my-1"  
            width="100" large  @click="getBuyer"
            color="secondary">
              search
            </v-btn>
          </v-col>
        </v-row>
      </div>
      <v-row>
        <h2 class="my-4 mx-3">Total buyers: {{allBuyers.length}}</h2>
        <h2 class="mx-4 my-4">Average age: {{meanAge}}</h2>
      </v-row>
      <!--List all Buyers -->
      <v-row col="12" class="mt-5">
          <div v-for="(item,index) of allBuyers" :key="index">
              <h4 class="buyers">
                name: 
                <router-link
                  :to="{name:'Buyer', params:{buyerId: item.id}}" 
                  class="mx-1 my-1">
                    {{item.name}}
                </router-link>         
                <br> age: {{item.age}}     
                <br> id: {{item.id}}     
              </h4>
          </div>  
      </v-row>  
    </v-container>
</template>

<script>
import {mapState, mapMutations} from 'vuex'
import axios from 'axios'
  export default {
    name: 'ListBuyers',
    data () {
      return {
        allBuyers: [],
        buyer: '',
        meanAge: 0,
        buyerID: '',
        }
    },
    computed:{
        ...mapState(['loading'])
    },
    methods:{
        ...mapMutations(['showDialog', 'quitDialog']),
        async getBuyers() {
          try {
              this.showDialog({title:"wait, get all buyers"})
              const response = await fetch("http://localhost:3000/buyers");
              const buyers = await response.json()
              this.allBuyers = buyers
              this.getMeanAge()
          } catch (error) {          
          } finally {
            this.quitDialog()
          }
        },
        deleteBuyers(){
          this.allBuyers = []
          this.meanAge = 0
        },
        async getBuyer(){
            try {
              this.showDialog({title:"wait, get buyer" + this.buyerID})
              const response = await fetch("http://localhost:3000/buyers");
              const buyers = await response.json()
              this.allBuyers = buyers.filter(buyer => buyer.id === this.buyerID)
          } catch (error) {          
          } finally {
            this.quitDialog()
          }
        },
        getMeanAge(){
          var totalAge = 0
          this.allBuyers.forEach(buyer =>{
              totalAge = totalAge+buyer.age
          });
          this.meanAge = totalAge/this.allBuyers.length | 0
        }
    }
  }
</script>



<style scoped>
  .buyers {
    width: auto;
    padding: 5px 5px;
    margin: 5px;
    color:rgb(87, 87, 87);
    box-shadow: 0 0 4px 0 rgba(53, 62, 139, 0.5)
  }
  a {
    text-decoration: none;
    color: rgb(9, 67, 153);
  }
</style>