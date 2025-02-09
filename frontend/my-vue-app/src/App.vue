<script setup>
import { ref, onMounted } from 'vue'

const message = ref("Loading...");
const apiUrl = "/api"; // Nginx will proxy to backend

onMounted(async () => {
  try {
    const response = await fetch(`${apiUrl}/`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
    });

    if (!response.ok) throw new Error("Network response was not ok");

    const data = await response.json();
    message.value = data.message;
  } catch (error) {
    console.error("Error fetching data:", error);
    message.value = "Failed to fetch data";
  }
});
</script>

<template>
  <div>
    <h1>Vue + Golang</h1>
    <p>{{ message }}</p>
  </div>
</template>

<style>
h1 {
  color: #42b983;
}
</style>
