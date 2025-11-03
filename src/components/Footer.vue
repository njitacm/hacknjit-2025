<template>
  <div class="Footer">
    <footer>
      <div class="socials-container">
        <!-- <a href="https://discord.com" target="_blank"
        ><img src="../assets/logos/discord.svg"
        /></a>
        <a href="https://github.com" target="_blank"
        ><img src="../assets/logos/github-white.svg"
        /></a> -->
        <a href="https://www.instagram.com/NJITACM/" target="_blank"><img src="../assets/logos/instagram.svg" /></a>
        <a href="https://www.linkedin.com/company/njit-acm" target="_blank"><img
            src="../assets/logos/linkedin.svg" /></a>
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
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
footer {
  background: linear-gradient(transparent, #00000075);
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
  display: grid;
  grid-template-columns: 100%;
  gap: 8px;
  align-content: center;
  align-items: center;
  width: 100%;
}

.links-container a {
  flex-grow: 1;
  flex-shrink: 1;
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

.socials-container a {
  transition: transform 250ms ease;
}

.socials-container img {
  width: 50px;
  height: 50px;
}

.links-container a {
  --hov-act-col: #ffffff25;
  --hov-act-bord: solid 1px #ffffff33;
  border: 1px transparent solid;
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
    border: var(--hov-act-bord);
  }

  .socials-container a:hover {
    transform: scale(var(--hover-scale));
  }
}

@media (pointer: coarse) {
  .links-container a:active {
    background-color: var(--hov-act-col);
    border: var(--hov-act-bord);
  }

  .socials-container a:active {
    transform: scale(var(--press-scale));
  }
}

@media (max-width: 400px) {
  .socials-container img {
    width: 40px;
    height: 40px;
  }
}
</style>
