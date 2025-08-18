<template>
  <div id="main-container">
    <!-- <MLHBanner /> -->
    <main class="body-container">
      <WhatIs></WhatIs>
      <div class="banner-container">
        <NotifSignup id="NotifSignup" />
        <TheCountdown />
        <div id="view-hint" ref="viewHint">
          <span>View last year's photos</span>
          <br />
          <span>&#129123;</span>
        </div>
      </div>
      <div id="past-pics" ref="past-pics">
        <img src="./assets/PastPictures/hacknjit2024_1.jpg" />
        <img src="./assets/PastPictures/hacknjit2024_2.jpg" />
        <img src="./assets/PastPictures/hacknjit2024_3.jpg" />
        <img src="./assets/PastPictures/hacknjit2024_4.jpg" />
      </div>
      <TheFAQ />
    </main>
    <TheFooter />
    <Modal />
  </div>
</template>

<script>
import MLHBanner from "./components/MLHBanner.vue";
import TheCountdown from "./components/TheCountdown.vue";
import TheFAQ from "./components/TheFAQ.vue";
import TheFooter from "./components/TheFooter.vue";
import TheHeader from "./components/TheHeader.vue";
import NotifSignup from "./components/NotifSignup.vue";
import TheSponsors from "./components/TheSponsors.vue";
import TheTeam from "./components/TheTeam.vue";
import Modal from "./components/Modal.vue";

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
    NotifSignup,
    Modal,
  }
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
  /* overflow-x: hidden; */
  overflow-x: auto;
  width: 100%;
  height: 100%;
}

body {
  height: 100%;
}

* {
  margin: 0;
  padding: 0;
  font-weight: 500;
}

button {
  -webkit-tap-highlight-color: transparent;
}

#main-container {
  display: grid;
  grid-template-rows: 1fr auto;
  place-content: center;
  place-items: center;
}

.banner-container {
  min-height: 100svh;
  height: fit-content;
  display: grid;
  gap: 1rem;
}

#past-pics {
  margin-top: 32px;
  padding: 20px;
  columns: 2 300px;
  column-gap: 1rem;
}

#past-pics img {
  border-radius: 0.5rem;
  display: block;
  margin-bottom: 1rem;
  width: 100%;
}

#view-hint {
  align-content: center;
}

#view-hint * {
  margin: 0 auto;
}
</style>
