<template>
  <div class="Banner" :style="{ background: gradient }">
    <div class="page-side-padding container">
      <div class="top-container">
        <h1>HackNJIT</h1>
        <p class="inner_text">
          HackNJIT is a 24-hour hackathon at the New Jersey Institute of
          Technology, run by its ACM student chapter in conjunction with the Ying
          Wu College of Computing.
        </p>
        <p>
          Stay tuned, we'll return in November 2025!
        </p>
        <NotifSignup id="NotifSignup" />
      </div>
      <div class="bottom-container">
        <TheCountdown />
        <div class="view-hint" ref="viewHint">
          <span>View last year's photos</span>
          <br />
          <br />
          <img src="../assets/down_arrow.svg" style="width: 25px; height: 25px" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TheCountdown from "./TheCountdown.vue";
import NotifSignup from "./NotifSignup.vue";

export default {
  props: {
    gradient: String
  },
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
  grid-template-rows: 1fr auto;
  gap: 1rem;
  margin: 0 auto;
  display: grid;
  gap: 30px;
  width: 90%;
  max-width: 1500px;
  justify-items: center;
  align-content: space-between;
  align-items: center;
}

p {
  font-size: 2em;
}

@media (max-width: 1000px) {
  p {
    font-size: 1.5em;
  }
}

@media (max-width: 425px) {
  h1 {
    font-size: 3em;
  }

  p {
    font-size: 1.25em;
  }

  .top-container {
    align-self: last baseline;
  }
}

.view-hint {
  align-content: center;
}

.view-hint * {
  margin: 0 auto;
}

.top-container,
.bottom-container {
  display: grid;
  gap: 32px;
}
</style>