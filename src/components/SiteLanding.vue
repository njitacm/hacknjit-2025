<template>
  <div class="outer-div">
    <img id="HackNJITLogo" src="../assets/HackNJIT2023/hacknjit2023_logo.png" />
    <h1>
      Thank you all for an amazing hackathon. <br />We'll see you again in 2024
    </h1>
    <div class="button-row">
      <!-- <h1>Registration is now OPEN!!!!</h1>
      TODO: After registration ends, say walk in registration still works
      <button type="button" class="btn border border-success"
          onclick="location.href='https://forms.gle/1hLYAR1JCCyti8kQ9'"
          style="background-color: #1D539F; color:#d9f0f9; font-size:xx-large"><strong>Register
              Here</strong>
      </button>
      <div class="col-md-12" style="padding-top: 5px; margin: 2vh auto;"> -->

      <button
        onclick="location.href='https://discord.gg/umQHUfhnMn'"
        style="background-color: #1d539f; color: white; font-size: xx-large"
      >
        <img src="../assets/discord.svg" height="50" />Join Our Discord!
      </button>

      <button onclick="location.href='https://hacknjit-2023.devpost.com/'">
        <img
          src="../assets/devpost_white.png"
          height="50"
          style="padding-right: 5px"
        />See the Devpost!
      </button>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.outer-div {
  width: 80%;
  border-bottom: black solid 4px;

  margin: 0 auto;
}
#HackNJITLogo {
  width: 37.5%;
}
h1 {
  margin: 0 auto;
  width: 65%;
  font-size: 2.5rem;
}
button {
  margin: 0.5rem 0.25rem;
  margin-top: 0.75rem;
  padding: 0.375rem 0.75rem;
  background-color: #1d539f;
  color: white;
  font-size: 2rem;
  cursor: pointer;
  line-height: 1.5;
  border: 1px solid transparent;
  border-radius: 0.25rem;
  transition: all 0.15s ease-in-out;
}
button:focus {
  box-shadow: 0 0 0 0.25rem rgba(13, 110, 253, 0.25);
}
img {
  vertical-align: middle;
  margin-right: 7.5px;
}
</style>