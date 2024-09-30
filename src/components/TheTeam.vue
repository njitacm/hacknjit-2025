<template>
  <div class="outer-container" ref="container">
    <img
      id="gear2"
      class="floating-gear"
      src="../assets/HackNJIT2024/gears/gear2.svg"
    />
    <h1 v-intersection-observer="[onIntersectionObserver]">Meet The Team!</h1>
    <main class="team">
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member1"
      >
        <img src="../assets/eboard/president.png" />
        <h2>Will Simkins</h2>
        <h3>ACM President</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member2"
      >
        <img src="../assets/eboard/vice_president.png" />
        <h2>Anthony Siracusa</h2>
        <h3>ACM Vice-President</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member3"
      >
        <img src="../assets/eboard/treasurer.png" />
        <h2>Logan Kriebel</h2>
        <h3>ACM Treasurer</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member4"
      >
        <img src="../assets/eboard/event_master.png" />
        <h2>Priya Patel</h2>
        <h3>ACM Event Master</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member5"
      >
        <img src="../assets/eboard/graphic_designer.png" />
        <h2>Kapila Mane</h2>
        <h3>ACM Graphic Designer</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member6"
      >
        <img src="../assets/eboard/webmaster.png" />
        <h2>Abdullah Imran</h2>
        <h3>ACM Webmaster</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member7"
      >
        <img src="../assets/eboard/secretary.png" />
        <h2>Ananya Tyagi</h2>
        <h3>ACM Secretary</h3>
      </div>
      <div
        class="member"
        v-intersection-observer="[onIntersectionObserver]"
        ref="member8"
      >
        <img src="../assets/eboard/public_relations.png" />
        <h2>Darren Bonifacio</h2>
        <h3>ACM Public Relations</h3>
      </div>
    </main>
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    console.log(target);
    target.classList.add("fade-in");
  }
}
</script>
<script>
export default {
  mounted() {
    for (let i = 1; i <= 8; i++) {
      var increase = i <= 4 ? i / 4 : (i - 4) / 4;
      this.$refs["member" + i].style["transition-delay"] = increase + "s";
    }
  },
};
</script>

<style scoped>
.floating-gear {
  position: absolute;
  z-index: -100;
  transform: rotate(30deg);
  top: 45%;
  right: -10%;
  width: 45vw;
  max-width: 700px;
  opacity: 0.55;
  animation-name: gear-fade-in;
  animation-duration: 0.75s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
.fade-in {
  animation-name: fade-in;
  animation-duration: 0.5s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
div.outer-container {
  width: 70%;
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
  padding: 1.5rem;
  position: relative;
}
h1 {
  font-size: 4rem;
  margin-bottom: 1.75rem;
}
main.team {
  display: flex;
  justify-content: center;
  gap: 1rem;
  flex-wrap: wrap;
}
.member {
  flex: 1;
  margin-top: 1rem;
}
.member img {
  border-radius: 2rem;
  border: 4px white solid;
  max-width: 300px;
  aspect-ratio: 1 / 1;
}
.member h2 {
  margin: 0.5rem;
}
@keyframes gear-fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 0.55;
  }
}
@media (max-width: 1350px) {
  .member img {
    max-width: 225px;
  }
}
@media (max-width: 750px) {
  .member img {
    max-width: 35vw;
  }
  div.outer-container {
    width: 85%;
  }
  h1 {
    font-size: 3.5rem;
  }
}
</style>