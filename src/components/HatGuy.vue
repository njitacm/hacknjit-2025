<template>
  <img ref="hatguyimg" src="../assets/HackNJIT2024/hatguy/hatguy0.png" />
</template>

<script>
export default {
  data() {
    return {
      hats: [
        require("../assets/HackNJIT2024/hatguy/hatguy0.png"),
        require("../assets/HackNJIT2024/hatguy/hatguy1.png"),
        require("../assets/HackNJIT2024/hatguy/hatguy2.png"),
        require("../assets/HackNJIT2024/hatguy/hatguy3.png"),
        require("../assets/HackNJIT2024/hatguy/hatguy4.png"),
      ],
    };
  },
  mounted() {
    // console.log(this.$refs.hatguyimg.src);
    this.$refs.hatguyimg.src = this.hats[Math.floor(5 * Math.random(0))];
  },
};
</script>

<style scoped>
img {
  z-index: -2;
  /* width: 40vh; */
}
@media (max-width: 1500px) {
  img {
    width: unset;
    max-height: 37.5vh;
  }
}
@media (max-width: 750px) {
  img {
    width: unset;
    max-height: 32.5vh;
  }
}
@media (max-width: 500px) {
  img {
    width: unset;
    max-height: 27.5vh;
  }
}
</style>