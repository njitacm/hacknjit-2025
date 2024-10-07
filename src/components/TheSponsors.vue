<template>
  <div ref="container" class="outer-container">
    <div class="title" v-intersection-observer="[onIntersectionObserver]">
      <!-- <img class="title-svg" src="../assets/HackNJIT2024/gears/gear4.svg" /> -->
      <h1 class="title">Our Sponsors</h1>
      <!-- @click="
          window.open(
            'https://cdn.discordapp.com/attachments/1193704838010253323/1275603312468557877/HackNJIT_2024_Sponsorship_Packet.pdf?ex=66d3aca9&is=66d25b29&hm=acccbe5ad986f1124a1da9bc6aa7c65d2666a07f36caae3209df5fbbb9644fdc&'
          )
        " -->
      <!-- <h2 @click="toggleSponsors">Sponsor Packet</h2> -->
      <!-- <img class="title-svg" src="../assets/HackNJIT2024/gears/gear4.svg" /> -->
    </div>
    <div class="sponsors" v-intersection-observer="[onIntersectionObserver]">
      <!-- <h1>Title Sponsor</h1>
        <h1>Gold Sponsor</h1>
        <h1>Silver Sponsor</h1>
        <h1>Bronze Sponsor</h1> -->
      <div class="sponsor-grid">
        <img
          class="sponsor"
          src="../assets/HackNJIT2024/Sponsors/Merck.png"
          alt="Merck"
        />
        <img
          class="sponsor"
          src="../assets/HackNJIT2024/Sponsors/ISACA_nobg.png"
          alt="Isaca"
        />
        <img
          class="sponsor"
          style="width: 85%; height: 85%"
          src="../assets/HackNJIT2024/Sponsors/AntlionAudio.png"
          alt="Antlion Audio"
        />
        <img
          class="sponsor"
          id="standoutStickers"
          src="../assets/HackNJIT2024/Sponsors/StandOut-Stickers-Logo.png"
          alt="Isaca"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    console.log(target);
    target.classList.add("fade-in");
  }
}
</script>
<script>
export default {
  data() {
    return {
      showSponsors: true,
    };
  },
  methods: {
    toggleSponsors() {
      this.showSponsors = !this.showSponsors;
    },
  },
};
</script>

<style scoped>
.sponsor-grid {
  display: grid;
  grid-template-columns: repeat(2, 50%);
  grid-template-rows: repeat(2, 400px);
  flex-wrap: wrap;
  align-content: center;
  justify-content: space-around;
  background: #8f3d0625;
  /* border: 8px solid var(--main-bg-color); */
  border-bottom: 4px solid var(--main-bg-color);
  /* border-radius: 32px; */
  gap: 0.5rem;
  padding: 1rem;
}
.sponsor {
  width: 100%;
  height: 100%;
  object-fit: contain;
  align-self: center;
  justify-self: center;
}
#standoutStickers {
  grid-column: 2;
}
div.outer-container {
  width: 100%;
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
}
h1.title {
  font-size: 4rem;
  font-weight: bold;
  /* border-bottom: 8px solid white; */
  margin-bottom: 0.75rem;
}
h2 {
  display: inline-block;
  cursor: pointer;
  width: fit-content;
  justify-self: center;
}
h2::after {
  content: "";
  width: 0px;
  height: 4px;
  display: block;
  background: white;
  transition: 300ms;
}
h2:hover::after {
  width: 100%;
}
div.title {
  padding-top: 1.5rem;
  border-top: 4px solid var(--main-bg-color);
  background: #8f3d0625;
  display: grid;
  grid-template-columns: 30% 1fr 30%;
  grid-template-rows: 100%;
  justify-content: center;
  column-gap: 1rem;
}
div.title h1 {
  grid-column: 2;
  grid-row: 1/2;
  align-self: center;
}
div.title h2 {
  grid-column: 2;
}
div.title img {
  width: 6rem;
  animation-name: gearSpin;
  animation-duration: 4s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  grid-row: 1 / 3;
}
div.title img:nth-of-type(1) {
  grid-column: 1;
  justify-self: right;
}
div.title img:nth-of-type(2) {
  grid-column: 3;
  justify-self: left;
  animation-direction: reverse;
}
h1 {
  font-size: 3rem;
}
table {
  width: 100%;
  border-collapse: collapse;
  border-radius: 1rem;
  border-style: hidden; /* hide standard table (collapsed) border */
  box-shadow: 0 0 0 2px white; /* this draws the table border  */
}
thead th:first-of-type {
  padding: none;
  margin: none;
  border: 2px white solid;
  border-radius: 0 1rem 0 5px;
}
tr.row-title {
  border: 2px white solid;
}
tr.row-title td:first-of-type {
  border: 2px white solid;
}
tr td:first-of-type {
  border-right: 2px white solid;
}
td {
  margin: 0.75rem;
  word-wrap: break-word;
  max-width: 75px;
}
td img {
  width: 2rem;
}
td:last-of-type {
  padding: 0.75rem;
}
@keyframes gearSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
  /* 100% {
    transform: rotate(360);
  } */
}
.title-gear {
  filter: brightness(0) saturate(100%) invert(17%) sepia(71%) saturate(3534%)
    hue-rotate(15deg) brightness(94%) contrast(108%);
}
.gold-gear {
  filter: brightness(0) saturate(100%) invert(71%) sepia(23%) saturate(3958%)
    hue-rotate(4deg) brightness(97%) contrast(102%);
}
.silver-gear {
  filter: brightness(0) saturate(100%) invert(62%) sepia(24%) saturate(151%)
    hue-rotate(355deg) brightness(88%) contrast(89%);
}
.bronze-gear {
  filter: brightness(0) saturate(100%) invert(25%) sepia(44%) saturate(548%)
    hue-rotate(5deg) brightness(100%) contrast(93%);
}
.v-enter-active,
.v-leave-active {
  transition: opacity 0.25s linear;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
.fade-in {
  animation-name: fade-in;
  animation-duration: 0.75s;
  /* animation-delay: 1.5s; */
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
@media (max-width: 1250px) {
  div.title img {
    width: 9rem;
  }
}
@media (max-width: 800px) {
  div.title img {
    width: 7rem;
  }
}
@media (max-width: 550px) {
  .title-svg {
    display: none;
  }
  .outer-container {
    width: 100%;
  }
  div.title h1.title {
    display: inherit;
  }
  h1.title {
    width: 90vw;
    justify-self: center;
    text-align: center;
    font-size: 3rem;
  }
  h1 {
    font-size: 2.5rem;
  }
}
</style>