import express from 'express';

import { CORSMiddleware } from './config/cors.js';
import { routes } from './routes/index.js';
import BD from './config/bd.js';

async function main() {
    await BD.getInstance().connect();
        
    const app = express();
    const PORT = 3000;
    
    // establecemos los middlewares 
    app.use([
        CORSMiddleware, 
        express.json(), 
        express.urlencoded({ extended: true }),
        routes // routes API
    ]);
    
    
    app.listen(PORT, () => console.log(`Server running in port ${PORT}`));
}

main().catch(console.error);

