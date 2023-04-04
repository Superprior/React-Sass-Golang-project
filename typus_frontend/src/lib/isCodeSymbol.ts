
/**
 * Determines if a character is a valid symbol for putting it into the code field.
 * 
 * @param {string} c - The character which need to be checked. 
 * @returns {boolean} - True for valid and false for invalid characters.
 * 
 */
const isCodeSymbol = (c: string): boolean => {
    if (c.length !== 1) return false;
    if (c === '[' || c === ']') return true;

    let regex = new RegExp("[a-zA-Z0-9\t\n\"'`~!@#№$;%^:&*()-_+={}<>,.?/ ]");
    return regex.test(c);
}

export default isCodeSymbol;