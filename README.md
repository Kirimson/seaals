# SEaa(L)S 🦭 SEals as a Service

A simple API/application, similar in spirit to `Cataas`. Just with seals instead.

## Why Seals?

Why not?

## Setup

Here's how to serve your own seals in a development environment. By default SEaa(L)S listens on port 3000. Override this
by setting your PORT environment variable.

0. Install node/npm: `brew install node`/`choco install nodejs`
1. Clone the repo
2. Install dependencies with `npm i`
3. Setup the .env file, using the defaults from .env.example
4. Ensure that your chosen seal directory exists. Default is in src/resources/seals
5. Init the DB with `npx prisma migrate dev`
6. Run the application with `npm run dev`

If you want to deploy this in a production environment, make some more production appropriate choices, E.G:

1. Change Prisma's DB provider to something other than SQLite in `prisma/prisma.schema`
2. Use `migrate deploy` rather than `migrate dev`
