<template>
  <div class="outer-container">
    <img
      v-for="i in 20"
      :key="i"
      :class="[i % 2 == 0 ? 'gearSpinInverted' : 'gearSpin']"
      src="../assets/HackNJIT2024/gears/gear4.svg"
    />
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.outer-container {
  display: flex;
  justify-content: center;
  margin: 0 auto;
  width: 67.5%;
  padding: 0;
  border: 4px var(--color4) solid;
  border-radius: 30px;
  padding: 0.25rem;
}
img {
  width: 5%;
  animation-duration: 4s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  filter: brightness(0) saturate(100%) invert(43%) sepia(9%) saturate(1781%)
    hue-rotate(342deg) brightness(89%) contrast(91%);
}
.gearSpin {
  animation-name: gearSpin;
}
.gearSpinInverted {
  animation-name: gearSpin;
  animation-direction: reverse;
}
@keyframes gearSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
  /* 100% {
    transform: rotate(360);
  } */
}
</style>