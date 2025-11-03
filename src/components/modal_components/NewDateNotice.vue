<template>
  <div class="NewDateNotice">
    <p>
      HackNJIT will now be held on <strong>March 7 - March 8, 2026</strong>.
      <br />
      <br />
      This was not an easy decision to make, as we know many of you were excited for this event, but we felt it
      necessary
      in order to ensure we can host the best possible experience for everyone.
      <br />
      <br />
      If you have any questions or concerns, email us at 
      <a href="mailto:acm@njit.edu" @click="copyToClipboard">acm@njit.edu</a> 
      for more information.
    </p>
    <p v-if="showCopiedMsg" class="copiedMsg">
      Email copied to clipboard!
    </p
    <hr />
    <div class="actions">
      <button class="pill" @click="closeForever">Don't show again</button>
      <button class="pill" @click="closeModal">OK</button>
    </div>
  </div>
</template>

<script>
import { useModal } from '../../composables/useModal';
const { closeModal } = useModal();

export default {
  data() {
    return {
      loaded: false,
      closeModal: closeModal,
      showCopiedMsg: false,
      timeout: null
    }
  },
  methods: {
    closeForever() {
      localStorage.setItem("hideNewDateNotice", true);
      closeModal();
    },
    copyToClipboard() {
      navigator.clipboard.writeText("acm@njit.edu");
      this.showCopiedMsg = true;
      this.timeout = () => {
        clearTimeout(this.timeout);
        this.showCopiedMsg = false;
      };
      setTimeout(this.timeout, 2500);
    }
  }
}
</script>

<style scoped>
.NewDateNotice {
  width: 100%;
  max-width: 500px;
  padding: 32px;
  display: grid;
  gap: 32px;
}

strong {
  font-weight: bolder;
}

.actions {
  display: flex;
  gap: 16px;
  justify-content: space-between;
}

.copiedMsg {
  position: absolute;
  bottom: 32px;
  color: white;
  text-align: center;
  left: 50%;
  transform: translateX(-50%);
}

a {
  font-size: 1em;
}

button.pill {
  font-size: 1.25em;
}

@media (max-width: 500px) {
  .actions {
    flex-direction: column;
  }
}
</style>