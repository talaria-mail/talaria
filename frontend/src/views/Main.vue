<template>
  <div id="main">
    <Navbar /> 
    <main class="container">
      <div class="row">
        <div class="col-sm">
          <Menu /> 
        </div>
        <div class="col-lg">
          <Composer/>
        </div>
      </div> 
    </main>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import Menu from '@/components/Menu.vue'
import Composer from '@/components/Composer.vue';

export default {
  name: "Main",
  components: {
    Navbar,
    Menu,
    Composer
  }
};
</script>
<style>
main.container {
  height: 100%;
}

.container {
  max-width: 1600px;
}
</style>
