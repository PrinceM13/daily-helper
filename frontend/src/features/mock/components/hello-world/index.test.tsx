describe("Basic math sanity check", () => {
  it("1 + 2 should equal 3", () => {
    expect(1 + 2).toBe(3);
  });
});

// import "@testing-library/jest-dom";
// import { render, screen, fireEvent } from "@testing-library/react";
// import HelloWorld from "./index";

// describe("HelloWorld UI", () => {
//   it("renders with initial count", () => {
//     render(<HelloWorld />);
//     expect(screen.getByTestId("count")).toHaveTextContent("Count: 0");
//   });

//   it("increments when button clicked", () => {
//     render(<HelloWorld />);
//     fireEvent.click(screen.getByText("Increment"));
//     expect(screen.getByTestId("count")).toHaveTextContent("Count: 1");
//   });

//   it("resets when button clicked", () => {
//     render(<HelloWorld />);
//     fireEvent.click(screen.getByText("Increment"));
//     fireEvent.click(screen.getByText("Reset"));
//     expect(screen.getByTestId("count")).toHaveTextContent("Count: 0");
//   });
// });
