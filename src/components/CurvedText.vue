<template>
  <div class="svg-curved-text-container">
    <svg :viewBox="`0 0 ${viewBoxSize} ${viewBoxSize}`">
      <defs>
        <path :id="pathId" :d="pathDefinition"></path>
      </defs>

      <text 
        v-for="(char, index) in text"
        :key="index"
        class="char-text"
        :style="{ animationDelay: index * 0.08 + 's' }"
      >
        <textPath :href="`#${pathId}`" :startOffset="getLetterOffset(index)">
          {{ char }}
        </textPath>
      </text>
    </svg>
  </div>
</template>

<script setup>
import { computed } from 'vue';

// Define props to make the component reusable and easily adjustable
const props = defineProps({
  text: { type: String, default: 'HackNJIT' },
  radius: { type: Number, default: 160 },
  arc: { type: Number, default: 100 },
});

// A unique ID for the path so the textPath can reference it
const pathId = `curved-text-path-${Math.random().toString(36).substring(2, 9)}`;

// Set the SVG canvas size based on the radius to ensure everything fits
const viewBoxSize = computed(() => props.radius * 2.2);
const centerX = computed(() => viewBoxSize.value / 2);
const centerY = computed(() => viewBoxSize.value / 2);

// Dynamically generate the SVG path 'd' attribute string
const pathDefinition = computed(() => {
  const startAngle = -props.arc / 2;
  const endAngle = props.arc / 2;

  // Convert angles to radians for trigonometry
  const startRad = (startAngle - 90) * (Math.PI / 180);
  const endRad = (endAngle - 90) * (Math.PI / 180);

  // Calculate start and end points of the arc
  const x1 = centerX.value + props.radius * Math.cos(startRad);
  const y1 = centerY.value + props.radius * Math.sin(startRad);
  const x2 = centerX.value + props.radius * Math.cos(endRad);
  const y2 = centerY.value + props.radius * Math.sin(endRad);

  // Large arc flag is 1 if the arc is > 180 degrees, 0 otherwise
  const largeArcFlag = props.arc > 180 ? 1 : 0;

  // M = Move to, A = Arc
  return `M ${x1} ${y1} A ${props.radius} ${props.radius} 0 ${largeArcFlag} 1 ${x2} ${y2}`;
});

// Calculate the percentage offset for each letter along the path
const getLetterOffset = (index) => {
  const totalChars = props.text.length;
  // Spread letters evenly across 100% of the path
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
  overflow: visible; /* Allows text to be seen even if path is near edge */
}

.char-text {
  font-size: 2.8rem;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  fill: #ffffff; /* Use 'fill' instead of 'color' for SVG text */
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