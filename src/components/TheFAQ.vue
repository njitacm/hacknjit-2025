<template>
  <div class="component-container" id="FAQ" ref="sectionRef">
    <h1>FAQ</h1>
    <Accordion class="faq-container">
      <div class="faq-topic-container" v-for="(topic_faqs, topic) in faqs" :key="topic">
        <h2 class="faq-topic">{{ topic }}</h2>
        <AccordionPanel v-for="(faq, index) in topic_faqs" :key="index" :value="faq.Answer">
          <AccordionHeader>{{ faq.Question }}</AccordionHeader>
          <AccordionContent>
            <p class="m-0">{{ faq.Answer }}</p>
          </AccordionContent>
        </AccordionPanel>
      </div>
    </Accordion>
  </div>
</template>

<script>
import faqs from "../data/faq.json"
import Accordion from 'primevue/accordion';
import AccordionPanel from 'primevue/accordionpanel';
import AccordionHeader from 'primevue/accordionheader';
import AccordionContent from 'primevue/accordioncontent';
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
const { observe, unobserve } = useIntersectionObserver();

export default {
  components: {
    Accordion,
    AccordionPanel,
    AccordionHeader,
    AccordionContent
  },
  data() {
    return {
      faqs: faqs
    };
  },
  mounted() {
    observe(this.$refs.sectionRef);
  },
  unmounted() {
    unobserve(this.$refs.sectionRef);
  }
};

</script>

<style scoped>
.TheFAQ {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}

.TheFAQ button {
  color: white;
}

.p-accordion {
  color: #274029;
  width: 100%;
}

.p-accordionpanel {
  border-radius: 8px;
  background-color: white;
  margin: 16px;
  border-bottom: 3px solid #8080807F
}


.p-accordioncontent .p-accordioncontent-content p {
  padding: 16px;
  padding-top: 0px;
}

.p-accordion button {
  padding: 16px;
  height: 100%;
}

.p-accordion,
.p-accordion p,
.p-accordion button {
  border: none;
  color: #274029;
  font-size: 1em;
  text-align: left;
  transition: all 0.25s ease-out;
}

.p-accordionheader {
  font-weight: 850;
  border: none;
  text-wrap: wrap;
}

.component-container {
  margin: 16px auto;
}

.faq-container {
  margin: auto;
  width: calc(100vw - 64px);
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}

.faq-container .faq-topic-container {
  flex: 40%;
}

.p-accordion {
  --p-accordion-header-active-color: #40BB4A;
  --p-accordion-header-active-hover-color: #40BB4A;
}

.p-accordionpanel-active {}

.p-accordionpanel-active button {
  color: #17641E;
}

h2.faq-topic {
  font-size: 3rem;
  color: white;
  text-align: center;
}

@media (max-width: 450px) {
  h1 {
    font-size: 2rem;
    margin-bottom: 16px;
  }

  h2.faq-topic {
    font-size: 1.5rem;
  }

  button.p-accordionheader {
    font-size: 1rem;
  }

  .faq-container .faq-topic-container {
    flex: 100%;
  }

  .faq-container {
    width: 100vw;
  }
}
</style>

<style>
/* I don't know why but this needs to be global to work */
.p-accordioncontent-content,
.p-accordioncontent .p-accordioncontent-content {
  border: none;
  border-style: none;
  border-color: none;
}
</style>
