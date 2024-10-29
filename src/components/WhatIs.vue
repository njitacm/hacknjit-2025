<template>
  <div class="outer-div">
    <carousel
      ref="myCarousel"
      class="carousel"
      @slide-start="handleSlideStart"
      @slide-end="handleSlideEnd"
      :items-to-show="1"
      :wrapAround="true"
    >
      <slide v-for="slide in slides" :key="slide">
        <img :src="slide" :alt="slide" />
      </slide>

      <template #addons>
        <Navigation :class="{ appear: clickedArrows }" />
        <Pagination :class="{ appear: clickedArrows }" />
      </template>
    </carousel>
    <div class="inner-div">
      <h1>What is HackNJIT? {{ currentSlide }}</h1>
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
import { Carousel, Slide, Navigation, Pagination } from "vue3-carousel";
export default {
  components: { Carousel, Slide, Navigation, Pagination },
  data() {
    return {
      startTime: 0,
      clickedArrows: false,
      slides: [
        require("../assets/Slides/hacknjit_1.jpg"),
        require("../assets/Slides/hacknjit_2.jpg"),
        require("../assets/Slides/hacknjit_3.jpg"),
        require("../assets/Slides/hacknjit_4.jpg"),
      ],
    };
  },
  methods: {
    handleSlideStart() {
      // console.log("Start!");
      this.startTime = new Date();
      this.clickedArrows = true;
    },
    handleSlideEnd() {
      setTimeout(() => {
        this.clickedArrows = false;
      }, 700);
    },
  },
};
</script>

<style scoped>
.outer-div {
  width: 80%;
  margin: 0.5rem auto;
  display: flex;
  flex-wrap: wrap;
}
h1 {
  font-size: 3.5rem;
}
p {
  font-size: 1.75rem;
  text-align: left;
}
.carousel {
  width: 50%;
  flex: 1;
}
.inner-div {
  margin: 25px 50px;
  flex: 1;
}
img {
  width: 100%;
}
@media (max-width: 1400px) {
  h1 {
    font-size: 3rem;
  }
}
@media (max-width: 1300px) {
  p {
    font-size: 1.6rem;
  }
}
@media (max-width: 1200px) {
  h1 {
    font-size: 2.5rem;
    margin-bottom: 5px;
  }
  p {
    font-size: 1.4rem;
  }
}
@media (max-width: 1100px) {
  .outer-div {
    flex-direction: column;
  }
  .carousel {
    width: 100%;
  }
  .inner-div {
    margin: 0;
    margin-bottom: 0.5rem;
  }
}
</style>
<style>
.carousel__icon {
  background-color: rgba(208, 208, 231, 0.175);
  border-radius: 0.5rem;
}
.carousel__pagination {
  margin: 0;
  z-index: 20;
  position: relative;
  background-color: rgba(208, 208, 231, 0.175);
  width: fit-content;
  margin-left: auto;
  margin-right: auto;
  bottom: 25px;
  border-radius: 0.5rem;
}
.appear {
  animation: appear 750ms ease-in-out;
  border-radius: 0.5rem;
}
@keyframes appear {
  0% {
    background-color: rgba(208, 208, 231, 0.175);
  }
  50% {
    background-color: rgba(208, 208, 231, 0.575);
  }
  100% {
    background-color: rgba(208, 208, 231, 0.175);
  }
}
</style>