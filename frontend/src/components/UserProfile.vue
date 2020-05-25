<template>
  <div id="user-profile">
    <pre>
      {{ getCurrentUser() }}
    </pre>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'

export default {
  name: 'user-profile',
  data() {
    return {
      user: {}
    }
  }, 
  methods: {
    ...mapGetters(['getCurrentUser']),
  },
  created() {
    this.$store.dispatch('fetchCurrentUser')
  } 
}
</script>
