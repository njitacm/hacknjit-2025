<template>
  <span class="container">
    <img src="../assets/Bulb.png" />
    <p class="number">{{ count }}</p>
  </span>
</template>

<script>
export default {
  props: ["count"],
};
</script>

<style scoped>
.container {
  position: relative;
}
.number {
  position: absolute;
  bottom: 110px;
  left: 35px;
  font-size: 4rem;
  background: -webkit-linear-gradient(#d39d4d, #e95c5c);
  text-shadow: 1px 1px 20px #d39d4d, -1px 1px 20px #e95c5c,
    -1px -1px 20px #d39d4d, -1px 1px 20px #e95c5c;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>