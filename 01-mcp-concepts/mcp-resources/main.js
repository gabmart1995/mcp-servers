const {McpServer, ResourceTemplate} = require('@modelcontextprotocol/sdk/server/mcp.js');
const {StdioServerTransport} = require('@modelcontextprotocol/sdk/server/stdio.js');
const { default: z } = require('zod');

async function main() {
    const server = new McpServer({
        name: 'mcp-example-calculator',
        version: '0.0.1'
    });

    server.tool(
        'add',
        { a: z.number(), b: z.number() },
        async function({a, b}) {
            return {
                content: [{type: 'text', text: String(a + b)}]
            }
        }
    );

    server.tool(
        'substract',
        { a: z.number(), b: z.number() },
        async function({a, b}) {
            return {
                content: [{type: 'text', text: String(a - b)}]
            }
        }
    );

    server.tool(
        'multiply',
        { a: z.number(), b: z.number() },
        async function({a, b}) {
            return {
                content: [{type: 'text', text: String(a * b)}]
            }
        }
    );

    server.tool(
        'divide',
        { a: z.number(), b: z.number() },
        async function({a, b}) {
            if (b === 0) {
                return {
                    isError: true,
                    content: [{type: 'Text', text: 'Error: dont divide by zero'}]
                };
            }
            
            return {
                content: [{type: 'text', text: String(a / b)}]
            }
        }
    );

    server.resource(
        'greeting',
        new ResourceTemplate(
            'greeting://{name}',
            { list: undefined } // descartara almacenar el archivo
        ),
        async function (uri, {name}) {
            return {
                contents: [{
                    uri: uri.href,
                    text: `Hello ${name}!`
                }]
            }
        }
    )
    
    const transport = new StdioServerTransport();
    await server.connect(transport);
}

main().catch(console.error);
