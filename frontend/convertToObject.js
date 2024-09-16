"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var ts = require("typescript");
var visit = function (node) {
    if (ts.isTypeLiteralNode(node)) {
        // Extract the properties of the type literal
        var properties = node.members.map(function (member) {
            var _a, _b;
            // Check if the member is a property signature
            if (ts.isPropertySignature(member)) {
                var name_1 = member.name.text;
                var type = member.type;
                // Check if type is defined before processing
                if (type) {
                    // Handle literal types
                    if (ts.isLiteralTypeNode(type) && ts.isStringLiteral(type.literal)) {
                        return _a = {}, _a[name_1] = type.literal.text, _a;
                    }
                    // Handle union types
                    else if (ts.isUnionTypeNode(type)) {
                        var types = type.types.map(function (t) {
                            if (ts.isLiteralTypeNode(t) && ts.isStringLiteral(t.literal)) {
                                return t.literal.text;
                            }
                            return "unknown";
                        });
                        return _b = {}, _b[name_1] = types, _b;
                    }
                }
            }
            return null;
        });
        // Combine the properties into a single object
        return properties.reduce(function (acc, prop) { return Object.assign(acc, prop); }, {});
    }
    return null;
};
var convertToObject = function (type) {
    var _a;
    var sourceFile = ts.createSourceFile("temp.ts", type, ts.ScriptTarget.ES2015, true);
    // Variables to hold the result and the type name
    var result = null;
    var typeName = null;
    // Function to visit each node in the AST
    var traverse = function (node) {
        // Extract the type name if it's a type alias
        if (ts.isTypeAliasDeclaration(node) && ts.isTypeLiteralNode(node.type)) {
            typeName = node.name.text;
        }
        var resultFromVisit = visit(node);
        if (resultFromVisit) {
            result = resultFromVisit;
        }
        ts.forEachChild(node, traverse);
    };
    // Start visiting nodes from the source file
    traverse(sourceFile);
    // Return the result with the type name as the top-level key
    return typeName ? (_a = {}, _a[typeName] = result, _a) : result;
};
// Example used for debugging
var example1 = "\n    type Button1 = {\n        variant: \"solid\";\n    };\n";
var example2 = "\n    type Button2 = {\n        variant: \"solid\";\n        color: \"blue\";\n    };\n";
// Example provided
var example3 = "\n    type Button = {\n        variant: \"solid\" | \"text\";\n    };\n";
console.log(convertToObject(example1));
console.log(convertToObject(example2));
console.log(convertToObject(example3));
