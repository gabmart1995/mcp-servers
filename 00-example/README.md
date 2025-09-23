# MCP Weather-Example
Este proyecto es una integración de modelos MCP (Model Context Protocol) para inteligencia artificial

El objetivo es mejorar la interaccion con el usuario con sus clientes MCP. Pueden acceder al sistema de archivos, bases de datos o consultas externas a servicios REST para obtener información y mejorar su resultado.

Este servicio consulta la información del clima de una ciudad especifica. 

Se incluye el codigo del cliente para el desarrollo de herramientas del lado del
frontend con el objetivo de hacer testing, para evitar la necesidad de instalar 
claude desktop

Se utiliza la tecnologia de SEA para transformar el codigo en un ejecutable

## Requerimientos
- NodeJS >= V.22
- Windsurf, Cursor o VS code con Copilot con IA incorporada
- Esbuild instalado en el SO configurado en el PATH.

### comandos
`npm install` - instala las dependencias<br/>
`npm run server` - corre el servicio MCP<br />
`npm run inspector` - corre el servicio de testing MCP abre una pestaña en el navegador.<br />
`npm run client` - corre un cliente mcp para usar la conexion sin el inspector.<br />
`npm run build` - ejecuta esbuild para compilar los modulos en un solo archivo. para SEA<br />
`npm run build-linux` - ejecuta la compilacion para sistemas linux x64 <br />
`npm run build-linux` - ejecuta la compilacion para sistemas windows x64

