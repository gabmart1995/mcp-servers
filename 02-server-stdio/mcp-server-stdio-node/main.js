const { Server } = require('@modelcontextprotocol/sdk/server/index.js');
const { StdioServerTransport } = require('@modelcontextprotocol/sdk/server/stdio.js');
const { ListToolsRequestSchema, CallToolRequestSchema } = require('@modelcontextprotocol/sdk/types.js');
const { default: z } = require('zod');

/**
 * este modulo configura el servidor 
 * permite construir y tener mas control acerca de los metodos STDIO
 */
async function main() {
    const server = new Server({
        name: 'example-server',
        version: '0.0.1'
    },
        {
            capabilities: {
                tools: {}
            }
        }
    );

    // establemos los esquemas
    const addArgsSchema = z.object({
        a: z.number().describe('First number'),
        b: z.number().describe('Secound number')
    });

    const multiplyArgsSchema = z.object({
        a: z.number().describe('First number'),
        b: z.number().describe('Secound number')
    });

    const greetingArgsSchema = z.object({
        name: z.string().describe('Name for the person to greet')
    });

    // lista los herramientas MCP disponibles
    server.setRequestHandler(ListToolsRequestSchema, async () => {
        return {
            tools: [
                {
                    name: 'add',
                    description: 'Add two numbers',
                    inputSchema: {
                        type: 'object',
                        properties: {
                            a: { type: 'number', description: 'first number' },
                            b: { type: 'number', description: 'second number' }
                        },
                        required: ['a', 'b']
                    }
                },
                {
                    name: 'multiply',
                    description: 'multiply two numbers',
                    inputSchema: {
                        type: 'object',
                        properties: {
                            a: { type: 'number', description: 'first number' },
                            b: { type: 'number', description: 'second number' }
                        },
                        required: ['a', 'b']
                    }
                },
                {
                    name: 'get_greeting',
                    description: 'Generate a personalized greeting',
                    inputSchema: {
                        type: 'object',
                        properties: {
                            name: { type: 'string', description: 'Name for a person to greet' }
                        },
                        required: ['name']
                    }
                },
                {
                    name: 'get_server_info',
                    description: 'Generate information about this mcp server',
                    inputSchema: {
                        type: 'object',
                        properties: {},
                    }
                },
            ],
        }
    });

    // manejador de recursos MCP
    server.setRequestHandler(CallToolRequestSchema, async request => {
        const { name, arguments: args } = request.params;

        // utilizamos un switch para determinar las acciones del LLM
        switch (name) {
            case 'add': {
                // validamos los datos
                const { a, b } = addArgsSchema.parse(args);
                const result = a + b;

                console.error(`Adding: ${a} + ${b} = ${result}`);

                return {
                    content: [
                        {
                            type: 'text',
                            text: `${a} + ${b} = ${result}`
                        }
                    ]
                }
            }

            case 'multiply': {
                // validamos los datos
                const { a, b } = multiplyArgsSchema.parse(args);
                const result = a * b;

                console.error(`multiply: ${a} * ${b} = ${result}`);

                return {
                    content: [
                        {
                            type: 'text',
                            text: `${a} * ${b} = ${result}`
                        }
                    ]
                }
            }

            case 'get_greeting': {
                // validamos los datos
                const { name } = greetingArgsSchema.parse(args);
                const greeting = `Hello, ${name}! Welcome to the MCP stdio server.`;

                console.error(`Generated greeting for ${name}`); // Log to stderr

                return {
                    content: [
                        {
                            type: 'text',
                            text: greeting
                        }
                    ]
                }
            }

            case 'get_server_info': {
                return {
                    content: [
                        {
                            type: 'text',
                            text: JSON.stringify({
                                server_name: "example-stdio-server",
                                version: "0.0.1",
                                transport: "stdio",
                                capabilities: ["tools"],
                                description: "Example MCP server using stdio transport (MCP 2025-06-18 specification)",
                            }, null, 2),
                        }
                    ]
                }
            }

        }
    });

    // arrancamos el servidor
    console.error("Starting MCP stdio server...");

    const transport = new StdioServerTransport();
    await server.connect(transport);

    console.error("Server connected via stdio transport");
}

// Handle process termination gracefully
process.on("SIGINT", () => {
    console.error("Received SIGINT, shutting down gracefully");
    process.exit(0);
});

process.on("SIGTERM", () => {
    console.error("Received SIGTERM, shutting down gracefully");
    process.exit(0);
});


main().catch((err) => {
    console.error("Server error :", err);
    process.exit(1);
});

