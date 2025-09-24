<template>
  <div class="NotifSignup">
    <a :href="vw <= SHOW_MODAL_THRESHOLD ? 'https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?usp=dialog' : null"
      target="_blank" class="email-signup">
      <button @click="openNewModal">
        Sign up for email updates
      </button>
    </a>
  </div>
</template>

<script setup>
import SignupForm from "./modal_components/SignupForm.vue";
import { ref, onMounted, onBeforeMount } from "vue";
import { useModal } from "../composables/useModal";

const SHOW_MODAL_THRESHOLD = 400;   // viewport width in px;
const getVw = () => Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0);
const vw = ref(getVw());
const { openModal, closeModal } = useModal();

const onWindowResize = () => {
  vw.value = getVw();

  if (vw.value <= SHOW_MODAL_THRESHOLD) {
    closeModal();
  }
};

onMounted(() => {
  window.addEventListener("resize", onWindowResize);
});

onBeforeMount(() => {
  window.removeEventListener("resize", onWindowResize);
});

const openNewModal = () => {
  if (vw.value > SHOW_MODAL_THRESHOLD) {
    openModal({
      title: "Receive Email Updates",
      component: SignupForm
    });
  }
};

</script>

<style scoped>
.email-signup {
  border-radius: 100px;
}

.email-signup button {
  border-radius: 100px;
  background-color: white;
  color: #170800;
  padding: 12px 24px;
  border: none;
  font-weight: bold;
  cursor: pointer;
}

h1 {
  margin: 5rem 0 1.75rem 0;
  font-weight: bold;
}

p {
  height: -moz-fit-content;
  height: fit-content;
  align-self: center;
  text-align: center;
}

img {
  max-width: 50%;
  border-radius: 30px;
  flex: 1;
}
</style>
