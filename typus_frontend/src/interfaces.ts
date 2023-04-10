export interface CodeCharacter {
    /**
     * This interface represents a single character of code.
     * 
     * @property {string}  c             - The character itself.
     * @property {boolean} wasTyped      - Represents if the character was already typed by a user. Should be false by default.
     * @property {boolean} isHighlighted - Determines if the character is highlighted. 
     *                                     This property is used for marking a character if it was typed incorrectly.
     * @interface
     */

    c: string;
    wasTyped: boolean;
    isHighlighted: boolean;
}

export interface CodeLine {
    /**
     * This interface represents a line of code.
     * 
     * @property {CodeCharacter[]} chars - Stores an array of characters to type.
     * 
     * @interface
     */

    chars: CodeCharacter[];
}

export interface Cursor {
    /**
     * This interface represents a cursor which is used in a code field.
     * 
     * @property {number} x - Coordinate by the X-axis
     * @property {number} y - Coordinate by the Y-axis
     * 
     * @interface
     */

    x: number;
    y: number;
}
