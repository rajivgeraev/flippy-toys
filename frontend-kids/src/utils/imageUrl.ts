// src/utils/imageUrl.ts
export function getOptimizedImageUrl(url: string, type: 'main' | 'thumbnail'): string {
    if (!url.includes('cloudinary.com')) return url;

    const transformations = {
        main: 'f_auto,q_auto,w_800,h_800,c_fill',
        thumbnail: 'f_auto,q_auto,w_112,h_112,c_fill'
    };

    return url.replace('/upload/', `/upload/${transformations[type]}/`);
}