<template>
  <div class="PastPics" ref="past-pics">
    <img src="../assets/PastPictures/hacknjit2024_1.jpg" />
    <img src="../assets/PastPictures/hacknjit2024_2.jpg" id="two" />
    <img src="../assets/PastPictures/hacknjit2024_3.jpg" />
    <img src="../assets/PastPictures/hacknjit2024_4.jpg" />
  </div>
</template>

<style scoped>
.PastPics {
  margin-top: 32px;
  gap: 16px;
  display: grid;
  grid-template-columns: 50% 50%;
  place-content: center;
  height: auto;
  
  @media(max-width: 600px) {
    & {
      grid-template-columns: 100%;
    }
  }
}

img {
  border-radius: 0.5rem;
  /* display: block; */
  top: 0;
  aspect-ratio: 16/9;
  object-fit: cover;
}

#two {
  /* adjust the numbers to crop */
  object-position: 100% 40%;
}
</style>