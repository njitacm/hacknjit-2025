<template>
  <body>
    <main>
      <img
        id="HackNJITLogo"
        src="../assets/HackNJIT2024/hacknjit2024_logo.png"
      />
      <section>
        <Transition mode="out-in">
          <div v-if="mode == 0" ref="mode0">
            <h1>Welcome to HackNJIT</h1>
            <p>
              HackNJIT is a 24-hour hackathon at the New Jersey Institute of
              Technology, run by its ACM student chapter in conjunction with the
              Ying Wu College of Computing. Students break into small teams of
              up to 4, work together over 24 hours to create a tech project, and
              at the end present it to our judges!
            </p>
          </div>
          <div v-else-if="mode == 1" ref="mode1">
            <h1>Registration</h1>
            <p>
              Registration is currently closed for HackNJIT 2024, check back in
              soon!
            </p>
          </div>
          <div v-else-if="mode == 2" ref="mode2" id="links">
            <SteampunkButton
              @click="openInNewTab('https://discord.gg/umQHUfhnMn')"
              isLarge="true"
            >
              Join Our Discord!
            </SteampunkButton>
            <SteampunkButton
              @click="openInNewTab('https://hacknjit-2023.devpost.com/')"
              isLarge="true"
            >
              See the Devpost!
            </SteampunkButton>
            <!-- <SteampunkButton
              @click="openInNewTab('https://hacknjit-2023.devpost.com/')"
              isLarge="true"
            >
              A Secret Third Thing
            </SteampunkButton> -->
          </div>
        </Transition>
      </section>
    </main>
    <nav>
      <SteampunkButton
        @click="changeMode(0)"
        :class="[mode == 0 ? 'activeButton' : 'unactiveButton']"
      >
        Who We Are
      </SteampunkButton>
      <SteampunkButton
        @click="changeMode(1)"
        :class="[mode == 1 ? 'activeButton' : 'unactiveButton']"
      >
        <div>Registration</div>
      </SteampunkButton>
      <SteampunkButton
        @click="changeMode(2)"
        :class="[mode == 2 ? 'activeButton' : 'unactiveButton']"
      >
        Links
      </SteampunkButton>
    </nav>
    <div class="division"></div>
  </body>
</template>

<script>
import SteampunkButton from "./SteampunkButton.vue";

export default {
  components: {
    SteampunkButton,
  },
  data() {
    return {
      mode: 2,
      sectionHeight: 0,
    };
  },
  computed: {
    heightOfSection() {
      return this.sectionHeight;
    },
  },
  methods: {
    changeMode(mode) {
      this.mode = mode;
    },
    openInNewTab(url) {
      window.open(url, "_blank").focus();
    },
  },
};
</script>

<style scoped>
nav {
  display: flex;
  margin: 0 auto;
  margin-bottom: 0.5rem;
  justify-content: center;
  gap: 1rem;
}
button {
  margin: 0.5rem 0.25rem;
  margin-top: 0.75rem;
  /* padding: 0.375rem 0.75rem; */
  background-color: var(--main-fg-color);
  color: white;
  font-size: 2rem;
  cursor: pointer;
  line-height: 1.5;
  /* border: 0.25rem solid #d4b490;
  border-radius: 1rem; */
  transition: all 0.5s ease-in-out;
  /* width: 250px; */
  height: fit-content;
}
.unactiveButton {
  scale: 0.875;
}
.activeButton {
  /* border: 0.25rem solid #bd9d78; */
  /* box-shadow: 0 0 0 0.25rem #bd9d78; */
  transform: scale(1.1);
}
button:focus {
  /* box-shadow: 0 0 0 0.25rem #bd9d78; */
}
.button-img {
  vertical-align: middle;
  margin-right: 7.5px;
}
.button-text {
  text-align: center;
  width: 100%;
  align-self: center;
}
.division {
  width: 50%;
  margin: auto;
  border-radius: 0.5rem;
  border-bottom: 0.25rem solid var(--color5);
}
main {
  width: 60%;
  margin: 0 auto;
  display: flex;
  align-content: center;
  justify-content: center;
  height: fit-content;
}
#HackNJITLogo {
  flex: 1;
}
section {
  max-height: 486px;
  height: 486px;
  /* border: 1px solid white; */
  border-radius: 1rem;
  flex: 1;
  align-content: start;
  align-self: center;
  padding: 0.5rem;
}
#links {
  /* align-content: center; */
  display: grid;
  /* grid-template-rows: 162px 162px 162px; */
  grid-template-rows: 300px 50%;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}
h1 {
  font-size: 4rem;
  /* text-shadow: 3px 3px 3px #a7793e, -3px 3px 4px #a7793e, 3px -3px 4px #a7793e,
    -3px -3px 4px #a7793e; */
  padding: 1rem;
  padding-bottom: 0;
}
p {
  font-size: 1.75rem;
  padding: 1rem;
  padding-top: 0;
  /* text-shadow: 3px 3px 3px #a7793e, -3px 3px 4px #a7793e, 3px -3px 4px #a7793e,
    -3px -3px 4px #a7793e; */
  text-align: start;
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
</style>