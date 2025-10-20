<template>
  <div>
    <!-- <Banner gradient="radial-gradient(yellow, green)" /> -->
    <Banner />
    <Tracks/>
    <About />
    <TheSponsors />
    <TheFAQ />
  </div>
</template>

<script>
import NavBar from "../components/NavBar.vue";
import Banner from "../components/Banner.vue";
import About from "../components/About.vue";
import TheSponsors from "../components/Sponsors.vue";
import TheFAQ from "../components/FAQ.vue";
import Tracks from "../components/Tracks.vue";
import Modal from "../components/Modal.vue";

export default {
  name: "HackNJIT",
  components: {
    NavBar,
    Banner,
    About,
    TheSponsors,
    TheFAQ,
    Tracks,
    Modal,
  }
};
</script>

<style scoped>

</style>
