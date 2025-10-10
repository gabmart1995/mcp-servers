/**
 * Cliente MCP personalizado para el ejercicio del clima
 * es una alternativa para todas aquellos paises que no 
 * pueden descargar Claude Desktop.
 */
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';

async function main() {
    // 1.- generamos el protocolo de transporte
    const transport = new StdioClientTransport({
        command: 'node', // or node server.mjs | node server.cjs
        args: ['server.mjs']
    });

    // 2.- generamos el cliente
    const client = new Client({
        name: 'fetch-weather-client',
        version: '0.0.1'
    }); 

    // 3.- realizamos la conexion 
    await client.connect(transport);

    // 4.- Ejecutamos la herramienta del MCP
    const result = await client.callTool({
        name: 'fetch-weather',
        arguments: {
            city: 'Barcelona'
        }
    });

    // Nota: evitar usar console.log ya que el protocolo STDIO
    // usa el flujo stdin del computador en su caso
    // usamos console.error
    console.error(result);

    // 5.- cerramos la conexión para gestionar los recursos
    await client.close();
}


main().catch(err => {
    console.error(err);
    process.exit(1);
});
