
<template>
  <header 
      class="header" 
    @mouseenter="isTitlebarActive = true" 
    @mouseleave="isTitlebarActive = false"
    :style="{ 
      width: isDropdownVisible ? `${dropdownWidth}px` : 'fit-content'
    }"
    >

    <!-- Title bar -->
    <span
      class="mainbar"
      :class="[ isTitlebarActive || isDropdownVisible ? 'active' : '' ]"
      :style="{ width: `${mainBarWidth}px` }"
    >
      <h1 class="text-lg"><a href="#Title">HackNJIT</a></h1>
    </span>

    <!-- Dropdown -->
    <nav
      class="dropdown"
      :style="{ 
        width: `${dropdownWidth}px`,
        top: isDropdownVisible ? `${dropdownTop}px` : '0px',
        height: isDropdownVisible ? `${dropdownHeight}px` : '0px'
      }"
    >
      <ul>
        <li
          v-for="(item, index) in navItems"
          :key="index"
          :style="getItemStyle(index)"
        >
          <a :href="item.href" class="nav-link">{{ item.label }}</a>
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
const mainBarDefault = 400
const mainBarShrunk = 200
const dropdownWidthExpanded = 750
const dropdownHeight = 75
const dropdownTop = 50

// Reactive widths
const mainBarWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? mainBarShrunk : mainBarDefault))
const dropdownWidth = computed(() => (isTitlebarActive.value || isHoveredNav.value ? dropdownWidthExpanded : 0))
const isDropdownVisible = computed(() => isTitlebarActive.value || isHoveredNav.value)

// Style for each list item
function getItemStyle(index) {
  if (isDropdownVisible.value) {
    return {
      opacity: 1,
      transform: 'translateY(0)',
      transition: `opacity 0.3s ease ${(index+1) * 0.15}s, transform 0.3s ease ${(index+1) * 0.15}s`
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
  margin-top: 8px;
  margin-left: auto;
  margin-right: auto;
  width: fit-content;
  position: fixed;
  top: 0;
  z-index: 50;
  display: flex;
  justify-content: center;
  justify-self: center;
}

.mainbar {
  height: 50px;
  background-color: black;
  cursor: pointer;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: width 0.3s ease, border-radius 0.3s ease;
}
.mainbar.active {
  border-radius: 16px 16px 0 0;
}

.mainbar h1 {
  font-size: 24px;
  color: white;
}

.dropdown {
  position: absolute;
  background-color: black;
  border-radius: 16px;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  transition: width 0.3s ease, top 0.3s ease, height 0.3s ease;
}

ul {
  display: flex;
  width: 100%;
  justify-content: space-around;
  padding: 0;
  margin: 0;
  list-style: none;
}
li {
  font-size: 18px;
}
li a {
  color: white;
  text-decoration: none;
}
nav {
  z-index: -50;
}
</style>

