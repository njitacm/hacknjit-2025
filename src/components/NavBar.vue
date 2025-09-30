<template>
  <header v-on="mouseListeners" @touchend.stop.passive="isTouch ? onTouch() : null" ref="header"
    :class="{ active: isNavActive }">
    <nav :style="{ width: navSize.width, height: navSize.height }">
      <ul ref="ul">
        <!-- visible even when nav is active -->
        <li :style="getItemStyle(0)">
          <!-- touch devices when nav is closed: disbable the link by turning it into a span -->
          <RouterLink to="/" class="principal nav-link" :class="{ disabled: isTouch && !isNavActive }">
            <span>{{ isNavActive ? "Home" : activeSectionId }}</span>
            <img :src="downArrow" class="icon" />
          </RouterLink>
        </li>
        <!-- hidden when nav is closed -->
        <li v-for="(item, index) in navItems" :key="index" :style="getItemStyle(index + 1)">
          <RouterLink :to="{ path: item.path ?? '/', hash: item.hash ?? '' }" class="nav-link"
            :class="{ disabled: isTouch && !isNavActive }">
            {{ item.label }}
          </RouterLink>
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
import downArrow from "../assets/icons/down_arrow.svg";

// composables and stores
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

// vars
const heightEmFactor = 4;             // em
const screenSizeThreshold = 800;      // px - manually adjust based on number of nav items
const navItems = [
  // required: label, optional: hash, path
  { label: 'Sponsors', hash: '#Sponsors' },
  { label: 'FAQ', hash: '#FAQ' },
  { label: 'Contact', hash: '#Contact' },
  { label: 'Register', path: '/registration' },
];

const navSizes = {
  // small screens
  small: {
    shrunk: {
      width: "250px",
      height: `${heightEmFactor}em`,

    },
    expanded: {
      width: "calc(100vw - 16px)",         // matches the nav's padding/top gap of 8px
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

const onTouch = () => {
  if (interactedWithNav.value === false) {
    interactedWithNav.value = true;
  }
};

const onScroll = () => {
  if (window.scrollY > 50) {
    scrollLock.value = false;
  } else {
    scrollLock.value = true;
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

// computed properties
const isNavActive = computed(() => (
  (screenSize.value === "large" && (scrollLock.value === true || interactedWithNav.value === true)))
  ||
  (screenSize.value === "small" && interactedWithNav.value === true)
);
const navSize = computed(() => (isNavActive.value ? navSizes[screenSize.value].expanded : navSizes[screenSize.value].shrunk));

const mouseListeners = computed(() => {
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

// Style for each list item
function getItemStyle(index) {
  if (isNavActive.value) {
    return {
      width: "auto",
      height: "auto",
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
      obj.transform = 'translateY(-10px)';

      if (screenSize.value === "large") {
        obj.width = "0px";
      } else {
        obj.height = "0px";
      }
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
  line-height: 1em;
  border-radius: 2lh;
  overflow: hidden;
  transition: width 300ms ease, height 300ms ease;
}

ul {
  height: 100%;
  display: flex;
  justify-content: center;
  padding: 0;
  margin: 0;
  line-height: inherit;
  list-style: none;
}

li {
  flex-grow: 1;
  line-height: inherit;
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

.disabled {
  touch-action: none;
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
