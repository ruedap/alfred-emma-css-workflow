import { parse } from "https://deno.land/std/flags/mod.ts";

export const getQuery = () => {
  return parse(Deno.args)._.join(" ");
};
