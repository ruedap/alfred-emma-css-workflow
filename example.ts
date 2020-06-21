import { readFileStrSync } from "https://deno.land/std/fs/read_file_str.ts";
import { parse as yamlParse } from "https://deno.land/std/encoding/yaml.ts";
import { parse } from "https://deno.land/std/flags/mod.ts";

// @deno-types="https://deno.land/x/fuse@v6.0.4/dist/fuse.d.ts"
import Fuse from 'https://deno.land/x/fuse@v6.0.4/dist/fuse.esm.min.js';

const yml = readFileStrSync('./emma-data.yml')
const data = yamlParse(yml) as any
console.log(data)
// console.log(data.rules.props[0])

const fuse = new Fuse(data.rules.props, {
  includeScore: true,
  useExtendedSearch: true,
  keys: ['name', 'values.name']
}, undefined)

const args = parse(Deno.args)
const res = fuse.search(args._.join(' '))
console.log(res)
console.log(String(args._.join(' ')))
