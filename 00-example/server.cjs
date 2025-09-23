/**
 * Conversion de servicio de MCP a un binario ejecutable
 * usando Single executable application (SEA) de NodeJS. 
 * 
 * Importante
 * 
 * 1.- se deben transformar los import a require dentro del contexto de SEA
 * 2.- los modulos deben transformarse a commonJS (Obligatorio)
 * 3.- Se debe usar un empaquetador para juntar todas las dependiencias en un solo
 * archivo puedes usar esbuild (recomendado) o webpack y transpilarlo a node22, es obligatorio
 * ya que los binarios deben funcionar sin dependencias externas.
 * 4.- ejecutar el make dependiendo de tu sistema operativo. crea una copia de ejecutable 
 * de node e inyecta el codigo dentro del archivo bundle.js generado por el empaquetador
 * 5.- listo ya tienes el ejecutable para tu sistema.
 */

const sea = require('node:sea');

if (sea.isSea()) {
    const {createRequire} = require('node:module');
    require = createRequire(__filename);
}

const {McpServer} = require('@modelcontextprotocol/sdk/server/mcp.js');
const {StdioServerTransport} = require('@modelcontextprotocol/sdk/server/stdio.js');
const {default: z} = require('zod');

const GEOCODING_API = 'https://geocoding-api.open-meteo.com/v1/';
const WEATHER_API = 'https://api.open-meteo.com/v1/';

async function main() {
    async function handleFetchWeather({ city }) {
        const response = await fetch(`${GEOCODING_API}search?name=${city}&count=10&language=en&format=json`);
        if (!response.ok) {
            return {
                content: [
                    { type: 'text', text: `El servicio de localizacion no se encuentra activo en estos momentos` }
                ],
            };
        }

        const data = await response.json();
        if (data.length === 0) {
            return {
                content: [
                    { type: 'text', text: `No se encontro el clima para la ciudad de ${city}` }
                ],
            };
        }

        // extraemos los datos de la consulta
        const { latitude, longitude } = data.results[0];

        // consulta el api del clima
        const responseWeather = await fetch(`${WEATHER_API}forecast?latitude=${latitude}&longitude=${longitude}&hourly=temperature_2m,precipitation,rain,is_day`)
        if (!responseWeather.ok) {
            return {
                content: [
                    { type: 'text', text: `El servicio del clima no se encuentra activo en estos momentos` }
                ],
            };
        }

        const weatherData = await responseWeather.json();

        // le retornamos a la IA los datos crudos para que los procese hacia 
        // el usuario
        return {
            content: [
                { type: 'text', text: JSON.stringify(weatherData, null, 2) }
            ],
        };
    }

    // 1.- crear el servidor:
    // es la interfaz pricipal con el repositorio MCP.
    const server = new McpServer({
        name: 'Demo',
        version: '0.0.1',
        capabilities: {
            resources: {},
            tools: {}
        },
    });

    // 2.- definir las herramientas LLM
    // permite al LLM realizar acciones dentro del servidor
    server.tool(
        'fetch-weather',
        'Tool a fetch weather of a city',
        { city: z.string().describe('City Name') },
        handleFetchWeather
    );

    // 3.- Escuchar las conexiones del cliente
    // lo dejamos abierto para que pueda recibir las entradas
    const transport = new StdioServerTransport();
    await server.connect(transport);
}

// ejecutamos la promesa principal
main().catch(err => {
    console.error(err);
    process.exit(1);
});
