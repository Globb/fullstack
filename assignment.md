# Part 1 - Backend

You have been tasked with building a Go app that:
1. Retrieves input in structured JSON (the format of which is to your choosing), 
1. Validates the input
1. Logs the input to STDOUT
1. Sends a request (the format of which is to your choosing) to a HTTP based API, with the input being part of the request message.

## Guidance

* Provide a detailed walkthrough of your thought process to understand architectural & other considerations, that help you arrive at a proposed solution
* Provide a descriptive overview of the the solution
* Document challenges, considerations, requirements

## Notes
* We don’t expect you to provide a compiled app, however, the codebase must be compilable
    * Feel free to provide the codebase as a zip file, or as a Git repository
    * You can also provide a Go Playground link, if you prefer
* Feel free to provide diagrams, and any other material, as relevant to support your response
* Do not use a mock API package, or any other mock shortcut approach
* We expect that you will use either a real API, or build your own API separately
    * If you use a public API that requires authentication, you will need to provide us with temporary credentials

# Part 2 - Frontend

Using the [TypeScript compiler API](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API), write a function that inputs a string containing a TypeScript type, and outputs an object literal representing the type.

## Requirements

* Must have a function called visit
* Must have a function called convertToObject that calls the specified visit function
* Must not use CommonJS in your source files (e.g. require())
* Must use the official version of https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API
    * `e.g. import * as ts from "typescript"`;
* We do not expect the convertToObject function to support all kinds of types, but at a bare minimum should work with the example input in the below code block (i.e. type Button…)
* (optional) Provide comments in your source code as much as you want, but keeping it concise

### Sample code

```typescript
const visit = (node: ts.Node) => {
    // implement the visit function here
};
    
const convertToObject = (type: string) => {
    // implement the convertToObject function here
    // this function must use the visit function above
};

// example usage
convertToObject(`
    type Button = {
        variant: "solid" | "text";
    };
`);
/***
outputs

{
    button: {
        variants: ["solid", "text"],
    },
}
*
*/
```