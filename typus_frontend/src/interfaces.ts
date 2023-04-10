export interface CodeCharacter {
    /**
     * This interface represents a single character of code.
     * 
     * @param {string}  c             - The character itself.
     * @param {boolean} wasTyped      - Represents if the character was already typed by a user. Should be false by default.
     * @param {boolean} isHighlighted - Determines if the character is highlighted. 
     *                                  This parameter is used for marking a character if it was typed incorrectly.
     */

    c: string;
    wasTyped: boolean;
    isHighlighted: boolean;
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

export interface CodeSample {
    id: number;
    title: string;
    content: string;
    langSlug: string;
}

export interface CodeSamples {
    samples: CodeSample[];
}