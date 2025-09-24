<template>
  <div ref="container" class="outer-container">
    <h1 class="title">Our Sponsors</h1>
    <div class="sponsors">
      <div v-for="(sponsor, index) in sponsors" :key="index" class="sponsor-grid">
        <a :href="sponsor.link" :title="sponsor.name" target="_blank">
          <img :src="sponsor.src" :alt="sponsor.alt" />
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import sponsors from "../data/sponsors";
export default {
  data() {
    return {
      showSponsors: true,
      sponsors: sponsors,
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
  grid-template-columns: repeat(6, 16.67%);
  grid-template-rows: repeat(3, 300px);
  flex-wrap: wrap;
  align-content: center;
  justify-content: space-around;
  row-gap: 0.5rem;
  padding: 1rem;
}

.sponsor-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  align-self: center;
  justify-self: center;
}

div.outer-container {
  width: 100%;
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
}

h1.title {
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

@media (max-width: 1500px) {
  .sponsor-grid {
    grid-template-rows: repeat(3, 200px);
  }
}

@media (max-width: 1000px) {
  .sponsor-grid {
    grid-template-columns: 100%;
    grid-template-rows: repeat(6, auto);
  }

  .sponsor-img {
    max-width: 75% !important;
  }
}

@media (max-width: 550px) {
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
  }
}
</style>