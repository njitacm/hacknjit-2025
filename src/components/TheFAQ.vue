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
    // console.log(target);
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
            "HackNJIT 2024 will be happening November 2nd-3rd; check in will start at 9am. It will be at the NJIT Campus Center (150 Bleeker St, Newark, NJ 07102).",
        },
        {
          Question: "Who can sign up to compete?",
          Answer:
            "Any undergraduate or graduate student at any college in the US, or any recent (within 12 months) graduate from a college, may register as long as they are 18 years or older.",
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
            "Bring a form of identification and whatever you need to hack. This includes your laptop, your charger, and any hardware you'll need for your project. Note that we'll be running a hardware table, so you can borrow hardware (Pis, Arduinos, sensors, etc.) from us at any time during the event.",
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
          Question: "How many people can be in one team?",
          Answer: "We allow up to four people per team.",
        },
        {
          Question: "Can a member of our team be at the event virtually?",
          Answer:
            "No, all team members are required to check in at the start of the event.",
        },
        {
          Question: "Do we need to stay the full duration of the event?",
          Answer:
            "You are not required to stay on-site for the full event, however you are required to be present at check-in and during judging if you are submitting a project.",
        },
        {
          Question: "How long does the Hackathon last?",
          Answer:
            "HackNJIT is unique in that it occurs during daylight savings time, which means it is technically a 25 hour Hackathon! Make good use of that extra hour!",
        },
        {
          Question: "What are the start and end times for the event?",
          Answer:
            "Check-in starts at 9:00 Saturday and hacking starts at noon the same day. You are required to submit your project by noon on Sunday.",
        },
        {
          Question: "Where do I submit my project?",
          Answer: "You submit your project on our Devpost! (linked above)",
        },
        {
          Question: "Where is the full rule list?",
          Answer:
            "You can find the full list of rules on our Devpost! (linked above)",
        },
        {
          Question:
            "What is the theme for the event? Do I need to make a project pertaining to that theme?",
          Answer:
            'Our theme is "Steampunk" but you aren\'t required to make a themed hack, there is a separate prize category for "Best Themed Hack".',
        },
        {
          Question: "What projects are allowed?",
          Answer:
            "Anything you can think of! We accept any software project from websites to phone app to video games and more!",
        },
        {
          Question:
            "I don't have any programming experience, can I still compete?",
          Answer:
            "Of course! HackNJIT is an amazing opportunity to learn more about programming and we heavily encourage you attend regardless of your skill level. Even if you don't submit a project, we will be hosting numerous educational workshops that you can attend to pick up new skills.",
        },
        {
          Question: "Do I need to be a programmer to be compete?",
          Answer:
            "Not at all! Every team needs a good business major to lead their pitch, or an artist to make project mockups! The best teams consist of individuals with a diverse array of skills working together to create something none could do alone!",
        },
        {
          Question: "How can I form a team?",
          Answer:
            "If you don't have a team already, join our discord to meet new people and find others looking for a team!",
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