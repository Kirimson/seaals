# SEAaLS API Server

<!--toc:start-->

- [SEAaLS API Server](#seaals-api-server)
  - [Sample endpoints](#sample-endpoints)
    - [/seal](#seal)
    - [/seal/says/{text}](#sealsaystext)
  - [Development](#development) - [Adding additional SQL Operations](#adding-additional-sql-operations) - [Adding additional API routes/parameters](#adding-additional-api-routesparameters) - [Running the server](#running-the-server) - [Docker](#docker)
  <!--toc:end-->

An API Server to serve seal images. Uses Imagemagick to add effects to images

## Sample endpoints

To see all endpoints, go to `/swagger/index.html` to view the full API spec

### /seal

Get a random seal image

Params:

- `tag (query)`: Type of seal to get by one of its tags
- `filter (query)`: Filter to apply to the image
- `permalink (query)`: Create a permalink to the seal, including filters applied to the image

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
- `permalink (query)`: Create a permalink to the seal, including filters applied to the image

## Development

### Adding additional SQL Operations

All interactions with the database are generated with sqlc. This handles the generation of the database client,
which interacts directly with the database.

Run `sqlc generate` to update the generated database client when the schema or queries SQL files have updated.

The seals and tags service uses this database client to add business logic to the database operations. The service still
only handles operations on the database layer. Anything at the OS file system level (saving/deleting image files) is done
at the controller level

### Adding additional API routes/parameters

The REST API server is generated and defined by the openapi spec file, located at `/openapi.yaml`. oapi-codegen is used
to generate the REST API server at `api/*.gen.go`. Files ending in `*.gen.go` should not be edited.

Run `go generate api/seaals.go` to update the generate API server. Under the hood, this runs `oapi-codegen` against the
models and server oapi-codegen files in the `api` directory.

The implementation of the REST API endpoints are defined in `api/seaals.go` which implements the generated API
`ServerInterface` interface.

### Running the server

Get all dependencies with `go get`, and `npm i` (For tailwind)

Seaals depends on Imagemagick 7.0 to run correctly. It must be installed to make use of the imagick
package, which uses the Imagemagick 7.x bindings.

Run the server with `go run main.go serve`

To update the output css, if adding new tailwind classes, run `npm run css`

> This runs `npx @tailwindcss/cli -i ./public/assets/input.css -o ./public/assets/output.css`

To live refresh the server, and regenerate css, use `go tool air`

#### Docker

Build a docker image with `docker build -t <image_name> .`

This is an alpine-based image, with imagemagick 7 and other dependencies pre-installed
