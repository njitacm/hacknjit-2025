import { defineStore } from 'pinia';

export const useNavigationStore = defineStore('navigation', {
    state: () => ({
        activeSectionId: 'HackNJIT', // Initial active section
    }),
    actions: {
        // An action to update the state
        setActiveSection(sectionId) {
            this.activeSectionId = sectionId;
        },
    },
});