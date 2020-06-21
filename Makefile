run:
	./deno run --allow-read ./src/mod.ts $(args)

test:
	./deno test

fmt:
	./deno fmt

lint:
	./deno lint --unstable
