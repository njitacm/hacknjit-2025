<template>
  <div class="Banner">
    <NotifSignup id="NotifSignup" />
    <TheCountdown />
    <div class="view-hint">
      <span>View last year's photos</span>
      <br />
      <span>&#129123;</span>
    </div>
  </div>
</template>

<script>
import NotifSignup from "./NotifSignup.vue";
import TheCountdown from "./TheCountdown.vue";

export default {
  components: {
    NotifSignup,
    TheCountdown
  }
}
</script>

<style scoped>
.Banner {
  min-height: 100svh;
  height: fit-content;
  display: grid;
  gap: 1rem;
}

.view-hint {
  align-content: center;
}

.view-hint * {
  margin: 0 auto;
}
</style>