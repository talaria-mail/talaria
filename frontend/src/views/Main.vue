<template>
  <div id="Main">
    <Navbar />
    <template>
        <router-view/>
    </template> 
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'

export default {
  name: "Main",
  components: {
    Navbar,
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
