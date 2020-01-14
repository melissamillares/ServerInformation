<template>
  <div > 
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
            <h1>Server Information</h1>        
            <b-card-text>
                <p>Enter the URL of the website you want to get the information from</p>
                <b-form inline id="input-group">                 
                    <b-form-input
                        id="input"
                        v-model="url"
                        required
                        placeholder="https://google.com"                        
                        class="w-75 mb-2 mr-sm-2 mb-sm-0"
                    ></b-form-input>
                    <b-button variant="info" v-on:click="getDomain">
                        Search
                    </b-button>
                    <p>{{ domain }}</p>
                    <p>{{ title }}</p>
                </b-form>
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
        name: 'Search', 
        data() {
            return {
                url: '',
                domain: '',
                title: '',
                errors: [],              
            }
        }, 
        methods: {
            getDomain: function(){
                //document.write(this.domain)
                axios.post(`http://127.0.0.1:3000/domain`, '"' + this.url + '"')                      
                .then(response => {
                    this.domain = response.data.url
                    this.title = response.data.title
                })              
                .catch(e => {
                    this.errors.push(e)
                })
            }
        }
    };
</script>
