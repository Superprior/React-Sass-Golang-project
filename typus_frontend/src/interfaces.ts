export interface CodeCharacter {
    /**
     * This interface represents a single character of code.
     * 
     * @param {string} c - The character itself.
     * @param {boolean} wasTyped - Represents if the character was already typed by a user. Should be false by default.
     * 
     */

    c: string;
    wasTyped: boolean;
}

export interface CodeLine {
    /**
     * This interface represents a line of code.
     * 
     * @param {CodeCharacter[]} chars - Stores an array of characters to type.
     * 
     */

    chars: CodeCharacter[];
}

export interface Cursor {
    /**
     * This interface represents a cursor which is used in a code field.
     * 
     * @param {number} x - Coordinate by the X-axis
     * @param {number} y - Coordinate by the Y-axis
     * 
     */

    x: number;
    y: number;
}