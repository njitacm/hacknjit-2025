<template>
  <div class="curved-text-wrapper">
    <div v-for="(char, index) in text" :key="index" class="char-container" :style="getCharContainerStyle(index)">
      <span class="char" :style="{ animationDelay: index * 0.08 + 's' }">
        {{ char }}
      </span>
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
  // Calculate the angle for each character, ensuring it's centered
  const anglePerChar = totalChars > 1 ? props.arc / (totalChars - 1) : 0;
  const startAngle = -anglePerChar * (totalChars / 2);
  // const x = props.radius * Math.cos(rotation * Math.PI/180 - Math.PI/2);
  // const y = props.radius * Math.sin(rotation * Math.PI/180 - Math.PI/2);
  
  // return {
    //   // Each container is rotated around the common origin point below
  //   transform: `rotate(${rotation}deg) translateX(${x}px) translateY(${y}px)`,
  //   // The transform-origin is set to the center of the circle, based on the radius
  //   // transformOrigin: `0 calc(100% + ${props.radius}px)`,
  // };
  
  // const rotationForChar = (props.arc / -2) + (anglePerChar * index);
  const rotation = startAngle + anglePerChar * index;
  
  // To position the text at the top of the circle, we must apply a -90 degree offset.
  // This makes the center of our arc (0°) align with the top of the circle (-90°).
  const angleOnCircle = rotation - 90;
  
  // Convert this final angle to radians for the trig functions.
  const angleInRad = angleOnCircle * (Math.PI / 180);

  // Calculate the X and Y coordinates using the corrected angle.
  const x = props.radius * Math.cos(angleInRad);
  const y = props.radius * Math.sin(angleInRad);

  return {
    // The order of CSS transforms is crucial:
    // 1. First, move (translate) the letter to its (x, y) position on the circle.
    // 2. Then, rotate the letter itself (using the original relative angle) to match the curve.
    transform: `translateX(${x}px) translateY(${y}px) rotate(${rotation}deg)`,
  };
};
</script>

<style scoped>
.curved-text-wrapper {
  /* display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
  min-width: 400px;
  position: absolute;
  border-radius: 10px;
  transform: translateY(-50%);
  top: calc(250px + var(--top-offset));
  left: 50%; */
  /* overflow: hidden; */
}

.char-container {
  position: absolute;
  top: 0;
  left: 0;
  /* height: var(--radius-px, 160px); */
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