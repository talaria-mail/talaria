<template>
  <div id="app">
    <Login/> 
  </div>
</template>

<script>

import Login from "./components/Login.vue";

export default {
  name: "app",
  components: {
    Login 
  }
};
</script>

<style>
#app {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-image: url('./assets/beach.jpg');
  background-position: center; /* Center the image */
  background-repeat: no-repeat; /* Do not repeat the image */
  background-size: cover; /* Resize the background image to cover the entire container */
  height: 100%;
}
</style>
