<template>
  <div ref="container">
    <Transition>
      <button v-show="containerIsVisible" @click="$emit('click-emit')">
        <slot></slot>
      </button>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useIntersectionObserver } from "@vueuse/core";

const container = ref(null);
const containerIsVisible = ref(false);

useIntersectionObserver(container, ([{ isIntersecting }]) => {
  containerIsVisible.value = isIntersecting;
});
</script>

<style scoped>
button {
  position: relative;
  z-index: 1;
  padding: 0.5rem;
  font-size: 2rem;
  background-color: var(--main-bg-color);
  border-radius: 0.5rem;
  border-color: var(--main-fg-color);
  border-width: 4px;
  z-index: 2;
}
button * {
  color: black;
  text-decoration: none;
}
@media (max-width: 1500px) {
  button {
    font-size: 1.8rem;
  }
}
@media (max-width: 1000px) {
  button {
    font-size: 1.6rem;
  }
}
@media (max-width: 750px) {
  button {
    font-size: 1.4rem;
  }
}
@media (max-width: 600px) {
  button {
    font-size: 1.2rem;
  }
}
@media (max-width: 500px) {
  button {
    font-size: 1rem;
  }
}
@media (max-width: 420px) {
  button {
    font-size: 0.75rem;
  }
}
.v-enter-active,
.v-leave-active {
  transition: opacity 5s linear;
  transition-delay: 2s;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
