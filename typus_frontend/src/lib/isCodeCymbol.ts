

const isCodeCymbol = (c: string): boolean => {
    if (c.length !== 1) {
        throw new Error('Got multiple characters instead of one');
    }
    if (c == '[' || c == ']') return true;

    let regex = new RegExp("[a-zA-Z0-9\t\n\"'`~!@#№$;%^:&*()-_+={}<>,.?/]");
    return regex.test(c);
}

export default isCodeCymbol;