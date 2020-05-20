<template>
  <div id="login" class="form-signin">
    <img class="mb-4" src="../assets/talaria.svg" alt="" width="72" height="72"> 
    <label for="inputEmail" class="sr-only">Email address</label>
    <input v-model="username" v-on:keyup.enter="login" class="form-control" placeholder="Username" required autofocus>
    <label for="inputPassword" class="sr-only">Password</label>
    <input v-model="password" v-on:keyup.enter="login" type="password" class="form-control" placeholder="Password" required>
    <button @click="login" class="btn btn-lg btn-primary btn-block">Sign in</button>
  </div>
</template>

<script>

import { LoginRequest } from "../proto/auth_pb";
import { AuthClient } from "../proto/auth_grpc_web_pb";

export default {
  name: "login",
  components: {},
  data: function() {
    return {
      token: "",
      username: "",
      password: ""
    };
  },
  created: function() {
    this.client = new AuthClient("http://localhost:8081", null, null);
  },
  methods: {
    login: function() {
        let req = new LoginRequest();
        req.setUsername(this.username);
        req.setPassword(this.password);
        this.client.login(req, {}, (err, resp) => {
            console.log(err); 
            this.token = resp.getToken();
            console.log(this.token); 
        });
    }
  }
};
</script>

<style>
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 30px;
  margin: auto;
  border-radius: 5px;
  background-color: #04151c;
  opacity: 0.9;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}

.form-signin .form-control:focus {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
