<template>
  <div class="TheFAQ">
    <div class="page-side-padding">
      <h1>Frequently Asked Questions</h1>
      <div class="faq">
        <div class="faq-card" v-for="(faq, index) in faqs" :key="index">
          <h1>{{ faq.Question }}</h1>
          <p>{{ faq.Answer }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import faqs from "../data/faq.js"

export default {
  data() {
    return {
      faqs: faqs
    };
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