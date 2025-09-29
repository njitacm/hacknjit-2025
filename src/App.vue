<template>
  <MLHBanner />
  <NavBar />
  <main class="body-container">
    <RouterView></RouterView>
  </main>
  <TheFooter />
</template>

<script>
import MLHBanner from "./components/MLHBanner.vue";
import TheFooter from "./components/TheFooter.vue";
import Banner from "./components/Banner.vue";
import NavBar from "./components/NavBar.vue";

/*
TODO:
* Mission statement
Tracks:
* Each page for tracks
* Each track
* Climate - solar punk
* Conservation - endangered species, biodiversity, oceanography, trash reef
* Public Health - medical, DNA strand, polymer chain, chemical, data analysis, hospital

TODO:
* Built-in registration form
* Gradient background
* Waiting for graphics
* Standard aspect ratio for images
* Add tracks
* Main site: pull events from sheets

TODO: NAV BAR
* About
  * Student chapter of ACM [link to main website] and YWCC
* Schedule
* Team?
* Contact
* Registration
* FAQ
* Sponsors

TODO: Double check FAQs
TODO: Log-in

Title banner:
* Universal-style vector GIF globe scroll animation, HackNJIT text rotates behind globe and "Hack the world" appears in its place
* Then it scrolls up
* Location / date would stay with the globe
* We need animators

*/

export default {
  name: "HackNJIT",
  components: {
    MLHBanner,
    TheFooter,
    NavBar,
    Banner,
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

  --max-content-width: 1000px;
}

html {
  scroll-behavior: smooth;
  font-family: sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: white;
  box-sizing: border-box;
  background: #274029;
  z-index: -100;
  overflow-x: auto;
  width: 100%;
  height: 100%;
}

body {
  height: 100%;
}

body.modal-open {
  overflow: hidden;
}

* {
  margin: 0;
  padding: 0;
  font-weight: 500;
  line-height: 1.5em;
}

p:not([.vf-registration > p]),
span:not([.vf-registration > p]),
a:not([.vf-registration > p]) {
  font-size: 1.25em;
}

h1:not([.vf-registration > p]),
h2:not([.vf-registration > p]),
h3:not([.vf-registration > p]) {
  font-weight: bold;
}

h1 {
  font-size: 6em;
}

h2 {
  font-size: 3em;
}

h3 {
  font-size: 1.5em;
}

@media (max-width: 1000px) {
  h1 {
    font-size: 4em;
  }

  h2 {
    font-size: 2em;
  }

  h3 {
    font-size: 1.25em;
  }

  p,
  span,
  a,
  button {
    font-size: 1em;
  }
}

button {
  font-size: 1.5em;
  -webkit-tap-highlight-color: transparent;
}

.page-side-padding {
  margin: 0 auto;
  width: calc(100vw - 64px);
}

main {
  display: grid;
  gap: 32px;
}

</style>
