// https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type TItem = Readonly<{
  uid?: string;
  title: string;
  subtitle: string;
  arg?: string; // recommended
  icon: {
    type: "fileicon" | "filetype";
    path: string;
  };
  valid?: boolean;
  match?: string;
  autocomplete?: string; // recommended
  type?: "default" | "file" | "file:skipcheck";
  mods?: object;
  text?: object;
  quicklookurl?: string;
}>;

export type TResponse = Readonly<{
  items: TItem[];
}>;

export const toJson = (response: TResponse): string => {
  return JSON.stringify(response);
};
