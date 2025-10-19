// composables/useIsTouch.js

import { ref, onMounted, onUnmounted } from 'vue';

/**
 * A Vue composable that detects if the user's primary input mechanism is a touch screen.
 * It uses the CSS media queries for hover capability and pointer accuracy.
 *
 * @returns {object} An object containing a reactive boolean `isTouch`.
 */
export function useIsTouch() {
    const isTouch = ref(false);

    // The media query to check for non-touch devices (those with hover and a fine pointer).
    const mediaQuery = window.matchMedia('(hover: hover) and (pointer: fine)');

    // This handler updates the reactive ref based on whether the media query matches.
    // If it matches, it's NOT a touch device.
    const updateIsTouch = (event) => {
        isTouch.value = !event.matches;
    };

    // onMounted is used to ensure this code only runs on the client,
    // where `window.matchMedia` is available.
    onMounted(() => {
        // Set the initial value
        isTouch.value = !mediaQuery.matches;

        // Listen for changes (e.g., user plugs in a mouse to a tablet)
        mediaQuery.addEventListener('change', updateIsTouch);
    });

    // Clean up the event listener when the component is unmounted to prevent memory leaks.
    onUnmounted(() => {
        mediaQuery.removeEventListener('change', updateIsTouch);
    });

    return { isTouch };
}