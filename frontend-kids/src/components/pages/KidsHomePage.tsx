"use client";

import { Child } from "@/types/child";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

interface KidsHomePageProps {
  child: Child;
}

export function KidsHomePage({ child }: KidsHomePageProps) {
  const [points] = useState(240);
  const router = useRouter();

  useEffect(() => {
    if (window.Telegram?.WebApp) {
      window.Telegram.WebApp.ready();
      window.Telegram.WebApp.expand();
    }
  }, []);

  const handleToyIslandClick = () => {
    router.push(`${child.id}/swipe`);
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-sky-400 via-blue-400 to-blue-500">
      {/* Floating Clouds */}
      <div className="fixed top-4 left-8 text-6xl animate-pulse">☁️</div>
      <div className="fixed top-12 right-12 text-5xl animate-pulse delay-300">
        ☁️
      </div>

      {/* Header with Character and Points */}
      <div className="pt-6 px-4">
        <div className="max-w-md mx-auto flex justify-between items-center">
          {/* Character Speech */}
          <div className="bg-white rounded-2xl p-3 shadow-lg relative">
            <div className="absolute -top-2 left-1/2 transform -translate-x-1/2 w-4 h-4 bg-white rotate-45" />
            <span className="text-xl">Hi, Victoria! 👋</span>
          </div>
          {/* Points Badge */}
          <div className="bg-yellow-400 rounded-xl px-4 py-2 shadow-lg">
            <div className="flex items-center gap-2">
              <span className="text-2xl">⭐</span>
              <span className="font-bold text-yellow-900">{points}</span>
            </div>
          </div>
        </div>
      </div>

      {/* Main Islands Grid */}
      <div className="px-4 pt-8">
        <div className="max-w-md mx-auto grid grid-cols-2 gap-6">
          {/* Treasure Island */}
          <div
            className="relative pb-[100%] cursor-pointer"
            onClick={handleToyIslandClick}
          >
            <div className="absolute inset-0">
              <div className="relative group h-full">
                <div className="absolute inset-0 bg-yellow-300 rounded-3xl blur opacity-50 group-hover:opacity-75 transition-opacity" />
                <div className="relative bg-gradient-to-br from-yellow-100 to-yellow-200 rounded-3xl h-full flex flex-col justify-between p-4 sm:p-6">
                  <div className="text-7xl sm:text-9xl">🏝️</div>
                  <div>
                    <h3 className="text-base sm:text-lg font-bold text-yellow-900">
                      Toy island
                    </h3>
                    <p className="text-xs sm:text-sm text-yellow-700 line-clamp-2">
                      Find new treasures
                    </p>
                  </div>
                </div>
                <div className="absolute top-2 right-2 animate-bounce">✨</div>
              </div>
            </div>
          </div>

          {/* Rest of the islands remain the same... */}
          {/* Magic Island */}
          <div className="relative pb-[100%]">
            <div className="absolute inset-0">
              <div className="relative group h-full">
                <div className="absolute inset-0 bg-purple-300 rounded-3xl blur opacity-50 group-hover:opacity-75 transition-opacity" />
                <div className="relative bg-gradient-to-br from-purple-100 to-purple-200 rounded-3xl h-full flex flex-col justify-between p-4 sm:p-6">
                  <div className="text-7xl sm:text-9xl">🔮</div>
                  <div>
                    <h3 className="text-base sm:text-lg font-bold text-purple-900">
                      Magic island
                    </h3>
                    <p className="text-xs sm:text-sm text-purple-700 line-clamp-2">
                      Tell your desire
                    </p>
                  </div>
                </div>
                <div className="absolute top-2 right-2 animate-bounce">✨</div>
              </div>
            </div>
          </div>

          {/* Collection Castle */}
          <div className="relative pb-[100%]">
            <div className="absolute inset-0">
              <div className="relative group h-full">
                <div className="absolute inset-0 bg-pink-300 rounded-3xl blur opacity-50 group-hover:opacity-75 transition-opacity" />
                <div className="relative bg-gradient-to-br from-pink-100 to-pink-200 rounded-3xl h-full flex flex-col justify-between p-4 sm:p-6">
                  <div className="text-7xl sm:text-9xl">🏰</div>
                  <div>
                    <h3 className="text-base sm:text-lg font-bold text-pink-900">
                      My castle
                    </h3>
                    <p className="text-xs sm:text-sm text-pink-700 line-clamp-2">
                      Your collection
                    </p>
                  </div>
                </div>
                <div className="absolute top-2 right-2 animate-bounce">✨</div>
              </div>
            </div>
          </div>

          {/* Game Island */}
          <div className="relative pb-[100%]">
            <div className="absolute inset-0">
              <div className="relative group h-full">
                <div className="absolute inset-0 bg-green-300 rounded-3xl blur opacity-50 group-hover:opacity-75 transition-opacity" />
                <div className="relative bg-gradient-to-br from-green-100 to-green-200 rounded-3xl h-full flex flex-col justify-between p-4 sm:p-6">
                  <div className="text-7xl sm:text-9xl">🎮</div>
                  <div>
                    <h3 className="text-base sm:text-lg font-bold text-green-900">
                      Game Island
                    </h3>
                    <p className="text-xs sm:text-sm text-green-700 line-clamp-2">
                      The time of fun!
                    </p>
                  </div>
                </div>
                <div className="absolute top-2 right-2 animate-bounce">✨</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Magic Tree */}
      <div className="px-4 pt-8 pb-24">
        <div className="max-w-md mx-auto bg-gradient-to-br from-green-100 to-green-200 rounded-3xl p-6 shadow-xl">
          <div className="flex items-center gap-4">
            <div className="text-6xl animate-pulse">🌳</div>
            <div className="flex-1">
              <div className="text-lg font-bold text-green-900">
                Your magic tree
              </div>
              <div className="mt-2 bg-white/50 rounded-full h-4 overflow-hidden">
                <div className="bg-green-500 h-full w-3/4 rounded-full" />
              </div>
            </div>
            <div className="text-4xl animate-bounce">🌟</div>
          </div>
        </div>
      </div>

      {/* Navigation Dock */}
      <div className="fixed bottom-6 left-1/2 transform -translate-x-1/2">
        <div className="bg-white/30 backdrop-blur-lg rounded-full p-4 shadow-lg">
          <div className="flex gap-8">
            <button className="text-3xl hover:scale-110 transition-transform">
              🎁
            </button>
            <button className="text-3xl hover:scale-110 transition-transform">
              🔍
            </button>
            <button className="text-3xl hover:scale-110 transition-transform">
              ⭐
            </button>
            <button className="text-3xl hover:scale-110 transition-transform">
              😊
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
