<template>
    <v-container>
        <!-- Buyer information -->
        <v-card>
            <v-img>
                <h1 class="mx-4 mt-5">{{buyerName}} - {{buyerAge}} years old</h1>
            </v-img>
                <v-card-title>
                    <v-row>
                        <h4 class="buyer">BuyerId: {{ $route.params.buyerId }}</h4>
                        <h4 class="buyer">No. products: {{totalProducts}}</h4> 
                        <h4 class="buyer">Total money: $ {{totalMoney}}</h4>
                    </v-row>
                </v-card-title>
            <v-divider/>
        <!-- Buyer's product information -->
            <v-card-text>
                <div>
                    <h1 class="my-3">Purchase History</h1>
                    <v-row>
                        <div v-for="(item, index) of buyerInformation.purchaseHistory" :key="index">
                            <h4 class="products">
                                {{item.name}} ${{item.price}}</h4>
                        </div>
                    </v-row>
                </div>
            </v-card-text>
            <v-divider></v-divider>
        <!-- Other people  -->
            <v-card-text>
                <div>
                    <h1 class="my-3">Other people who bought from the same place</h1>
                    <v-divider></v-divider>
                    <v-row :wrap="true">
                        <div  v-for="(item,index) of buyerInformation.otherBuyers" :key="index">
                            <h4 class="products">name: {{item.name}} 
                                <br> age: {{item.age}} 
                                <br> id: {{item.id}} 
                            </h4> 
                        </div>
                    </v-row>
                </div>
            </v-card-text>            
        <!-- Recomendations  -->
            <v-card-text>   
                <div>
                    <h1 class="my-3">Recommended Products</h1>
                        <v-row>
                            <div v-for="(item, index) of buyerInformation.recomendations" :key="index">
                                <h4 class="products">{{item.name}} ${{item.price}}</h4>
                            </div>
                        </v-row>
                </div>
            </v-card-text>
            <v-divider></v-divider>
        </v-card>
    </v-container>
</template>

<script>
import {mapState, mapMutations} from 'vuex'
export default {
    name: 'BuyerInfo',
    data () {
      return {
          buyerName: 'Name',
          buyerAge: 0,
          buyerInformation: [],
          totalMoney: 0,
          totalProducts: 0
      }
    },
    computed:{
        ...mapState(['loading'])
    },
    methods:{
        ...mapMutations(['showDialog', 'quitDialog']),
        async getBuyerInfo(buyerId){
            try {
                this.showDialog({title:"get buyer information " + buyerId})
                const response = await fetch(`http://localhost:3000/consultBuyer/${buyerId}`);
                const buyerConsult = await response.json()
                this.buyerInformation = buyerConsult
                this.buyerName = this.buyerInformation.buyer.name
                this.buyerAge = this.buyerInformation.buyer.age
                this.totalProducts = buyerConsult.purchaseHistory.length
                this.getTotalMoney()
            } catch (error) {
                console.log(error)
            } finally {
                this.quitDialog()
            }
        },
        getTotalMoney(){
            this.buyerInformation.purchaseHistory.forEach(product =>{
                 this.totalMoney = this.totalMoney+product.price
            });
        }
    },
    created(){
        this.getBuyerInfo(this.$route.params.buyerId)
    }
}
</script>


<style scoped>
    h4.buyer{
        margin-top: 10px;
        color:rgb(48, 48, 48);
        margin-left: 15px;
    }
    h4.products {
        width: auto;
        padding: 5px 15px;
        margin: 5px;
        margin-top: 10px;
        color:rgb(87, 87, 87);
        box-shadow: 0 0 5px 0 rgba(53, 62, 139, 0.5)}
</style>