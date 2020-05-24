<template>
  <div class="dropdown">
    <button
      class="btn btn-outline-light dropdown-toggle"
      type="button"
      id="dropdownMenuButton"
      data-toggle="dropdown"
      aria-haspopup="true"
      aria-expanded="false"
    >
      <i class="fa fa-cog"></i>
    </button>
    <div
      class="dropdown-menu dropdown-menu-right"
      aria-labelledby="dropdownMenuButton"
    >
      <button class="dropdown-item">Profile</button>
      <button class="dropdown-item" v-on:click="handleLogout">Log out</button>
      <div class="dropdown-divider"></div>
      <a class="dropdown-item" href="#">Admin</a>
    </div>
  </div>
</template>
<script>
import { mapActions } from 'vuex'

export default {
  name: 'user-settings',
  methods: {
    ...mapActions(['logout']),
    handleLogout() {
      this.logout(this)
    }
  }  
}
</script>
