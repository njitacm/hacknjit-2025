<template>
  <div class="outer-container">
    <h1>Welcome to HackNJIT!</h1>
    <main>
      <img src="../assets/Slides/hacknjit_1.jpg" />
      <p>
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing. Students break into small teams of up to 4,
        work together over 24 hours to create a tech project, and at the end
        present it to our judges!
      </p>
    </main>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.outer-container {
  width: 60%;
  margin: 0rem auto;
  margin-bottom: 5rem;
  /* color: var(--color3); */
}
h1 {
  font-size: 4rem;
  margin-bottom: 1.75rem;
  font-weight: bold;
}
main {
  display: flex;
  gap: 2rem;
  justify-content: space-between;
  align-content: center;
}
p {
  flex: 1;
  font-size: 2rem;
  height: fit-content;
  align-self: center;
  font-weight: bold;
  text-align: left;
}
img {
  max-width: 50%;
  /* border: 4px var(--color3) solid; */
  border: 4px white solid;
  border-radius: 30px;
  flex: 1;
}
</style>