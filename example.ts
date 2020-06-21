import { readFileStrSync } from "https://deno.land/std/fs/read_file_str.ts";
import { parse } from "https://deno.land/std/encoding/yaml.ts";

const yml = readFileStrSync('./emma-data.yml')
const data = parse(yml) as any
console.log(data.rules.props[0])
