<template>
  <div class="outer-container">
    <div class="gear-container">
      <div class="gear" ref="days-gear">
        <img
          src="../assets/HackNJIT2024/gears/gear5.svg"
          :style="dayRotation"
        />
        <p>{{ format(days) }}</p>
      </div>
      <div class="gear" ref="hours-gear">
        <img
          src="../assets/HackNJIT2024/gears/gear5.svg"
          :style="hourRotation"
        />
        <p>{{ format(hours) }}</p>
      </div>
      <div class="gear" ref="minutes-gear">
        <img
          src="../assets/HackNJIT2024/gears/gear5.svg"
          :style="minuteRotation"
        />
        <p>{{ format(minutes) }}</p>
      </div>
      <div class="gear" ref="seconds-gear">
        <img
          src="../assets/HackNJIT2024/gears/gear5.svg"
          class="seconds"
          :style="secondRotation"
        />
        <p>{{ format(seconds) }}</p>
      </div>
      <h1>Days</h1>
      <h1>Hours</h1>
      <h1>Minutes</h1>
      <h1>Seconds</h1>
    </div>
    <div class="text-container">
      <h1 class="title">
        Hack <br />
        NJIT <br />
        2024
      </h1>
    </div>
    <img class="clocktower-img" src="../assets/HackNJIT2024/clocktower.png" />
  </div>
</template>


<script>
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
  components: {},
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
  height: fit-content;
  width: 75%;
  margin: 0 auto;
  margin-top: 0rem;
  display: flex;
  justify-content: center;
  align-content: center;
  gap: 7.5rem;
}
.gear-container {
  align-self: center;
  display: grid;
  grid-template-columns: 42.5% 1fr;
  grid-template-rows: 25% 25% 25% 25%;
  row-gap: 3.5rem;
  width: fit-content;
  position: relative;
  align-content: center;
}
.gear-container > h1 {
  grid-column: 2;
  align-self: center;
  font-size: 4rem;
  text-align: left;
}
.gear-container > h1:first-of-type {
  grid-row: 1;
}
.gear-container > h1:nth-of-type(2) {
  grid-row: 2;
}
.gear-container > h1:nth-of-type(3) {
  grid-row: 3;
}
.gear-container > h1:nth-of-type(4) {
  grid-row: 4;
}
.gear {
  position: relative;
  width: fit-content;
  transform: scale(1.25);
  display: flex;
  align-content: center;
  gap: 1rem;
  justify-self: start;
  grid-column: 1;
}
.gear p {
  position: absolute;
  font-size: 3rem;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.seconds {
  animation-name: secondRotation;
  animation-duration: 1s;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
}
.minutes {
  animation-name: minuteRotation;
  animation-duration: 60s;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
}
.text-container {
  align-self: center;
  height: fit-content;
}
.text-container h1 {
  text-align: left;
  font-size: 7.5rem;
  margin: 1rem 0;
}
.text-container h1.title {
  font-size: 12.5rem;
  border: 4px white solid;
  border-radius: 50px;
  padding: 1rem 1rem;
}
.temp-block {
  width: 300px;
  height: 50%;
  align-self: center;
  background: black;
}
.clocktower-img {
  width: 275px;
}
@keyframes secondRotation {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@keyframes minuteRotation {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>