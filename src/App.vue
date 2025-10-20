<template>
  <main class="body-container">
    <NavBar />
    <MLHBanner />
    <div class="gradient"></div>
    <RouterView></RouterView>
    <TheFooter />
  </main>
</template>

<script>
import MLHBanner from "./components/MLHBanner.vue";
import TheFooter from "./components/TheFooter.vue";
import Banner from "./components/Banner.vue";
import NavBar from "./components/NavBar.vue";

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
  --hacknjit-primary-transparent: #4a3728e5;
  --hacknjit-primary: #4a3728;
  --hacknjit-secondary: #8B6B4F;
  --hacknjit-tertiary: #57cc99;
  --hacknjit-fourth: #5BA36A;
  --hacknjit-fifth: #3990b5;
  --hacknjit-sixth: #307c7a;

  --main-bg-color: #274029;
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
  --semi-transparent-white: #ffffffaa;

  --border-radius: 10px;
  --max-content-width: 1000px;
}

html {
  scroll-behavior: smooth;
  font-family: sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: white;
  box-sizing: border-box;
  z-index: -100;
  overflow-x: hidden;
  width: 100%;
}

@keyframes gradient-fade-in {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

body.modal-open {
  overflow: hidden;
}

.body-container {
  min-height: 100svh;
  position: relative;
  display: grid;
  grid-template-rows: 1fr auto;
}

.gradient {
  width: 100%;
  height: 100%;
  background-color: black;
  position: absolute;
  z-index: -9999;
}

.gradient::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: linear-gradient(to bottom, black 50svh, var(--hacknjit-sixth),
      var(--hacknjit-fifth),
      var(--hacknjit-fourth),
      var(--hacknjit-tertiary),
      var(--hacknjit-secondary),
      var(--hacknjit-primary));
  animation: gradient-fade-in 2.5s linear forwards;
}

* {
  margin: 0;
  padding: 0;
  font-weight: 500;
  line-height: 1.5em;
  -webkit-tap-highlight-color: transparent;
}

p,
span,
a {
  font-size: 1.25em;
}

h1,
h2,
h3 {
  font-weight: bold;
  text-align: center;
}

h1 {
  font-size: 6em;
}

h2 {
  font-size: 4em;
}

h3 {
  font-size: 2.5em;
}

.View {
  position: relative;
  top: 80px;
  margin-bottom: 80px;
}

@media (max-width: 1000px) {
  h1 {
    font-size: 4em;
  }

  h2 {
    font-size: 3em;
  }

  h3 {
    font-size: 2em;
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

.section {
  margin: 32px auto;
  width: calc(100vw - 64px);
  max-width: var(--max-content-width);
}

.section-title {
  margin-bottom: 16px;
}

a.pill,
.router-link {
  color: white;
  text-decoration: none;
}

.pill {
  --bkg-hov-act: radial-gradient(circle at top, var(--hacknjit-tertiary), transparent),
    conic-gradient(var(--hacknjit-fourth),
      var(--hacknjit-fifth) 135deg 225deg,
      var(--hacknjit-fourth));
  --bkg: conic-gradient(var(--hacknjit-fourth),
      var(--hacknjit-fifth) 135deg 225deg,
      var(--hacknjit-fourth));
  height: fit-content;
  border-radius: 1000px;
  border: none;
  padding: 8px 24px;
  cursor: pointer;
  background: var(--bkg);
  font-size: 1em;
  transition: transform 250ms ease, filter 500ms ease;
}

.pill,
.pill span {
  font-weight: bold;
}

@media(hover: hover) and (pointer: fine) {
  .pill:hover {
    background: var(--bkg-hov-act);
    filter: drop-shadow(0px 10px 10px #00000033);
    transform: scale(1.05);
  }
}

@media(pointer: coarse) {
  .pill:active {
    transform: scale(0.95);
    background: var(--bkg-hov-act);
  }
}
</style>
