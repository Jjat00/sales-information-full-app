<template>
    <v-container>
        <!-- Selecte Date -->
        <h1>Select date to sync</h1>
        <v-row>
            <v-col xs="12">
                <v-card>
                    <v-date-picker 
                        class="my-2" 
                        v-model="date" 
                        :full-width="true"
                        :min="minDate"   
                        :max="maxDate"   
                    ></v-date-picker>
                </v-card>
            </v-col>
            <v-col class="text-center align-center">
                <h1 class="my-8">Upload data on selected date </h1>
                <h1 class="mb-8">{{date}}</h1>
                <p>
                    You will upload all the information of the sales rached today: buyers, products and transactions
                </p>
                <v-btn @click="dialogLoadData=true" 
                    large color="secondary">Load Data
                </v-btn>
            </v-col>
        <div class="my-2">
        </div>
        </v-row>
        <!-- Dialog -->
        <v-dialog v-model="dialogLoadData" width="400" persistent>
            <v-card>
                <v-card-title>Are you sure?</v-card-title>
                <v-card-text>The database will be alterated</v-card-text>
                <v-card-actions>
                    <v-btn class="ml-15" width="100" 
                        small color="secondary " 
                        @click="dialogLoadData=false">
                        cancel
                    </v-btn>
                    <v-btn  class="ml-15" width="100" 
                        small color="secondary" 
                        @click="postData(date)">
                        yes
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <!-- dialog OK response  -->
        <v-dialog v-model="dialogResponse" width="500">
            <v-card color="primary" dark>
                <v-card-title>
                    <h1>Response</h1>
                </v-card-title>
                <v-divider color="white" ></v-divider>
                <v-card-text class="text-center mt-5">
                    <h2>{{responseSyncData.BuyerResponse}}!</h2>
                    <h2>{{responseSyncData.ProductResponse}}!</h2>
                    <h2>{{responseSyncData.TransactionResponse}}!</h2>                            
                </v-card-text>
                <v-card-actions>
                    <v-btn block class="mr-15px secondary" @click="dialogResponse=false">Ok</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>        
        <!-- dialog error response  -->
        <v-dialog v-model="dialogError" width="400">
            <v-card color="primary" dark>
                <v-card-title>
                    <h1>Error!</h1>
                </v-card-title>
                <v-divider color="white" ></v-divider>
                <v-card-text class="text-center mt-5">
                    <h4>Error connect database</h4>
                </v-card-text>
                <v-card-actions>
                    <v-btn block class="mr-15px secondary" @click="dialogError=false">Ok</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>        
    </v-container>
</template>

<script>
import {mapState, mapMutations} from 'vuex'
export default {
    name: 'LoadData',
    data () {
      return {
            date: new Date().toISOString().substring(0,10),
            minDate: '2018',
            maxDate: new Date().toISOString().substring(0,10),
            dialogLoadData: false,
            dialogResponse: false,
            dialogError: false,
            responseSyncData: ""
      }
    },
    computed:{
        ...mapState(['loading'])
    },
    methods:{
        ...mapMutations(['showDialog', 'quitDialog']),
        async postData(date){
            try {
                this.showDialog({title:"zync all data on "+date})
                this.dialogLoadData = false
                const dateUnix = new Date(date).getTime()/1000
                const dateObject = {
                    method: 'POST',
                    body:  JSON.stringify({
                        date: dateUnix.toString(10)
                    })
                }
                const res = await fetch('http://localhost:3000/loadData',dateObject)
                console.log(res.status)
                if (res.status==200) {
                    this.dialogResponse=true
                }else{
                    this.dialogError = true
                }
                const response = await res.json()
                this.responseSyncData = response
            } catch (error) {
                console.log(error)
            } finally {
                this.quitDialog()
            }
        }
    }
}
</script>