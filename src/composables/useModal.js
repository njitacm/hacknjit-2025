import { ref, readonly, shallowRef, markRaw } from "vue";

const modalStack = ref([]);

export function useModal() {
    const openModal = (options) => {
        // make ids more reliable
        document.querySelector('body').classList.add('modal-open');
        const id = Date.now() + Math.random();
        
        modalStack.value.push({
            id,
            title: options.title,
            component: markRaw(options.component),
            props: options.props || {},
        });
    };
    
    const closeModal = () => {
        if (modalStack.value.length > 0) {
            modalStack.value.pop();
            
            if (modalStack.value.length === 0) {
                document.querySelector('body').classList.remove('modal-open');
            }
        }
    };

    return {
        modalStack: readonly(modalStack),
        openModal,
        closeModal,
    };
}