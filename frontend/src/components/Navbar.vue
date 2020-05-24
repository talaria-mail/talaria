<template>
  <nav class="navbar navbar-dark bg-dark">
    <!-- Navbar content -->
    <div class="container">
      <a class="navbar-brand" href="#">
        <img
          class="d-inline-block align-top"
          loading="lazy"
          src="../assets/talaria-light.svg"
          height="30px"
        />
        Talaria
      </a>
      <UserSettings/>
    </div>
  </nav>
</template>
<script>
import UserSettings from './UserSettings.vue'
export default {
  name: 'navbar',
  components: {
    UserSettings 
  }
}
</script>
<style>
a.navbar-brand {
  font-family: "Playfair Display", serif;
}
</style>
