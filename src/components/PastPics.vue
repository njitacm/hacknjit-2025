<template>
  <div class="PastPics">
    <div class="page-side-padding container">
      <img src="../assets/PastPictures/hacknjit2024_1.jpg" />
      <img src="../assets/PastPictures/hacknjit2024_2.jpg" id="two" />
      <img src="../assets/PastPictures/hacknjit2024_3.jpg" />
      <img src="../assets/PastPictures/hacknjit2024_4.jpg" />
    </div>
  </div>
</template>

<style scoped>
.container {
  gap: 16px;
  display: grid;
  grid-template-columns: auto auto;
  place-content: center;
  /* max-width: var(--max-content-width); */

  @media(max-width: 600px) {
    & {
      grid-template-columns: 100%;
    }
  }
}

img {
  width: 100%;
  height: auto;
  border-radius: 0.5rem;
  display: block;
  aspect-ratio: 16/9;
  object-fit: cover;
}

#two {
  /* adjust the numbers to crop */
  object-position: 100% 40%;
}
</style>