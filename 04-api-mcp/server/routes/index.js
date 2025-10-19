import { Router } from 'express';
import ProjectController from '../controllers/projects.js';

const routes = Router();

// rutas de la API 
routes.post('/project/save', ProjectController.save);
routes.get('/project/list', ProjectController.list);
routes.get('/project/list/:id', ProjectController.item);
routes.delete('/project/delete/:id', ProjectController.deleteProject);
routes.put('/project/update', ProjectController.update);

export { routes };