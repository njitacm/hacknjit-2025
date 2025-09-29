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
p,
span,
a {
  font-size: 1.25em;
}

h1,
h2,
h3 {
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


</style>
