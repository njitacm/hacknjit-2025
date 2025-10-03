<template>
  <div class="outer-container gradient" ref="container">
    <Transition>
      <header v-show="containerIsVisible">
        <p class="countdown">
          {{ formattedDays }} : {{ formattedHours }} : {{ formattedMinutes }} : {{ formattedSeconds }}
        </p>
      </header>
    </Transition>
    <p class="date">
      {{ tweened_month.toFixed(0) }} /
      {{ tweened_day.toFixed(0) }} /
      {{ tweened_year.toFixed(0) }}
    </p>
  </div>
</template>

<script setup>
import gsap from "gsap";
import { ref, computed, onMounted } from "vue";
import { useIntersectionObserver } from "@vueuse/core";

const container = ref(null);
const containerIsVisible = ref(false);

useIntersectionObserver(container, ([{ isIntersecting }]) => {
  containerIsVisible.value = isIntersecting;
});

const targetDate = new Date(`2025/11/8 12:00:00`);
const days = ref(0);
const hours = ref(0);
const minutes = ref(0);
const seconds = ref(0);

const tweened_year = ref(0);
const tweened_month = ref(0);
const tweened_day = ref(0);

const formattedDays = computed(() => String(Math.floor(days.value)).padStart(2, "0"));
const formattedHours = computed(() => String(Math.floor(hours.value)).padStart(2, "0"));
const formattedMinutes = computed(() => String(Math.floor(minutes.value)).padStart(2, "0"));
const formattedSeconds = computed(() => String(Math.floor(seconds.value)).padStart(2, "0"));

function updateCountdown() {
      var timeLeft = targetDate - new Date();
      days.value = timeLeft / 8.64e7;
      hours.value = (timeLeft  / 3.6e6) % 24;
      minutes.value = (timeLeft / 60000) % 60;
      seconds.value = (timeLeft / 1000) % 60;
}

onMounted(() => {
  setInterval(updateCountdown, 1000);
  
  gsap.to(tweened_year, { value: targetDate.getFullYear(), duration: 4.5, ease: "expo" });
  gsap.to(tweened_month, { value: targetDate.getMonth() + 1, duration: 4.5, ease: "expo" });
  gsap.to(tweened_day, { value: targetDate.getDate(), duration: 4.5, ease: "expo" });
});

</script>

<style scoped>
* {
  transition: all 0.1s linear;
  box-sizing: border-box;
}

.outer-container {
  color: white;
  height: fit-content;
  width: 75%;
  /* height: calc(100vh); */
  margin: 0 auto;
  width: fit-content;
  position: relative;
  overflow: hidden;
}

.date {
  font-size: 2em;
}

.countdown {
  font-size: 3em;
}

@media (max-width: 1000px) {
  .date {
    font-size: 1.5em;
  }

  .countdown {
    font-size: 2.5em;
  }
}

@media (max-width: 425px) {
  .date {
    font-size: 1.25em;
  }

  .countdown {
    font-size: 2em;
  }
}

.gradient {
  animation-name: fade-in;
  animation-duration: 10s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  width: 100%;
}

header {
  margin: 0 auto;
  display: flex;
  gap: 0.5rem;
  align-content: center;
  justify-content: center;
}

h3 {
  animation-name: fade-in;
  opacity: 0;
  animation-delay: 2s;
  animation-duration: 2.5s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}

@keyframes fade-in {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

.v-enter-active,
.v-leave-active {
  transition: opacity 5s linear;
  transition-delay: 0.5s;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>