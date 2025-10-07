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
  all: initial;
  width: 75%;
  margin: 7.5% auto;
  background-color: white;
  padding: 16px;
  color: black;
  border-radius: 8px;
  text-align: unset;
}

.vf-registration p,
.vf-registration span,
.vf-registration a {
  font-size: var(--vf-font-size);
}
</style>

