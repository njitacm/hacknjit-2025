<template>
  <!--<Transition>-->
  <!--<Transition>-->
    <header ref="pageHeader" v-show="showHeader">
      <nav class="main">
        <!-- <h1>HACKNJIT</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Registration</h1>
        <!-- <h1 class="alternate">Welcome to HackNJIT!</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Schedule</h1>
        <!-- <h1 class="tertiary"><a href="#Team">Meet The Team</a></h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Sponsors</h1>
        <!-- <h1 class="secondary"><a href="#Intro">About Us</a></h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">FAQ</h1>
        <!-- <h1 class="primary" @click="toTop()">HackNJIT</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Contact Us</h1>
        <!-- <h1 class="secondary">
        <!-- <h1>HACKNJIT</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Registration</h1>
        <!-- <h1 class="alternate">Welcome to HackNJIT!</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Schedule</h1>
        <!-- <h1 class="tertiary"><a href="#Team">Meet The Team</a></h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Sponsors</h1>
        <!-- <h1 class="secondary"><a href="#Intro">About Us</a></h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">FAQ</h1>
        <!-- <h1 class="primary" @click="toTop()">HackNJIT</h1> -->
         <h1 class="primary" style="border-radius: 35px; padding:10px 20px;">Contact Us</h1>
        <!-- <h1 class="secondary">
          <a href="#Sponsors">Sponsors</a>
        </h1>
        <h1 class="tertiary">
          <a href="#FAQ">FAQ</a>
        </h1> -->
      </nav>
    </header>
  <!---</Transition>-->
  <!---</Transition>-->
</template>
<button></button>
<button></button>
<script>
export default {
  data() {
    return {
      showHeader: true,
      //fixedHeader: false,
      fixedHeader: true,
      //fixedHeader: false,
      fixedHeader: true,
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
  background-color: #5a6d5a;
  background-color: #5a6d5a;
  color: white;
  display: flex;
  width: 100vw;
  z-index: 100;
  padding: 0.5rem 0;
  position: absolute;
  top: 0;
  display: flex;
  justify-content: center;
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
  transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}
h1.alternate {
  display: none;
}
h1.primary {
  font-size: 2.75rem;
  font-weight: bold;
}
h1.primary:hover{
  background: #4a5d4ab3; /* Translucent version of the original green (70% opacity) */
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
@media (max-width: 1500px) {
  header {
    margin-top: 30px;
  }
}
@media (max-width: 1100px) {
  h1.primary {
    font-size: 2.5rem;
    font-weight: bold;
  }
  h1.secondary {
    font-size: 1.5rem;
  }
  h1.tertiary {
    font-size: 1.1rem;
  }
}
@media (max-width: 1000px) {
  h1.tertiary {
    display: none;
  }
}
@media (max-width: 750px) {
  h1.primary,
  h1.secondary {
    display: none;
  }
  h1.alternate {
    display: inherit;
    font-size: 3rem;
  }
}
@media (max-width: 650px) {
  h1 {
    width: 100%;
    text-align: center;
  }
  h1::after {
    content: unset;
    width: unset;
    height: unset;
    display: block;
    background: white;
  }
  h1:hover::after {
    width: unset;
  }
}
@media (max-width: 600px) {
  h1.alternate {
    display: inherit;
    font-size: 2.75rem;
  }
}
@media (max-width: 550px) {
  h1.alternate {
    display: inherit;
    font-size: 2.5rem;
  }
}
@media (max-width: 500px) {
  header {
    margin-top: 0;
  }
  h1.alternate {
    width: 75%;
    font-size: 2.25rem;
  }
}
@media (max-width: 375px) {
  h1.alternate {
    display: inherit;
    font-size: 2rem;
  }
}
</style>