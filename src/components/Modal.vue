<template>
  <div :class="'Modal' + (open === true ? ' open' : '')">
    <div class="top">
      <slot name="title"></slot>
      <button class="close" @click="$emit('modalClosed'); open = false">X</button>
    </div>
    <slot name="body"></slot>
  </div>
</template>

<script>
export default {
  name: "Modal",
  props: ["isOpen"],
  emits: ["modalClosed"],
  data() {
    return {
      open: this.isOpen,
    };
  },
  created() {
    console.log("Mounted!");
  }
}
</script>

<style scoped>
.Modal:not(.open) {
  display: none;
  background-color: white;
  color: black;
}

.Modal.open {
  display: block;
}

.Modal {
  position: absolute;
  align-content: center;
  margin: 0 auto;
}
</style>