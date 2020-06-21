import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { search } from "./search.ts";

Deno.test("search", () => {
  const list = [
    {
      title: "Old Man's War",
      author: "John Scalzi",
      tags: ["fiction"],
    },
    {
      title: "The Lock Artist",
      author: "Steve",
      tags: ["thriller"],
    },
  ];

  const actual = search(list, ["author", "tags"], "tion");
  const expected = [
    {
      item: {
        title: "Old Man's War",
        author: "John Scalzi",
        tags: ["fiction"],
      },
      refIndex: 0,
      score: 0.17320508075688773,
    },
  ];
  assertEquals(actual, expected);
});
