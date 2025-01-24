// src/app/not-found.tsx
import Link from "next/link";

export default function NotFound() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-purple-100 to-white flex items-center justify-center p-4">
      <div className="bg-white rounded-2xl p-6 shadow-xl max-w-md w-full text-center">
        <div className="text-6xl mb-4">🔍</div>
        <h2 className="text-xl font-bold text-gray-800 mb-2">
          Страница не найдена
        </h2>
        <p className="text-gray-600 mb-4">
          Похоже, эта страница потерялась в игре
        </p>
        <Link
          href="/"
          className="inline-block bg-purple-500 text-white px-6 py-2 rounded-full hover:bg-purple-600 transition-colors"
        >
          Вернуться домой
        </Link>
      </div>
    </div>
  );
}
