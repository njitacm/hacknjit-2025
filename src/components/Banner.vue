<template>
  <div class="Banner container">
    <NotifSignup id="NotifSignup" />
    <TheCountdown />
    <div class="view-hint" ref="viewHint">
      <span>View last year's photos</span>
      <br />
      <br />
      <img src="../assets/down_arrow.svg" style="width: 25px; height: 25px" />
    </div>
  </div>
</template>

<script>
import TheCountdown from "./TheCountdown.vue";
import NotifSignup from "./NotifSignup.vue";

export default {
  components: {
    TheCountdown,
    NotifSignup
  }
}

</script>

<style scoped>
.container {
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