<template>
  <div>    
       
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
            <h1>History</h1>
            <b-card-text>                                      
                <p>History of last domains consulted</p>                
                <b-button v-on:click="getAll" variant="info">Consult</b-button>
            </b-card-text> 

            <b-container fluid>                                                                                       
                <div class="mt-3" v-for="domain in domains" :key="domain.url"> 
                    <div >  
                        <p>{{domain.url}}</p>                                                                                                
                    </div>                    
                    <b-button block href="#" v-b-toggle.accordion-1 variant="dark">Domain Info</b-button>
                    <b-collapse id="accordion-1" visible accordion="my-accordion" role="tabpanel">                                                                  
                        <b-table-simple striped hover responsive>                        
                            <b-tbody>  
                                <b-tr>
                                    <b-th class="text-right" variant="dark" colspan="1">URL</b-th>
                                    <b-td class="text-left text-light" text-light colspan="5">{{ domain.url }}</b-td>
                                </b-tr>                                    
                                <b-tr>
                                    <b-th class="text-right" variant="dark" colspan="1">Servers Changed</b-th> 
                                    <b-td class="text-left text-light" colspan="5">{{ domain.changed }}</b-td>
                                </b-tr>
                                <b-tr>
                                    <b-th class="text-right" variant="dark" colspan="1">SSL Grade</b-th> 
                                    <b-td class="text-left text-light" colspan="5">{{ domain.ssl_grade }}</b-td>
                                </b-tr>
                                <b-tr>
                                    <b-th class="text-right" variant="dark" colspan="1">Previous SSL</b-th> 
                                    <b-td class="text-left text-light" colspan="5">{{ domain.previous }}</b-td>
                                </b-tr>
                                <b-tr> 
                                    <b-th class="text-right" variant="dark" colspan="1">Logo</b-th> 
                                    <b-td class="text-left text-light" colspan="5">{{ domain.logo }}</b-td>
                                </b-tr>                           
                                <b-tr> 
                                    <b-th class="text-right" variant="dark" colspan="1">Title</b-th> 
                                    <b-td class="text-left text-light" colspan="5">{{ domain.title }}</b-td>
                                </b-tr> 
                                <b-tr>
                                    <b-th class="text-right" variant="dark" colspan="1">Is Down</b-th>     
                                    <b-td class="text-left text-light" colspan="5">{{ domain.down }}</b-td>
                                </b-tr>                                                                                                                                                                                                  
                            </b-tbody>
                        </b-table-simple>  
                    </b-collapse> 
                </div> 
                <div>
                    <!-- Table for the servers -->
                    <b-button block href="#" v-b-toggle.accordion-2 variant="dark">Servers</b-button>
                    <b-collapse id="accordion-2" visible accordion="my-accordion" role="tabpanel">
                        <b-table class="text-light" variant="dark" hover :items="servers"></b-table>
                    </b-collapse> 
                </div>                    
            </b-container>            
            <b-pagination                
                v-model="currentPage"  
                :total-rows="rows"              
                :per-page=1
                aria-controls="my-table"                
            ></b-pagination>  
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
                currentPage: 1,
                rows: 0,
                domain: {  
                    url: '',                                      
                    changed: '',
                    ssl_grade: '',
                    title: '',
                    logo: '',
                    previous: '',
                    down: '',                    
                },
                domains: [],
                servers: [{
                    address: '',
                    ssl_grade: '',
                    country: '',
                    owner: '',
                }],
                errors: [],
                error: false,
            }
        },
        methods: {  
            getAll: function(){                              
                axios.get('http://127.0.0.1:3000/getalldomains')                                   
                .then(response => {
                    if (response.data == "error"){
                        this.error = true
                    } else {
                        this.rows = response.data['items'].length
                        for (var i = 0; i < this.rows; i++) {                            
                            // guardar esta info en domain, y luego agregar a domains
                            this.error = false
                            this.servers = response.data.items[i].info.servers
                            this.domain.url = response.data.items[i].url
                            this.domain.changed = response.data.items[i].info.servers_changed
                            this.domain.ssl_grade = response.data.items[i].info.ssl_grade
                            this.domain.title = response.data.items[i].info.title
                            this.domain.logo = response.data.items[i].info.logo
                            this.domain.previous = response.data.items[i].info.previous_ssl_grade
                            this.domain.down = response.data.items[i].info.is_down 

                            //this.$set(this.domain, this.domain.url, domainUrl)

                            this.domains.push(this.domain)                             
                            //document.write(this.domains[i].url, "* ", i)                                                                                                                                                                                     
                        }                                  
                    }                   
                })                                 
                .catch(e => {
                    this.errors.push(e)
                })                                                    
            }                        
        }         
    }      
</script>
