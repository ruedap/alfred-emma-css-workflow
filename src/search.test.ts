import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { search } from "./search.ts";
import { dummyList1 } from "./emma.test.ts";

Deno.test("search", () => {
  const actual = search(dummyList1, ["name", "values.name"], "auto");
  const expected = [
    dummyList1[1],
    dummyList1[2],
    dummyList1[0],
  ];
  assertEquals(actual, expected);
});
