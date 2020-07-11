import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { search } from "./search.ts";
import { sampleList1 } from "./emma.test.ts";

Deno.test("search", () => {
  const actual = search(sampleList1, ["name", "values.name"], "auto");
  const expected = [
    sampleList1[1],
    sampleList1[2],
    sampleList1[0],
  ];
  assertEquals(actual, expected);
});
