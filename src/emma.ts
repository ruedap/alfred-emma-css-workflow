type TVar = Readonly<{
  name: string;
  value: string;
}>;

type TProp = Readonly<{
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
