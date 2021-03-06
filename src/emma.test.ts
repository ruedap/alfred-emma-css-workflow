import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { generateAbbr, isNumeric, propsToResponse, TProp } from "./emma.ts";

export const sampleProps: { [key: string]: TProp } = {
  position: {
    name: "position",
    abbr: "pos",
    group: "display",
    values: [
      { name: "static", abbr: "s" },
      { name: "relative", abbr: "r" },
      { name: "absolute", abbr: "a" },
      { name: "sticky", abbr: "sk" },
      { name: "fixed", abbr: "f" },
    ],
  },
  top: {
    name: "top",
    abbr: "t",
    group: "display",
    values: [
      { name: "auto", abbr: "a" },
      { name: "0", abbr: "0" },
    ],
  },
  right: {
    name: "right",
    abbr: "r",
    group: "display",
    values: [
      { name: "auto", abbr: "a" },
      { name: "0", abbr: "0" },
    ],
  },
};

export const sampleList1: TProp[] = [
  sampleProps.position,
  sampleProps.top,
  sampleProps.right,
];

Deno.test("isNumeric", () => {
  assertEquals(isNumeric("1"), true);
  assertEquals(isNumeric("0"), true);
  assertEquals(isNumeric("1.0"), true);
  assertEquals(isNumeric(""), false);
  assertEquals(isNumeric("foo"), false);
});

Deno.test("generateAbbr", () => {
  assertEquals(generateAbbr("a", "0"), "a0");
  assertEquals(generateAbbr("a", "b"), "a-b");
});

Deno.test("propsToResponse", () => {
  const actual = propsToResponse([sampleProps.top, sampleProps.right]);
  const expected = {
    items: [
      {
        title: "t-a { top: auto; }",
        subtitle: "Paste class name: t-a",
        arg: "t-a",
        icon: { type: "fileicon", path: "./icon.png" },
      },
      {
        title: "t0 { top: 0; }",
        subtitle: "Paste class name: t0",
        arg: "t0",
        icon: { type: "fileicon", path: "./icon.png" },
      },
      {
        title: "r-a { right: auto; }",
        subtitle: "Paste class name: r-a",
        arg: "r-a",
        icon: { type: "fileicon", path: "./icon.png" },
      },
      {
        title: "r0 { right: 0; }",
        subtitle: "Paste class name: r0",
        arg: "r0",
        icon: { type: "fileicon", path: "./icon.png" },
      },
    ],
  };
  assertEquals(actual, expected);
});
