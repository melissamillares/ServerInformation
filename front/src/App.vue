<template>
  <div id="app">   
    <b-navbar toggleable="lg" type="dark" variant="dark">      
      <b-container>        
        <b-navbar-brand href="#">Server Info</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>  
          <b-navbar-nav>
            <b-nav-item :to="{name: 'search'}" exact>Search</b-nav-item>
            <b-nav-item :to="{name: 'history'}">History</b-nav-item>
          </b-navbar-nav>   
        </b-collapse> 
      </b-container>           
    </b-navbar> 

    <router-view/>

  </div>
</template>

<script>

</script> 
