<template>
  <div v-if="isOpen" class="Modal">
    <div class="container" ref="target">
      <div class="top">
        <h1 className="title">
          <slot name="title"></slot>
        </h1>
        <button class="close" @click.stop="$emit('modalClose')">&#10005;</button>
      </div>
      <div class="body">
        <slot name="body"></slot>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, useTemplateRef } from "vue";
import { onClickOutside } from "@vueuse/core";

const emit = defineEmits(["modalClose"]);
const props = defineProps({
  isOpen: Boolean,
});

const target = useTemplateRef("target");
onClickOutside(target, () => emit("modalClose"));

</script>

<style scoped>
.Modal {
  --anim-dur: 250ms;
  backdrop-filter: blur(5px);
  align-content: center;
  position: absolute;
  z-index: 100001;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  animation: fade-in linear var(--anim-dur);
}

.Modal .container {
  width: 95%;
  height: 95%;
  max-width: fit-content;
  max-height: fit-content;
  margin: 0 auto;
  background-color: white;
  color: black;
  border-radius: 1rem;
  animation: drop-in ease-out var(--anim-dur);
  padding-bottom: 12px;
}

.Modal .top {
  display: grid;
  grid-template-columns: 1fr auto;
}

.Modal .title * {
  font-size: 2rem;
  display: block;
  padding: 16px;
  font-weight: bold;
}

.Modal button.close {
  color: gray;
  font-size: 2rem;
  cursor: pointer;
  aspect-ratio: 1;
  background-color: transparent;
  border-top-right-radius: 1rem;
  border: none;
}

.Modal .body {
  overflow-y: auto;
}

@keyframes fade-in {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

@keyframes drop-in {
  from {
    transform: translateY(-50px);
  }

  to {
    transform: none;
  }
}
</style>