import { useState } from "react";

export function useHelloWorld() {
  const [count, setCount] = useState(0);

  const increment = () => setCount((c) => c + 1);
  const reset = () => setCount(0);

  return { count, increment, reset };
}
