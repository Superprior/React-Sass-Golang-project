

const isCodeSymbol = (c: string): boolean => {
    if (c.length !== 1) return false;
    if (c == '[' || c == ']') return true;

    let regex = new RegExp("[a-zA-Z0-9\t\n\"'`~!@#№$;%^:&*()-_+={}<>,.?/ ]");
    return regex.test(c);
}

export default isCodeSymbol;