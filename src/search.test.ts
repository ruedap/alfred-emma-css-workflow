import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { search } from "./search.ts";
import { dummyList1 } from "./emma.test.ts";

Deno.test("search", () => {
  const actual = search(dummyList1, ["name", "values.name"], "auto");
  const expected = [
    {
      item: dummyList1[1],
      refIndex: 1,
      score: 1.0536712127723509e-8,
    },
    {
      item: dummyList1[2],
      refIndex: 2,
      score: 1.4901161193847656e-8,
    },
    {
      item: dummyList1[0],
      refIndex: 0,
      score: 0.3712142238654118,
    },
  ];
  assertEquals(actual, expected);
});
