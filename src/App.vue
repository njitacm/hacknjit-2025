<template>
  <!-- <div class="gradient"></div> -->
  <MLHBanner />
  <TheHeader />
  <TheCountdownWillVersion />
  <TheIntro id="Intro" />
  <GearDivider />
  <TheRegistration id="Registration" />
  <GearDivider />
  <TheSponsors />
  <TheTeam id="Team" />
  <TheFAQ id="FAQ" />
  <TheFooter />
  <!-- <GearDivider />
  <NewCountdown />
  <TheOldIntro /> -->
</template>

<script>
import GearDivider from "./components/GearDivider.vue";
import MLHBanner from "./components/MLHBanner.vue";
// import NewCountdown from "./components/NewCountdown.vue";
import TheCountdownWillVersion from "./components/TheCountdownWillVersion.vue";
import TheFAQ from "./components/TheFAQ.vue";
import TheFooter from "./components/TheFooter.vue";
import TheHeader from "./components/TheHeader.vue";
import TheIntro from "./components/TheIntro.vue";
// import TheOldIntro from "./components/TheOldIntro.vue";
import TheRegistration from "./components/TheRegistration.vue";
import TheSponsors from "./components/TheSponsors.vue";
import TheTeam from "./components/TheTeam.vue";
export default {
  name: "HackNJIT",
  components: {
    TheHeader,
    MLHBanner,
    // TheOldIntro,
    TheTeam,
    TheFAQ,
    TheFooter,
    TheRegistration,
    TheSponsors,
    TheIntro,
    GearDivider,
    // NewCountdown,
    TheCountdownWillVersion,
  },
};
</script>

<style>
:root {
  --main-bg-color: #dab473;
  --main-fg-color: #8f5f49;
  --text-color: white;
  --secondary-bg-color: #3f2b12;
  --color5: #231b18;
  --color4: #856045;
  --color3: #6b4131;
  --color2: #a7793e;
  --color1: #887972;
  --mlh-banner: #f7f7f7;
  --mlh-banner-text: #122f4c;
  --mlh-banner-transparent1: #f7f7f77c;
  --mlh-banner-transparent2: #f7f7f7ab;
  --edge-colors: #d3b28eb2;
}
html {
  scroll-behavior: smooth;
  font-family: sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /* color: var(--color3); */
  color: white;
  box-sizing: border-box;
  font-size: 16px;
  /* background: #f8f6f5; */
  background: #1d1710;
  z-index: -100;
  /* background: url("./assets/HackNJIT2024/repeating-bg.jpg"); */
  padding-top: 40px;
}
/* .gradient {
  background: linear-gradient(
    to right,
    #dbbe9b70,
    #dbbe9ba9 10%,
    #dbbe9be1 20%,
    #dbbe9be1 80%,
    #dbbe9ba9 90%,
    #dbbe9b70
  );
  width: 100%;
  height: 100%;
  position: fixed;
  z-index: -10;
  top: 0;
} */
* {
  margin: 0;
  padding: 0;
  font-weight: 500;
}
</style>
