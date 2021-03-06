import { parse } from "https://deno.land/std/encoding/yaml.ts";
import { TResponse } from "./alfred/response.ts";

type TVar = Readonly<{
  name: string;
  value: string;
}>;

export type TProp = Readonly<{
  name: string;
  abbr: string;
  group: string;
  values: {
    name: string;
    abbr: string;
  }[];
}>;

type TMixin = Readonly<{
  name: string;
  abbr: string;
  desc?: string;
  group: string;
  decls: {
    prop: string;
    value: string;
  }[];
}>;

export type TEmma = Readonly<{
  ver: string;
  vars: TVar[];
  rules: {
    props: TProp[];
    mixins: TMixin[];
  };
}>;

export const getData = (path = "./emma-data.yml") => {
  const yml = Deno.readTextFileSync(path);
  return parse(yml) as TEmma;
};

export const isNumeric = (str: string) => {
  return !isNaN(parseInt(str));
};

export const generateAbbr = (propAbbr: string, valueAbbr: string) => {
  return isNumeric(valueAbbr)
    ? `${propAbbr}${valueAbbr}`
    : `${propAbbr}-${valueAbbr}`;
};

export const propsToResponse = (props: TProp[]) => {
  const res = props.map((prop) => {
    return prop.values.map((value) => {
      const abbr = generateAbbr(prop.abbr, value.abbr);
      return ({
        title: `${abbr} { ${prop.name}: ${value.name}; }`,
        subtitle: `Paste class name: ${abbr}`,
        arg: `${abbr}`,
        icon: {
          type: "fileicon",
          path: "./icon.png",
        },
      });
    });
  }).flat();

  return { items: res } as TResponse;
};

export const responseToJson = (response: TResponse) => {
  return JSON.stringify(response);
};
