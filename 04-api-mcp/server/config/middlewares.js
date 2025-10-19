/**
 * CORS Middleware
 * @param {express.Request} request 
 * @param {express.Response} response 
 * @param {express.NextFunction} next 
 */
function CORSMiddleware(request, response, next) {
    response.setHeader('Access-Control-Allow-Origin', '*');
    response.setHeader('Access-Control-Allow-Credentials', 'true');
    response.setHeader('Access-Control-Allow-Headers', 'Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With');
    response.setHeader('Access-Control-Allow-Methods', 'POST, OPTIONS, GET, PUT, DELETE');

    // -- esta cabecera permite durante el tiempo indicado almacenar en cache los
    // -- datos de las peticiones preflight CORS
    // -- que son las peticiones OPTIONS que se mandan antes de cada peticion
    response.setHeader('Access-Control-Max-Age', '300');

    if (request.method === 'OPTIONS') {
        response.status(204);
        return;
    }

    next(); // ejecuta la siguiente funcion
};

export {
    CORSMiddleware
}