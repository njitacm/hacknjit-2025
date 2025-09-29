<template>
  <div class="TheFAQ">
    <Accordion value="0" >
     <AccordionPanel v-for="faq in faqs" :key="faq.Question">
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
