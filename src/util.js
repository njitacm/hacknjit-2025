const IMAGES = import.meta.glob(
    [
        // add as needed for more image extensions
        '/src/assets/**/*.jpg',
        '/src/assets/**/*.png',
        '/src/assets/**/*.svg',
    ],
    // The 'eager: true' option loads all images at once,
    // which might cause performance issues for many images.
    { eager: false, as: 'url' }
);

async function getImageUrl(name) {
    const imagePath = `/src/assets/${name}`;

    // Check if the image exists in our glob import
    if (!IMAGES[imagePath]) {
        throw new Error(`Image not found: ${name}`);
    }

    // Call the importer function. This returns a Promise.
    const importer = IMAGES[imagePath];
    const module = await importer();

    return module;
}

export { getImageUrl };