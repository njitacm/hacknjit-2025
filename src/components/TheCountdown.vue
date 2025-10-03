<template>
  <div class="outer-container countdown-container" ref="container">
    <Transition>
      <header v-show="containerIsVisible">
        <p class="countdown">
          {{ format(days) }}&nbsp;:&nbsp;{{ format(hours) }}&nbsp;:&nbsp;{{ format(minutes) }}&nbsp;:&nbsp;{{
            format(seconds) }}
        </p>
      </header>
    </Transition>
    <p class="date">
      {{ format(tweened_month.toFixed(0)) }} /
      {{ format(tweened_day.toFixed(0)) }} /
      {{ format(tweened_year.toFixed(0), 4) }}
    </p>
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
      tweened_year: 0,
    };
  },
  watch: {
    date_year(n) {
      gsap.to(this, { duration: 4.5, tweened_year: Number(n) || 0, ease: "expo" });
    },
    date_day(n) {
      gsap.to(this, { duration: 4.5, tweened_day: Number(n) || 0, ease: "expo" });
    },
    date_month(n) {
      gsap.to(this, { duration: 4.5, tweened_month: Number(n) || 0, ease: "expo"});
    },
  },
  methods: {
    setTime() {
      var timeLeft = new Date(`${this.date.year}/${this.date.month}/${this.date.day} 12:00:00`) - new Date();

      this.days = timeLeft / 8.64e7;
      this.hours = (timeLeft / 3.6e6) % 24;
      this.minutes = (timeLeft / 60000) % 60;
      this.seconds = (timeLeft / 1000) % 60;
      this.milliseconds = (this.seconds % 1) * 1000;
    },
    format(number, digitsToPad = 2) {
      if (number < 0) {
        number = 0;
      }
      var s = "" + Math.floor(number);
      return s.padStart(digitsToPad, "0");
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
    }, 500);
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

.countdown-container {
  animation-name: fade-in;
  animation-duration: 5s;
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
  animation-delay: 2.5s;
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
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
