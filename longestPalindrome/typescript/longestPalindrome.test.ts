// import { describe, expect, test } from "@jest/globals";
import { longestPalindrome, isPalindrome } from "./main";

describe("Is", () => {
  test("Is Palindrome", () => {
    expect(isPalindrome("bab")).toBe(true);
  });
  test("Is not Palindrome", () => {
    expect(isPalindrome("babad")).toBe(false);
  });
});

describe("Longest Palindrome", () => {
  test("Base example", () => {
    expect(longestPalindrome("babad")).toBe("bab");
  });
  test("Second example", () => {
    expect(longestPalindrome("sbbd")).toBe("bb");
  });
});
