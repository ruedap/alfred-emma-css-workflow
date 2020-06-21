import { assertEquals } from "https://deno.land/std/testing/asserts.ts";
import { toJson, TResponse } from "./response.ts";

const response: TResponse = {
  items: [
    {
      uid: "pos-s",
      title: "pos-s { position: static; }",
      subtitle: "Paste class name: pos-s",
      arg: "pos-s",
      icon: {
        type: "fileicon",
        path: "icon.png",
      },
    },
    {
      uid: "pos-r",
      title: "pos-r { position: relative; }",
      subtitle: "Paste class name: pos-r",
      arg: "pos-r",
      icon: {
        type: "fileicon",
        path: "icon.png",
      },
    },
  ],
};

Deno.test("toJson", () => {
  const actual = toJson(response);
  const expected =
    `{"items":[{"uid":"pos-s","title":"pos-s { position: static; }","subtitle":"Paste class name: pos-s","arg":"pos-s","icon":{"type":"fileicon","path":"icon.png"}},{"uid":"pos-r","title":"pos-r { position: relative; }","subtitle":"Paste class name: pos-r","arg":"pos-r","icon":{"type":"fileicon","path":"icon.png"}}]}`;
  assertEquals(actual, expected);
});
