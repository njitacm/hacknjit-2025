<template>
  <div class="TheFAQ">
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
  components: { 
    Accordion, 
    AccordionPanel, 
    AccordionHeader,
    AccordionContent
  },
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
  width: 80%;
}

.p-accordionpanel,
.p-accordioncontent .p-accordioncontent-content {
  border: none; 
}

.p-accordioncontent .p-accordioncontent-content p {
  padding: 8px 0
}

.p-accordion p,
.p-accordion button {
  font-size: 1em;
  text-align: left;
}


.p-accordionpanel {
  margin: 16px;
  border-bottom: 3px solid #8080807F
}

.p-accordionheader {
  font-weight: 850;
  border: none;
  text-wrap: wrap;
}

</style>
