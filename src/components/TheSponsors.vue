<template>
  <div ref="sectionRef" class="component-container" id="Sponsors">
    <h1 class="title">Our Sponsors</h1>
    <div class="sponsors-container">
      <div class="sponsor-category">
        <div class="sponsor-item" v-for="(sponsor, index) in sponsors.bronze" :key="index">
          <a :href="sponsor.link" :title="sponsor.name" target="_blank">
            <img :src="sponsor.imgSrc" :alt="sponsor.imgAlt" class="sponsor-image" />
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

const showSponsors = ref(true);

function toggleSponsors() {
  this.showSponsors.value = !this.showSponsors.value;
}
</script>

<style scoped>
.sponsor-category {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  gap: 4rem;
}

div.sponsor-item {
  height: 100%;
  width: 40%;
  padding: 16px;
  height: 320px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #B7EBDA;
  border-radius: 8px;
}

.sponsor-image {
  width: 100%;
  object-fit: contain;
  filter: brightness(0.95);
  transition: transform 0.2s ease, filter 0.2s ease;
}

.sponsor-image:hover {
  transform: scale(1.05);
  filter: brightness(1.2);
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
