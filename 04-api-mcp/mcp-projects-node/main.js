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
            title: 'Gaurdar un proyecto',
            inputSchema: {
                name: z.string(),
                description: z.description(),
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
                    { type: 'text', text: JSON.stringify(data, null, 2)}
                ]
            }
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
                    { type: 'text', text: JSON.stringify(data, null, 2)}
                ]
            }
        }
    );


    const transport = new StdioServerTransport();
    await server.connect(transport);
}

main().catch(console.error);