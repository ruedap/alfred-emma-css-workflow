// @deno-types="https://deno.land/x/fuse@v6.0.4/dist/fuse.d.ts"
import Fuse from "https://deno.land/x/fuse@v6.0.4/dist/fuse.esm.min.js";

export const search = (list: any, keys: string[], query: string) => {
  const fuse = new Fuse(list, {
    includeScore: true,
    useExtendedSearch: true,
    keys: keys,
  }, undefined);

  return fuse.search(query);
};
