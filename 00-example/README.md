# MCP Weather-Example
Este proyecto es una integración de modelos MCP (Model Context Protocol) para inteligencia artificial

El objetivo es mejorar la interaccion con el usuario con sus clientes MCP. Pueden acceder al sistema de archivos, bases de datos o consultas externas a servicios REST para obtener información y mejorar su resultado.

Este servicio consulta la información del clima de una ciudad especifica. 

Se incluye el codigo del cliente para el desarrollo de herramientas del lado del
frontend con el objetivo de hacer testing, para evitar la necesidad de instalar 
claude desktop

## Requerimientos
- NodeJS >= V.18
- Windsurf, Cursor o VS code con Copilot con IA incorporada

### comandos
`npm install` - instala las dependencias<br/>
`npm run server` - corre el servicio MCP<br />
`npm run inspector` - corre el servicio de testing MCP abre una pestaña en el navegador.<br />
`npm run client` - corre un cliente mcp para usar la conexion sin el inspector.
