<template>
  <div class="outer-div" ref="root">
    <h1 id="section-title">Frequently Asked Questions:</h1>
    <div class="faq">
      <div
        class="faq-card"
        v-for="FAQ in FAQs"
        :key="FAQ.question"
        ref="questions"
        v-intersection-observer="[onIntersectionObserver]"
      >
        <h1>{{ FAQ.Question }}</h1>
        <p>{{ FAQ.Answer }}</p>
      </div>
    </div>
    <img
      id="gear5"
      class="floating-gear gear-fade-in"
      src="../assets/HackNJIT2024/gears/gear5.svg"
    />
    <img
      id="gear4"
      class="floating-gear gear-fade-in"
      src="../assets/HackNJIT2024/gears/gear4.svg"
    />
    <img
      id="gear3"
      class="floating-gear gear-fade-in"
      src="../assets/HackNJIT2024/gears/gear3.svg"
    />
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
  data() {
    return {
      FAQs: [
        {
          Question: "When and Where is HackNJIT?",
          Answer:
            "HackNJIT 2023 will be happening November 4th-5th; check in will start at 9am. It will be at the NJIT Campus Center (150 Bleeker St, Newark, NJ 07102).",
        },
        {
          Question: "Who can sign up to compete?",
          Answer:
            "Any undergraduate or graduate student at any college in the US, or any recent (within 6 months) graduate from a college, may register as long as they are 18 years or older.",
        },
        {
          Question: "Do I have to register as a team?",
          Answer:
            "No! You're welcome to register with your friends and compete with a team, but you can also register by yourself or even find a team at the event. ",
        },
        {
          Question: "Can I use any of my old projects?",
          Answer:
            "NO! The point of a hackathon is to build something original in a short time frame. That said, you are allowed to use open-source frameworks and third-party APIs, with proper citations of course.",
        },
        {
          Question: "Will there be food or entertainment?",
          Answer:
            "Yes, and yes! Students attending HackNJIT enjoy free food, free swag from sponsors, tech workshops, and fun events — expect a bunch of fun in person events this year!",
        },
        {
          Question: "Will there be prizes?",
          Answer:
            "Yes! Top teams are awarded a variety of prizes, including cash prizes from NJIT and special prizes from sponsors.",
        },
        {
          Question: "How big is HackNJIT?",
          Answer:
            "We're pretty big! Each year, 400 - 500 students attend HackNJIT!",
        },
        {
          Question: "Do I have to pay to compete?",
          Answer:
            "Nope! There is no fee required to signup or attend. It's totally free!",
        },
        {
          Question: "How many students do you accept each year?",
          Answer:
            "As many students as apply! At this time, we do not accept/reject students on the basis of their registration.",
        },
        {
          Question: "Do you offer parking?",
          Answer:
            "Yes, free overnight parking will be available to hackers. The parking deck is located at 42 Wilsey Street, Newark, NJ.",
        },
        {
          Question: "What should I bring?",
          Answer:
            "Bring a form of identification, proof of COVID vaccination and whatever you need to hack. This includes your laptop, your charger, and any hardware you'll need for your project. Note that we'll be running a hardware table, so you can borrow hardware (Pis, Arduinos, sensors, etc.) from us at any time during the event.",
        },
        {
          Question:
            "If I am not vacinnated or cannot find my COVID card can I still attend?",
          Answer:
            "Yes, as long you provide proof of a negative result within 3 days of the event (earliest being November 1st) and show it during entry, you are all set to go! ",
        },
        {
          Question: "What shouldn't I bring?",
          Answer:
            "We ask that you not bring your desktop or any other non-portable hardware to HackNJIT, as that will make it difficult for us to rearrange tables for judging, meals, and events. Of course, weapons, firearms, alcohol, and drugs are strictly forbidden. If you can't take it on a flight, you can't take it to HackNJIT!",
        },
        {
          Question: "Do you offer travel reimbursements?",
          Answer:
            "Unfortunately, we do not offer travel reimbursements at this time.",
        },
        {
          Question: "You guys haven’t answered my question.",
          Answer:
            "If you still have questions, or want more information, feel free to email us at hacknjit@njit.edu! ",
        },
      ],
    };
  },
};
</script>
  
<style scoped>
@keyframes gear-fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 0.55;
  }
}
#gear5 {
  top: 30%;
  right: 43%;
  width: 47vw;
  max-width: 782px;
}
#gear4 {
  top: 10%;
  left: 63%;
  width: 37vw;
  max-width: 482px;
}
#gear3 {
  bottom: 0%;
  right: -11.5%;
  width: 68vw;
  max-width: 402px;
  transform: rotate(248deg);
}
.floating-gear {
  position: absolute;
  z-index: -100;
  opacity: 0.55;
  animation-name: gear-fade-in;
  animation-duration: 0.75s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
.outer-div {
  position: relative;
  width: 60%;
  /* border-top: black solid 4px; */
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
  /* background-color: var(--color4); */
  padding: 1.5rem;
}
#section-title {
  font-size: 4rem;
  margin-bottom: 1.75rem;
}
.faq {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}
.faq-card {
  width: 100%;
  margin: 1rem;
  text-align: left;
}

h1 {
  margin-bottom: 0.25rem;
  font-size: 1.5rem;
}
p {
  font-size: 1.125rem;
}
.fade-in {
  animation-name: fade-in;
  animation-duration: 0.5s;
  /* animation-delay: 1.5s; */
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

@media (max-width: 750px) {
  .outer-div {
    width: 80%;
  }
  #section-title {
    font-size: 2.5rem;
  }
}
</style>