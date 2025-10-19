import fs from 'node:fs/promises';

import '../models/models.js';

class BD {
    /** @type {BD} */
    static instance
    
    /** @type {{[uuid: string]: ProjectScheme}} */
    data = {};
    
    static getInstance() {
        if (!BD.instance) BD.instance = new BD(); 

        return BD.instance;
    }

    async connect() {
        await BD.getInstance().read();
    }
    
    async write() {
        try {
            await fs.writeFile('bd.json', JSON.stringify(BD.getInstance().data));

        } catch (error) {
            console.error(error);
        }
    }
    
    async read() {
        try {
            const data = await fs.readFile('bd.json', { encoding: 'utf-8' });
            BD.getInstance().data = JSON.parse(data);
    
        } catch (error) {
            console.error(error);
        }
    }
}

export default BD;