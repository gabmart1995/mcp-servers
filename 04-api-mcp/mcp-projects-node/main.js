import path from 'node:path';

import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import z from 'zod';

const BASE_URL = 'http://localhost:3000/project';

async function main() {
    const server = new McpServer({
        name: 'mcp-projects',
        version: '1.0.0'
    });

    server.registerTool(
        'save_project', 
        {
            description: 'Crea un nuevo proyecto',
            title: 'Guarda un proyecto',
            inputSchema: {
                name: z.string(),
                description: z.string(),
                state: z.string()
            }
        },
        async ({ name, description, state }) => {
            const response = await fetch(`${BASE_URL}/save`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, description, state })
            });

            if (!response.ok) throw new Error('Error al guardar el proyecto');

            const data = await response.json();

            return {
                content: [
                    { type: 'text', text: JSON.stringify(data, null, 2) }
                ]
            };
        }
    );


    server.registerTool(
        'list_projects', 
        {
            description: 'Devuelve la lista de proyectos',
            title: 'Listar proyectos'
        },
        async () => {
            const response = await fetch(`${BASE_URL}/list`);

            if (!response.ok) throw new Error('Error al consultar el proyecto');

            const data = await response.json();

            return {
                content: [
                    { type: 'text', text: JSON.stringify(data, null, 2) }
                ]
            }
        }
    );

    server.registerTool(
        'list_project_id', 
        {
            description: 'Busca un proyecto usando un identificador',
            title: 'Busca un proyecto usando el id',
            inputSchema: {
                id: z.string(),
            }
        },
        async ({ id }) => {
            const response = await fetch(`${BASE_URL}/list/${id}`);

            if (!response.ok) throw new Error('Error al consultar los proyectos');

            const data = await response.json();

            return {
                content: [
                    { type: 'text', text: JSON.stringify(data, null, 2) }
                ]
            };
        }
    );

    server.registerTool(
        'update_project', 
        {
            description: 'Actualiza un proyecto usando un identificador',
            title: 'Actualiza un proyecto usando el id',
            inputSchema: {
                id: z.string(),
                name: z.string(),
                description: z.string(),
                state: z.string()
            }
        },
        async ({ id, name, description, state }) => {
            const response = await fetch(`${BASE_URL}/update`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name, description, state, id })
            });

            if (!response.ok) throw new Error('Error al actualizar el proyecto');

            const data = await response.json();

            return {
                content: [
                    { type: 'text', text: JSON.stringify(data, null, 2) }
                ]
            };
        }
    );

    server.registerTool(
        'delete_project', 
        {
            description: 'Borra un proyecto usando un identificador',
            title: 'Borra un proyecto usando el id',
            inputSchema: {
                id: z.string(),
            }
        },
        async ({ id }) => {
            const response = await fetch(`${BASE_URL}/delete/${id}`, {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ id })
            });

            if (!response.ok) throw new Error('Error al eliminar el proyecto');

            const data = await response.json();

            return {
                content: [
                    { type: 'text', text: JSON.stringify(data, null, 2) }
                ]
            };
        }
    );

    server.registerTool(
        'get_image_project', 
        {
            description: 'Obtiene la imagen del proyecto',
            title: 'Obtiene la imagen del proyecto',
            inputSchema: {
                file: z.string(),
            }
        },
        async ({ file }) => {
            const response = await fetch(`${BASE_URL}/image/${file}`);

            if (!response.ok) throw new Error('Error en la consulta');

            // extraemos la imagen del buffer
            const data = await response.arrayBuffer();
            const buffer = Buffer.from(data);
            const imageBase64 = buffer.toString('base64');
            const formats = {
                jpg: 'image/jpeg',
                jpeg: 'image/jpeg',
                png: 'image/png',
                git: 'image/gif'
            };

            // obtenemos la extension del archivo
            const extension = path
                .extname(file)
                .toLowerCase()
                .replace('.', '');

            if (!(extension in formats)) throw new Error('Error: formato no valido'); 

            return {
                content: [
                    { 
                        type: 'image', 
                        data:  imageBase64, // las imagenes deben venir en string base64
                        mimeType: formats[extension] // se debe especificar la extension
                    }
                ]
            };
        }
    );

    const transport = new StdioServerTransport();
    await server.connect(transport);
}

main().catch(console.error);