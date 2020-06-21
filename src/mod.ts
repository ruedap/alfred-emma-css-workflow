import { getData } from "./emma.ts";
import { getQuery } from "./query.ts";
import { search } from "./search.ts";

const data = getData();
const query = getQuery();
const searchResult = search(
  data.rules.props,
  ["name", "values.name"],
  query,
);

console.log(searchResult);
