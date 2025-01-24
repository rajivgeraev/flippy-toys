// src/components/pages/toys/ToySwiperPage.tsx
'use client';
import { useEffect, useState } from 'react';
import { getToys } from '@/lib/api/toys';
import Image from 'next/image';

interface Props {
    childId: string;
}

interface Toy {
    id: string;
    title: string;
    description: string;
    photos: Array<{
        url: string;
        is_main: boolean;
    }>;
}

export function ToySwiperPage({ childId }: Props) {
    const [toys, setToys] = useState<Toy[]>([]);
    const [currentIndex, setCurrentIndex] = useState(0);
    const [currentPhotoIndex, setCurrentPhotoIndex] = useState(0);

    useEffect(() => {
        const loadToys = async () => {
            try {
                // Используем childId при загрузке игрушек
                console.log('Loading toys for child:', childId);
                const toysData = await getToys();
                setToys(toysData);
            } catch (error) {
                console.error('Failed to load toys:', error);
            }
        };
        loadToys();
    }, [childId]); // Добавляем childId в зависимости


    const handlePhotoChange = (index: number) => {
        setCurrentPhotoIndex(index);
    };

    const handleSwipe = (direction: 'left' | 'right') => {
        if (direction === 'right' && currentIndex < toys.length - 1) {
            setCurrentIndex(prev => prev + 1);
        } else if (direction === 'left' && currentIndex > 0) {
            setCurrentIndex(prev => prev - 1);
        }
    };

    const currentToy = toys[currentIndex];

    return (
        <div className="min-h-screen bg-gradient-to-tr from-orange-400 via-red-400 to-pink-400">
            {currentToy && (
                <div className="min-h-screen flex flex-col">
                    <div className="flex-1 relative">
                        <div className="absolute inset-6 bg-white/20 backdrop-blur-xl rounded-[3rem] p-3">
                            <div className="relative h-full rounded-[2rem] overflow-hidden bg-white/30 backdrop-blur">
                                <div className="relative w-full h-full">
                                    <Image
                                        src={currentToy.photos[currentPhotoIndex]?.url}
                                        alt={currentToy.title}
                                        fill
                                        className="object-cover"
                                    />
                                </div>

                                <div className="absolute left-0 right-0 bottom-0 p-4 bg-gradient-to-t from-black/30 to-transparent backdrop-blur-sm">
                                    <div className="flex justify-center gap-3">
                                        {currentToy.photos.map((photo, index) => (
                                            <button
                                                key={index}
                                                onClick={() => handlePhotoChange(index)}
                                                className={`relative w-14 h-14 rounded-2xl overflow-hidden transition-transform
                         ${index === currentPhotoIndex ? 'ring-4 ring-white/60 scale-110' : 'ring-2 ring-white/30 hover:ring-white/60 hover:scale-105'}`}
                                            >
                                                <Image
                                                    src={photo.url}
                                                    alt={`View ${index + 1}`}
                                                    fill
                                                    className="object-cover"
                                                />
                                            </button>
                                        ))}
                                    </div>
                                </div>

                                <div className="absolute top-4 left-4 right-4 flex justify-between items-center">
                                    <div className="bg-white/30 backdrop-blur-xl px-4 py-2 rounded-2xl">
                                        <span className="text-white font-bold">{currentToy.title}</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="p-6 pb-10">
                        <div className="bg-white/20 backdrop-blur-xl rounded-3xl p-6">
                            <div className="flex justify-around">
                                {[
                                    { emoji: '👎', text: 'Не хочу' },
                                    { emoji: '⭐', text: 'Избранное' },
                                    { emoji: '👍', text: 'Хочу' }
                                ].map((btn, i) => (
                                    <button
                                        key={i}
                                        className="group"
                                        onClick={() => handleSwipe(i === 0 ? 'left' : 'right')}
                                    >
                                        <div className="w-20 h-20 bg-white/20 backdrop-blur rounded-2xl flex items-center justify-center mb-2 transform group-hover:scale-110 transition-all">
                                            <span className="text-4xl group-hover:animate-bounce">{btn.emoji}</span>
                                        </div>
                                        <div className="text-white text-center text-sm font-bold">{btn.text}</div>
                                    </button>
                                ))}
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
}