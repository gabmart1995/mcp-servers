# script que permite la creacion de bianrios SEA en NodeJS en Linux
# Version Node 22

# Genera el blob
node --experimental-sea-config sea-config.json

# Creamos una copia del binario instalado de node y lo reinstalamos
cp $(command -v node) mcp-resources

# Inject el blob script en el binario de node
npx postject mcp-resources NODE_SEA_BLOB main.blob --sentinel-fuse NODE_SEA_FUSE_fce680ab2cc467b6e072b8b5df1996b2