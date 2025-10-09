<template>
  <div class="TheFAQ section" id="FAQ" ref="sectionRef">
    <h2 class="section-title">FAQ</h2>
    <Accordion class="faq-container">
      <div v-for="(topicFaqs, topic, index1) in faqs" :key="index1" class="faq-topic-container">
        <h3 class="faq-topic">{{ topic }}</h3>
        <AccordionPanel v-for="(faq, index2) in topicFaqs" :key="index2" :value="index1 * topicFaqs.length + index2">
          <AccordionHeader>{{ faq.Question }}</AccordionHeader>
          <AccordionContent>
            <p class="m-0">{{ faq.Answer }}</p>
          </AccordionContent>
        </AccordionPanel>
      </div>
    </Accordion>
  </div>
</template>


<script setup>
import { onMounted, onBeforeUnmount, useTemplateRef } from 'vue';
import faqs from '../data/faq.json';
import Accordion from 'primevue/accordion';
import AccordionPanel from 'primevue/accordionpanel';
import AccordionHeader from 'primevue/accordionheader';
import AccordionContent from 'primevue/accordioncontent';
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();
const sectionRef = useTemplateRef("sectionRef");

// life cycle hooks
onMounted(() => {
  if (sectionRef.value) observe(sectionRef.value);
});

onBeforeUnmount(() => {
  if (sectionRef.value) unobserve(sectionRef.value);
});

</script>

<style scoped>
.TheFAQ {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}

.p-accordion,
.p-accordion-panel,
.p-accordion button {
  border-radius: var(--border-radius);
}

.p-accordionpanel-active button {
  border-radius: var(--border-radius) var(--border-radius) 0 0;
}

.p-accordionpanel-active .p-accordioncontent {
  border-radius: 0 0 var(--border-radius) var(--border-radius);
  overflow: hidden;
}

.p-accordion {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}

.p-accordionpanel {
  margin: 8px 0;
  border: none;
}

.p-accordioncontent-content p {
  font-size: 1em;
}

.p-accordion button {
  color: white;
  padding: 16px;
  height: 100%;
}

.p-accordion,
.p-accordion p,
.p-accordion button {
  text-align: left;
  transition: all 0.25s ease-out;
}

.p-accordionheader {
  font-weight: bold;
}

.faq-container .faq-topic-container {
  flex: 45%;
}

.p-accordion {
  --p-accordion-header-active-color: #40BB4A;
  --p-accordion-header-active-hover-color: #40BB4A;
}

.p-accordionpanel-active button {
  color: #17641E;
}

@media (max-width: 750px) {
  .faq-container .faq-topic-container {
    flex: 100%;
  }
}
</style>
