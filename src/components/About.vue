<template>
  <div class="About section" ref="sectionRef" id="About">
    <h2 class="section-title">About</h2>
    <div class="container">
      <img src="../assets/past_pics/hacknjit2022_1.jpg" class="main-img" />
      <div class="text-content">
        <p>HackNJIT is a 24-hour hackathon hosted anually at the New Jersey Institute of Technology. We're open to all, whether you're a beginner or pro. Our goal is to empower students to be creative, learn new things outside the classroom, and share memorable experiences.</p>
        <div class="key-info-container">
          <p class="key-info-header"><img src="../assets/icons/date.svg" class="icon" /><span>When?</span></p>
          <p class="key-info">November 8, 2025 - November 9, 2025</p>
        </div>
        <div class="key-info-container">
          <p class="key-info-header"><img src="../assets/icons/location.svg" class="icon" /><span>Where?</span></p>
          <p class="key-info">150 Bleeker St #1982, Newark, NJ 07102 (Campus Center Ballrooms)</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useIntersectionObserver } from "../composables/useIntersectionObserver";
const { observe, unobserve } = useIntersectionObserver();

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
.container {
  display: flex;
  gap: 16px;
}

.text-content {
  display: grid;
  align-content: space-between;
  gap: 16px;
  text-align: left;
}

.main-img {
  width: 500px;
  border-radius: 10px;
}

.key-info-container {
  display: grid;
  gap: 8px;
}

.key-info-header {
  background-color: white;
  color: var(--hacknjit-primary);
  border-radius: 10px;
  padding: 0 12px;
  display: flex;
  gap: 8px;
  align-items: center;
  width: fit-content;
}

.key-info-header span {
  font-weight: bold;
}

.icon {
  height: 1em;
}

.key-info {
  padding: 0 8px;
}

@media(max-width: 900px) {
  .container {
    flex-direction: column;
    align-items: center;
  }

  .main-img {
    width: 100%;
  }
}
</style>