<template>
  <div class="About" ref="sectionRef" id="About">
    <h2>About</h2>
    <div class="container">
      <img />
      <div class="text-content">
        <p>HackNJIT is a 24-hour hackathon hosted anually at the New Jersey Institute of Technology. Our mission is to
          empower students to create...</p>
        <p><img />Where? </p>
        <p><img />When? </p>
      </div>
    </div>
  </div>
</template>

<script>
import { observe, unobserve } from "../composables/useIntersectionObserver";

export default {
  mounted() {
    observe(this.$refs.sectionRef);
  },
  unmounted() {
    unobserve(this.$refs.sectionRef);
  }
}

</script>

<style scoped>
.content {
  display: flex;
  gap: 16px
}
</style>