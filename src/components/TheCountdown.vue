<template>
  <div class="outer-container" ref="container">
    <Transition appear>
      <div>
        <p class="countdown">
          {{ formattedDays }} : {{ formattedHours }} : {{ formattedMinutes }} : {{ formattedSeconds }}
        </p>
        <p class="date">
          {{ tweened_month.toFixed(0) }} /
          {{ tweened_day.toFixed(0) }} /
          {{ tweened_year.toFixed(0) }}
        </p>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import gsap from "gsap";
import { ref, computed, onMounted, useTemplateRef } from "vue";
import { useIntersectionObserver } from "@vueuse/core";

const container = useTemplateRef("container");
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
  hours.value = (timeLeft / 3.6e6) % 24;
  minutes.value = (timeLeft / 60000) % 60;
  seconds.value = (timeLeft / 1000) % 60;
}

onMounted(() => {
  updateCountdown();
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
  left: 50%;
  transform: translateX(-50%);
  width: fit-content;
  position: relative;
  overflow: hidden;
}

.countdown, .date {
  text-align: center;
}

.countdown {
  font-size: 3em ;
  white-space: nowrap;
}

.date {
  font-size: 2em;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 1s linear;
  transition-delay: 0.5s;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.v-enter-to,
.v-leave-from {
  opacity: 1;
}

@media(max-width: 600px) {
  .outer-container {
    font-size: 0.75em;
  }
}
</style>
