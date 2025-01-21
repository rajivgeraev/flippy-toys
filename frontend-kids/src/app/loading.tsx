// src/app/loading.tsx
export default function Loading() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-sky-400 via-blue-400 to-blue-500 flex items-center justify-center">
      <div className="bg-white/30 backdrop-blur-lg rounded-full p-8 animate-pulse">
        <div className="text-6xl">🎮</div>
      </div>
    </div>
  );
}
