<template>
  <header v-on="mouseListeners" @touchstart.passive="isTouch ? onTouch() : null" ref="nav-bar"
    :class="{ active: isNavActive }">
    <nav :style="{ width: `${navWidth}` }">
      <ul>
        <!-- Visible even when nav is active -->
        <li :style="getItemStyle(0)">
          <RouterLink to="/" class="principal nav-link">
            <span>{{ isNavActive ? "Home" : activeSectionId }}</span>
            <img :src="downArrow" class="icon" />
          </RouterLink>
        </li>
        <!-- Hidden when nav is unactive -->
        <li v-for="(item, index) in navItems" :key="index" :style="getItemStyle(index + 1)">
          <RouterLink :to="{ path: item.path ?? '/', hash: item.hash ?? '' }" class="nav-link">{{ item.label }}</RouterLink>
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
import downArrow from "../assets/icons/down_arrow.svg";

// composables and stores
const navBarRef = useTemplateRef("nav-bar");
const { isTouch } = useIsTouch();

const navigationStore = useNavigationStore();
const { activeSectionId } = storeToRefs(navigationStore);

// refs
const isHoveredNav = ref(false);
const scrollLock = ref(true);     // if true, will force the nav bar open

// vars
const navItems = [
  // required: label, optional: hash, path
  { label: 'Sponsors', hash: '#Sponsors' },
  { label: 'FAQ', hash: '#FAQ' },
  { label: 'Contact', hash: '#Contact' },
  { label: 'Register', path: '/registration' },
];

const navWidths = {
  shrunk: "250px",
  expanded: "750px",
};

// handlers
const onTouchStartOutside = () => {
  console.log("touch start outside");
  isHoveredNav.value = false;
  // Perform actions when touch starts outside
};

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
  if (isHoveredNav.value === false) {
    isHoveredNav.value = true;
  }
};

const onScroll = () => {
  if (window.scrollY > 50) {
    scrollLock.value = false;
  } else {
    scrollLock.value = true;
  }
};

// computed properties
const isNavActive = computed(() => (scrollLock.value === true || isHoveredNav.value === true));
const navWidth = computed(() => (isNavActive.value ? navWidths.expanded : navWidths.shrunk));

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

// life cycle hooks
onMounted(() => {
  window.addEventListener("scroll", onScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", onScroll);
});

// link touch start outside handler to the composable
useTouchStartOutside(navBarRef, onTouchStartOutside);

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
  width: fit-content;
}

header.active {
  width: 100%;
}

nav {
  background-color: #00250475;
  backdrop-filter: blur(25px);
  align-content: center;
  border: 1px solid #ffffff33;
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

.nav-link.principal {
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
}

header:not(.active) .nav-link.principal {
  padding: 0 24px;
}

.nav-link.principal span {
  flex-grow: 1;
  font-size: 1em;
}

.icon {
  width: 20px;
  height: 20px;
  opacity: 0.5;
}

header.active .nav-link.principal span {
  flex-grow: 0;
}

header.active .icon {
  width: 0px;
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
