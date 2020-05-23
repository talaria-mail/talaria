<template>
  <div id="login" class="text-center"> 
    <div class="form-signin">
      <h1>Talaria</h1>
      <img class="form-logo mb-4" src="../assets/talaria-light.svg"> 
      <label for="inputEmail" class="sr-only">Email address</label>
      <input v-model="username" v-on:keyup.enter="onLogin" class="form-control" placeholder="Username" required autofocus>
      <label for="inputPassword" class="sr-only">Password</label>
      <input v-model="password" v-on:keyup.enter="onLogin" type="password" class="form-control" placeholder="Password" required>
      <button @click="onLogin" class="btn btn-lg btn-primary btn-block">Sign in</button>
    </div>
  </div>
</template>

<script>

import { mapActions } from "vuex";

export default {
  name: "login",
  components: {},
  data: function() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    ...mapActions(['login']),
    onLogin() {
      this.login(this)
    }
  }
};
</script>

<style>

#login {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-image: url('./../assets/beach.webp');
  background-position: center; /* Center the image */
  background-repeat: no-repeat; /* Do not repeat the image */
  background-size: cover; /* Resize the background image to cover the entire container */
  height: 100%;
}

.form-logo {
  width: 120px;
}

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 30px;
  margin: auto;
  border-radius: 5px;
  background-color: #04151c;
  opacity: 0.9;
}

.form-signin > h1 {
  color: white;
  font-family: 'Playfair Display', serif;
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
