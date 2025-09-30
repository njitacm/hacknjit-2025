<template>
  <header class="header" @mouseenter="isTitlebarActive = true" @mouseleave="isTitlebarActive = false" :style="{
    width: isTitlebarActive ? '100%' : 'fit-content',
  }">

    <!-- Title bar
    <RouterLink to="/" class="mainbar" :style="{ width: `${mainBarWidth}` }">
      HackNJIT
    </RouterLink> -->

    <!-- Dropdown -->
    <nav class="dropdown" :style="{
      // opacity: isDropdownVisible ? 1 : 0,
      width: `${dropdownWidth}`,
      // fontSize: isDropdownVisible ? `${dropdownFontSize}` : '0px',
    }">
      <ul>
        <li :style="getItemStyle(0)">
          <RouterLink to="/" class="nav-link">
            {{ isTitlebarActive ? "Home" : "HackNJIT" }}
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
// import HomeView from '../views/HomeView.vue';

const isTitlebarActive = ref(false)
const isHoveredNav = ref(false)

const navItems = [
  { label: 'Sponsors', href: '#Sponsors' },
  { label: 'FAQ', href: '#About' },
  { label: 'Contact', href: '#Contact' }
]

// Sizes
// const mainBarClosedWith = "150px";
// const mainBarOpenWith = "300px";
const dropdownWidthExpanded = "700px";
// const dropdownFontSize = "1em";

// Reactive widths
// const mainBarWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? mainBarOpenWith : mainBarClosedWith))
const dropdownWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? dropdownWidthExpanded : "300px"))
const isDropdownVisible = computed(() => isTitlebarActive.value || isHoveredNav.value)

// Style for each list item
function getItemStyle(index) {
  if (isTitlebarActive.value) {
    return {
      width: "100px",
      flexGrow: "1",
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
      obj.opacity = 0;
      obj.width = 0;
      obj.flexGrow = 0;
      obj.transform = 'translateY(-10px)';
    }

    return obj;
  }
}
</script>

<style scoped>
.header {
  justify-self: center;
  position: fixed;
  top: 8px;
  z-index: 50;
  display: grid;
  justify-items: center;
  transition: gap 250ms ease;
}

.mainbar,
.dropdown {
  background-color: #00250475;
  backdrop-filter: blur(25px);
  align-content: center;
}

.mainbar {
  font-weight: bold;
  padding: 0.5em;
  font-size: 1.5em;
  line-height: 1.5em;
  border-radius: 1lh;
  transition: width 250ms ease, border-radius 100ms ease;
}

.dropdown {
  position: relative;
  height: 4em;
  border-radius: 1000px;
  overflow: hidden;
  transition: width 300ms ease-out, height 300ms ease-out, opacity 250ms ease;
}

nav {
  z-index: -50;
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
  transition: width 300ms ease-out, flex-grow 300ms ease-out;
}

.mainbar,
.nav-link {
  color: white;
  text-decoration: none;
}

.nav-link {
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
