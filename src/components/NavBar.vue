<template>
  <header @mouseenter="isNavActive = true" @mouseleave="isNavActive = false" :style="{
    width: isNavActive ? '100%' : 'fit-content',
  }">
    <nav :style="{ width: `${navWidth}` }">
      <ul>
        <li :style="getItemStyle(0)">
          <RouterLink to="/" class="nav-link">
            {{ isNavActive ? "Home" : "HackNJIT" }}
          </RouterLink>
        </li>
        <li v-for="(item, index) in navItems" :key="index" :style="getItemStyle(index + 1)">
          <a :href="item.href" class="nav-link">{{ item.label }}</a>
        </li>
        <li :style="getItemStyle(navItems.length)">
          <RouterLink to="/registration" class="nav-link">
            Register
          </RouterLink>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script setup>
import { ref, computed } from 'vue'

const isNavActive = ref(true)
const isHoveredNav = ref(false)

const navItems = [
  { label: 'Sponsors', href: '#Sponsors' },
  { label: 'FAQ', href: '#About' },
  { label: 'Contact', href: '#Contact' }
];

// Sizes
const navWidths = {
  shrunk: "175px",
  expanded: "750px",
};

// Reactive widths
const navWidth = computed(() => (isNavActive.value || isHoveredNav.value ? navWidths.expanded : navWidths.shrunk));

// Style for each list item
function getItemStyle(index) {
  if (isNavActive.value) {
    return {
      width: `auto`,
      flexGrow: 1,
      opacity: 1,
      transform: 'translateY(0)',
      transition: `opacity 0.3s ease ${(index + 1) * 0.1}s, transform 0.3s ease ${(index + 1) * 0.1}s`
    }
  } else {
    const obj = {
      transition: 'opacity 0s, transform 0s'
    };

    // index 0 ("HackNJIT") is not hidden
    if (index > 0) {
      obj.flexGrow = 0;
      obj.opacity = 0;
      obj.width = "0px";
      obj.transform = 'translateY(-10px)';
    }

    return obj;
  }
}
</script>

<style scoped>
header {
  justify-self: center;
  position: fixed;
  z-index: 50;
  display: grid;
  justify-items: center;
  transition: gap 250ms ease;
  padding: 8px 0;
  border-radius: 1000px;
}

nav {
  background-color: #00250475;
  backdrop-filter: blur(25px);
  align-content: center;
}

nav {
  position: relative;
  height: 4em;
  border-radius: 1000px;
  overflow: hidden;
  transition: width 300ms ease;
}

ul {
  height: 100%;
  display: flex;
  justify-content: center;
  padding: 0;
  margin: 0;
  list-style: none;
}

li {
  flex-grow: 1;
  line-height: 1.5em;
  font-size: 1.5em;
  display: block;
}

.nav-link {
  color: white;
  text-decoration: none;
  width: 100%;
  height: 100%;
  align-content: center;
  display: block;
}

/* desktop */
@media(hover: hover) and (pointer: fine) {
  .nav-link {
    transition: background-color 250ms ease;

  }

  .nav-link:hover {
    background-color: #ffffff10
  }
}

/* mobile */
@media(pointer: coarse) {
  .nav-link:active {
    background-color: #ffffff10
  }
}
</style>
