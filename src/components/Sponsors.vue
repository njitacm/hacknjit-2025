<template>
  <div ref="sectionRef" class="Sponsors component-container page-side-padding" id="Sponsors">
    <h2 class="title">Our Sponsors</h2>
    <div class="container">
      <div v-for="(sponsors, category, index1) in sponsors" :key="index1" class="category-container">
        <h3 v-if="sponsors.length">{{ category.at(0).toUpperCase() + category.substring(1) }}</h3>
        <div class="category-sponsors-container">
          <a v-for="(sponsor, index2) in sponsors" :key="index2" :href="sponsor.link" :title="sponsor.name" class="sponsor-item"
            target="_blank">
            <img :src="sponsor.imgSrc" :alt="sponsor.imgAlt" class="image" />
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import sponsorsData from "../data/sponsors";
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
import { getImageUrl } from "../util";

// convert sponsor imageSrcs to URLs
const sponsors = ref({});

const sectionRef = ref(null);
const { observe, unobserve } = useIntersectionObserver();

onMounted(async () => {
  if (sectionRef.value) observe(sectionRef.value);

  const newSponsors = JSON.parse(JSON.stringify(sponsorsData));

  // companies in each tier
  for (const companies of Object.values(newSponsors)) {
    for (const company of companies) {
      company.imgSrc = await getImageUrl(company.imgSrc);
    }
  }

  sponsors.value = newSponsors;
});
onBeforeUnmount(() => {
  if (sectionRef.value) unobserve(sectionRef.value);
});

</script>

<style scoped>
.category-container {
  display: grid;
  gap: 8px;
  padding: 16px;
}

.category-sponsors-container {
  display: flex;
  justify-content: center;
  align-content: center;
  flex-wrap: wrap;
  gap: 32px;
}

.sponsor-item {
  max-width: 500px;
  flex-basis: 200px;
  flex-grow: 1;
  flex-shrink: 0;
  padding: 16px;
  align-content: center;
  background: #B7EBDA;
  border-radius: 8px;
}

.image {
  max-height: 100%;
  max-width: 100%;
  object-fit: contain;
  filter: brightness(0.95);
  transition: transform 0.2s ease, filter 0.2s ease;
}

@media(hover: hover) and (pointer: fine) {
  .sponsor-item:hover .image {
    transform: scale(1.05);
    filter: brightness(1.2);
  }
}

@media (max-width: 768px) {
  .sponsor-logo {
    max-height: 60px;
  }
}

@media (max-width: 480px) {
  .sponsor-logo {
    max-height: 50px;
  }
}
</style>
