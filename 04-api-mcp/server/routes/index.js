import { Router } from 'express';

import upload from '../config/multer.js';
import ProjectController from '../controllers/projects.js';

const routes = Router();

// rutas de la API 
routes.post('/project/save', ProjectController.save);
routes.get('/project/list', ProjectController.list);
routes.get('/project/list/:id', ProjectController.item);
routes.delete('/project/delete/:id', ProjectController.deleteProject);
routes.put('/project/update', ProjectController.update);
routes.post('/project/upload/:id', upload.single('file_0'), ProjectController.upload);
routes.get('/project/image/:file', ProjectController.getImage);

export { routes };