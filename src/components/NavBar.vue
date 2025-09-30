<template>
  <header class="header" @mouseenter="isTitlebarActive = true" @mouseleave="isTitlebarActive = false" :style="{
    width: isTitlebarActive ? '100%' : 'fit-content',
  }">

    <!-- Dropdown -->
    <nav class="dropdown" :style="{ width: `${dropdownWidth}` }" ref="nav">
      <ul ref="nav-links">
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
import { ref, computed, useTemplateRef, onMounted } from 'vue'

const isTitlebarActive = ref(true)
const isHoveredNav = ref(false)
const navLinks = useTemplateRef("nav-links");
const nav = useTemplateRef("nav");
const autoWidths = [];

const navItems = [
  { label: 'Sponsors', href: '#Sponsors' },
  { label: 'FAQ', href: '#About' },
  { label: 'Contact', href: '#Contact' }
];

onMounted(() => {
  const resizeObserver = new ResizeObserver(entries => {
    for (const entry of entries) {
      console.log("resized");
      const width = entry.contentRect.width;
      if (dropdownWidths.expanded === "auto" && dropdownWidths.shrunk === "auto") {
        console.log("both auto");
        dropdownWidths.expanded = width;
      } else if (dropdownWidths.shrunk === "auto") {
        console.log("shrunk auto");
        if (width <= dropdownWidths.expanded) {
          console.log("less than or equal to expanded");
          dropdownWidths.shrunk = width;
        } else {
          dropdownWidths.shrunk = dropdownWidths.expanded;
          console.log("greater than expanded");
          dropdownWidths.expanded = width;
        }
        dropdownWidths.expanded = `${dropdownWidths.expanded}px`;
        dropdownWidths.shrunk = `${dropdownWidths.shrunk}px`;
        resizeObserver.disconnect();
      } else {
        console.log("Unexpected case");
      }
      console.log("new widths freshly resized", dropdownWidths);
    }
  });
  resizeObserver.observe(nav.value);
  setTimeout(() => {
    isTitlebarActive.value = false;
    setTimeout(() => {
      isTitlebarActive.value = true;
    }, 0);
  }, 0);

  for (const child of navLinks.value.children) {
    autoWidths.push(child.offsetWidth);
  }
});

// Sizes
const dropdownWidths = {
  shrunk: "auto",
  expanded: "auto",
};

// Reactive widths
const dropdownWidth = computed(() => {
    console.log("new width at computed", dropdownWidths)
    return (isTitlebarActive.value || isHoveredNav.value ? dropdownWidths.expanded : dropdownWidths.shrunk)
});

// Style for each list item
function getItemStyle(index) {
  if (isTitlebarActive.value) {
    return {
      width: `${autoWidths[index]}px`,
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
      obj.width = "0px";
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
  transition: width 300ms ease-out;
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
