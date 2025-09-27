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

<script>
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