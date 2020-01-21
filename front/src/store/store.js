import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {    
        domains: [],
        //servers: [],
    },     
    mutations: {
        setDomains (state, domain) {                        
            state.domains.push(domain)
            //state.servers.push(domain.info.servers)            
        },        
    }   
})
