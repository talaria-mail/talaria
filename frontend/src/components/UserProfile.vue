<template>
  <div id="user-profile">
    <pre>
      {{ user }}
    </pre>
  </div>
</template>
<script>

export default {
  name: 'user-profile',
  computed: {
    user() {
      return this.$store.state.users.current
    }
  },
  created() {
    this.$store.dispatch('fetchCurrentUser')
  } 
}
</script>
