<template>
  <Teleport to="body">
    <div v-for="(modal, index) in modalStack" :key="index" class="Modal" :style="{ zIndex: 1000 + index }">
      <Transition appear>
        <div class="modal-backdrop">
          <div class="container" ref="target">
            <header class="header">
              <h2>{{ modal.title }}</h2>
              <button @click.stop="closeModal" class="close">&times;</button>
            </header>
            <section class="body">
              <component :is="modal.component" v-bind="modal.props" />
            </section>
          </div>
        </div>
      </Transition>
    </div>
  </Teleport>
</template>

<script setup>
import { useTemplateRef } from "vue";
import { onClickOutside } from "@vueuse/core";
import { useModal } from "../composables/useModal";

const { modalStack, closeModal } = useModal();

const target = useTemplateRef("target");
onClickOutside(target, () => closeModal());
</script>

<style scoped>
.Modal,
.modal-backdrop {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.modal-backdrop {
  backdrop-filter: blur(5px);
  background-color: rgba(0, 0, 0, 0.5);
  align-content: center;
  z-index: 100001;
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

.Modal .header {
  display: grid;
  grid-template-columns: 1fr auto;
  height: fit-content;
}

.Modal .header * {
  font-size: 2rem;
  display: block;
  padding: 16px;
  font-weight: bold;
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
  .Modal .header * {
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
  transition: background-color 250ms linear;
}

.v-enter-from,
.v-leave-to {
  background-color: transparent;
}

.v-enter-active .container,
.v-leave-active .container {
  transition: transform 250ms ease-out, opacity 250ms ease-out;
}

.v-enter-from .container,
.v-leave-to .container {
  opacity: 0;
  transform: translateY(-50px);
}
</style>