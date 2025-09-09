"use client";

import { useHelloWorld } from "./use-hello-world";

export default function HelloWorld() {
  const { count, increment, reset } = useHelloWorld();

  return (
    <div className="mx-auto flex max-w-sm flex-col items-center gap-4 rounded-2xl border border-teal-100 bg-white p-6 shadow-md">
      <h1 className="text-2xl font-bold text-teal-700">Hello World</h1>

      <p
        data-testid="count"
        className="rounded-lg bg-teal-50 px-4 py-2 text-lg font-medium text-teal-600"
      >
        Count: {count}
      </p>

      <div className="flex gap-3">
        <button
          onClick={increment}
          className="rounded-lg bg-teal-600 px-4 py-2 font-medium text-white transition hover:bg-teal-700"
        >
          Increment
        </button>
        <button
          onClick={reset}
          className="rounded-lg bg-teal-100 px-4 py-2 font-medium text-teal-700 transition hover:bg-teal-200"
        >
          Reset
        </button>
      </div>
    </div>
  );
}
