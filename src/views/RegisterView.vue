<template>
  <div class="RegisterView View">
    <Registration></Registration>
  </div>
</template>

<script>
import Registration from "../components/Registration.vue";

export default {
  name: "HackNJIT",
  components: {
    Registration,
  }
};
</script>

<style scoped>
* {
  all: initial;
}

</style>

<style>
.RegisterView {
  --top-offset: 80px;
  width: 75%;
  height: 100%;
  background-color: white;
  padding: 16px;
  margin: auto;
  margin-bottom: var(--top-offset);
  top: var(--top-offset);
  position: relative;
  color: black;
  border-radius: 8px;
  text-align: unset;
  display: block;
}

.vf-registration p,
.vf-registration span,
.vf-registration a {
  font-size: var(--vf-font-size);
}
</style>

