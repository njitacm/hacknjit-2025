<template>
  <div class="svg-curved-text-container">
    <!-- <svg :viewBox="`0 0 ${VIEWBOX_SIZE} ${VIEWBOX_SIZE}`" style="transform: rotate(-2deg)"> -->
    <svg :viewBox="`0 0 ${VIEWBOX_SIZE} ${VIEWBOX_SIZE}`">
      <defs>
        <path :id="pathId" :d="pathDefinition"></path>
      </defs>
      <text class="char-text">
        <textPath :href="`#${pathId}`" startOffset="50%" text-anchor="middle">
          {{ props.text }}
        </textPath>
      </text>
      <path v-if="props.debug" id="debug-path" :d="debugPath" stroke="red" stroke-width="1" fill="none"></path>
      <circle v-if="props.debug" id="debug-circle" :cx="props.centerX" :cy="props.centerY" :r="props.radius"
        z-index="1000" fill="none" stroke="blue" stroke-width="1"></circle>
    </svg>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  text: { type: String, default: 'HackNJIT' },
  radius: { type: Number, default: 160 },
  arc: { type: Number, default: 100 },
  rotation: { type: Number, default: -2 },
  debug: { type: Boolean, default: false },
  centerX: { type: String, default: "50%" },
  centerY: { type: String, default: "50%" },
  centerOffsetX: { type: Number, default: 0 },
  centerOffsetY: { type: Number, default: 0 },
});

const pathId = `curved-text-path-${Math.random().toString(36).substring(2, 9)}`;

const VIEWBOX_SIZE = 400;
const centerX = VIEWBOX_SIZE / 2;

const centerY = computed(() => {
  const angleInRad = (props.arc / 2) * (Math.PI / 180);
  const yTop = -props.radius; // The very top of the circle
  const yBottom = -props.radius * Math.cos(angleInRad); // The lowest point where text appears

  const arcMidpoint = (yTop + yBottom) / 2;

  const viewboxCenter = 0;
  return viewboxCenter - arcMidpoint;
});

const pathDefinition = computed(() => {
  const startAngle = -props.arc / 2;
  const endAngle = props.arc / 2;

  const startRad = (startAngle - 90) * (Math.PI / 180);
  const endRad = (endAngle - 90) * (Math.PI / 180);

  // Note that we now use .value because centerY is a computed ref
  const x1 = centerX + props.radius * Math.cos(startRad);
  const y1 = centerY.value + props.radius * Math.sin(startRad);
  const x2 = centerX + props.radius * Math.cos(endRad);
  const y2 = centerY.value + props.radius * Math.sin(endRad);

  const largeArcFlag = props.arc > 180 ? 1 : 0;
  // const largeArcFlag = 0;

  console.log(props.centerOffsetY);
  //d="M 336 6 A 150 150 20 0 0 85 17" 
  return `M ${x1 + props.centerOffsetX} ${y1 + props.centerOffsetY} A ${props.radius} ${props.radius} 0 ${largeArcFlag} 1 ${x2 + props.centerOffsetX} ${y2 + props.centerOffsetY}`;
});

const debugPath = computed(() => {
  return pathDefinition.value;
});

const getLetterOffset = (index) => {
  const totalChars = props.text.length;
  const offset = (index / (totalChars - 1)) * 100;
  return `${offset}%`;
};
</script>

<style scoped>
svg {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: visible;
  animation: rotate-in 1s ease forwards;
}

.char-text {
  transform-origin: 50% 100%;
  font-size: var(--title-font-size);
  font-weight: bold;
  fill: #ffffff;
  font-family: monospace;
  text-shadow: 0 0 5px #00aaff, 0 0 10px #00aaff;
  /* text-anchor: middle; */
  opacity: 0;
  animation: fade-in 2s linear forwards;
}

@keyframes rotate-in {
  from {
    transform: rotate(-10deg);
  }

  to {
    transform: none;
  }
}

@keyframes fade-in {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}
</style>