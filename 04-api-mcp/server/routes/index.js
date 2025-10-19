import { Router } from 'express';
import ProjectController from '../controllers/projects.js';

const routes = Router();

// rutas de la API 
routes.post('/save', ProjectController.save);

export { routes };