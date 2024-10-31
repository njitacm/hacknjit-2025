<template>
  <!-- <div class="gradient"></div> -->
  <MLHBanner />
  <TheHeader />
  <TheCountdown />
  <TheIntro id="Intro" />
  <TheSponsors id="Sponsors" />
  <TheTeam id="Team" />
  <TheFAQ id="FAQ" />
  <TheFooter />
  <!-- <GearDivider />
  <NewCountdown />
  <TheOldIntro /> -->
</template>

<script>
if (window.location.pathname == "/register") {
  window.location.href = "https://forms.gle/MnPFuMQoVA5PCphp7";
} else if (window.location.pathname == "/discord") {
  window.location.href = "https://discord.gg/jDErKCXvjP";
} else if (window.location.pathname == "/devpost") {
  window.open("https://hacknjit-2024.devpost.com");
} else if (window.location.pathname != "/") {
  window.location.href = "";
  console.log(window.location.pathname);
}
import MLHBanner from "./components/MLHBanner.vue";
import TheCountdown from "./components/TheCountdown.vue";
import TheFAQ from "./components/TheFAQ.vue";
import TheFooter from "./components/TheFooter.vue";
import TheHeader from "./components/TheHeader.vue";
import TheIntro from "./components/TheIntro.vue";
import TheSponsors from "./components/TheSponsors.vue";
import TheTeam from "./components/TheTeam.vue";
export default {
  name: "HackNJIT",
  components: {
    TheHeader,
    TheCountdown,
    MLHBanner,
    TheTeam,
    TheFAQ,
    TheFooter,
    TheSponsors,
    TheIntro,
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
  color: white;
  box-sizing: border-box;
  font-size: 16px;
  background: #170800;
  z-index: -100;
  overflow-x: hidden;
  width: 100%;
}

* {
  margin: 0;
  padding: 0;
  font-weight: 500;
}
</style>
