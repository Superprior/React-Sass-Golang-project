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
                ["i", "m", "p", "o", "r", "t", " ", "h", "a", "s", "h", "l", "i", "b"],
                [],
                ["d", "e", "f", " ", "h", "a", "s", "h", "_", "f", "i", "l", "e", "(", "f", "i", "l", "e", "n", "a", "m", "e", ")", ":"],
                [" ", " ", " ", " ", "h", " ", "=", " ", "h", "a", "s", "h", "l", "i", "b", ".", "s", "h", "a", "1", "(", ")"]
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
        if (isCodeSymbol(event.key)) {
            const currentSymbolToType = this.state.lines[this.state.cursor.y].chars[this.state.cursor.x].c;
            if (event.key === currentSymbolToType) {
                this.state.lines[this.state.cursor.y].chars[this.state.cursor.x].wasTyped = true;
                this.setState({ cursor: {x: this.state.cursor.x + 1, y: this.state.cursor.y }});
            }   
        }

        switch (event.key) {
            case "ArrowRight": {
                if (this.state.cursor.x < this.state.lines[this.state.cursor.y].chars.length) {
                    this.setState({ cursor: { x: this.state.cursor.x + 1, y: this.state.cursor.y }})
                }
                break;
            }
            case "ArrowLeft": {
                if (this.state.cursor.x > 0) {
                    this.setState({ cursor: { x: this.state.cursor.x - 1, y: this.state.cursor.y }})
                }
                break;
            }
            case "ArrowUp": {
                if (this.state.cursor.y > 0) {
                    this.setState({ cursor: { x: this.state.cursor.x, y: this.state.cursor.y - 1 }});
                    if (this.state.lines[this.state.cursor.y - 1].chars.length <= this.state.cursor.x) {
                        this.setState({ cursor: { x: this.state.lines[this.state.cursor.y - 1].chars.length, y: this.state.cursor.y - 1 }})
                    }
                }
                break;
            }
            case "ArrowDown": {
                if (this.state.cursor.y < this.state.lines.length - 1) {
                    this.setState({ cursor: { x: this.state.cursor.x, y: this.state.cursor.y + 1 }});
                    if (this.state.lines[this.state.cursor.y + 1].chars.length <= this.state.cursor.x) {
                        this.setState({ cursor: { x: this.state.lines[this.state.cursor.y + 1].chars.length, y: this.state.cursor.y + 1 }});
                    }
                } 
                break;
            }
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