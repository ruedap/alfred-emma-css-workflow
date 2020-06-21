run:
	./deno run --allow-read ./src/index.ts $(args)

test:
	./deno test

fmt:
	./deno fmt

lint:
	./deno lint --unstable
