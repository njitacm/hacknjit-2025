<template>
  <Transition>
    <header ref="pageHeader" v-show="showHeader">
      <nav class="main">
        <h1 class="tertiary"><a href="#Team">Meet The Team</a></h1>
        <h1 class="secondary"><a href="#Intro">About Us</a></h1>
        <h1 class="primary" @click="toTop()">HackNJIT</h1>
        <h1 class="secondary"><a href="#Registration">Registration</a></h1>
        <h1 class="tertiary">
          <a href="#Sponsors">Sponsors</a>
        </h1>
      </nav>
      <nav></nav>
    </header>
  </Transition>
</template>

<script>
export default {
  data() {
    return {
      showHeader: true,
      fixedHeader: false,
    };
  },
  methods: {
    handleScroll() {
      this.fixedHeader =
        window.scrollY > this.$refs.pageHeader.clientHeight + 400;
      this.showHeader = window.scrollY < 10;
      // ||
      // window.scrollY > this.$refs.pageHeader.clientHeight + 750;
    },
    toTop() {
      window.scrollTo({ top: 0, behavior: "smooth" });
    },
  },
  created() {
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  mounted() {
    this.handleScroll();
  },
};
</script>

<style scoped>
* {
  transition: all linear 0.25s;
}
.fake-header {
  height: 40px;
}

header {
  color: white;
  display: flex;
  width: 100%;
  /* background: var(--edge-colors); */
  /* background: #d3b28ed8; */
  z-index: 100;
  padding: 0.5rem 0;
  position: absolute;
  top: 0;
  display: grid;
  grid-template-columns: 20% auto 20%;
  position: fixed;
}

.v-enter-active,
.v-leave-active {
  transition: all 0.5s ease-in-out;
}
.v-enter-from,
.v-leave-to {
  opacity: 0;
  top: -60px;
}
nav {
  padding: 4px;
  padding-bottom: 0;
  /* margin-left: 2rem; */
  display: flex;
  justify-content: center;
  align-content: baseline;
  cursor: pointer;
  gap: 2rem;
}
nav.main {
  grid-column: 2;
  border-bottom: 4px white none;
}
a {
  text-decoration: none;
  color: inherit;
}

h1 {
  align-self: baseline;
  font-size: 0.75rem;
  padding: 5px;
  text-align: center;
}
h1.primary {
  font-size: 2.75rem;
  font-weight: bold;
}
h1.secondary {
  font-size: 1.75rem;
}
h1.tertiary {
  font-size: 1.25rem;
}
h1.title:hover {
  text-decoration: none;
}
h1 {
  display: inline-block;
}
h1::after {
  content: "";
  width: 0px;
  height: 4px;
  display: block;
  background: white;
  transition: 300ms;
}
h1:hover::after {
  width: 100%;
}
</style>