<template>
  <div class="outer-div">
    <h1>Meet Our Sponsors!</h1>
    <h2 id="sponsors-gold">Gold Sponsors</h2>
    <div class="sponsor-row">
      <a target="_blank" href="https://www.verizon.com/"
        ><img
          alt="Verizon"
          src="../assets/HackNJIT2023/Sponsors/verizon.svg"
          width="50%"
      /></a>
    </div>
    <h2 id="sponsors-silver">Silver Sponsors</h2>
    <div class="sponsor-row">
      <a target="_blank" href="https://www.patientsafetytech.com/"
        ><img
          alt="Patient Safety Technology"
          src="../assets/HackNJIT2023/Sponsors/PSTC.png"
          width="100%"
      /></a>
      <a target="_blank" href="https://www.databank.com/"
        ><img
          alt="Databank"
          src="../assets/HackNJIT2023/Sponsors/databank.png"
          width="60%"
      /></a>
    </div>
    <h2 id="sponsors-bronze">Bronze Sponsors</h2>
    <div class="sponsor-row">
      <a target="_blank" href="https://www.scm-lp.com/"
        ><img
          id="SCM"
          alt="Stevens Captial Management"
          src="../assets/HackNJIT2023/Sponsors/SCM.svg"
      /></a>
    </div>
    <h2 id="sponsors-special">Special Thanks</h2>
    <div class="sponsor-grid">
      <a target="_blank" href="https://cloud.google.com/">
        <img
          alt="Google Cloud"
          src="../assets/HackNJIT2023/Sponsors/google-cloud.svg"
          style="max-width: 30%"
        />
      </a>
      <a target="_blank" href="https://github.com/">
        <img
          alt="Github"
          src="../assets/HackNJIT2023/Sponsors/github.png"
          style="max-width: 30%"
        />
      </a>
      <a target="_blank" href="https://aws.amazon.com/">
        <img
          alt="AWS"
          src="../assets/HackNJIT2023/Sponsors/aws.svg"
          style="max-width: 30%"
        />
      </a>
      <a
        target="_blank"
        href=" http://hackp.ac/mlh-StandOutStickers-hackathons"
      >
        <img
          alt="StandoutStickers"
          src="../assets/HackNJIT2023/Sponsors/standoutStickers.png"
          style="max-width: 30%"
        />
      </a>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.outer-div {
  margin: 0 auto;
  width: 80%;
  border-top: black solid 4px;
  padding: 0.5rem;
}
h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
}
h2 {
  margin-bottom: 0.5rem;
  font-size: 1.75rem;
}
#sponsors-gold {
  background: radial-gradient(
    rgb(245, 222, 51) 30%,
    rgba(255, 255, 255, 0) 69%
  );
}
#sponsors-silver {
  background: radial-gradient(
    rgb(185, 193, 201) 30%,
    rgba(255, 255, 255, 0) 69%
  );
}
#sponsors-bronze {
  background: radial-gradient(
    rgb(205, 127, 50) 30%,
    rgba(255, 255, 255, 0) 69%
  );
}
#sponsors-special {
  background: radial-gradient(
    rgb(255, 109, 109) 30%,
    rgba(255, 255, 255, 0) 69%
  );
}
.sponsor-row {
  justify-content: space-around;
  align-content: center;
  display: flex;
  width: 50%;
  margin-right: auto;
  margin-left: auto;
  flex-wrap: wrap;
  margin-bottom: 1rem;
}
.sponsor-row a {
  flex: 1 0 0;
  align-self: center;
  display: block;
}
.sponsor-row a img {
  max-height: 16.625rem;
  width: auto;
  height: auto;
}
.sponsor-grid {
  width: 80%;
  margin-right: auto;
  margin-left: auto;
  display: grid;
  grid-template-columns: auto auto;
}
</style>