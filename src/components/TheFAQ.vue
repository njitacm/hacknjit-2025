<template>
  <div class="TheFAQ" ref="sectionRef" id="FAQ">
    <Accordion value="0">
      <AccordionPanel v-for="(faq, index) in faqs" :key="index" :value="index">
        <AccordionHeader>{{ faq.Question }}</AccordionHeader>
        <AccordionContent>
          <p class="m-0">{{ faq.Answer }}</p>
        </AccordionContent>
      </AccordionPanel>
    </Accordion>
  </div>
</template>

<script>
import faqs from "../data/faq.js"
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
  beforeUnmount() {
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
  width: 80%;
}

.p-accordionpanel {
  border-radius: 8px;
  background-color: white;
  margin: 16px;
  border-bottom: 3px solid #8080807F
}


.p-accordioncontent .p-accordioncontent-content p {
  padding: 16px;
}

.p-accordion button {
  padding: 16px;
  height: 100%;
}

.p-accordion p,
.p-accordion button {
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

</style>

<style>
.p-accordion {
  --p-accordion-header-active-color: #40BB4A;
  --p-accordion-header-active-hover-color: #40BB4A;
}
.p-accordionpanel,
.p-accordioncontent .p-accordioncontent-content {
  border: none; 
  border-style: none;
}

.p-accordionpanel-active {
}

.p-accordionpanel-active button {
  color: #17641E;
}

</style>
