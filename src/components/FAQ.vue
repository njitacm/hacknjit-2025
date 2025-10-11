<template>
  <div class="FAQ section" id="FAQ" ref="sectionRef">
    <h2 class="section-title">FAQ</h2>
    <Accordion class="faq-container">
      <div v-for="(topicFaqs, topic, index1) in faqs" :key="index1" class="faq-topic-container">
        <h3 class="faq-topic">{{ topic }}</h3>
        <AccordionPanel v-for="(faq, index2) in topicFaqs" :key="index2" :value="faq.Answer">
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
.FAQ {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}

.p-accordion {
  --p-accordion-header-background: var(--semi-transparent-white);
  --p-accordion-header-color: var(--hacknjit-primary);
  --p-accordion-content-background: transparent;
  --p-accordion-content-padding: 16px;
  --p-accordion-header-hover-background: white;
  --p-accordion-header-hover-color: var(--hacknjit-primary);
  --p-accordion-header-active-hover-background: var(--hacknjit-primary);
  --p-accordion-header-active-background: var(--hacknjit-primary);
  --p-accordion-header-active-color: white;
  --p-accordion-header-toggle-icon-color: var(--hacknjit-primary);
  --p-accordion-header-toggle-icon-hover-color: var(--hacknjit-primary);

  @media(pointer: coarse) {
    & {
      --p-accordion-header-hover-background: var(--p-accordion-header-background);
    }
  }
}

.p-accordion button {
  border-radius: var(--border-radius) !important;
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
  font-size: 1.1em;
  font-weight: bold;
}

.p-accordion button {
  padding: 16px;
  height: 100%;
}

.p-accordion,
.p-accordion p,
.p-accordion button {
  text-align: left;
  transition: all 0.25s ease-out, background-color 0s, color 0s;
}

.p-accordionheader {
  font-weight: bold;
}

.faq-container .faq-topic-container {
  flex: 45%;
}

@media (max-width: 750px) {
  .faq-container .faq-topic-container {
    flex: 100%;
  }
}
</style>
