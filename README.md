# SEAaLS API Server

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
only handles opertions on the database layer. Anything at the OS file system level (saving/deleting image files) is done
at the controller level

### Adding additional API routes/paramters

The REST API server is generated and defined by the openapi spec file, located at `/openapi.yaml`. oapi-codgen is used
to generate the REST API server at `api/*.gen.go`. Files ending in `*.gen.go` should not be edited.

Run `go generate api/seaals.go` to update the generate API server. Under the hood, this runs `oapi-codegen` against the
models and server oapi-codegen files in the `api` directory.

The implementation of the REST API endpoints are defined in `api/seaals.go` which implements the generated API
`ServerInterface` interface.