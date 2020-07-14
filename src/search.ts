// @deno-types="https://github.com/krisk/Fuse/raw/v6.2.1/dist/fuse.d.ts"
import Fuse from "https://github.com/krisk/Fuse/raw/v6.2.1/dist/fuse.esm.js";
import { TProp } from "./emma.ts";

type TSearchResult = Readonly<{
  item: TProp;
  refIndex: number;
  score: number;
}>;

export const search = (list: TProp[], keys: string[], query: string) => {
  if (!query) return list

  const options = {
    includeScore: true,
    useExtendedSearch: true,
    keys: keys,
  } as const;

  const fuse = new Fuse(list, options, undefined);
  const sr = fuse.search(query) as TSearchResult[];

  return sr.map((v) => v.item) as TProp[];
};
