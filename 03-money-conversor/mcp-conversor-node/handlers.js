require('./types/jsdoc_types'); // types jsdoc

const API_URL = 'https://cdn.moneyconvert.net/api/latest.json';
const CURRENCY_BASE = 'USD';

/**
 * Consulta el valor de la moneda frente al dolar
 * @param {string} currency
 * @returns {Promise<number>}
 */
async function getMoneyValue(currency) {
    const response = await fetch(API_URL);

    if (!response.ok) throw new Error('Error al acceder a la url del API');

    /** @type {ExchangeRequest} */
    const { rates } = await response.json();

    /** @type {number | undefined} */
    const value = rates[currency.toUpperCase()];

    if (!value) throw new Error('No se encontro la moneda solicitada: ' + currency);

    return value;
}

/**
 * manejador de conversion de tipos
 * @param {{origin: string, destination: string, amount: number}} data
 * @returns {Promise<{rate: number, value: number}>}
 */
async function getConvertion({ origin, destination, amount }) {
    const response = await fetch(API_URL);

    if (!response.ok) throw new Error('Error al acceder a la url del API');

    /** @type {ExchangeRequest} */
    const data = await response.json();
    let { base = CURRENCY_BASE, rates } = data;

    if (!base || !rates) {
        throw new Error('No se encontro la moneda solicitada: ' + currency);
    } 

    let rate = 0.00;

    // si el origen es igual al destino
    if (origin === base) {
        rate = rates[destination];
    
    } else {
        // conversion indirecta se divide entre la tasa de origen
        // y destino
        rate = rates[destination] / rates[origin];
    }
    
    const value = amount * rate;

    return { value, rate };
}

module.exports = {
    API_URL,
    CURRENCY_BASE,
    getMoneyValue,
    getConvertion
};