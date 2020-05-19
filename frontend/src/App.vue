<template>
  <div id="app">
    <section>
      <span class="title-text">Login</span>
      <div class="row justify-content-center mt-4">
        <input v-model="username" v-on:keyup.enter="login" class="mr-1" placeholder="Todo Item">
        <input v-model="password" v-on:keyup.enter="login" class="mr=1" placeholder="Passw3rd">
        <button @click="login" class="btn btn-primary">Sign In</button>
      </div>
    </section>
    <section>
      <div class="row">
        <div class="offset-md-3 col-md-6 mt-3">
          <ul class="list-group justify-content-center">
            <li
              class="row list-group-item border mt-2 col-xs-1"
              v-for="todo in todos"
              v-bind:key="todo.id"
            >
              <div>
                <span>{{todo.task}}</span>
                <span @click="deleteTodo(todo)" class="offset-sm-1 col-sm-2 delete text-right">X</span>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

import { LoginRequest } from "./proto/auth_pb";
import { AuthClient } from "./proto/auth_grpc_web_pb";

export default {
  name: "app",
  components: {},
  data: function() {
    return {
      inputField: "",
      todos: [],
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
            this.token = resp.ToObject().Token;
            console.log(this.token); 
        });
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.title-text {
  font-size: 22px;
}
</style>
