<template>
  <div class="elegant-carousel" @mouseenter="stopAutoplay" @mouseleave="startAutoplay">
    <div class="carousel-viewport">
      <div class="carousel-items-container">
        <slot></slot>
      </div>
    </div>

    <button class="carousel-nav prev" @click="prev" :disabled="isPrevDisabled" aria-label="Previous slide">
      ‹
    </button>
    <button class="carousel-nav next" @click="next" :disabled="isNextDisabled" aria-label="Next slide">
      ›
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, useSlots } from 'vue';

// --- Props ---

const props = defineProps({
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

// The index of the first visible item
const currentIndex = ref(0);
let autoplayTimer = null;

// --- Slot & Item Counting ---

const slots = useSlots();

// Computes the total number of items passed into the slot.
// This handles items generated with v-for (which creates a Fragment).
const totalItems = computed(() => {
  if (!slots.default) return 0;
  const nodes = slots.default();
  let len = 0;
  console.log(nodes);

  // Handle v-for fragment
  for (const node of nodes) {
    if (node.type.toString() === 'Symbol(v-fgt)') {
      // Filter out comments from fragment children
      len += node.children.filter(child => child.type.toString() !== 'Symbol(v-cmt)').length;
    } else if (node.type.toString() !== 'Symbol(v-cmt)') {
      len += 1;
    }
  }

  console.log(len);
  return len;
});

// --- Computed Properties for Logic & Styling ---

// The width of a single item as a percentage
const itemWidthPercent = computed(() => 100 / props.numVisible);

// The CSS value for the transform
const translateX = computed(() => {
  // We move the container left by the width of `currentIndex` items
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

const next = () => {
  if (!canScroll.value) return;

  const nextIndex = currentIndex.value + props.numScroll;

  if (nextIndex > maxIndex.value) {
    // We are at or past the end
    if (props.circular) {
      currentIndex.value = 0; // Wrap to start
    } else {
      currentIndex.value = maxIndex.value; // Clamp to end
    }
  } else {
    // We can scroll normally
    currentIndex.value = nextIndex;
  }
};

const prev = () => {
  if (!canScroll.value) return;

  const prevIndex = currentIndex.value - props.numScroll;

  if (prevIndex < 0) {
    // We are at or past the beginning
    if (props.circular) {
      currentIndex.value = maxIndex.value; // Wrap to end
    } else {
      currentIndex.value = 0; // Clamp to start
    }
  } else {
    // We can scroll normally
    currentIndex.value = prevIndex;
  }
};

// --- Autoplay ---

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
// Create CSS-usable strings from computed properties
const itemWidthCSS = computed(() => `${itemWidthPercent.value}%`);
const translateXCSS = computed(() => `${translateX.value}%`);

</script>

<style scoped>
.elegant-carousel {
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

  /* THE ANIMATION:
    We use a transform for hardware-accelerated sliding.
    The cubic-bezier(0.23, 1, 0.32, 1) is an "easeOutQuint" curve.
    It starts fast and decelerates smoothly, feeling very responsive.
  */
  transform: translateX(v-bind(translateXCSS));
  transition: transform 0.6s cubic-bezier(0.23, 1, 0.32, 1);
}

/* ::v-slotted is a "deep" selector that can style content 
  passed into <slot> from the parent component.
*/
:slotted(*) {
  /* Use v-bind to dynamically set the width based on numVisible */
  flex: 0 0 v-bind(itemWidthCSS);
  width: v-bind(itemWidthCSS);

  /* Add some padding for gutters */
  padding: 0 8px;
  box-sizing: border-box;

  /* Ensure children don't shrink or break layout */
  min-width: 0;
}

.carousel-nav {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);

  /* Styling */
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

.carousel-nav:hover:not(:disabled) {
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

.carousel-nav:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}
</style>