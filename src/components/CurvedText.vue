<template>
  <div class="CurvedText">
    <svg :viewBox="`0 0 ${props.viewBoxSizeX} ${props.viewBoxSizeY}`" :style="`width: ${2 * props.viewBoxSizeX}; margin-left: -${props.viewBoxSizeX}px`" :class="props.svgClass">
      <defs>
        <path :id="pathId" :d="pathDefinition"></path>
      </defs>
      <text class="text" :class="props.textClass">
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
  viewBoxSizeX: { type: Number, default: 400 },
  viewBoxSizeY: { type: Number, default: 400 },
  textClass: { type: String, default: "" },
  svgClass: { type: String, default: "" },
  debug: { type: Boolean, default: false },
  centerX: { type: String, default: "50%" },
  centerY: { type: String, default: "50%" },
  centerOffsetX: { type: Number, default: 0 },
  centerOffsetY: { type: Number, default: 0 },
});

const pathId = `curved-text-path-${Math.random().toString(36).substring(2, 9)}`;

const centerX = computed(() => {
  return props.viewBoxSizeX / 2;
})

const centerY = computed(() => {
  const angleInRad = (props.arc / 2) * (Math.PI / 180);
  const yTop = -props.radius;
  const yBottom = -props.radius * Math.cos(angleInRad);

  const arcMidpoint = (yTop + yBottom) / 2;

  const viewboxCenter = props.viewBoxSizeY / 2;
  return viewboxCenter - arcMidpoint;
});

const pathDefinition = computed(() => {
  const startAngle = -props.arc / 2;
  const endAngle = props.arc / 2;

  const startRad = (startAngle - 90) * (Math.PI / 180);
  const endRad = (endAngle - 90) * (Math.PI / 180);

  const x1 = centerX.value + props.radius * Math.cos(startRad);
  const y1 = centerY.value + props.radius * Math.sin(startRad);
  const x2 = centerX.value + props.radius * Math.cos(endRad);
  const y2 = centerY.value + props.radius * Math.sin(endRad);

  const largeArcFlag = props.arc > 180 ? 1 : 0;

  return `M ${x1 + props.centerOffsetX} ${y1 + props.centerOffsetY} A ${props.radius} ${props.radius} 0 ${largeArcFlag} 1 ${x2 + props.centerOffsetX} ${y2 + props.centerOffsetY}`;
});

const debugPath = computed(() => {
  return pathDefinition.value;
});
</script>

<style scoped>
svg {
  left: 50%;
  transform: translateX(50%);
  width: 100%;
  height: 100%;
  position: absolute;
  transform-origin: center;
  overflow: hidden;
}
</style>