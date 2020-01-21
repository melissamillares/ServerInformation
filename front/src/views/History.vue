<template>
  <div>    
       
    <div ><br></div>          
    <b-container fluid>               
        <b-card  
            overlay            
            img-src="https://picsum.photos/900/125/?image=3"
            img-top
            text-variant="white"                                              
            style="max-width: 100rem;"           
        >  
            <h1>History</h1>
            <b-card-text>                                      
                <p>History of last domains consulted</p>                
                <b-button v-on:click="getAll" variant="info">Consult</b-button>
            </b-card-text> 
            <div>
               <br>
            </div>            
            <b-card-text>
                <b-container fluid border-variant="secondary">                                   
                    <div v-if="!rows" class="w-25 bg-info text-light">                    
                        <p class="mt-1">There are no results</p>
                    </div>                                       

                    <b-list-group v-if="rows" class="text-center text-dark" id="result">                                            
                        <b-list-group-item>
                            <b-table-simple hover small caption-top responsive>
                                <b-thead class="text-center text-dark">
                                    <b-tr>
                                        <b-th colspan="6" class="text-center text-light" variant="dark">{{ domains[currentPage].url }}</b-th>        
                                    </b-tr>
                                    <b-tr>
                                        <b-th>Servers Changed</b-th>
                                        <b-th>SSL Grade</b-th>
                                        <b-th>Previous SSL</b-th>
                                        <b-th>Logo</b-th>
                                        <b-th>Title</b-th>
                                        <b-th>Is Down</b-th>                                    
                                    </b-tr>
                                </b-thead>
                                <b-tbody class="text-center text-dark">
                                    <b-tr>
                                        <b-td>{{ domains[currentPage].info.changed }}</b-td>
                                        <b-td>{{ domains[currentPage].info.ssl_grade }}</b-td>
                                        <b-td>{{ domains[currentPage].info.previous }}</b-td>
                                        <b-td>{{ domains[currentPage].info.logo }}</b-td>
                                        <b-td>{{ domains[currentPage].info.title }}</b-td>
                                        <b-td>{{ domains[currentPage].info.down }}</b-td>                                
                                    </b-tr>
                                </b-tbody>
                            </b-table-simple>    
                        </b-list-group-item> 

                        <b-list-group-item>
                            <b-table-simple hover small caption-top responsive>
                                <b-thead class="text-center text-dark">
                                    <b-tr>
                                        <b-th colspan="6" variant="info">
                                            Servers
                                        </b-th>        
                                    </b-tr>
                                    <b-tr>
                                        <b-th>Address</b-th>
                                        <b-th>SSL Grade</b-th>
                                        <b-th>Country</b-th>
                                        <b-th>Owner</b-th>                                   
                                    </b-tr>
                                </b-thead>
                                <b-tbody class="text-center text-dark" v-for="ss in domains[currentPage].info.servers" :key="ss">
                                    <b-tr>
                                        <b-td>{{ ss.address }}</b-td>
                                        <b-td>{{ ss.ssl_grade }}</b-td>
                                        <b-td>{{ ss.country }}</b-td>
                                        <b-td>{{ ss.owner }}</b-td>                                                              
                                    </b-tr>
                                </b-tbody>
                            </b-table-simple>
                        </b-list-group-item>
                    </b-list-group> 
                    <!-- <div class="mt-3" v-for="d in domains" :key="d.url">
                        <div>                                                                                                                                          
                            <b-table-simple hover small caption-top responsive>
                                <b-thead class="text-center text-dark">
                                    <b-tr>
                                        <b-th colspan="6" class="text-center text-light" variant="dark">{{ d.url }}</b-th>        
                                    </b-tr>
                                    <b-tr>
                                        <b-th>Servers Changed</b-th>
                                        <b-th>SSL Grade</b-th>
                                        <b-th>Previous SSL</b-th>
                                        <b-th>Logo</b-th>
                                        <b-th>Title</b-th>
                                        <b-th>Is Down</b-th>                                    
                                    </b-tr>
                                </b-thead>
                                <b-tbody class="text-center text-dark">
                                    <b-tr>
                                        <b-td>{{ d.info.changed }}</b-td>
                                        <b-td>{{ d.info.ssl_grade }}</b-td>
                                        <b-td>{{ d.info.previous }}</b-td>
                                        <b-td>{{ d.info.logo }}</b-td>
                                        <b-td>{{ d.info.title }}</b-td>
                                        <b-td>{{ d.info.down }}</b-td>                                
                                    </b-tr>
                                </b-tbody>
                            </b-table-simple> 
                        
                            <b-table-simple hover small caption-top responsive>
                                <b-thead class="text-center text-dark">
                                    <b-tr>
                                        <b-th colspan="6" variant="info">
                                            Servers
                                        </b-th>        
                                    </b-tr>
                                    <b-tr>
                                        <b-th>Address</b-th>
                                        <b-th>SSL Grade</b-th>
                                        <b-th>Country</b-th>
                                        <b-th>Owner</b-th>                                   
                                    </b-tr>
                                </b-thead>
                                <b-tbody class="text-center text-dark" v-for="ss in d.info.servers" :key="ss">
                                    <b-tr>
                                        <b-td>{{ ss.address }}</b-td>
                                        <b-td>{{ ss.ssl_grade }}</b-td>
                                        <b-td>{{ ss.country }}</b-td>
                                        <b-td>{{ ss.owner }}</b-td>                                                              
                                    </b-tr>
                                </b-tbody>
                            </b-table-simple>                             

                        </div>                                                                                                                              
                    </div> -->   
                    <b-pagination
                        v-if="rows"
                        v-model="currentPage"
                        :total-rows="rows" 
                        :per-page="perPage"                   
                        aria-controls="result"
                        align="center"                        
                    ></b-pagination>                                                                                                                                                                                                                                                           
                </b-container>                                                
            </b-card-text>

        </b-card>                            
    </b-container> 
    <div ><br></div>
  </div>      
</template>

<script>  
    import axios from 'axios'  
    export default {
        name: 'History', 
        data() {
            return {
                currentPage: 0,
                rows: 0, 
                perPage: 1,                             
                errors: [],                
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
                        this.error = false        

                        for (var i = 0; i < this.rows; i++) {
                            this.$store.commit('setDomains', response.data.items[i])
                            //this.$store.commit('setDomains', response.data.items[i].info.servers)
                        }                                 
                    }                   
                })                                 
                .catch(e => {
                    this.errors.push(e)
                })                                                    
            }                        
        },
        computed: {
            domains() {
                return this.$store.state.domains;
            },
            servers() {
                return this.$store.state.servers;
            }
        }         
    }      
</script>
