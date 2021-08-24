<template>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
  <div class="container-fluid">
    <router-link to="/" class="navbar-brand">Home</router-link>
    
    <div >
      <ul class="navbar-nav me-auto mb-2 mb-md-0" v-if="!auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link" >Login</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/register" class="nav-link" >Register</router-link>
        </li>
      </ul>
      <ul class="navbar-nav me-auto mb-2 mb-md-0" v-if="auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link" @click="logout" >Logout</router-link>
        </li>
       
      </ul>
    </div>
  </div>
</nav>
</template>

<script lang="ts">
import { computed } from '@vue/runtime-core';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';
export default {
    name: "Nav",
    setup() {
        const store = useStore();
       
        const auth = computed(()=>store.state.authenticated)
        const logout = async()=>{
        const res = await fetch('http://localhost:3000/api/logout', {
          method:'POST',
          headers:{'Content-Type':'application/json'},
          credentials:"include",
        });
       store.state.authenticated=false
        }

        return {
            auth,
            logout
        }
    },
}
</script>