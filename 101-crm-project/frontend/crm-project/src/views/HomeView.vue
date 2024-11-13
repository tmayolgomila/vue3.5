<template>

  <div class="container">
    <h1 class="text-3xl font-bold underline">Create User</h1>

    <form @submit.prevent="createUser" class="mt-4">
      <div>
        <label for="name" class="block">Name: </label>
        <input type="text" v-model="user.name" id="name" placeholder="Name here" class="mt-2 p-2 border rounded"
          required />
      </div>
      <div class="mt-4">
        <label for="email" class="block">Email:</label>
        <input v-model="user.email" type="email" id="email" placeholder="Email here" class="mt-2 p-2 border rounded"
          required />
      </div>
      <button type="submit" class="mt-4 p-2 bg-blue-500 text-white rounded">Create User</button>
    </form>

    <!-- Success Message-->
    <div v-if="createdUser" class="mt-4">
      <h2 class="text-xl font-semibold">User created:</h2>
      <p>Name: {{ createdUser.name }}</p>
      <p>Email: {{ createdUser.email }}</p>
    </div>

    <!-- Error Message -->
    <div v-if="errorMessage" class="mt-4 text-red-500">
      <p>{{ errorMessage }}</p>
    </div>

  </div>

</template>

<script setup>
import axios from 'axios';
import { ref } from 'vue';


const user = ref({
  name: '',
  email: ''
})

const createdUser = ref(null)
const errorMessage = ref(null)

const createUser = async () => {
  try {
    const response = await axios.post('http://192.168.10.86:8080/users', user.value);
    createdUser.value = response.data;
    errorMessage.value = null
  } catch (error) {
    errorMessage.value = 'Error creating user: ' + (error.response?.data?.message || 'Server error');
    createdUser.value = null;

  }
}

</script>

<style scoped>
.container {
  max-width: 400px;
  margin: auto;
  padding: 20px;
}
</style>
