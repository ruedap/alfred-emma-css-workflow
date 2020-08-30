import { getData, propsToResponse, responseToJson } from "./emma.ts";
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
const json = responseToJson(response);


const outputToAlfred = (json: string): void => {
  console.log(json);
};

outputToAlfred(json);
