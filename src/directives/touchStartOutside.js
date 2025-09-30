export default {
    beforeMount(el, binding) {
        el.clickOutsideEvent = function (event) {
            // Check if the clicked element is outside the component's element
            if (!(el === event.target || el.contains(event.target))) {
                // Call the provided handler function
                binding.value(event);
            }
        };
        document.addEventListener('touchstart', el.clickOutsideEvent);
    },
    unmounted(el) {
        document.removeEventListener('touchstart', el.clickOutsideEvent);
    },
};