<template>
  <footer>
    <p>We Abide by the MLH Code of Conduct</p>
    <p>HackNJIT is hosted by</p>
    <div class="footer-images">
      <a href="http://njit.acm.org/" target="_blank"
        ><img id="acm-logo" src="../assets/Footer/ACM-LOGO.png"
      /></a>
      <a href="https://computing.njit.edu/" target="_blank"
        ><img id="ywcc-logo" src="../assets/Footer/ywcc_logo.png"
      /></a>
    </div>
  </footer>
</template>

<script>
export default {};
</script>

<style scoped>
footer {
  background: #1e262a;
  color: rgb(170, 170, 170);
  display: flex;
  flex-direction: column;
  padding: 16px 0;
}
p:first-of-type {
  text-decoration: underline;
  margin-top: 0.5rem;
  margin-bottom: 1rem;
}
.footer-images {
  display: flex;
  align-content: center;
  justify-content: center;
}
.footer-images a {
  align-content: center;
}
#acm-logo {
  height: 100px;
}
#ywcc-logo {
  height: 60px;
}
</style>