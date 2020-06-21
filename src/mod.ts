import { getData, propsToResponse } from "./emma.ts";
import { getQuery } from "./query.ts";
import { search } from "./search.ts";

const data = getData();
const query = getQuery();
const searchResult = search(
  data.rules.props,
  ["name", "values.name"],
  query,
);
const response = propsToResponse(searchResult);

console.log(response);
