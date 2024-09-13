<template>
  <footer>
    <nav>
      <hr />
      <img src="../assets/Logos/discord.svg" />
      <img src="../assets/Logos/github-white.svg" />
      <img src="../assets/Logos/instagram.svg" />
      <img src="../assets/Logos/linkedin.svg" />
      <img src="../assets/Logos/twitter.svg" />
      <img src="../assets/Logos/youtube.svg" />
      <hr />
    </nav>
    <div class="NJIT_logos">
      <img id="acm-logo" src="../assets/Logos/acm_logo_white.svg" />
      <img id="ywcc-logo" src="../assets/Logos/NJIT_white.png" />
    </div>
    <aside>
      <hr />
      <a href="acm.njit.edu" target="_blank">Main ACM Site</a>
      <hr />
      <a href="https://www.acm.org/" target="_blank">ACM National</a>
      <hr />
      <a href="https://computing.njit.edu/" target="_blank"
        >NJIT Ying Wu College of Computing</a
      >
      <hr />
      <a href="acm.njit.edu" target="_blank">Sponsor Packet</a>
      <hr />
      <a href="mailto:hacknjit@njit.edu" target="_blank">Email Us</a>
      <hr />
      <a
        href="https://static.mlh.io/docs/mlh-code-of-conduct.pdf"
        target="_blank"
        >MLH Code of Conduct</a
      >
      <hr />
    </aside>
  </footer>
</template>
  
  <script>
export default {};
</script>
  
  <style scoped>
footer {
  /* background: var(--edge-colors); */
  color: var(--text-color);
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 2rem;
  padding-bottom: 2rem;
}
nav,
aside {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 1rem;
  padding: 0.75rem 0;
}
aside {
  flex-wrap: wrap;
}
div.line {
  display: flex;
  align-content: center;
}
nav hr {
  border: none;
  border-top: 4px solid white;
  border-radius: 1000rem;
  /* border-top: 3px solid black;
  border-bottom: 3px solid black; */
  width: 25%;
  height: 1px;
  align-self: center;
}
.NJIT_logos {
  display: flex;
  align-content: center;
  justify-content: center;
  gap: 0.5rem;
}
#acm-logo {
  width: 100px;
  align-self: center;
}
#ywcc-logo {
  width: 200px;
}
nav img {
  width: 50px;
  height: 50px;
}
a {
  color: inherit;
}
@media (max-width: 600px) {
  nav img {
    width: 40px;
    height: 40px;
  }
  aside {
    font-size: 0.75rem;
  }
}
@media (max-width: 500px) {
  nav img {
    width: 35px;
  }
}
</style>