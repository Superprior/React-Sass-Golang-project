import React from 'react'
import './styles/code-area.sass'
import { CodeCharacter, CodeLine, Cursor } from '../interfaces';
import isCodeSymbol from '../lib/isCodeSymbol';



interface Props {}

interface State {
    /**
     * This interface represents the state of the CodeArea. 
     * 
     * @param {Cursor} cursor - The cursor.
     * @param {CodeLine[]} lines - An array of the code lines. Gets its value in componentDidMount().
     * 
     */

    cursor: Cursor;
    lines: CodeLine[];
}


class CodeArea extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            cursor: { x: 0, y: 0 },
            lines: [],
        }
    }

    async componentDidMount() {
        const initialHardcodedCodeLines = [
                ['i', 'm', 'p', 'o', 'r', 't', ' ', 'h', 'a', 's', 'h', 'l', 'i', 'b'], 
                [], 
                ['d', 'e', 'f', ' ', 'h', 'a', 's', 'h', '_', 'f', 'i', 'l', 'e', '(', 'f', 'i', 'l', 'e', 'n', 'a', 'm', 'e', ')', ':'], 
                [' ', ' ', ' ', 'h', ' ', '=', ' ', 'h', 'a', 's', 'h', 'l', 'i', 'b', '.', 's', 'h', 'a', '1', '(', ')'], 
                [], 
                [' ', ' ', ' ', 'w', 'i', 't', 'h', ' ', 'o', 'p', 'e', 'n', '(', 'f', 'i', 'l', 'e', 'n', 'a', 'm', 'e', ',', "'", 'r', 'b', "'", ')', ' ', 'a', 's', ' ', 'f', 'i', 'l', 'e', ':'],
                [' ', ' ', ' ', ' ', ' ', ' ', ' ', 'c', 'h', 'u', 'n', 'k', ' ', '=', ' ', '0'], 
                [' ', ' ', ' ', ' ', ' ', ' ', ' ', 'w', 'h', 'i', 'l', 'e', ' ', 'c', 'h', 'u', 'n', 'k', ' ', '!', '=', ' ', 'b', "'", "'", ':'], 
                [' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'c', 'h', 'u', 'n', 'k', ' ', '=', ' ', 'f', 'i', 'l', 'e', '.', 'r', 'e', 'a', 'd', '(', '1', '0', '2', '4', ')'], 
                [' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'h', '.', 'u', 'p', 'd', 'a', 't', 'e', '(', 'c', 'h', 'u', 'n', 'k', ')', '\n'], ['\n'], [' ', ' ', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', 'h', '.', 'h', 'e', 'x', 'd', 'i', 'g', 'e', 's', 't', '(', ')'],
                [], 
                ['m', 'e', 's', 's', 'a', 'g', 'e', ' ', '=', ' ', 'h', 'a', 's', 'h', '_', 'f', 'i', 'l', 'e', '(', '"', 't', 'r', 'a', 'c', 'k', '1', '.', 'm', 'p', '3', '"', ')'], 
                ['p', 'r', 'i', 'n', 't', '(', 'm', 'e', 's', 's', 'a', 'g', 'e', ')']
        ];
        
        let hardcodedCodeLines: CodeLine[] = [];

        for (let i = 0; i < initialHardcodedCodeLines.length; i++) {
            let codeLine: CodeLine = { chars: [] };
            for (let j = 0; j < initialHardcodedCodeLines[i].length; j++) {
                codeLine.chars.push({ c: initialHardcodedCodeLines[i][j], wasTyped: false });
            }
            hardcodedCodeLines.push(codeLine);
        }

        this.setState({ lines: hardcodedCodeLines })
        document.addEventListener("keydown", this.handleKeyboard);
    }

    handleKeyboard = (event: KeyboardEvent): void => {
        if (isCodeSymbol(event.key) && this.state.cursor.x < this.state.lines[this.state.cursor.y].chars.length) {
            const currentSymbolToType = this.state.lines[this.state.cursor.y].chars[this.state.cursor.x].c;
            if (event.key === currentSymbolToType) {
                this.state.lines[this.state.cursor.y].chars[this.state.cursor.x].wasTyped = true;
                this.setState({ cursor: {x: this.state.cursor.x + 1, y: this.state.cursor.y }});
            }   
            return;
        } else if (event.key === "Enter" && this.state.cursor.y < this.state.lines.length - 1 && 
                   this.state.cursor.x === this.state.lines[this.state.cursor.y].chars.length) {
            this.setState({ cursor: { x: 0, y: this.state.cursor.y + 1 }});
            return;
        }  
    }

    render() {
        return (
            <>
                <div>
                    <div className='code-area-wrapper'>
                        {
                            this.state.lines?.map((line: CodeLine, lineNumber: number) => {
                                return (
                                    <div className='line' key={lineNumber}>
                                        <div className='line-number-wrapper'>
                                            <span className='line-number'>
                                                { lineNumber + 1 }
                                            </span>
                                        </div>
                                        <div className='line-code-wrapper'>
                                            {    
                                                this.state.lines[lineNumber].chars.map((char: CodeCharacter, charIndex: number) => {
                                                    return (
                                                        <div style={{ display: 'flex' }} key={`${lineNumber}:${charIndex}`}>
                                                            <div style={{ display: 'flex'}}>
                                                                <span className='line-code' style={{ opacity: `${char.wasTyped ? '1' : '0.5'}` }}>
                                                                    { char.c }
                                                                </span>
                                                                { this.state.cursor.x === charIndex && this.state.cursor.y === lineNumber ? (
                                                                    <span className='cursor'></span>
                                                                ) : null }
                                                            </div>
                                                            { this.state.cursor.x === this.state.lines[lineNumber].chars.length && 
                                                              charIndex + 1 === this.state.lines[lineNumber].chars.length &&
                                                              this.state.cursor.y === lineNumber ? (
                                                                <span className='cursor' style={{ position: 'relative' }}></span>
                                                            ) : null }
                                                        </div>
                                                    )
                                                })
                                            }
                                            { this.state.lines[lineNumber].chars.length === 0 && this.state.cursor.y === lineNumber ? (
                                                <span className='cursor' style={{ position: 'relative' }}></span>
                                            ) : null }
                                        </div>
                                    </div>
                                )
                            })
                        }
                    </div>
                </div>
            </>
        )
    }
}

export default CodeArea;