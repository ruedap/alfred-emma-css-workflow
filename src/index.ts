import { readFileStrSync } from "https://deno.land/std/fs/read_file_str.ts";
import { parse as yamlParse } from "https://deno.land/std/encoding/yaml.ts";
import { parse } from "https://deno.land/std/flags/mod.ts";

import { search } from "./search.ts";

const yml = readFileStrSync("./emma-data.yml");
const data = yamlParse(yml) as any;

const query = parse(Deno.args)._.join(" ");
const result = search(
  data.rules.props,
  ["name", "values.name"],
  query,
);

console.log(result);
