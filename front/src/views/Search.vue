<template>
  <div > 
    <div >
        <p >

        </p>
    </div>       
    <b-container fluid>         
        <b-card  
            overlay
            img-src="https://cdn.pixabay.com/photo/2016/03/26/13/09/notebook-1280538_960_720.jpg" 
            text-variant="white"                               
            style="max-width: 100rem;"            
        >    
            <h1>Server Information</h1>        
            <b-card-text>
                <p>Enter the URL of the website you want to get the information from</p>
                <b-container fluid>
                    <b-form inline id="input-group">                                         
                        <b-form-input
                            id="input"
                            v-model="url"
                            required
                            placeholder="https://google.com"                                                   
                            class="w-75 mb-2 mr-1"                                                                                                               
                        ></b-form-input>

                        <b-button variant="info" class="mb-2" v-on:click="getDomain">
                            Search
                        </b-button>                    
                        <b-popover v-if="urlEmpty" target="input" triggers="hover" placement="bottomright">
                            <template v-slot:title>Attention!</template>   
                            This field cannot be empty                     
                        </b-popover>                                          
                    </b-form>
                </b-container>

                <b-container fluid>                    
                    <!-- Table for the result -->
                    <b-table-simple striped hover responsive>
                        <b-thead head-variant="dark"> 
                            <b-tr><b-th colspan="6">Result</b-th></b-tr>                                                                                         
                        </b-thead> 
                        <b-tbody>  
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">URL</b-th>
                                <b-td class="text-left text-light" text-light colspan="5">{{ url }}</b-td>
                            </b-tr>
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">Servers</b-th>
                                <b-td class="text-left text-light" text-light colspan="5">{{ servers }}</b-td>
                            </b-tr>
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">Servers Changed</b-th> 
                                <b-td class="text-left text-light" colspan="5">{{ changed }}</b-td>
                            </b-tr>
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">SSL Grade</b-th> 
                                <b-td class="text-left text-light" colspan="5">{{ ssl_grade }}</b-td>
                            </b-tr>
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">Previous SSL</b-th> 
                                <b-td class="text-left text-light" colspan="5">{{ previous }}</b-td>
                            </b-tr>
                            <b-tr> 
                                <b-th class="text-right" variant="dark" colspan="1">Logo</b-th> 
                                <b-td class="text-left text-light" colspan="5">{{ logo }}</b-td>
                            </b-tr>                           
                            <b-tr> 
                                <b-th class="text-right" variant="dark" colspan="1">Title</b-th> 
                                <b-td class="text-left text-light" colspan="5">{{ title }}</b-td>
                            </b-tr> 
                            <b-tr>
                                <b-th class="text-right" variant="dark" colspan="1">Is Down</b-th>     
                                <b-td class="text-left text-light" colspan="5">{{ down }}</b-td>
                            </b-tr>                                                                                                                                                                                                  
                        </b-tbody>
                    </b-table-simple>                    
                </b-container>

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
                servers: [],
                changed: '',
                ssl_grade: '',
                title: '',
                logo: '',
                previous: '',
                down: '',
                errors: [], 
                urlEmpty: false,                  
            }
        }, 
        methods: {            
            getDomain: function(){                                
                if (this.url != "") {
                    this.urlEmpty = false
                    axios.post(`http://127.0.0.1:3000/domain`, '"' + this.url + '"')                      
                    .then(response => {
                        this.servers = response.data.servers
                        this.changed = response.data.servers_changed
                        this.ssl_grade = response.data.ssl_grade
                        this.title = response.data.title
                        this.logo = response.data.logo
                        this.previous = response.data.previous_ssl_grade
                        this.down = response.data.is_down
                    })              
                    .catch(e => {
                        this.errors.push(e)
                    })
                } else {
                    return this.urlEmpty = true                    
                }                
            }
        }
    };
</script>
