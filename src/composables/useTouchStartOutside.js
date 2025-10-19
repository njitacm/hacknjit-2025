import { onMounted, onUnmounted } from 'vue';

export function useTouchStartOutside(elementRef, callback) {
    const handleClickOutside = (event) => {
        if (elementRef.value && !(elementRef.value === event.target || elementRef.value.contains(event.target))) {
            callback(event);
        }
    };

    onMounted(() => {
        document.addEventListener('touchstart', handleClickOutside);
    });

    onUnmounted(() => {
        document.removeEventListener('touchstart', handleClickOutside);
    });
}