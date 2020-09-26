<template>
    <v-container>
        <div class="home">
            <h2>Welcome</h2>
            <h1>What do you want to do?</h1>
            <!-- Buttons -->
            <v-btn 
                class="mx-2 mt-10" 
                width="150" 
                large 
                dark
                color="secondary"
                :to="{name:'Upload'}">
                    Zync Data
            </v-btn>              
            <v-btn 
                class="mx-2 mt-10" 
                width="150" 
                large 
                dark
                color="secondary"
                :to="{name:'Buyers'}">
                    View buyers
            </v-btn>             
        </div>
    </v-container>
</template>

<script>
  export default {
    name: 'Home'
  }
</script>


<style scoped>
    h2{
        color:rgb(0, 0, 0);
        padding: 25px;
    }
    h1{
        margin: 20px;
        color:rgb(14, 0, 37);
    }
    div .home{
        text-align: center;
        width: 100%;
        height: 400px;
        box-shadow: 0 0 20px 0 rgba(53, 62, 139, 0.3);
    }
</style>