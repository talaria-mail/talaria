<template>
  <div id="main">
    <Navbar/>
    <Composer/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Composer from '@/components/Composer.vue';

export default {
  name: "main",
  components: {
    Navbar,
    Composer
  }
};
</script>
