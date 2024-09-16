import * as ts from "typescript";

// Define the type properties for result
type Property = { [key: string]: any };

const visit = (node: ts.Node): any => {
  if (ts.isTypeLiteralNode(node)) {
    // Extract the properties of the type literal
    const properties = node.members.map((member) => {
      // Check if the member is a property signature
      if (ts.isPropertySignature(member)) {
        const name = (member.name as ts.Identifier).text;
        const type = member.type;

        // Check if type is defined before processing
        if (type) {
          // Handle literal types
          if (ts.isLiteralTypeNode(type) && ts.isStringLiteral(type.literal)) {
            return { [name]: type.literal.text };
          }
          // Handle union types
          else if (ts.isUnionTypeNode(type)) {
            const types = type.types.map((t) => {
              if (ts.isLiteralTypeNode(t) && ts.isStringLiteral(t.literal)) {
                return t.literal.text;
              }
              return "unknown";
            });
            return { [name]: types };
          }
        }
      }
      return null;
    });

    // Combine the properties into a single object
    return properties.reduce<Property>(
      (acc, prop) => Object.assign(acc, prop),
      {},
    );
  }
  return null;
};

const convertToObject = (type: string): any => {
  const sourceFile = ts.createSourceFile(
    "temp.ts",
    type,
    ts.ScriptTarget.ES2015,
    true,
  );

  // Variables to hold the result and the type name
  let result: any = null;
  let typeName: string | null = null;

  // Function to visit each node in the AST
  const traverse = (node: ts.Node) => {
    // Extract the type name if it's a type alias
    if (ts.isTypeAliasDeclaration(node) && ts.isTypeLiteralNode(node.type)) {
      typeName = (node.name as ts.Identifier).text;
    }

    const resultFromVisit = visit(node);
    if (resultFromVisit) {
      result = resultFromVisit;
    }

    ts.forEachChild(node, traverse);
  };

  // Start visiting nodes from the source file
  traverse(sourceFile);

  // Return the result with the type name as the top-level key
  return typeName ? { [typeName]: result } : result;
};

// Example used for debugging
const example1 = `
    type Button1 = {
        variant: "solid";
    };
`;

const example2 = `
    type Button2 = {
        variant: "solid";
        color: "blue";
    };
`;

// Example provided
const example3 = `
    type Button = {
        variant: "solid" | "text";
    };
`;

console.log(convertToObject(example1));
console.log(convertToObject(example2));
console.log(convertToObject(example3));
