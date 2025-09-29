<template>
  <header class="header" @mouseenter="isTitlebarActive = true" @mouseleave="isTitlebarActive = false" :style="{
    width: isDropdownVisible ? `${dropdownWidth}` : 'fit-content',
    gap: isDropdownVisible ? '8px' : '0px'
  }">

    <!-- Title bar -->
    <a class="mainbar" :style="{ width: `${mainBarWidth}` }" href="#Title">
      HackNJIT
    </a>

    <!-- Dropdown -->
    <nav class="dropdown" :style="{
      opacity: isDropdownVisible ? 1 : 0,
      width: `${dropdownWidth}`,
      fontSize: isDropdownVisible ? `${dropdownFontSize}` : '0px',
    }">
      <ul>
        <li v-for="(item, index) in navItems" :key="index" :style="getItemStyle(index)">
          <a :href="item.href" class="nav-link">{{ item.label }}</a>
        </li>
        <li>
            <RouterLink to="/registration">Register</RouterLink>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script setup>
import { ref, computed } from 'vue'

const isTitlebarActive = ref(false)
const isHoveredNav = ref(false)

const navItems = [
  { label: 'Sponsors', href: '#Sponsors' },
  { label: 'FAQ', href: '#About' },
  { label: 'Contact', href: '#Contact' }
]

// Sizes
const mainBarClosedWith = "150px";
const mainBarOpenWith = "300px";
const dropdownWidthExpanded = "600px";
const dropdownFontSize = "1em";

// Reactive widths
const mainBarWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? mainBarOpenWith : mainBarClosedWith))
const dropdownWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? dropdownWidthExpanded : "500px"))
const isDropdownVisible = computed(() => isTitlebarActive.value || isHoveredNav.value)

// Style for each list item
function getItemStyle(index) {
  if (isDropdownVisible.value) {
    return {
      opacity: 1,
      transform: 'translateY(0)',
      transition: `opacity 0.3s ease ${(index + 1) * 0.1}s, transform 0.3s ease ${(index + 1) * 0.1}s`
    }
  } else {
    return {
      opacity: 0,
      transform: 'translateY(-10px)',
      transition: 'opacity 0s, transform 0s'
    }
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
  justify-content: space-around;
  padding: 0;
  margin: 0;
  list-style: none;
}

li {
  width: 100%;
  line-height: 1.5em;
  font-size: 1.5em;
  display: block;
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
