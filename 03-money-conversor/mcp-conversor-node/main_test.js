const test = require('node:test');

const { Client } = require('@modelcontextprotocol/sdk/client/index.js');
const { StdioClientTransport } = require('@modelcontextprotocol/sdk/client/stdio.js');

const { getMoneyValue, getConvertion, CURRENCY_BASE } = require('./handlers');

const transport = new StdioClientTransport({
    command: 'node', // or node server.mjs | node server.cjs
    args: ['main.js']
});

const client = new Client({
    name: 'calculator-client',
    version: '0.0.1'
});


test.describe('testing-mcp-money-conversor', () => {
    test.before(async () => await client.connect(transport));

    test.it('obtener valor entre la moneda base', async t => {
        // para poder comparar los resultados debemos 
        // realizar 2 consultas una al mcp 
        // y la segunda hacia el API para comparar ambos resultados
        const input = { currency: 'VES' };
        const results = await Promise.all([
            client.callTool({
                name: 'valor_moneda',
                arguments: input
            }),
            getMoneyValue(input.currency),
        ]);

        const [ result, value ] = results;
        const expected = `el valor actual de la moneda ${input.currency.toUpperCase()} frente al es: ${CURRENCY_BASE} es ${value.toFixed(6)}`;
        const [{ text }] = result.content; 

        // console.error({text, expected});

        t.assert.equal(text, expected);        
    });

    test.it('convertir moneda al tipo cambio', async t => {
        // para poder comparar los resultados debemos 
        // realizar 2 consultas una al mcp 
        // y la segunda hacia el API para comparar ambos resultados
        const input = { origin: 'EUR', destination: 'VES', amount: 1.00 };
        const results = await Promise.all([
            client.callTool({
                name: 'conversor_tipo_cambio',
                arguments: input
            }),
            getConvertion(input),
        ]);    
        
        const [ result, { value, rate } ] = results;
        const expected = `${input.amount} ${input.origin} = ${value.toFixed(2)} ${input.destination} (Tasa: ${rate.toFixed(6)}, Moneda Base ${CURRENCY_BASE})`;
        const [{ text }] = result.content; 

        // console.error({text, expected});

        t.assert.equal(text, expected);   
    });

    test.after(async () => await client.close());
});