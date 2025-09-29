<template>
  <main class="body-container">
    <!-- <Banner gradient="radial-gradient(yellow, green)" /> -->
    <Banner />
    <!-- <Tracks/> -->
    <PastPics />
    <!-- <TheSponsors /> -->
    <TheFAQ />
  </main>
</template>

<script>
import TheFAQ from "../components/TheFAQ.vue";
import TheSponsors from "../components/TheSponsors.vue";
import Modal from "../components/Modal.vue";
import Banner from "../components/Banner.vue";
import Tracks from "../components/Tracks.vue";
import PastPics from "../components/PastPics.vue";
import NavBar from "../components/NavBar.vue";

export default {
  name: "HackNJIT",
  components: {
    TheFAQ,
    TheSponsors,
    Modal,
    Banner,
    Tracks,
    PastPics,
    NavBar,
  }
};
</script>

<style scoped>

</style>
