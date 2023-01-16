# Frontend

An UI for displaying the data from [evaluation](../evaluation/README.md) and converting it to nice graphs.

## Build args

Unfortunately, it isn't possible to use environment variables that can be changed at runtime. But Flutter allows for setting build arguments at compiletime. These can be set by adding `--dart-define=<ARG>=<VALUE>` to the build command.

| ARG       | Default                     | Optional | Description                                                                                |
|:----------|:----------------------------|----------|:-------------------------------------------------------------------------------------------|
| `API_URL` | _http://localhost:8080/api_ | Yes      | The URL that will be used for GraphQL requests. Must be accessible from the users browser. |