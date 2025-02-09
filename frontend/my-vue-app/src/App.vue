<script setup>
import { ref } from "vue";

const message = ref("Click the button to fetch data...");
const apiUrl = "/api"; // Nginx will proxy to backend

const fetchData = async () => {
  try {
    const response = await fetch(`${apiUrl}/hello`, {
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
};
</script>

<template>
  <div class="container">
    <h1>Vue + Golang Communication</h1>
    <p>{{ message }}</p>
    <button @click="fetchData">Fetch Data from Backend</button>
  </div>
</template>

<style>
.container {
  text-align: center;
  margin-top: 50px;
}

button {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}

button:hover {
  background-color: #36956d;
}
</style>
