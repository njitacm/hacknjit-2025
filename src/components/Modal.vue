<template>
  <Teleport to="body">
    <div v-for="(modal, index) in modalStack" :key="index" class="Modal" :style="{ zIndex: 1000 + index }">
      <div class="modal-backdrop">
        <Transition appear>
          <div class="container" v-on-click-outside="closeModal">
            <header class="header">
              <h2>{{ modal.title }}</h2>
              <button @click="closeModal" class="close">&times;</button>
            </header>
            <section class="body">
              <component :is="modal.component" v-bind="modal.props" />
            </section>
          </div>
        </Transition>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { vOnClickOutside } from "@vueuse/components"
import { useModal } from "../composables/useModal";
const { modalStack, closeModal } = useModal();
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
  transform-origin: top center;
  max-width: 90%;
  max-height: 95%;
  width: fit-content;
  height: fit-content;
  margin: 0 auto;
  background-color: white;
  color: black;
  border-radius: var(--border-radius);
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

@media (prefers-reduced-motion: no-preference) {

  .v-enter-active,
  .v-leave-active {
    animation: enter;
    animation-duration: 250ms;
    animation-timing-function: ease;
  }

  @keyframes enter {
    from {
      transform: perspective(100cm) rotateX(30deg) translateY(0px);
    }

    to {
      transform: perspective(none) rotateX(0deg) translateY(0px);
    }
  }
}
</style>