import { parse } from "https://deno.land/std/flags/mod.ts";
import { getData } from "./emma.ts";
import { search } from "./search.ts";

const data = getData();
const query = parse(Deno.args)._.join(" ");
const result = search(
  data.rules.props,
  ["name", "values.name"],
  query,
);

console.log(result);
