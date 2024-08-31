<template>
  <button class="steampunk-button" :class="{ large: isLarge }">
    <img class="svg" src="../assets/HackNJIT2024/conveyorbelt.svg" />
    <p><slot></slot></p>
  </button>
</template>

<script>
export default {
  props: ["isLarge"],
};
</script>

<style scoped>
.steampunk-button {
  /* background-color: var(--main-bg-color); */
  background-color: #ad8e70;
  color: var(--color4);
  cursor: pointer;
  /* border: 0.25rem solid #d4b490; */
  border: none;
  border-radius: 1000rem;
  min-width: 15.625rem;
  position: relative;
  display: flex;
  align-content: center;
  padding: 0rem 0rem;
  gap: 0.5rem;
}
.svg {
  width: 400px;
  height: 65px;
}
.large .svg {
  width: 500px;
  height: 75px;
}
p {
  display: block;
  position: absolute;
  font-size: 2rem;
  left: 5rem;
  align-self: center;
  /* color: #361200; */
  color: var(--color4);
  width: 240px;
}
.large p {
  width: 340px;
}
</style>