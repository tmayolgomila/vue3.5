<template>
<nav>

  <RouterLink to="/">Home</RouterLink>
  <RouterLink to="/features">Features</RouterLink>
</nav>

<main class="bg-slate-100">
  <RouterView/>
</main>

</template>


<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';


</script>

