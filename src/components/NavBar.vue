<!-- values that need to be manually depending on nav size adjusted are denoted with `TOCHANGE` -->
<template>
  <header v-on="mouseEventListeners" ref="header"
    :class="{ active: isNavActive }">
    <nav :style="{ width: navSize.width, height: navSize.height }" @touchmove.passive="onTouchMove"
      @touchend.passive="onTouchEnd">
      <ul ref="ul" :style="{ height: navSize.height }">
        <!-- visible even when nav is active -->
        <li :style="getItemStyle(0)">
          <!-- touch devices when nav is closed: disbable the link by turning it into a span -->
          <div class="section-indicator">
            <span>{{ getSectionHeader() }}</span>
            <img :src="downArrow" class="icon" />
          </div>
        </li>
        <!-- hidden when nav is closed -->
        <li v-for="(item, index) in navItems" :key="index" :style="getItemStyle(index + 1)">
          <button tabindex="-1" class="nav-link" @click="!isTouch ? goToPage(item) : null"
            @touchend.passive="() => !scrolling && goToPage(item)">
            {{ item.label }}
          </button>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, useTemplateRef } from 'vue';
import { storeToRefs } from "pinia";
import { useNavigationStore } from '../stores/navigation';
import { useIsTouch } from '../composables/useIsTouch';
import { useTouchStartOutside } from '../composables/useTouchStartOutside';
import { useRouter, useRoute } from 'vue-router';
import downArrow from "../assets/icons/down_arrow.svg";

// composables and stores
const router = useRouter();
const { isTouch } = useIsTouch();

const navigationStore = useNavigationStore();
const { activeSectionId } = storeToRefs(navigationStore);

// template refs
const headerRef = useTemplateRef("header");
const ulRef = useTemplateRef("ul");

// refs
const interactedWithNav = ref(false); // mouse enter, mouse leave, touch start affects nav open/closed
const scrollLock = ref(true);         // if true, will force the nav bar to be open
const screenSize = ref("large");      // small or large depending on screen size
const currentPage = ref("Home");
const scrolling = ref(false);         // if user is scrolling the nav on touch screens

// vars
const heightEmFactor = 4;             // em
const screenSizeThreshold = 925;      // px - TOCHANGE - manually adjust based on number of nav items
let wentToPage = false;               // a "lock" - ensures the nav closes when touch user taps on nav link
const navItems = [
  // required: label, optional: hash, path
  { label: "Home" },
  { label: 'About', hash: '#About' },
  { label: 'Sponsors', hash: '#Sponsors' },
  { label: "FAQ", hash: '#FAQ' },
  // { label: 'Contact', hash: '#Contact' },
  { label: "Register", path: "/registration" },
];

const navSizes = {
  // small screens
  small: {
    shrunk: {
      width: "250px",
      height: `${heightEmFactor}em`,

    },
    expanded: {
      width: "var(--expanded-touch-width)",         // matches the nav's padding/top gap of 8px
      height: `${(navItems.length + 1) * heightEmFactor}em`,
    },
  },
  // large screens
  large: {
    shrunk: {
      width: "250px",
      height: `${heightEmFactor}em`,
    },
    expanded: {
      width: "750px",
      height: `${heightEmFactor}em`,
    },
  }
};

// handlers
const onTouchStartOutside = () => {
  interactedWithNav.value = false;
  // Perform actions when touch starts outside
};

const onMouseEnter = () => {
  interactedWithNav.value = true;
};

const onMouseLeave = () => {
  interactedWithNav.value = false;
};

const onScroll = () => {
  // scroll lock functionality (force nav open if at top)
  if (window.scrollY > 50) {
    scrollLock.value = false;
  } else {
    scrollLock.value = true;
  }
};

const onTouchMove = () => {
  if (isTouch.value === true && scrolling.value === false) {
    scrolling.value = true;
  }
};

const onTouchEnd = () => {
  if (isTouch.value === true && scrolling.value === true) {
    scrolling.value = false;
  }

  if (interactedWithNav.value === false && isNavActive.value === false && !wentToPage) {
    interactedWithNav.value = true;
  } else if (wentToPage) {
    wentToPage = false;
  }
};

const onMediaQueryChange = (e) => {
  const isSmallScreen = e.matches;

  if (isSmallScreen) {
    screenSize.value = "small";
    ulRef.value.style.flexDirection = "column";
  } else {
    screenSize.value = "large";
    ulRef.value.style.flexDirection = "row";
  }
};

router.beforeEach((to) => {
  if (to.path === "/registration") {
    currentPage.value = "Registration";
  } else if (to.path === "/") {
    currentPage.value = "Home";
  }
});

// computed properties
const isNavActive = computed(() => (
  ((screenSize.value === "large" && (scrollLock.value === true || interactedWithNav.value === true)))
  ||
  (screenSize.value === "small" && interactedWithNav.value === true)
));
const navSize = computed(() => (isNavActive.value ? navSizes[screenSize.value].expanded : navSizes[screenSize.value].shrunk));

const mouseEventListeners = computed(() => {
  // If it's a touch device, return an empty object. The @touch.passive will be used
  if (isTouch.value === true) {
    return {};
  }

  // If it's not a touch device, return an object with the event handlers.
  return {
    mouseenter: onMouseEnter,
    mouseleave: onMouseLeave,
  };
});

// query for for large vs small screens
const mqList = window.matchMedia(`(max-width: ${screenSizeThreshold}px)`);

// life cycle hooks
onMounted(() => {
  if (mqList.matches) {
    onMediaQueryChange({ matches: true });
  }
  window.addEventListener("scroll", onScroll);
  mqList.addEventListener("change", onMediaQueryChange);
});

onUnmounted(() => {
  window.removeEventListener("scroll", onScroll);
  mqList.removeEventListener("change", onMediaQueryChange);
});

// link touch start outside handler to the composable
useTouchStartOutside(headerRef, onTouchStartOutside);

function getSectionHeader() {
  if (currentPage.value === "Home") {
    return activeSectionId.value.replace("-", " ");
  }
  return currentPage.value;
}

function goToPage(item) {
  if (isNavActive.value === true) {
    router.push({ path: item.path ?? '/', hash: item.hash ?? '' });

    if (isTouch.value === true) {
      wentToPage = true;
      interactedWithNav.value = false;
    }
  }
}

// style for each list item
function getItemStyle(index) {
  if (isNavActive.value === true) {
    const obj = {
      opacity: 1,
      transform: 'translateY(0)',
    };

    if (window.matchMedia('(prefers-reduced-motion: no-preference)').matches) {
      obj.transition = `opacity 0.3s ease ${(index + 1) * 50}ms, transform 0.3s ease ${(index + 1) * 50}ms`;
    }

    // section header that shows even with nav bar closed
    if (index === 0) {
      obj.flexGrow = 0;

      if (screenSize.value === "large") {
        obj.width = "0px";
        obj.height = "0px";
      } else {
        obj.width = "0px";
        obj.height = "0px";
      }
    } else {
      obj.flexGrow = 1;
      obj.width = "auto";
      obj.height = "auto";
    }

    return obj;
  } else {
    const obj = {};

    // section header that shows even with nav bar closed
    if (index === 0) {
      obj.width = "auto";
      obj.height = "100%";

    } else {
      obj.flexGrow = 0;
      obj.opacity = 0;
      obj.transform = 'translateY(-10px)';

      if (screenSize.value === "large") {
        obj.width = "0px";
        obj.height = "0px";
      } else {
        obj.width = "0px";
        obj.height = "0px";
      }
    }

    return obj;
  }
}
</script>

<style scoped>
header {
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  z-index: 50;
  padding: 8px;
  border-radius: 1000px;
  width: fit-content;
  max-height: 50svh;
}

header.active {
  width: 100%;
}

nav {
  --expanded-touch-width: calc(90vw - 50px - 16px);
  --border-width: 1px;
  border: var(--border-width) solid #ffffff33;
  background-color: #00250475;
  backdrop-filter: blur(25px);
  align-content: center;
  position: relative;
  line-height: 1em;
  border-radius: 2lh;
  margin-inline: auto;
  max-height: 100svh;
  overflow: auto;

  @media(prefers-reduced-transparency: reduce) {
    & {
      background-color: #002504;
      backdrop-filter: none;
    }
  }
}

ul {
  height: 100%;
  display: flex;
  justify-content: center;
  padding: 0;
  margin: 0;
  line-height: inherit;
  list-style: none;
  overflow: hidden;
}

li {
  flex-grow: 1;
  line-height: inherit;
  font-size: 1.5em;
}

.nav-link {
  color: white;
  font-size: 1em;
  cursor: pointer;
  background: transparent;
  border: none;
  width: 100%;
  height: 100%;
  align-content: center;
  display: block;
}

.disabled {
  touch-action: none;
}

.section-indicator {
  text-align: center;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
}

header:not(.active) .section-indicator {
  padding: 0 24px;
}

.section-indicator span {
  flex-grow: 1;
  font-size: 1em;
}

.icon {
  width: 20px;
  height: 20px;
  opacity: 0.5;
}

header.active .section-indicator span {
  flex-grow: 0;
  font-size: 0;
}

header.active .icon {
  width: 0px;
}

@media(prefers-reduced-motion: no-preference) {
  nav {
    transition: width 375ms ease, height 375ms ease;
  }
}

@media(prefers-reduced-motion: no-preference) and (hover: hover) and (pointer: fine) {
  .nav-link {
    transition: background-color 250ms ease;
  }
}

/* move nav toward the left to avoid collision with MLH banner */
/* TOCHANGE - adjust max-width depending on nav bar overlap with MLH banner */
@media(max-width: 1100px) {
  header {
    left: 0;
    transform: none;
  }

  nav {
    margin-inline: 0;
  }
}

@media(max-width: 500px) {
  nav {
    /* accounts for padding on sides and in the middle (along the MLH banner) */
    --expanded-touch-width: calc(90vw - 25px - 24px);
  }
}

/* desktop */
@media(hover: hover) and (pointer: fine) {
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
