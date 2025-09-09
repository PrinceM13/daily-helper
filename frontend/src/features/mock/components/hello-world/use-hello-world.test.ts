import { act, renderHook } from "@testing-library/react";
import { useHelloWorld } from "./use-hello-world";

describe("useHelloWorld", () => {
  test("should initialize with count = 0", () => {
    const { result } = renderHook(() => useHelloWorld());
    expect(result.current.count).toBe(0);
  });

  test("should increment the count", () => {
    const { result } = renderHook(() => useHelloWorld());

    act(() => {
      result.current.increment();
    });

    expect(result.current.count).toBe(1);
  });

  test("should reset the count", () => {
    const { result } = renderHook(() => useHelloWorld());

    act(() => {
      result.current.increment();
      result.current.reset();
    });

    expect(result.current.count).toBe(0);
  });
});
