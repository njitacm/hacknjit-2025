import { ref, readonly, shallowRef } from "vue";

const modalStack = ref([]);

export function useModal() {
    const openModal = (options) => {
        // make ids more reliable
        const id = Date.now() + Math.random();

        modalStack.value.push({
            id,
            title: options.title,
            component: shallowRef(options.component),
            props: options.props || {},
        });
    };

    const closeModal = () => {
        if (modalStack.value.length > 0) {
            modalStack.value.pop();
        }
    };

    return {
        modalStack: readonly(modalStack),
        openModal,
        closeModal,
    };
}