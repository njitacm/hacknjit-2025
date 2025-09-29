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
  all: initial;
  width: 75%;
  margin: 5% auto 5% auto;
  background-color: white;
  padding: 16px;
  color: black;
  border-radius: 8px;
  text-align: unset;
}



</style>

