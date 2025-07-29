<template>
  <div v-if="isOpen" class="Modal">
    <div class="container"  ref="target">
      <div class="top">
        <slot name="title"></slot>
        <button class="close" @click.stop="$emit('modalClose')">X</button>
      </div>
      <slot name="body"></slot>
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
  position: fixed;
  z-index: 99;
  top: 0;
  left: 0;
  width: 100% !important;
  height: 100% !important;
  background-color: rgba(0, 0, 0, 0.5);
}

.Modal .container {
  margin: 0 auto;
  background-color: white;
  color: black;
}
</style>