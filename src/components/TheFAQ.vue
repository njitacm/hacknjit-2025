<template>
  <div class="TheFAQ" :style="{ background: gradient }">
    <div class="page-side-padding">
      <h2>Frequently Asked Questions</h2>
      <div class="faq">
        <div class="faq-card" v-for="(faq, index) in faqs" :key="index">
          <h3>{{ faq.Question }}</h3>
          <p>{{ faq.Answer }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import faqs from "../data/faq.js"

export default {
  props: {
    gradient: String
  },
  data() {
    return {
      faqs: faqs
    };
  }
}
</script>

<style scoped>
@keyframes gear-fade-in {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 0.55;
  }
}

#gear5 {
  top: 30%;
  right: 43%;
  width: 47vw;
  max-width: 782px;
}

#gear4 {
  top: 10%;
  left: 63%;
  width: 37vw;
  max-width: 482px;
}

#gear3 {
  bottom: 0%;
  right: -11.5%;
  width: 68vw;
  max-width: 402px;
  transform: rotate(248deg);
}

.floating-gear {
  position: absolute;
  z-index: -100;
  opacity: 0.55;
  animation-name: gear-fade-in;
  animation-duration: 0.75s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}

.outer-div {
  position: relative;
  width: 60%;
  /* border-top: black solid 4px; */
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
  /* background-color: var(--color4); */
  padding: 1.5rem;
}

#section-title {
  font-size: 4rem;
  margin-bottom: 1.75rem;
}

.faq {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}

.faq-card {
  width: 100%;
  margin: 1rem;
  text-align: left;
}

h1 {
  margin-bottom: 0.25rem;
  font-size: 1.5rem;
}

p {
  font-size: 1.125rem;
}

.fade-in {
  animation-name: fade-in;
  animation-duration: 0.5s;
  /* animation-delay: 1.5s; */
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}

@keyframes fade-in {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

@media (max-width: 750px) {
  .outer-div {
    width: 80%;
  }

  #section-title {
    font-size: 2.5rem;
  }
}
</style>