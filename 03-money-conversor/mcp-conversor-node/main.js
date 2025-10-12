const { McpServer } = require('@modelcontextprotocol/sdk/server/mcp.js');
const { StdioServerTransport } = require('@modelcontextprotocol/sdk/server/stdio.js');
const { default: z } = require('zod');

const { getMoneyValue, CURRENCY_BASE, getConvertion } = require('./handlers');

async function main() {
    const server = new McpServer({
        name: 'mcp-money-conversor',
        version: '0.0.1'
    });

    server.tool(
        'valor_moneda',
        'Devuelve el valor actual de una moneda que necesites (USD, EUR, etc)',
        { currency: z.string().min(1, 'Debes indicar la moneda, EUR, USD') },
        async ({ currency }) => {
            const value = await getMoneyValue(currency);
            return {
                content: [
                    {
                        type: 'text',
                        text: `el valor actual de la moneda ${currency.toUpperCase()} frente al es: ${CURRENCY_BASE} es ${value.toFixed(6)}`,
                    }
                ]
            };
        }
    );

    server.tool(
        'conversor_tipo_cambio',
        'Devuelve el valor actual de una moneda frente a otra',
        { 
            origin: z.string().length(3, 'Debe ser un codigo ISO de 3 letras que represente a una moneda'),
            destination: z.string().length(3, 'Debe ser un codigo ISO de 3 letras que represente a una moneda'),
            amount: z.number(),
        },
        async ({ origin, destination, amount }) => {
            const { value, rate } = await getConvertion({ origin, destination, amount });
            return {
                content: [
                    { 
                        type: 'text', 
                        text: `${amount} ${origin} = ${value.toFixed(2)} ${destination} (Tasa: ${rate.toFixed(6)}, Moneda Base ${CURRENCY_BASE})`,
                    }
                ]
            }
        }
    );

    const transport = new StdioServerTransport();
    await server.connect(transport);
}


main().catch(err => { 
    console.error(err);
    process.exit(1);
});