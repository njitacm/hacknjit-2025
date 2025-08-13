<template>
  <div v-if="isOpen" class="Modal">
    <Transition>
      <div class="container" ref="target" v-if="isOpen" appear>
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
    </Transition>
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
  backdrop-filter: blur(5px);
  background-color: rgba(0, 0, 0, 0.5);
  align-content: center;
  position: absolute;
  z-index: 100001;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.Modal .container {
  max-width: 90%;
  max-height: 95%;
  width: fit-content;
  height: fit-content;
  margin: 0 auto;
  background-color: white;
  color: black;
  border-radius: 1rem;
  padding-bottom: 12px;
  overflow: hidden;
  padding: 10px;
}

.Modal .top {
  display: grid;
  grid-template-columns: 1fr auto;
  height: fit-content;
}

.Modal .title * {
  font-size: 2rem;
  display: block;
  padding: 16px;
  font-weight: bold;
}

.Modal .title {
  height: fit-content;
}

.Modal button.close {
  color: gray;
  font-size: 2rem;
  max-height: 32px;
  cursor: pointer;
  aspect-ratio: 1;
  background-color: transparent;
  border-top-right-radius: 1rem;
  border: none;
}

.Modal .body {
  max-height: 70vh;
  overflow-y: auto;
  height: inherit;
}

@media(max-width: 550px) {
  .Modal .title * {
    font-size: 1.5rem;
  }

  .Modal .container {
    box-sizing: border-box;
    max-width: none;
    width: 100%;
  }
}

.v-enter-active,
.v-leave-active {
  transition: all 1s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>