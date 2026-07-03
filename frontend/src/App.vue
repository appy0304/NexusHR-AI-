<!-- <script setup>

import { ref } from "vue"
import axios from "axios"

const name=ref("")
const age=ref("")



async function createUser() {

  try {

    const payload={
      name:name.value,
      age:Number(age.value),
    }

    console.log("payload",payload)

    const response = await axios.post(
      "http://localhost:8080/users",
      payload
    )

    console.log(response.data)

    alert(JSON.stringify(response.data))

  } catch (error) {

    console.error(error)

  }

}
async function fetchUsers() {

  try {

    const response = await axios.get(
      "http://localhost:8080/users"
    )

    console.log(response.data)

    alert(JSON.stringify(response.data))

  } catch (error) {

    console.error(error)

  }

}


</script>

<template>

  <div>

    <h1>Simple API Testing</h1>

    <input 
    v-model="name"
    placeholder="enter the name" />

    <br /><br />

    <input 
    v-model="age"
    placeholder="enter the age"
    />
    <br /><br />

    <button @click="fetchUsers">
      Fetch Users
    </button>

    <button @click="createUser">
      Create User
    </button>

  </div>

</template> -->

<template>
  <router-view />
</template>
<script setup>
import { onMounted } from 'vue'
import { useThemeStore } from './stores/theme'

const themeStore = useThemeStore()
onMounted(() => {
  themeStore.init()
})
</script>
