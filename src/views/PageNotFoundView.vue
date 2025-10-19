<template>
  <div class="PageNotFoundView View">
    <p class="input-path">{{ route.fullPath }}</p>
    <h1>Page Not Found</h1>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();
</script>

<style scoped>
.input-path {
  font-family: monospace;
  font-size: 5em;
  border-radius: 25px;
  border: 1px solid white;
  width: fit-content;
  margin: auto;
  padding: 0 16px;
}
</style>