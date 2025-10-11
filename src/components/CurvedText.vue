<template>
  <div class="curved-text-container">
    <div class="curved-text-wrapper">
      <div v-for="(char, index) in text" :key="index" class="char-container" :style="getCharContainerStyle(index)">
        <span class="char" :style="{ animationDelay: index * 0.08 + 's' }">
          {{ char }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';

// Define props to make the component reusable and easily adjustable
const props = defineProps({
  // The text to display along the curve
  text: {
    type: String,
    default: 'HackNJIT'
  },
  // The radius of the circular path in pixels
  radius: {
    type: Number,
    default: 160
  },
  // The total angle (in degrees) that the text should span
  arc: {
    type: Number,
    default: 100
  },
});

/**
 * Calculates the rotational style for each character's container.
 * This function places each character along the defined arc.
 * @param {number} index - The index of the character in the text string.
 */
const getCharContainerStyle = (index) => {
  const totalChars = props.text.length;
  // Calculate the starting angle to center the text
  const startAngle = -props.arc / 2;
  // Calculate the angle for each character, ensuring it's centered
  const anglePerChar = totalChars > 1 ? props.arc / (totalChars - 1) : 0;
  const rotation = startAngle + anglePerChar * index;

  return {
    // Each container is rotated around the common origin point below
    transform: `rotate(${rotation}deg)`,
    // The transform-origin is set to the center of the circle, based on the radius
    transformOrigin: `0 ${props.radius}px`,
  };
};
</script>

<style scoped>
.curved-text-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
  min-width: 400px;
  position: relative;
  /* A dark background to make the white text pop */
  background-color: #0d1117;
  border-radius: 10px;
  overflow: hidden;
}

.center-image {
  width: 240px;
  height: 240px;
  border-radius: 50%;
  position: absolute;
  object-fit: cover;
}

.curved-text-wrapper {
  position: absolute;
  /* Center the wrapper and pull it up by the radius to align the rotation axis */
  transform: translateY(-50%);
  top: 50%;
  left: 50%;
}

.char-container {
  position: absolute;
  top: 0;
  left: 0;
  height: var(--radius-px, 160px);
  /* Match radius prop */
  width: 1ch;
  /* Give it a small width to align characters */
}

.char {
  display: inline-block;
  font-size: 2.8rem;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  color: #ffffff;
  text-shadow: 0 0 5px #00aaff, 0 0 10px #00aaff;

  /* Initial state for animation */
  opacity: 0;

  /* Animation definition */
  animation: fadeInRotate 0.6s cubic-bezier(0.68, -0.55, 0.27, 1.55) forwards;
}

/**
 * Keyframes for the letter-by-letter animation.
 * Each letter rotates and fades into its final position.
 */
@keyframes fadeInRotate {
  0% {
    opacity: 0;
    transform: rotate(-25deg) scale(0.8);
  }

  100% {
    opacity: 1;
    transform: rotate(0deg) scale(1);
  }
}
</style>