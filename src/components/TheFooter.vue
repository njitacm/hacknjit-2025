<template>
  <footer class="Footer section">
    <div class="socials-container">
      <!-- <a href="https://discord.com" target="_blank"
        ><img src="../assets/logos/discord.svg"
      /></a>
      <a href="https://github.com" target="_blank"
        ><img src="../assets/logos/github-white.svg"
      /></a> -->
      <a href="https://www.instagram.com/NJITACM/" target="_blank"><img src="../assets/logos/instagram.svg" /></a>
      <a href="https://www.linkedin.com/company/njit-acm" target="_blank"><img src="../assets/logos/linkedin.svg" /></a>
      <a href="https://x.com/NJITACM" target="_blank"><img src="../assets/logos/twitter.svg" /></a>
      <a href="https://https://www.youtube.com/channel/UC3stSn7UK5-IX6mNnlWlRLw" target="_blank"><img
          src="../assets/logos/youtube.svg" /></a>
    </div>
    <div class="NJIT_logos">
      <img id="acm-logo" src="../assets/logos/acm_logo_white.svg" />
      <img id="ywcc-logo" src="../assets/logos/NJIT_white.png" />
    </div>
    <div class="links-container">
      <a href="https://acm.njit.edu" target="_blank">Main ACM Site</a>
      <a href="https://www.acm.org/" target="_blank">ACM National</a>
      <a href="https://computing.njit.edu/" target="_blank">NJIT Ying Wu College of Computing</a>
      <!-- <hr />
      <a href="https://acm.njit.edu" target="_blank">Sponsor Packet</a> -->
      <a href="mailto:hacknjit@njit.edu" target="_blank">Email Us</a>
      <a href="https://static.mlh.io/docs/mlh-code-of-conduct.pdf" target="_blank">MLH Code of Conduct</a>
    </div>
  </footer>
</template>

<script>
export default {};
</script>

<style scoped>
.Footer {
  margin-inline: auto;
  margin-bottom: 32px;
  border-radius: var(--border-radius);
  background-color: var(--hacknjit-secondary);
  color: var(--text-color);
  display: grid;
  justify-content: center;
  gap: 16px;
  padding: 32px;
}

.socials-container {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 32px;
}

.links-container {
  --gap: 8px;
  --cols: 2;
  display: grid;
  grid-template-columns: repeat(var(--cols), calc(100% / var(--cols) - var(--gap) / (var(--cols) - 1)));
  gap: var(--gap);
  align-content: center;
  align-items: center;
  width: 100%;
}

.NJIT_logos {
  display: flex;
  align-content: center;
  justify-content: center;
  flex-wrap: wrap;
  gap: 0.5rem;
}

#acm-logo {
  width: 100px;
  align-self: center;
}

#ywcc-logo {
  width: 200px;
}

.socials-container img {
  width: 50px;
  height: 50px;
}

.links-container a {
  --hov-act-col: var(--hacknjit-primary);
  width: 100%;
  border-radius: var(--border-radius);
  color: white;
  font-weight: bold;
  text-decoration: none;
  flex-grow: 1;
  text-align: center;
  padding: 8px 16px;
}

@media (hover: hover) and (pointer: fine) {
  .links-container a:hover {
    background-color: var(--hov-act-col);
  }
}

@media (pointer: coarse) {
  .links-container a:active {
    background-color: var(--hov-act-col);
  }
}

@media(max-width: 1000px) {
  .Footer {
    width: 100%;
    border-radius: 0;
  }
}

@media (max-width: 600px) {
  .socials-container img {
    width: 40px;
    height: 40px;
  }

  .links-container {
    --cols: 1;
    grid-template-columns: repeat(var(--cols), calc(100% / var(--cols)));
  }
}

@media (max-width: 500px) {
  .socials-container img {
    width: 35px;
  }
}
</style>
