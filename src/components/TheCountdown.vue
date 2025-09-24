<template>
  <div class="outer-container gradient" ref="container">
    <h1>
      {{ format(tweened_month.toFixed(0)) }} /
      {{ format(tweened_day.toFixed(0)) }} /
      {{ tweened_year.toFixed(0) }}
    </h1>
    <Transition>
      <header v-show="containerIsVisible">
        <h3>
          {{ format(days) }}&nbsp;:&nbsp;{{ format(hours) }}&nbsp;:&nbsp;{{ format(minutes) }}&nbsp;:&nbsp;{{
            format(seconds) }}
        </h3>
      </header>
    </Transition>
  </div>
</template>

<script setup>
import gsap from "gsap";
import { ref } from "vue";
import { useIntersectionObserver } from "@vueuse/core";

const container = ref(null);
const containerIsVisible = ref(false);

useIntersectionObserver(container, ([{ isIntersecting }]) => {
  containerIsVisible.value = isIntersecting;
});
</script>

<script>
export default {
  data() {
    return {
      date: { year: 2025, month: 11, day: 8 },
      milliseconds: 0,
      seconds: 0,
      minutes: 0,
      ten_minutes: 0,
      hours: 0,
      ten_hours: 0,
      days: 0,
      ten_days: 0,
      date_month: null,
      date_day: null,
      date_year: null,
      tweened_month: 0,
      tweened_day: 0,
      tweened_year: 2000,
    };
  },
  watch: {
    date_year(n) {
      gsap.to(this, { duration: 5, tweened_year: Number(n) || 0 });
    },
    date_day(n) {
      gsap.to(this, { duration: 5, tweened_day: Number(n) || 0 });
    },
    date_month(n) {
      gsap.to(this, { duration: 5, tweened_month: Number(n) || 0 });
    },
  },
  methods: {
    setTime() {
      var timeLeft = new Date(`${this.date.year}/${this.date.month}/${this.date.day} 12:00:00`) - new Date();

      var daysLeft = timeLeft / 8.64e7;
      var hoursLeft = (timeLeft / 3.6e6) % 24;
      var minLeft = (timeLeft / 60000) % 60;
      var secsLeft = (timeLeft / 1000) % 60;
      var millisecondsLeft = (secsLeft % 1) * 1000;

      this.milliseconds = millisecondsLeft;
      this.seconds = secsLeft;
      this.minutes = minLeft;
      this.hours = hoursLeft;
      this.days = daysLeft;
    },
    format(number) {
      if (number < 0) {
        number = 0;
      }
      var s = "" + Math.floor(number);
      return s.padStart(2, "0");
    },
  },
  mounted() {
    setInterval(() => {
      this.setTime();
    }, 1000);
    setTimeout(() => {
      this.date_day = this.date.day;
      this.date_month = this.date.month;
      this.date_year = this.date.year;
    }, 2000);
  },
};
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