<template>
  <div class="outer-container">
    <h1>Meet The Team!</h1>
    <main class="team">
      <div class="member">
        <img src="../assets/eboard/president.png" />
        <h2>Will Simkins</h2>
        <h3>ACM President</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/vice_president.png" />
        <h2>Anthony Siracusa</h2>
        <h3>ACM Vice-President</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/treasurer.png" />
        <h2>Logan Kriebal</h2>
        <h3>ACM Treasurer</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/event_master.png" />
        <h2>Priya Patel</h2>
        <h3>Event Master</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/graphic_designer.png" />
        <h2>Kapila Mane</h2>
        <h3>ACM Graphic Designer</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/webmaster.png" />
        <h2>Abdullah Imran</h2>
        <h3>ACM Webmaster</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/secretary.png" />
        <h2>Ananya Tyagi</h2>
        <h3>ACM Secretary</h3>
      </div>
      <div class="member">
        <img src="../assets/eboard/public_relations.png" />
        <h2>Darren Bonifacio</h2>
        <h3>ACM Public Relations</h3>
      </div>
    </main>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
div.outer-container {
  width: 70%;
  margin-top: 2rem;
  margin-left: auto;
  margin-right: auto;
  padding: 1.5rem;
}
h1 {
  font-size: 4rem;
  margin-bottom: 1.75rem;
}
main.team {
  display: flex;
  justify-content: center;
  gap: 1rem;
  flex-wrap: wrap;
}
.member {
  flex: 1;
  margin-top: 1rem;
}
.member img {
  border-radius: 2rem;
  border: 4px white solid;
  max-width: 300px;
  aspect-ratio: 1 / 1;
}
.member h2 {
  margin: 0.5rem;
}
</style>