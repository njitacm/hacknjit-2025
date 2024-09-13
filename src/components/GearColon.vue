<template>
  <div class="colon-container">
    <img class="top-gear" src="../assets/HackNJIT2024/gears/gear5.svg" />
    <img class="bottom-gear" src="../assets/HackNJIT2024/gears/gear5.svg" />
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.colon-container {
  width: fit-content;
  display: flex;
  flex-direction: column;
  gap: 0;
  margin: 0;
  height: fit-content;
  align-self: center;
}
img {
  margin: -4px 0;
  width: 55px;
  animation: rotate 4s linear 0s infinite;
}
.bottom-gear {
  animation-delay: 0.2s;
  animation-direction: reverse;
}
@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@media (max-width: 1000px) {
  img {
    width: 40px;
  }
}
@media (max-width: 750px) {
  img {
    width: 32.5px;
    margin: 1px;
  }
}
@media (max-width: 600px) {
  img {
    width: 25.5px;
  }
}
@media (max-width: 500px) {
  img {
    width: 20.5px;
  }
}
</style>