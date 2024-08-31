<template>
  <div ref="gradient" class="gradient"></div>
  <div class="background"></div>
  <div class="outer-container">
    <!-- <div class="text-container">
      <h1 class="title">
        Hack <br />
        NJIT <br />
        2024
      </h1>
    </div> -->
    <header>
      <h1>{{ format(days) }}</h1>
      <GearColon class="gear-colon" />
      <h1>{{ format(hours) }}</h1>
      <GearColon />
      <h1>{{ format(minutes) }}</h1>
      <GearColon />
      <h1>{{ format(seconds) }}</h1>
    </header>
    <div class="countdown-gradient"></div>
  </div>
</template>

<script>
import GearColon from "./GearColon.vue";

export default {
  data() {
    return {
      milliseconds: 0,
      seconds: 0,
      minutes: 0,
      ten_minutes: 0,
      hours: 0,
      ten_hours: 0,
      days: 0,
      ten_days: 0,
    };
  },
  components: {
    GearColon,
  },
  methods: {
    setTime() {
      var timeLeft = new Date("2024/11/02 10:00:00") - new Date();

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
      var s = "" + Math.floor(number);
      return s.padStart(2, "0");
    },
  },
  computed: {
    secondRotation() {
      //   var rotationPercent = 360 - (this.milliseconds / 1000) * 360;
      //   return "transform: rotate(" + rotationPercent + "deg);";
      return "";
    },
    minuteRotation() {
      var rotationPercent = 360 - (this.seconds / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
    hourRotation() {
      var rotationPercent = 360 - (this.minutes / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
    dayRotation() {
      var rotationPercent = 360 - (this.hours / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
  },
  mounted() {
    setInterval(() => {
      this.setTime();
    }, 1000);
  },
};
</script>
<style scoped>
* {
  transition: all 0.1s linear;
}
.outer-container {
  color: white;
  height: fit-content;
  width: 75%;
  height: 100vh;
  margin: 0 auto;
  margin-top: 2.5rem;
  width: fit-content;
}
.background {
  background: black;
  width: 100%;
  height: 100vh;
  position: absolute;
  z-index: -11;
  top: 0;
}
.gradient {
  /* background: linear-gradient(
    to right,
    #dbbe9b70,
    #dbbe9ba9 10%,
    #dbbe9be1 20%,
    #dbbe9be1 80%,
    #dbbe9ba9 90%,
    #dbbe9b70
  ); */
  background: url("../assets/HackNJIT2024/repeating-bg.jpg");
  animation-name: fade-in;
  animation-duration: 10s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  width: 100%;
  height: 100vh;
  position: absolute;
  z-index: -10;
  top: 0;
  opacity: 0.25;
}
header {
  margin: 0 auto;
  display: flex;
  gap: 0.5rem;
  align-content: center;
}
.gear-colon {
  align-self: center;
}
h1 {
  font-size: 12rem;
}

.text-container {
  align-self: center;
  height: fit-content;
}
.text-container h1 {
  font-size: 10rem;
  border: 4px white solid;
  border-radius: 50px;
  padding: 1rem 1rem;
  width: fit-content;
  margin: 0 auto;
}

.countdown-gradient {
  background: linear-gradient(
    to right,
    #00000090,
    #00000060 10%,
    #00000035 20%,
    #00000035 80%,
    #00000060 90%,
    #00000090
  );
  width: 100%;
  height: 100vh;
  position: absolute;
  left: 0;
  z-index: -10;
  top: 0;
}
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 0.25;
  }
}
</style>