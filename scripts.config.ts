import { DenonConfig } from "https://deno.land/x/denon@2.4.7/mod.ts";

const DENO = "deno";
const DIST_FILE = "Emma.alfredworkflow";
const BIN_FILE = "./emma-css-workflow";
const PLIST_FILE = "./info.plist";
const ICON_FILE = "./icon.png";
const DATA_FILE = "./emma-data.yml";
const ENTRY_FILE = "./src/mod.ts";

const SCRIPTS = {
  build: {
    cmd:
      `zip -r ${DIST_FILE} ${BIN_FILE} ${PLIST_FILE} ${ICON_FILE} ${DATA_FILE}`,
  },
  compile: {
    cmd:
      `${DENO} compile --unstable --lite --allow-read --target x86_64-apple-darwin --output ${BIN_FILE} ${ENTRY_FILE}`,
  },
  clean: {
    cmd: `rm -rf ${DIST_FILE}`,
  },
  open: {
    cmd: `open ${DIST_FILE}`,
  },
  run: {
    cmd: `${DENO} run --allow-read ${ENTRY_FILE}`,
  },
  test: {
    cmd: `${DENO} test`,
  },
  fmt: {
    cmd: `${DENO} fmt --check`,
  },
  lint: {
    cmd: `${DENO} lint --unstable`,
  },
};

const config: DenonConfig = {
  watch: false,
  scripts: {
    ...SCRIPTS,
    make: [
      { cmd: SCRIPTS.clean.cmd },
      { cmd: SCRIPTS.compile.cmd },
      { cmd: SCRIPTS.build.cmd },
    ],
    check: [
      { cmd: SCRIPTS.fmt.cmd },
      { cmd: SCRIPTS.lint.cmd },
    ],
  },
};

export default config;
