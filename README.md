# SEAaLS - Seals of Approval

<!--toc:start-->

- [SEAaLS - Seals of Approval](#seaals-seals-of-approval)
  - [Sample endpoints](#sample-endpoints)
    - [/seal](#seal)
    - [/seal/says/{text}](#sealsaystext)
  - [Development](#development)
    - [Adding additional SQL Operations](#adding-additional-sql-operations)
    - [Adding additional API routes/parameters](#adding-additional-api-routesparameters)
    - [Running the server](#running-the-server)
      - [Docker](#docker)
  - [Administration](#administration)
  <!--toc:end-->

An API Server to serve seal images. Uses Imagemagick to add effects to images

## Sample endpoints

To see all endpoints, go to `/swagger/index.html` to view the full API spec

### /seal

Get a random seal image

Params:

- `tag (query)`: Type of seal to get by one of its tags
- `filter (query)`: Filter to apply to the image
- `permalink (query)`: Create a permalink to the seal, including filters applied
  to the image

### /seal/says/{text}

Get a random seal image with a caption

Params:

- `text (path)`: Text of the caption
- `tag (query)`: Type of seal to get by one of its tags
- `filter (query)`: Filter to apply to the image
- `position (query)`: Position of the text
- `fontSize (query)`: Size of the text
- `fontColour (query)`: Colour of the text
- `borderColour (query)`: Colour of the text border
- `borderSize (query)`: Size of the text border
- `permalink (query)`: Create a permalink to the seal, including filters applied
  to the image

## Development

### Adding additional SQL Operations

All interactions with the database are generated with sqlc. This handles the
generation of the database client, which interacts directly with the database.

Run `sqlc generate` to update the generated database client when the schema or
queries SQL files have updated.

The seals and tags service in `service/` uses the database client, adding business
logic to operations. The services only handle operations at the database layer.
File handling, such as writing a new Seal to disk is done by the controllers

### Adding additional API routes/parameters

The REST API server is generated and defined by the openapi spec file, located
at `/openapi.yaml`. oapi-codegen generates the server at `api/*.gen.go`.
Files ending in `*.gen.go` should not be edited.

Run `go generate api/seaals.go` to update the generate API server. This runs
`oapi-codegen` against the models and server oapi-codegen configuration files
in the `api` directory.

The implementation of the REST API endpoints are defined in `api/seaals.go`
which implements the generated API `ServerInterface` interface.

### Running the server

Get all dependencies with `go get`, and `npm i` (For tailwind)

Seaals depends on Imagemagick 7.0 to run correctly. It must be installed to make
use of the `imagick` package, which uses Imagemagick 7.x bindings, before the
server can be started.

The database path and data path can be defined by the `--db` and `--base-path`/`-b`
flags, or the `DB_PATH` and `BASE_PATH` environment variables. Env files can
also be used.

If starting the server for the first time, migrate the database schema with
`seaals db migrate`

Start the server locally with `go run main.go serve`

To update the output css, if adding new tailwind classes, run `npm run css`

> This runs `npx @tailwindcss/cli -i ./public/assets/input.css -o ./public/assets/output.css`

To live refresh the server, and regenerate css, use `go tool air`. This also re-generates
CSS by running `tailwinscss/cli`

#### Docker

Build a docker image with `docker build -t <image_name> .`

This is an alpine-based image, with Imagemagick 7 and other dependencies pre-installed

### Administration

Currently, SEAals can only be administered from the command line, using the
`seaals admin` command:

- `seaals admin list-seals` - Lists all Seals added to the database
- `seaals admin list-tags` - List all Tags added to the database
- `seaals admin add-seal` - Add a Seal to the database. Can be provided via a
  `--path` or `--url`
- `seaals admin add-tag` - Add a Tag to the database
- `seaals admin delete-seal` - Removes a seal from the database and disk
