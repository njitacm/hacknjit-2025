<template>
  <div 
    class="prime-carousel"
    @mouseenter="stopAutoplay"
    @mouseleave="startAutoplay"
  >
    <div class="carousel-viewport">
      <div class="carousel-items-container">
        <div 
          v-for="(item, index) in items" 
          :key="index" 
          class="carousel-item"
          role="group"
          :aria-label="`Item ${index + 1} of ${totalItems}`"
        >
          <slot name="item" :data="item" :index="index">
            <div classclass="default-item">
              {{ item }}
            </div>
          </slot>
        </div>
      </div>
    </div>

    <button 
      class="carousel-nav prev" 
      @click="prev" 
      :disabled="isPrevDisabled"
      aria-label="Previous slide"
    >
      ‹
    </button>
    <button 
      class="carousel-nav next" 
      @click="next" 
      :disabled="isNextDisabled"
      aria-label="Next slide"
    >
      ›
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';

// --- Props ---

const props = defineProps({
  // The array of data to display
  items: {
    type: Array,
    default: () => [],
  },
  // Number of items to show at once
  numVisible: {
    type: Number,
    default: 3,
  },
  // Number of items to scroll by
  numScroll: {
    type: Number,
    default: 1,
  },
  // Autoplay interval in milliseconds. 0 to disable.
  autoplayInterval: {
    type: Number,
    default: 0,
  },
  // Wrap around to the beginning/end
  circular: {
    type: Boolean,
    default: false,
  },
});

// --- State ---

const currentIndex = ref(0); // The index of the first visible item
let autoplayTimer = null;

// --- Computed Properties for Logic & Styling ---

// Total items is now much simpler!
const totalItems = computed(() => props.items.length);

// The width of a single item as a percentage
const itemWidthPercent = computed(() => 100 / props.numVisible);

// The CSS value for the transform
const translateX = computed(() => {
  return -currentIndex.value * itemWidthPercent.value;
});

// The maximum index we can scroll to
const maxIndex = computed(() => totalItems.value - props.numVisible);

// --- Button Disabled State ---

const canScroll = computed(() => totalItems.value > props.numVisible);

const isPrevDisabled = computed(() => {
  return !canScroll.value || (!props.circular && currentIndex.value === 0);
});

const isNextDisabled = computed(() => {
  return !canScroll.value || (!props.circular && currentIndex.value >= maxIndex.value);
});

// --- Navigation Functions ---
// This logic remains identical to the previous version.

const next = () => {
  if (!canScroll.value) return;
  const nextIndex = currentIndex.value + props.numScroll;
  
  if (nextIndex > maxIndex.value) {
    if (props.circular) {
      currentIndex.value = 0;
    } else {
      currentIndex.value = maxIndex.value;
    }
  } else {
    currentIndex.value = nextIndex;
  }
};

const prev = () => {
  if (!canScroll.value) return;
  const prevIndex = currentIndex.value - props.numScroll;

  if (prevIndex < 0) {
    if (props.circular) {
      currentIndex.value = maxIndex.value;
    } else {
      currentIndex.value = 0;
    }
  } else {
    currentIndex.value = prevIndex;
  }
};

// --- Autoplay ---
// This logic also remains identical.

const startAutoplay = () => {
  if (props.autoplayInterval <= 0 || autoplayTimer) return;
  autoplayTimer = setInterval(next, props.autoplayInterval);
};

const stopAutoplay = () => {
  clearInterval(autoplayTimer);
  autoplayTimer = null;
};

// --- Lifecycle Hooks ---

onMounted(startAutoplay);
onUnmounted(stopAutoplay);

// --- CSS v-bind ---
const itemWidthCSS = computed(() => `${itemWidthPercent.value}%`);
const translateXCSS = computed(() => `${translateX.value}%`);

</script>

<style scoped>
.prime-carousel {
  position: relative;
  width: 100%;
  box-sizing: border-box;
}

.carousel-viewport {
  overflow: hidden;
  width: 100%;
}

.carousel-items-container {
  display: flex;
  /* The same elegant animation */
  transform: translateX(v-bind(translateXCSS));
  transition: transform 0.6s cubic-bezier(0.23, 1, 0.32, 1);
}

/* We no longer use :slotted, as we control the item wrapper.
  We style our .carousel-item class directly.
*/
.carousel-item {
  flex: 0 0 v-bind(itemWidthCSS);
  width: v-bind(itemWidthCSS);
  padding: 0 8px; /* Gutters */
  box-sizing: border-box;
  min-width: 0; /* Prevents flex shrinking issues */
}

/* Fallback item styling */
.default-item {
  height: 100px;
  background: #f4f4f4;
  border: 1px solid #ddd;
  border-radius: 4px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #888;
}

/* Navigation buttons (unchanged) */
.carousel-nav {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background-color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 50%;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  width: 40px;
  height: 40px;
  font-size: 24px;
  font-weight: bold;
  color: #333;
  cursor: pointer;
  transition: all 0.2s ease;
  z-index: 10;
}

.carousel-nav:hover:not(:disabled) {
  background-color: #fff;
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
  transform: translateY(-50%) scale(1.05);
}

.carousel-nav.prev { left: 12px; }
.carousel-nav.next { right: 12px; }
.carousel-nav:disabled { opacity: 0.3; cursor: not-allowed; }
</style>