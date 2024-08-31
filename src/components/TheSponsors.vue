<template>
  <div class="outer-container">
    <div class="title">
      <img src="../assets/HackNJIT2024/gears/gear4.svg" />
      <h1 class="title">Our Sponsors</h1>
      <!-- @click="
          window.open(
            'https://cdn.discordapp.com/attachments/1193704838010253323/1275603312468557877/HackNJIT_2024_Sponsorship_Packet.pdf?ex=66d3aca9&is=66d25b29&hm=acccbe5ad986f1124a1da9bc6aa7c65d2666a07f36caae3209df5fbbb9644fdc&'
          )
        " -->
      <h2 @click="toggleSponsors">Sponsor Packet</h2>
      <img src="../assets/HackNJIT2024/gears/gear4.svg" />
    </div>
    <Transition>
      <div class="sponsors" v-if="showSponsors">
        <h1>Title Sponsor</h1>
        <h1>Gold Sponsor</h1>
        <h1>Silver Sponsor</h1>
        <h1>Bronze Sponsor</h1>
      </div>
      <div class="sponsors" v-else>
        <table>
          <thead>
            <th></th>
            <th>Bronze</th>
            <th>Silver</th>
            <th>Gold</th>
            <th>Title</th>
          </thead>
          <tbody>
            <tr class="row-title">
              <td>Event</td>
              <td class="row-span"></td>
            </tr>
            <tr>
              <td>Company table at the event</td>
              <td>
                <img
                  class="bronze-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="silver-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="gold-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="title-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
            </tr>
            <tr>
              <td>Provide a challenge and/or an API to be hacked on</td>
              <td></td>
              <td></td>
              <td>
                <img
                  class="gold-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="title-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
            </tr>
            <tr>
              <td>Provide two judges to serve in the final presentations</td>
              <td></td>
              <td></td>
              <td>
                <img
                  class="gold-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="title-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
            </tr>
            <tr>
              <td>Provide two judges to serve in the final presentations</td>
              <td></td>
              <td>
                <img
                  class="silver-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="gold-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
              <td>
                <img
                  class="title-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
            </tr>
            <tr>
              <td>Side event or workshop</td>
              <td></td>
              <td></td>
              <td></td>
              <td>
                <img
                  class="title-gear"
                  src="../assets/HackNJIT2024/gears/gear3.svg"
                />
              </td>
            </tr>
            <tr>
              <td>Present a tech talk or demo during the event</td>
            </tr>
            <tr>
              <td>Speak at the opening ceremony</td>
            </tr>
            <tr>
              <td>Bring mentors to the event</td>
            </tr>
          </tbody>
          <tbody>
            <tr class="row-title">
              <td>Recruiting</td>
              <td class="row-span"></td>
            </tr>
            <tr>
              <td>
                Access to email addresses, LinkedIn and Github profiles of all
                attendees
              </td>
            </tr>
            <tr>
              <td>Access to virtual breakout rooms</td>
            </tr>
            <tr>
              <td>Access to resumes</td>
            </tr>
          </tbody>
          <tbody>
            <tr class="row-title">
              <td>Branding</td>
              <td class="row-span"></td>
            </tr>
            <tr>
              <td>Social media promotion</td>
            </tr>
            <tr>
              <td>Logo on site, program, and presentation</td>
            </tr>
            <tr>
              <td>Logo on t-shirt</td>
            </tr>
            <tr>
              <td>Distribute promo items</td>
            </tr>
            <tr>
              <td>
                Title sponsorship of HackNJIT (ex: “Powered by [Company Name]”)
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </Transition>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showSponsors: false,
    };
  },
  methods: {
    toggleSponsors() {
      this.showSponsors = !this.showSponsors;
    },
  },
};
</script>

<style scoped>
div.outer-container {
  width: 70%;
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
  padding: 1.5rem;
}
h1.title {
  font-size: 4rem;
  font-weight: bold;
  /* border-bottom: 8px solid white; */
  margin-bottom: 0.75rem;
}
h2 {
  display: inline-block;
  cursor: pointer;
  width: fit-content;
  justify-self: center;
}
h2::after {
  content: "";
  width: 0px;
  height: 4px;
  display: block;
  background: white;
  transition: 300ms;
}
h2:hover::after {
  width: 100%;
}
div.title {
  display: grid;
  grid-template-columns: 30% 1fr 30%;
  grid-template-rows: 1fr 1fr;
  justify-content: center;
  gap: 1rem;
}
div.title h1 {
  grid-column: 2;
}
div.title h2 {
  grid-column: 2;
}
div.title img {
  animation-name: gearSpin;
  animation-duration: 4s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
  grid-row: 1 / 3;
}
div.title img:nth-of-type(1) {
  grid-column: 1;
  justify-self: right;
}
div.title img:nth-of-type(2) {
  grid-column: 3;
  justify-self: left;
  animation-direction: reverse;
}
h1 {
  font-size: 3rem;
}
table {
  width: 100%;
  border-collapse: collapse;
  border-radius: 1rem;
  border-style: hidden; /* hide standard table (collapsed) border */
  box-shadow: 0 0 0 2px white; /* this draws the table border  */
}
thead th:first-of-type {
  padding: none;
  margin: none;
  border: 2px white solid;
  border-radius: 0 1rem 0 5px;
}
tr.row-title {
  border: 2px white solid;
}
tr.row-title td:first-of-type {
  border: 2px white solid;
}
tr td:first-of-type {
  border-right: 2px white solid;
}
td {
  margin: 0.75rem;
  word-wrap: break-word;
  max-width: 75px;
}
td img {
  width: 2rem;
}
td:last-of-type {
  padding: 0.75rem;
}
@keyframes gearSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
  /* 100% {
    transform: rotate(360);
  } */
}
.title-gear {
  filter: brightness(0) saturate(100%) invert(17%) sepia(71%) saturate(3534%)
    hue-rotate(15deg) brightness(94%) contrast(108%);
}
.gold-gear {
  filter: brightness(0) saturate(100%) invert(71%) sepia(23%) saturate(3958%)
    hue-rotate(4deg) brightness(97%) contrast(102%);
}
.silver-gear {
  filter: brightness(0) saturate(100%) invert(62%) sepia(24%) saturate(151%)
    hue-rotate(355deg) brightness(88%) contrast(89%);
}
.bronze-gear {
  filter: brightness(0) saturate(100%) invert(25%) sepia(44%) saturate(548%)
    hue-rotate(5deg) brightness(100%) contrast(93%);
}
</style>