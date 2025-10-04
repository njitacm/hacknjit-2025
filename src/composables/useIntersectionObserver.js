// composables/useIntersectionObserver.js
import { useNavigationStore } from '../stores/navigation';

// This 'observer' variable is created only once, when the module is imported.
// It will be shared by every component that uses this composable.
let observer;

// The callback function that the IntersectionObserver will execute.
const handleIntersection = (entries) => {
    const navigationStore = useNavigationStore();

    entries.forEach(entry => {
        // When a section is intersecting, call the Pinia action.
        if (entry.isIntersecting) {
            navigationStore.setActiveSection(entry.target.id);
        }
    });
};

// This is the exported composable function.
export function useIntersectionObserver() {

    // This function initializes the observer if it hasn't been already.
    // It's "lazy-loaded" — it only runs when the first component needs it.
    const initObserver = () => {
        if (observer) return; // Don't create it more than once.

        const options = {
            root: null, // relative to the viewport
            rootMargin: "0% 0px -90% 0px",          // place tripwire 10% down from the top
            threshold: 0, // trigger when 50% of the element is visible
        };

        observer = new IntersectionObserver(handleIntersection, options);
    };

    // A function to add a target element to the observer.
    const observe = (element) => {
        initObserver(); // Ensure the observer instance exists.
        if (element) {
            observer.observe(element);
        }
    };

    // A function to remove a target element from the observer (for cleanup).
    const unobserve = (element) => {
        if (observer && element) {
            observer.unobserve(element);
        }
    };

    // Return the functions so components can use them.
    return {
        observe,
        unobserve,
    };
}