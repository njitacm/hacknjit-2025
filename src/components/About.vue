<template>
  <div class="About section" ref="sectionRef" id="About">
    <h2 class="section-title">About</h2>
    <div class="container">
      <div class="text-content">
        <p>
          HackNJIT is a 24-hour hackathon hosted anually at the New Jersey Institute of Technology. We're open to all,
          whether you're a beginner or pro. Our goal is to empower students to be creative, learn new things outside the
          classroom, and share memorable experiences.
        </p>
        <div class="key-info-container">
          <div class="key-info-sub-container">
            <div class="key-info-header">
              <div class="title">
                <img src="../assets/icons/date.svg" class="icon" /><span>When?</span>
              </div>
              <a target="_blank" class="action pill"
                href="https://www.google.com/calendar/render?action=TEMPLATE&text=HackNJIT%202025&dates=20251108T090000/20251109T170000&ctz=America/New_York&details=HackNJIT%20is%20NJIT's%20annual%20student-run%20hackathon.%20Join%20us%20for%20a%20weekend%20of%20innovation,%20learning,%20and%20fun!&location=Campus%20Center,%20150%20Bleecker%20St%20%231982,%20Newark,%20NJ%2007102,%20USA">
                <span>Add to calendar</span>
                <img class="icon" src="../assets/icons/external_link.svg" />
              </a>
            </div>
            <p class="key-info">November 8, 2025 - November 9, 2025</p>
          </div>
          <div class="key-info-sub-container">
            <div class="key-info-header">
              <div class="title">
                <img src="../assets/icons/location.svg" class="icon" /><span>Where?</span>
              </div>
              <a target="_blank" class="action pill" href="https://maps.app.goo.gl/zpCuGThyr5TocsMe6">
                <span>View map</span>
                <img class="icon" src="../assets/icons/external_link.svg" />
              </a>
            </div>
            <p class="key-info">150 Bleeker St #1982, Newark, NJ 07102 (Campus Center Ballrooms)</p>
          </div>
        </div>
      </div>
      <PastPics />
    </div>
  </div>
</template>

<script>
import { useIntersectionObserver } from "../composables/useIntersectionObserver";
import PastPics from "./PastPics.vue";
const { observe, unobserve } = useIntersectionObserver();

export default {
  components: { PastPics },
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
  display: grid;
  gap: 32px;
}

.text-content {
  display: grid;
  align-content: space-between;
  gap: 16px;
  text-align: left;
}

.icon {
  height: 1.25em;
}

.key-info-header {
  display: flex;
  gap: 16px;
  justify-content: space-between;
}

.key-info-container {
  display: grid;
  grid-template-columns: auto auto;
  gap: 32px;
  align-items: start;
}

.key-info-sub-container {
  display: grid;
  gap: 16px;
}

.title {
  background-color: white;
  color: var(--hacknjit-primary);
  border-radius: var(--border-radius);
  padding: 8px 12px;
  display: flex;
  gap: 8px;
  align-items: center;
  font-weight: bold;
}

.action {
  display: flex;
  gap: 8px;
  align-items: center;
  justify-content: space-between;
}

.title span {
  font-weight: bold;
}

.key-info {
  padding: 0 8px;
}

@media(max-width: 750px) {
  .key-info-container {
    grid-template-columns: auto;
  }
}

@media(max-width: 400px) {
  .key-info-header {
    flex-direction: column;
  }
}
</style>