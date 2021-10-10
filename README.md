# pokemonTranslationAPI

Small API project that provide 2 endpoints:

- /pokemon/<pokemon_name>
- /pokemon/translated/<pokemon_name>
  
The first endpoint return a small set of attributes of the pokemon provided as parameter.
The second endpoint return the same set of attributes ot the previous one but with the description translated with the yoda or shakespeare translation.

# Getting Started

## Docker

The easy way to get started with this project is to use Docker. I tested this with the Docker Engine 19.03.1 on Mac OS. First of all download the git repository.

```bash
git clone https://github.com/succoDiPompelmo/pokemonTranslationAPI
cd pokemonTranslationAPI
```

Build the docker image locally, using a custom image name

```bash
sudo docker build -t <image_name> .
```

Now you can run your application by exposing the internal port 3000 (https://www.whitesourcesoftware.com/free-developer-tools/blog/docker-expose-port/)

```bash
sudo docker -p 3000:3000 <image_name>
```

Once finished you can use any HTTP client (curl, postman, etc...) that you want to simulate request to the docker container.

```bash
curl http://localhost:3000/pokemon/mewtwo
{"Is_legendary":true,"description":"It was created by a scientist after years of horrific gene splicing and DNA engineering experiments.","habitat":"rare","name":"mewtwo"}
```

## Build from source

You can also decide to build directly the application from source. As before you need to download the source code.

```bash
git clone https://github.com/succoDiPompelmo/pokemonTranslationAPI
cd pokemonTranslationAPI
```

Now ask go to build the binaries (version of go used is go 1.16.4 on Mac OS, in case you need to install go https://golang.org/doc/install) and launch the app with two simple commands

```bash
go build
./pokemonTranslationAPI
```

Now you can use a HTTP client to simulate the request to the adress http://localhost:3000.

# Testing

To run all the test suite you only need to execute the command

```bash
go test -v
```

# API Definition

The service provide 2 endpoint

- **/pokemon/<pokemon_name>** that given a valid pokemon name returns a JSON response with the fields related to the pokemon description, pokemon name, if the pokemon is legendary and the habitat name.
- **/pokemon/translated/<pokemo_name>** returns exactly the same output of the first route but with the description translated using the Yoda or Shakespeare translation (more on this here https://api.funtranslations.com/). In case the description cannot be translated for whatever reason the standard description is provided.

# DESIGN CHOICES

- I used the fiber framework to implement the API, because is well documented and faster to code than many competitors (e.g gorilla/mux).
- The requirements specify that in case the translated description cannot be provided, for any reason, that the standard should be used. This is reflected in code that returns always a description even when it encounters an error.
- Most of the pokemon is provided with multiple english/non-english description. I decided that the english description choice would be the first one provided by the server (no particular reason) and in case no english description (or description at all) is present an empty string will be used.
- To simplify the process of developing the API and allowing the test of corner case (e.g timeouts, errors calling the pokemon api, etc..) I decided to mock any request to external API. This was made in two steps. First implement everything without mocks. Then add mocks and keep test. In this way the test also verify that the mocks are responding in a similar way to respet the server. (For translation API I skipped the no-mock part casue the really low rate limit cap).
- Caching was provided at the route level by Fiber framework. It's a first draft and a better approach in this case would be to cache also the external API calls if time allows.
- It's possible to run the API server with TLS enabled, but since the certificate are self signed some problems could arise. I decided to keep it switched off for now.
- No rate limit was implemented on the routes. Since we are caching the response and the number of pokemon is pretty low (max 1000), we are confident that even if the load increase the cache will provide support for the increase number of request. Fun Translation has a rate limit of 5 request/hour, but it's not a problem because when that limit is reached we simply start to return the standard description (or the cached one).

# WHAT'S MISSING

Below a list of the things that are missing before going to be fully operation:

- Metrics/tracing are not present in the project. There are some libraries that allow fiber to enable a Prometheus endpoint.
- More powerful and complete logging system. For now only basic logging when error occour.
- Detailed error description as response. Now only the status code is returned
- Enhanced security with valid certificate and a API Key or other forms of Authentication/Authorization
- Heavy refactoring, lot of code smells and repeated code.
- CI/CD. Tested github actions but not bridged that on this project. A CI/CD pipeline that automatically build docker images/binaries should be useful.
- Integration/Stress automated test.
