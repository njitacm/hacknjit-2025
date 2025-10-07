<template>
  <div class="registration-container">
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
.registration-container {
  width: 75%;
  height: 100%;
  background-color: white;
  padding: 16px;
  margin: auto;
  top: 80px;
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

