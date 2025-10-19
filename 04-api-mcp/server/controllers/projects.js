import fs from 'node:fs';
import path from 'node:path';
import { randomUUID } from 'node:crypto';

import BD from '../config/bd.js';

/**
 * Salva los cambios
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
async function save(request, response) {
    /** @type {ProjectScheme} */
    let body = request.body;
 
    if (!body || !body.name || !body.description || !body.state) {
        return response.status(400).json({
            ok: false,
            status: 400,
            message: 'Faltan datos por enviar'
        });
    }

    let uuid = randomUUID();
    
    /** @type {ProjectScheme} */
    let project = {
        ...body,
        id: uuid,
        created_at: new Date().toISOString(),
        image: body.image || undefined
    };

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

/**
 * Lista los proyectos en la colección
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
function list(request, response) {
    const data = BD.getInstance().data;

    if ((Object.keys(data)).length === 0) {
        return response.status(404).json({
            ok: false,
            status: 404,
            message: 'No existen proyectos que mostrar'
        });
    }

    return response.status(200).json({
        ok: true,
        status: 200,
        projects: Object.values(data)
    });
}

/**
 * Retorna un proyecto en especifico usando el UUID
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
function item(request, response) {
    /** @type {string} */
    let id = request.params.id;
    
    const data = BD.getInstance().data;
    const find = id in data;

    if (!find) {
        return response.status(404).json({
            ok: false,
            status: 404,
            message: 'El proyecto solicitado no existe'
        });
    }

    return response.status(200).json({
        ok: true,
        status: 200,
        project: data[id]
    });
}

/**
 * Borra un proyecto en especifico
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
async function deleteProject(request, response) {
    /** @type {string} */
    let id = request.params.id;
    
    const data = { ...BD.getInstance().data };
    const find = id in data;

    if (!find) {
        return response.status(404).json({
            ok: false,
            status: 404,
            message: 'El proyecto solicitado no existe'
        });
    }

    // eliminamos el proyecto en memoria
    delete (BD.getInstance()).data[id];

    // actualizamos el archivo en BD
    try {
        await (BD.getInstance()).write();
        
        return response.status(200).json({
            ok: true,
            status: 200,
            message: 'Operación exitosa',
            deleted: data[id],
        });

    } catch (error) {
        console.error(error);
        
        return response.status(500).json({
            ok: false,
            status: 500,
            message: 'error al escribir en BD',
        }); 
    }
}


/**
 * Actualiza un proyecto en especifico
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
async function update(request, response) {
    /** @type {ProjectScheme} */
    let body = request.body;
    
    // verificamos los campos
    if (!body || !body.name || !body.description || !body.state || !body.id) {
        return response.status(400).json({
            ok: false,
            status: 400,
            message: 'Faltan datos por enviar'
        });
    }

    // verificamos la informacion en base de datos
    const data = { ...BD.getInstance().data };
    const find = body.id in data;

    if (!find) {
        return response.status(404).json({
            ok: false,
            status: 404,
            message: 'El proyecto solicitado no existe'
        });
    }

    /** @type {ProjectScheme} */
    let project = {
        ...body,
        created_at: (data[body.id]).created_at || undefined,
        image: body.image || undefined
    }

    // borramos el id
    delete project.id;

    // actualizamos en memoria
    (BD.getInstance()).data[body.id] = project;

    // actualizamos el archivo en BD
    try {
        await (BD.getInstance()).write();
        
        return response.status(200).json({
            ok: true,
            status: 200,
            message: 'Operación exitosa',
            updated: data[body.id],
        });

    } catch (error) {
        console.error(error);
        
        return response.status(500).json({
            ok: false,
            status: 500,
            message: 'error al escribir en BD',
        }); 
    }
}

/**
 * permite la carga de archivos usando form data
 * hacia carpeta en especifico y actualiza el registro
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
async function upload(request, response) {
    /** @type {string} */
    let id = request.params.id;
    
    if (!request.file) {
        return response.status(404).json({
            ok: false,
            status: 404,
            message: 'No se ha subido ningun archivo'
        });
    }

    // filtramos las extensiones
    const filePath = request.file.path;
    const extension = path
        .extname(request.file.originalname)
        .toLowerCase()
        .replace('.', '');

    const validExtensions = ['png', 'jpg', 'jpeg', 'gif'];

    if (!validExtensions.includes(extension)) {
        fs.unlinkSync(filePath);
    
        return response.status(400).json({
            ok: true,
            status: 400,
            message: 'Extensión no valida',     
        });
    }
    
    const projectUpdate = (BD.getInstance()).data[id];

    // si existe el registro de la imagen
    if (projectUpdate.image) {
        const oldImagePath = `./uploads/images/${projectUpdate.image}`;

        // verifica si existe en el disco y la borra
        if (fs.existsSync(oldImagePath)) fs.unlinkSync(oldImagePath);
    }

    // actualizamos la imagen del proyecto
    (BD.getInstance()).data[id].image = request.file.filename;

    // actualizamos el archivo en BD
    try {
        await (BD.getInstance()).write();
        
        return response.status(200).json({
            ok: true,
            status: 200,
            message: 'Proyecto actualizado',
            project: projectUpdate,
            newFile: request.file.filename 
        });

    } catch (error) {
        console.error(error);
        
        return response.status(500).json({
            ok: false,
            status: 500,
            message: 'error al escribir en BD',
        }); 
    }
}

/**
 * permite visualizar las imagenes
 * @param {Express.Request} request 
 * @param {Express.Response} response 
 */
function getImage(request, response) {
    /** @type {string} */
    let fileName = request.params.file;
    let filePath = `./uploads/images/${fileName}`;

    fs.stat(filePath, (error, stat) => {
        if (!error && stat) return response.sendFile(path.resolve(filePath));

        return response.status(404).json({
            ok: true,
            status: 404,
            message: 'La imagen no existe'
        });
    });
}

export default {
    save,
    list,
    item,
    update,
    deleteProject,
    upload,
    getImage                
}