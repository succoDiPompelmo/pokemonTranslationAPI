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
