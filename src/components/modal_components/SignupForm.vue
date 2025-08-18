<template>
  <iframe title="Signup Google Form" :onload="onLoad" :class="{ loaded: loaded }"
    src="https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?embedded=true"
    frameborder="0" marginheight="0" marginwidth="0">
    Loading…
  </iframe>
</template>

<script>
export default {
  data() {
    return {
      loaded: false,
    }
  },
  methods: {
    onLoad() {
      this.loaded = true;
    }
  }
}
</script>

<style scoped>
iframe {
  max-width: 100%;
  max-height: 100%;
  width: 750px;
  height: 550px;
  overflow: hidden;
  opacity: 0;
}

iframe.loaded {
  opacity: 1;
}

@media(max-width: 650px) {
  iframe {
    height: 650px;
  }
}

@media(prefers-reduced-motion: no-preference) {
  iframe {
    transition: opacity 150ms linear;
    transition-delay: 50ms;
  }
}
</style>