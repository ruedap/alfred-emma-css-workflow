import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { search } from "./search.ts";

Deno.test("search", () => {
  const _position = {
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
  };

  const _top = {
    name: "top",
    abbr: "t",
    group: "display",
    values: [
      { name: "auto", abbr: "a" },
      { name: "0", abbr: "0" },
    ],
  };

  const _right = {
    name: "right",
    abbr: "r",
    group: "display",
    values: [
      { name: "auto", abbr: "a" },
      { name: "0", abbr: "0" },
    ],
  };

  const list = [
    _position,
    _top,
    _right,
  ];

  const actual = search(list, ["name", "values.name"], "auto");
  const expected = [
    {
      item: _top,
      refIndex: 1,
      score: 1.0536712127723509e-8,
    },
    {
      item: _right,
      refIndex: 2,
      score: 1.4901161193847656e-8,
    },
    {
      item: _position,
      refIndex: 0,
      score: 0.3712142238654118,
    },
  ];
  assertEquals(actual, expected);
});
