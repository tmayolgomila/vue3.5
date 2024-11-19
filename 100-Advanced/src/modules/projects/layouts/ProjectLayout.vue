<template>
  <div class="flex flex-col h-screen bg-white dark:bg-black">

    <!-- Header -->
    <header class="flex items-center h-14 px-4 border-b border-gray-300 sm:h-16 md:px-6 lg:px-8">
      <div>
        <a class="flex items-center gap-2 font-semibold" href="#">
          <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="40" height="40" />
        </a>
      </div>

      <nav class="flex items-center ml-auto space-x-4 h-10 sm:space-x-6">
        <div class="flex items-center space-x-4">
          <RouterLink :to="{ name: 'projectList' }">All Projects</RouterLink>
          <RouterLink :to="{ name: 'projectNew' }">New Project</RouterLink>

          <label class="relative inline-flex cursor-pointer items-center">
            <input id="switch" type="checkbox" class="peer sr-only" @change="toggleDarkMode" />
            <label for="switch" class="hidden"></label>
            <div
              class="peer h-6 w-11 rounded-full border bg-slate-200 after:absolute after:left-[2px] after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-gray-300 after:bg-white after:transition-all after:content-[''] peer-checked:bg-slate-800 peer-checked:after:translate-x-full peer-checked:after:border-white peer-focus:ring-green-300">
            </div>
          </label>
        </div>
      </nav>
    </header>
    <!-- Fin Header -->

    <!-- Main -->
    <main class="flex-1 flex items-center justify-center py-6">
      <RouterView />
    </main>
    <!-- Fin Main -->

    <!-- Footer -->
    <footer class="flex items-center h-14 px-4 border-t border-gray-300 sm:h-16 md:px-6 lg:px-8">
      <p class="flex-1 text-sm text-gray-500 text-center">
        © {{ new Date().getFullYear() }} Acme Corporation. Derechos reservados
      </p>
    </footer>
    <!-- Fin Footer -->

  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';

const isDarkMode = ref(false)

const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value
  if (isDarkMode.value) {
    document.body.classList.add('dark')
    localStorage.setItem('darkMode', 'true')
  } else {
    document.body.classList.remove('dark')
    localStorage.setItem('darkMode', 'false')
  }
}

onMounted(() => {

  if (localStorage.getItem('darkMode') == 'true') {
    document.body.classList.add('dark')
    isDarkMode.value = true

  } else {
    document.body.classList.remove('dark')
    isDarkMode.value = false

  }
})

</script>
