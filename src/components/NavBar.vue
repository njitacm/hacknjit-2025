<template>
  <header v-on="mouseListeners" @touchstart.passive="isTouch ? onTouch() : null" ref="nav-bar"
    :style="{ width: isNavActive ? '100%' : 'fit-content' }">
    <nav :style="{ width: `${navWidth}` }">
      <ul>
        <li :style="getItemStyle(0)">
          <RouterLink to="/" class="nav-link">
            {{ isNavActive ? "Home" : activeSectionId }}
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
import { ref, computed, onMounted, onUnmounted, useTemplateRef } from 'vue'
import { storeToRefs } from "pinia";
import { useNavigationStore } from '../stores/navigation';
import { useIsTouch } from '../composables/useIsTouch';
import { useTouchStartOutside } from '../composables/useTouchStartOutside';

const navBarRef = useTemplateRef("nav-bar");

const handleOutsideTouch = () => {
  console.log("touch start outside");
  isHoveredNav.value = false;
  // Perform actions when touch starts outside
};

useTouchStartOutside(navBarRef, handleOutsideTouch);

const navigationStore = useNavigationStore();
const { activeSectionId } = storeToRefs(navigationStore);

const { isTouch } = useIsTouch();

const isHoveredNav = ref(false);
const scrollLock = ref(true);     // if true, will force the nav bar open

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
const isNavActive = computed(() => (scrollLock.value === true || isHoveredNav.value === true));
const navWidth = computed(() => (isNavActive.value ? navWidths.expanded : navWidths.shrunk));

const onMouseEnter = () => {
  console.log("enter");
  isHoveredNav.value = true;
};

const onMouseLeave = () => {
  console.log("leave");
  isHoveredNav.value = false;
};

const onTouch = () => {
  console.log("touch");
  isHoveredNav.value = !isHoveredNav.value;
};

const mouseListeners = computed(() => {
  // If it's a touch device, return an empty object. The @touch.passive will be used
  if (isTouch.value === true) {
    console.log("value true");
    return {};
  }

  // If it's not a touch device, return an object with the event handlers.
  return {
    mouseenter: onMouseEnter,
    mouseleave: onMouseLeave,
  };
});

const scrollHandler = () => {
  if (window.scrollY > 50) {
    scrollLock.value = false;
  } else {
    scrollLock.value = true;
  }
};

onMounted(() => {
  window.addEventListener("scroll", scrollHandler);
});

onUnmounted(() => {
  window.removeEventListener("scroll", scrollHandler);
});

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
