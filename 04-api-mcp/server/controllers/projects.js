import { randomUUID } from 'node:crypto';

import BD from '../config/configuration.js';

/**
 * Salva los cambios
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
async function save(request, response) {
    /** @type {ProjectScheme} */
    let body = request.body;
 
    if (!body.name || !body.description || !body.state) {
        return response.status(400).json({
            ok: false,
            status: 400,
            message: 'Faltan datos por enviar'
        });
    }

    /** @type {ProjectScheme} */
    let project = {
        name: body.name,
        description: body.description,
        state: body.state,
        created_at: new Date().toISOString(),
        image: body.image || undefined
    }

    let uuid = randomUUID();

    // establecemos el nuevo valor
    (BD.getInstance()).data[uuid] = project;

    try {
        await (BD.getInstance()).write((BD.getInstance()).data);

        return response.status(200).json({
            ok: true,
            status: 200,
            message: 'datos insertado con éxito',
            project
        });

    } catch (error) {
        console.error(error);
        
        return response.status(500).json({
            ok: false,
            status: 500,
            message: 'error al almacenar en BD',
        });
    }

}

export default {
    save
}