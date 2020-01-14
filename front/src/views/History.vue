<template>
  <div>    
       
    <div >
        <p >

        </p>
    </div>       
    <b-container fluid="md">         
        <b-card  
            overlay
            img-src="https://cdn.pixabay.com/photo/2016/03/26/13/09/notebook-1280538_960_720.jpg" 
            text-variant="white"                                              
            style="max-width: 100rem;"            
        >  
            <h1>History</h1>
            <b-card-text>                                      
                <p>History of last domains consulted</p>
                <p>{{ domains }}</p> 
                <b-button v-on:click="getAll" variant="info">Clic</b-button>
            </b-card-text>  
        </b-card>        
    </b-container> 
    <div >
        <p >

        </p>
    </div>
  </div>      
</template>

<script>  
    import axios from 'axios'  
    export default {
        name: 'History', 
        data() {
            return {
                domains: '',
                errors: []
            }
        },
        methods: {  
            getAll(){                
                axios.get('http://127.0.0.1:3000/getalldomains')
                .then((result) => {
                    this.domains = result.body
                })
                .catch(e => {
                    this.errors.push(e)
                })                
            }                        
        } 
        /* mounted () {
            axios.get('http://127.0.0.1:3000/getalldomains')
                .then(function (response) {
                    this.info = response.data;
                });
        } */
    }      
</script>
