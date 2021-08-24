<template>
 <form @submit.prevent="submit">
    
    <h1 class="h3 mb-3 fw-normal">Please Register</h1>
     <div class="form-floating">
      <input  v-model="data.name" class="form-control" placeholder="Names">
      <label for="floatingInput">Name</label>
    </div>    
    <div class="form-floating">
      <input  v-model="data.username" class="form-control" placeholder="Names">
      <label for="floatingInput">Username</label>
    </div> 
    <div class="form-floating">
      <input v-model="data.email" type="email" class="form-control" placeholder="name@example.com">
      <label for="floatingInput">Email address</label>
    </div>
    <div class="form-floating">
      <input v-model="data.password" type="password" class="form-control" placeholder="Password">
      <label for="floatingPassword">Password</label>
    </div>
    <button class="w-100 btn btn-lg btn-primary" type="submit">Register</button>
    <p class="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
  </form>
</template>

<script lang="ts">
import { reactive} from "vue";
import { useRouter } from 'vue-router';

export default {
    name:"Register",
    setup(){
        const data = reactive({
            email:'',
            name:'',
            username:'',
            password:''
        });
        const router = useRouter();

        const submit =async ()=>{
            await fetch('http://localhost:3000/api/register', {
                method:'POST',
                headers:{'Content-Type':'application/json'},
                body:JSON.stringify(data)
            });
            
            await  router.push('/login');
        }
        return{
            data,
            submit
        }
    }
}

</script>
