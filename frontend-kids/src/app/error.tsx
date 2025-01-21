// src/app/error.tsx
"use client";

import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error("Error:", error);
  }, [error]);

  return (
    <div className="min-h-screen bg-gradient-to-b from-red-100 to-white flex items-center justify-center p-4">
      <div className="bg-white rounded-2xl p-6 shadow-xl max-w-md w-full text-center">
        <div className="text-6xl mb-4">😢</div>
        <h2 className="text-xl font-bold text-gray-800 mb-2">
          Что-то пошло не так
        </h2>
        <p className="text-gray-600 mb-4">
          Не переживай, давай попробуем еще раз
        </p>
        <button
          onClick={reset}
          className="bg-blue-500 text-white px-6 py-2 rounded-full hover:bg-blue-600 transition-colors"
        >
          Попробовать снова
        </button>
      </div>
    </div>
  );
}
