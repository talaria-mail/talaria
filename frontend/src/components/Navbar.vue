<template>
  <nav class="navbar navbar-dark bg-dark">
    <!-- Navbar content -->
    <div class="container">
      <a class="navbar-brand" href="#">
        <img class="d-inline-block align-top" loading="lazy" src="../assets/talaria-light.svg" height="30px"> 
        Talaria
      </a>
      <div class="dropdown">
        <button class="btn btn-outline-light dropdown-toggle" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          <i class="fa fa-cog"></i>
        </button>
        <div class="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenuButton">
          <a class="dropdown-item" href="#">Profile</a>
          <div class="dropdown-divider"></div>
          <a class="dropdown-item" href="#">Admin</a>
        </div>
      </div>
    </div>
  </nav>
</template>
<script>

</script>
<style>
a.navbar-brand {
 font-family: 'Playfair Display', serif;
}
</style>
