<template>
  <SiteLanding />
  <WhatIs />
  <HackNJITSponsors />
  <HackNJITFAQ />
  <SiteFooter />
  <MLHBanner />
</template>

<script>
import SiteFooter from "./components/SiteFooter.vue";
import SiteLanding from "./components/SiteLanding.vue";
import MLHBanner from "./components/MLHBanner.vue";
import HackNJITSponsors from "./components/HackNJITSponsors.vue";
import HackNJITFAQ from "./components/HackNJITFAQ.vue";
import WhatIs from "./components/WhatIs.vue";
export default {
  name: "HackNJIT",
  components: {
    SiteLanding,
    SiteFooter,
    MLHBanner,
    HackNJITSponsors,
    HackNJITFAQ,
    WhatIs,
  },
};
</script>

<style>
html {
  font-family: sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #d9f0f9;
  box-sizing: border-box;
  font-size: 16px;
}
* {
  margin: 0;
  padding: 0;
  font-weight: 500;
}
</style>
