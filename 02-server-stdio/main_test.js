const test = require('node:test');

const { Client } = require('@modelcontextprotocol/sdk/client/index.js');
const { StdioClientTransport } = require('@modelcontextprotocol/sdk/client/stdio.js');

// 1.- generamos el protocolo de transporte
const transport = new StdioClientTransport({
    command: 'node', // or node server.mjs | node server.cjs
    args: ['main.js']
});

// 2.- generamos el cliente
const client = new Client({
    name: 'server-stdio-client',
    version: '0.0.1'
});

test.describe('test mcp server stdio', () => {
    test.before(async () => await client.connect(transport));

    test.it('list tools', async t => {
        const {tools} = await client.listTools();
        
        t.assert.ok(
            Array.isArray(tools) && tools.length > 0, 
            'Tools is empty or not is array'
        );
    });

    test.it('add two numbers', async t => {
        const input = {
            a: 25.00,
            b: 25.00
        };

        const result = await client.callTool({
            name: 'add',
            arguments: input
        });
            
        const [{text}] = result.content;

        // verificamos los resultados
        t.assert.equal(
            text, 
            `${input.a} + ${input.b} = ${String(input.a + input.b)}`
        );
    });

    test.it('multiply two numbers', async t => {
        const input = {
            a: 25.00,
            b: 5.00
        };

        const result = await client.callTool({
            name: 'multiply',
            arguments: input
        });
            
        const [{text}] = result.content;

        // verificamos los resultados
        t.assert.equal(text, `${input.a} * ${input.b} = ${String(input.a * input.b)}`);
    });


    test.it('greeting', async t => {
        let name = 'gabriel'
        
        const result = await client.callTool({
            name: 'get_greeting',
            arguments: {
                name
            }
        });

        const [{text}] = result.content;

        t.assert.equal(text, `Hello, ${name}! Welcome to the MCP stdio server.`);
    });

    test.after(async () => await client.close());
});