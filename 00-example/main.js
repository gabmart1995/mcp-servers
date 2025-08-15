import {McpServer} from '@modelcontextprotocol/sdk/server/mcp.js';
import {StdioServerTransport} from "@modelcontextprotocol/sdk/server/stdio.js";
import * as z from 'zod';

const GEOCODING_API = 'https://geocoding-api.open-meteo.com/v1/';
const WEATHER_API = 'https://api.open-meteo.com/v1/';

async function main() {
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
        {city: z.string().describe('City Name')},
        async function({city}) {  // funcion manejadora del servicio
            // solicitamos la informacion a la geolocalizacion
            const response = await fetch(`${GEOCODING_API}search?name=${city}&count=10&language=en&format=json`);
            const data = await response.json();

            // sino llega informacion sobre la ciudad desde el API
            if (data.length === 0) {
                return {
                    content: [
                        {type: 'text', text: `No se encontro el clima para la ciudad de ${city}`}
                    ],
                };
            }

            // extraemos los datos de la consulta
            const {latitude, longitude} = data.results[0];
            
            // consulta el api del clima
            const responseWeather = await fetch(`${WEATHER_API}forecast?latitude=${latitude}&longitude=${longitude}&hourly=temperature_2m,precipitation,rain,is_day`)
            const weatherData = await responseWeather.json();

            // le retornamos a la IA los datos crudos para que los procese hacia 
            // el usuario
            return {
                content: [
                    {type: 'text', text: JSON.stringify(weatherData, null, 2)}
                ],
            };
        }
    );

    // 3.- Escuchar las conexiones del cliente
    const transport = new StdioServerTransport();
    await server.connect(transport);
}

// ejecutamos la promesa principal
main().catch(err => {
    console.error(err);
    process.exit(1);
});
