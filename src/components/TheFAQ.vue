<template>
  <div class="TheFAQ" ref="sectionRef" id="FAQ">
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

<script>
import faqs from "../data/faq.js"
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();

export default {
  data() {
    return {
      faqs: faqs
    };
  },
  mounted() {
    observe(this.$refs.sectionRef);
  },
  beforeUnmount() {
    unobserve(this.$refs.sectionRef);
  }
}
</script>

<style scoped>
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

.page-side-padding {
  @media(min-width: 900px) {
    & {
      width: calc(100vw - 256px);
    }
  }
}
</style>
