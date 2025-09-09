import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import HelloWorld from "./index";

describe("HelloWorld UI", () => {
  let user: ReturnType<typeof userEvent.setup>;
  let incrementButton: HTMLElement;
  let resetButton: HTMLElement;
  let count: HTMLElement;

  beforeEach(() => {
    render(<HelloWorld />);
    user = userEvent.setup({ delay: 100 }); // 100ms between each event
    incrementButton = screen.getByText("Increment");
    resetButton = screen.getByText("Reset");
    count = screen.getByTestId("count");
  });

  it("shows initial UI state", () => {
    expect(count).toHaveTextContent("Count: 0");
  });

  it("updates UI when increment button clicked", async () => {
    await user.click(incrementButton);
    expect(count).toHaveTextContent("Count: 1");
  });

  it("updates UI when reset button clicked", async () => {
    await user.click(incrementButton); // bump count to 1 first
    await user.click(resetButton); // then reset
    expect(count).toHaveTextContent("Count: 0");
  });
});
