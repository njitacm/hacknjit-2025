<template>
  <div class="deque-carousel" @mouseenter="stopAutoplay" @mouseleave="startAutoplay">
    <div class="carousel-viewport">
      <TransitionGroup tag="div" :name="transitionName" class="carousel-items-container">
        <div v-for="item in visibleItems" :key="item.__carousel_key ?? item" class="carousel-item">
          <slot name="item" :data="item" :index="getItemIndex(item)">
            <div class="default-item">{{ item }}</div>
          </slot>
        </div>
      </TransitionGroup>
    </div>

    <button class="carousel-nav prev" @click="prev" aria-label="Previous slide">
      ‹
    </button>
    <button class="carousel-nav next" @click="next" aria-label="Next slide">
      ›
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';

// --- Props ---
const props = defineProps({
  items: {
    type: Array,
    default: () => [],
    // Add a simple unique key to each item if it's an object,
    // which is critical for TransitionGroup.
    validator: (items) => {
      items.forEach((item, index) => {
        if (typeof item === 'object' && item !== null && !item.__carousel_key) {
          item.__carousel_key = Symbol(`carousel-item-${index}`);
        }
      });
      return true;
    }
  },
  numVisible: { type: Number, default: 3 },
  numScroll: { type: Number, default: 1 },
  autoplayInterval: { type: Number, default: 0 },
  // The 'circular' prop is no longer needed; this design is inherently circular.
});

// --- State ---
const localItems = ref([...props.items]); // A mutable copy of the items
const direction = ref('next'); // Controls animation direction
let autoplayTimer = null;

// --- Computed Properties ---

// The items we actually render in the <TransitionGroup>
const visibleItems = computed(() => {
  return localItems.value.slice(0, props.numVisible);
});

// The name of the CSS transition to use
const transitionName = computed(() => {
  return direction.value === 'next' ? 'slide-next' : 'slide-prev';
});

// Helper to get the original index for the slot
// Creates a map for quick lookups
const originalIndexMap = computed(() => {
  const map = new Map();
  props.items.forEach((item, index) => {
    map.set(item, index);
  });
  return map;
});

const getItemIndex = (item) => {
  return originalIndexMap.value.get(item);
};

// --- Navigation Functions ---

const next = () => {
  direction.value = 'next';
  // 1. Take 'numScroll' items from the beginning of the array
  const itemsToMove = localItems.value.splice(0, props.numScroll);
  // 2. Push them onto the end
  localItems.value.push(...itemsToMove);
};

const prev = () => {
  direction.value = 'prev';
  // 1. Take 'numScroll' items from the end of the array
  const itemsToMove = localItems.value.splice(-props.numScroll);
  // 2. Unshift them onto the beginning
  localItems.value.unshift(...itemsToMove);
};

// --- Autoplay (Unchanged) ---
const startAutoplay = () => {
  if (props.autoplayInterval <= 0 || autoplayTimer) return;
  autoplayTimer = setInterval(next, props.autoplayInterval);
};

const stopAutoplay = () => {
  clearInterval(autoplayTimer);
  autoplayTimer = null;
};

// --- Lifecycle (Unchanged) ---
onMounted(startAutoplay);
onUnmounted(stopAutoplay);

// --- CSS v-bind ---
const itemWidthCSS = computed(() => `${100 / props.numVisible}%`);

</script>

<style scoped>
.deque-carousel {
  position: relative;
  width: 100%;
}

.carousel-viewport {
  overflow: hidden;
  width: 100%;
}

.carousel-items-container {
  display: flex;
  position: relative;
  /* Required for TransitionGroup animations */
}

/* This styles the items passed in via the slot */
.carousel-item {
  flex: 0 0 v-bind(itemWidthCSS);
  width: v-bind(itemWidthCSS);
  padding: 0 8px;
  /* Gutters */
  box-sizing: border-box;
  min-width: 0;
}

/* --- TransitionGroup Animations --- */

/* This is the magic.
  'move' handles the items that are *staying* in view.
  'enter'/'leave' handles items *coming in* or *going out*.
*/
.slide-next-move,
.slide-next-enter-active,
.slide-next-leave-active,
.slide-prev-move,
.slide-prev-enter-active,
.slide-prev-leave-active {
  /* Use the same elegant animation! */
  transition: transform 0.6s cubic-bezier(0.23, 1, 0.32, 1);
}

/* 'leave-active' is crucial. It makes the item
  get taken out of the layout flow so the
  new item can slide into its place.
*/
.slide-next-leave-active,
.slide-prev-leave-active {
  position: absolute;
  /* Ensure leaving items are behind entering ones */
  z-index: 0;
}

/*
  'enter-active' items should be on top
*/
.slide-next-enter-active,
.slide-prev-enter-active {
  z-index: 1;
}

/* NEXT (sliding left) */
.slide-next-enter-from {
  transform: translateX(100%);
}

.slide-next-leave-to {
  transform: translateX(-100%);
}

/* PREV (sliding right) */
.slide-prev-enter-from {
  transform: translateX(-100%);
}

.slide-prev-leave-to {
  transform: translateX(100%);
}


/* --- Fallback & Nav (Unchanged) --- */

.default-item {
  height: 100px;
  background: #f4f4f4;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.carousel-nav {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background-color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 50%;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  width: 40px;
  height: 40px;
  font-size: 24px;
  font-weight: bold;
  color: #333;
  cursor: pointer;
  transition: all 0.2s ease;
  z-index: 10;
}

.carousel-nav:hover {
  background-color: #fff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  transform: translateY(-50%) scale(1.05);
}

.carousel-nav.prev {
  left: 12px;
}

.carousel-nav.next {
  right: 12px;
}
</style>