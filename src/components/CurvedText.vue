<template>
  <div class="svg-curved-text-container">
    <svg :viewBox="`0 0 ${VIEWBOX_SIZE} ${VIEWBOX_SIZE}`" style="transform: rotate(-2deg)">
      <defs>
        <path :id="pathId" :d="pathDefinition"></path>
      </defs>

      <text v-for="(char, index) in text" :key="index" class="char-text"
        :style="{ animationDelay: index * 0.08 + 's' }">
        <textPath :href="`#${pathId}`" :startOffset="getLetterOffset(index)">
          {{ char }}
        </textPath>
      </text>
    </svg>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  text: { type: String, default: 'HackNJIT' },
  radius: { type: Number, default: 160 },
  arc: { type: Number, default: 100 },
});

const pathId = `curved-text-path-${Math.random().toString(36).substring(2, 9)}`;

const VIEWBOX_SIZE = 400;
const centerX = VIEWBOX_SIZE / 2;

// --- THE FIX ---
// Convert centerY into a computed property to keep the text arc vertically centered.
const centerY = computed(() => {
  // Calculate the vertical positions of the arc's highest and lowest points
  // relative to the circle's center.
  const angleInRad = (props.arc / 2) * (Math.PI / 180);
  const yTop = -props.radius; // The very top of the circle
  const yBottom = -props.radius * Math.cos(angleInRad); // The lowest point where text appears

  // Find the midpoint of just the text arc
  const arcMidpoint = (yTop + yBottom) / 2;

  // To center the arc, we shift the circle's center (centerY) by the
  // opposite of the arc's own midpoint.
  const viewboxCenter = VIEWBOX_SIZE / 2;
  return viewboxCenter - arcMidpoint;
});

// This path definition now uses the new dynamic `centerY`
const pathDefinition = computed(() => {
  const startAngle = -props.arc / 2 + (props.arc / props.text.length);
  const endAngle = props.arc / 2;

  const startRad = (startAngle - 90) * (Math.PI / 180);
  const endRad = (endAngle - 90) * (Math.PI / 180);

  // Note that we now use .value because centerY is a computed ref
  const x1 = centerX + props.radius * Math.cos(startRad);
  const y1 = centerY.value + props.radius * Math.sin(startRad);
  const x2 = centerX + props.radius * Math.cos(endRad);
  const y2 = centerY.value + props.radius * Math.sin(endRad);

  const largeArcFlag = props.arc > 180 ? 1 : 0;

  return `M ${x1} ${y1} A ${props.radius} ${props.radius} 0 ${largeArcFlag} 1 ${x2} ${y2}`;
});

// getLetterOffset remains the same
const getLetterOffset = (index) => {
  const totalChars = props.text.length;
  const offset = (index / (totalChars - 1)) * 100;
  return `${offset}%`;
};
</script>

<style scoped>
.svg-curved-text-container {
  /* display: flex;
  justify-content: center;
  align-items: center; */
  /* position: relative; */
  /* border-radius: 10px; */
  overflow: hidden;
}

svg {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: visible;
  /* Allows text to be seen even if path is near edge */
}

.char-text {
  font-size: 2.8rem;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  fill: #ffffff;
  /* Use 'fill' instead of 'color' for SVG text */
  text-shadow: 0 0 5px #00aaff, 0 0 10px #00aaff;

  /* Set properties for centering each letter on its offset point */
  text-anchor: middle;

  /* Initial animation state */
  opacity: 0;
  transform: translateY(20px);

  /* Animation definition */
  animation: fadeIn 0.6s cubic-bezier(0.68, -0.55, 0.27, 1.55) forwards;
}

/* The animation is simpler now as we don't need to handle rotation */
@keyframes fadeIn {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>