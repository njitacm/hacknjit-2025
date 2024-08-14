<template>
  <div class="outer-div">
    <carousel class="carousel" :items-to-show="1" :wrapAround="true">
      <slide v-for="slide in slides" :key="slide">
        <img :src="slide" :alt="slide" />
      </slide>

      <template #addons>
        <navigation />
      </template>
    </carousel>
    <div class="inner-div">
      <h1>What is HackNJIT?</h1>
      <p>
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing. Students break into small teams of up to 4,
        work together over 24 hours to create a tech project, and at the end
        present it to our judges!
      </p>
    </div>
  </div>
</template>

<script>
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide, Navigation } from "vue3-carousel";
export default {
  components: { Carousel, Slide, Navigation },
  data() {
    return {
      slides: [
        require("../assets/Slides/hacknjit_1.jpg"),
        require("../assets/Slides/hacknjit_2.jpg"),
        require("../assets/Slides/hacknjit_3.jpg"),
        require("../assets/Slides/hacknjit_4.jpg"),
      ],
    };
  },
};
</script>

<style scoped>
.outer-div {
  width: 80%;
  margin: 0.5rem auto;
  display: flex;
}
.inner-div {
  margin: 25px 50px;
}
h1 {
  font-size: 3.5rem;
}
p {
  font-size: 1.75rem;
  text-align: left;
}
.carousel {
  width: 750px;
}
img {
  width: 750px;
}
</style>
<style>
.carousel__icon {
  background-color: rgba(208, 208, 231, 0.575);
  border-radius: 8px;
}
</style>