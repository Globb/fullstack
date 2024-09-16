# Fullstack Assignment

**Completed by Grayson Blanks**

<!-- TOC -->
* [Backend](#backend)
  * [Usage](#usage)
  * [Instructions](#instructions)
  * [Solution](#solution)
  * [Directory](#directory)
  * [Considerations](#considerations)
  * [Requirements](#requirements)
  * [Learnings](#learnings)
* [Frontend](#frontend)
  * [Challenges](#challenges)
<!-- TOC -->

[Assignment can be found here](/assignment.md)

## Backend

### Usage

This section lists out the cli usage, sub commands, and flags

* [backend](#backend-1)
  * [client](#client)
  * [server](#server)


#### backend

```sh
Usage:
  backend [flags]
  backend [command]

Available Commands:
  client      backend client
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  server      backend server

Flags:
      --debug           enable debug logging
      --domain string   Domain url for server. (default "http://localhost")
  -h, --help            help for backend
      --port int        Port used to connect to domain url for server. (default 8080)
```

##### client

```sh
Usage:
  backend client [flags]

Flags:
      --file string   Provide input as a json file. If used with json flag, this input file will be ignored. 
  -h, --help          help for client
      --json string   Provide input as raw json. NOTE: you might need to escape with \"

Global Flags:
      --debug           enable debug logging
      --domain string   Domain url for server. (default "http://localhost")
      --port int        Port used to connect to domain url for server. (default 8080)
```

##### server

```sh
Usage:
  backend server [flags]

Flags:
  -h, --help   help for server

Global Flags:
      --debug           enable debug logging
      --domain string   Domain url for server. (default "http://localhost")
      --port int        Port used to connect to domain url for server. (default 8080)
```

### Instructions

1. Clone this repo
1. Go to `./fullstack/backend` directory
1. Run `go build`
1. Using the new build, run `./backend server`
    * You should see the follow dialog: `{"time":"<TIMESTAMP>","level":"INFO","msg":"server is starting"}`
    * By default the application runs on `localhost:8080`
1. Open a new terminal and run `./backend client --file ./api/jsontests/success.json`
    * You should see `{"time":<TIMESTAMP>,"level":"INFO","msg":"successfully sent post request","response":"received message: {\"adCampaignId\":1,\"customerId\":2\"gameName\":\"halo\",\"imageName\":\"haloImage1\",\"validAccount\":true}"}`
1. Go back to the terminal with the server running
    * You should see `{"time":<TIMESTMAP>,"level":"INFO","msg":"server has received a request","endpoint":"/submit/input","requestType":"POST","responseBody":"{\"adCampaignId\":1,\"customerId\":2,\"gameName\":\"halo\",\"imageName\":\"haloImage1\",\"validAccount\":true}`

### Solution

I have two packages
1. `cmd` - handles cli args and sub commands
1. `api` - handles api server and client

For the `cmd` package, I wanted a simple way to start the server and send API calls to the server via a client without using curl.
* The `cmd/root.go` handles the execute function which is called by main as well as being the root command for the function.
* The `cmd/client.go` handles validating the input and passes the input to the `api/client.go` to be sent to the server. The file also handles figuring out if it should use the `--json` flag or the `--file` flag, where the `--json` flag wins if both are present. It also passes the `api.client.go` what address and port to send the request to.
* The `cmd/server.go` handles passing the address and port for creating the server from `api/server.go`

For the `api` package, I wanted a simple break down of the client, server, json structure and client unit testing
* The `api/client.go` contains the logic for validating JSON input and putting it into the `input` structure as well as sending the input to the server
* The `api/server.go` contains the logic for setting up the server mux, mux handlers, and serving 
* The `api/jsonstruct.go` simply contains the struct with json struct tags
* The `api/client_test.go` is unit testing for the 2 validation functions I wrote
* The `api/jsontests/*.json` files are used in the `api/client_test.go` tests

#### Directory

<!-- Used tree.nathanfriend.io -->
```
fullstack/
└── backend/
    ├── api/
    │   ├── jsontests/
    │   │   ├── failOnBool.json
    │   │   ├── failOnInt.json
    │   │   ├── failOnString.json
    │   │   ├── invalid.json
    │   │   └── success.json
    │   ├── client_test.go
    │   ├── client.go
    │   ├── jsonstruct.go
    │   └── server.go
    ├── cmd/
    │   ├── client.go
    │   ├── root.go
    │   └── server.go
    └── main.go
```

### Considerations

I wanted to use a RESTful API system. Go and RESTful APIs are pretty seamless and easy to create. I considered gRPC, but figured it'd be easier to not worry about HTTP2 and would make testing via `curl` a lot easier.

I originally wrote the code to handle files that contained json but after working on it, considered that maybe that wasn't allowed. Since I wrote the code, I left it in but users have the option to type in raw or pass a json file in.

I used the JSON package to marshal and unmarshal the json strings. This acted as my validator. I made a design decision that when marshalling and unmarshalling, if there was a missing file, I would still return the `Input` structure with the missing fields. It might be more desirable to return an empty string so that if someone doesn't check the error, they don't accidently send a bad request. This was done so my unit testing could be more thorough. 

I went with Go version 1.23.1, really for no other reason than that was what was already installed on my machine.

The logging done in the application is printed out in JSON format, I figured this was the desired output given we're working with JSON through this assignment.

### Requirements

I want to use minimal third party libraries and only ended up using two while creating this application:
1. [Cobra CLI](https://github.com/spf13/cobra) - this is to make the application easier to launch the server and to run a client to send requests
1. [Gotest](https://pkg.go.dev/gotest.tools/v3) - this is used for unit testing, specially for asserts

This means using Go's:
1. `log/slog` - Introduced in go 1.21, structure logging so I don't need to use Zerolog (what I have experience using before)
1. `net/http` - Recent update in 1.22 makes it easier to create servers that don't need logic to switch between `GET` and `POST` requests
1. `encoding/json` - Will handle json marshalling and unmarshalling, which will also act as a validator for json
1. `fmt` & `strings` - Used to handle string formatting
1. `os` - Used to open the json file
1. `io` & `bytes` - Used to convert message bodies to strings

Unit testing seemed important for the validation checks. I purposely kept the validation functions outside of the client interface so that I did not need to mock the interface when running the unit tests. Because I allow the validation functions to fail but also still send the `input` structure with what it was able to unmarshal, I was able to provide meaningful unit tests that show how the functions could fail.

### Learnings

I didn't really have any challenges. So I figured I'd write about what I learned. 

1. Go has really made RESTful APIs easier with the new server mux and being able to add `GET` or `POST` to mux handlers. Typically I would use something like [gorilla/mux](https://github.com/gorilla/mux).
1. I've typically used [Zerolog](https://github.com/rs/zerolog) for logging, but Go's new `log/slog` has been amazingly easy to use.

## Frontend

The code can be found in [frontend/convertToObject.ts](/frontend/convertToObject.ts)

### Challenges
I don't know if this part is required but this section was much tougher for me. I have never worked with the typescript compiler api before and don't have a firm grasp on the typescript syntax.

I had to learn the follow:
* Install typescript
* Understand how typescript syntax works
* Lots of research on the internet to understand how the api compiler works
* Create JavaScript from TypeScript
* Run the JavaScript file after changes were made to the TypeScript file
* Understand how to use prettier to format code, as I didn't know best practices for styling the code

I also had a huge challenge with the union type. So I broke down the example type into smaller pieces. I was able to get the properties of the type to print, but wasn't able to get the type name to print out in the json string. Eventually I added a check at the beginning of the traversal for the type name and stored it. There's probably a better way to do this, but it seems to work. Lastly, once these were all done, my code was able to run but my IDE was throwing a warning about the properties might get overloaded, so I figured out how to define the properties to prevent that message from popping up.

I was able to get my 2 additional examples correct before getting the provided example done. Those examples plus the one provided are printed to the console when running `node convertToObject.js`. I don't know if it was appropriate to keep those in or if that affects how you might test this code, but wanted to make sure it was in there in case it was required.

I also noticed that typescript created a lot of files. I'm not sure what files should be excluded as best practice, so I've included them all, including the generated `convertedToObject.js` file.