import React from 'react'
import './styles/code-area.sass'
import { CodeLine, Cursor } from '../interfaces';


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
    lines: CodeLine[] | null;
}


class CodeArea extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            cursor: { x: 0, y: 0 },
            lines: null,
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
        // document.addEventListener("keydown", this.handleCursorMovement); 
    }

    // handleCursorMovement = (event: KeyboardEvent): void => {
    //     alert(event.key);
    //     switch (event.key) {
    //         case "ArrowRight": {
    //             if (this.state.cursorCoords[0] < this.state.lines[this.state.cursorCoords[1]].length) {
    //                 this.setState({ cursorCoords: [this.state.cursorCoords[0] + 1, this.state.cursorCoords[1]]});
    //             }
    //             break;
    //         }
    //         case "ArrowLeft": {
    //             if (this.state.cursorCoords[0] > 0) {
    //                 this.setState({ cursorCoords: [this.state.cursorCoords[0] - 1, this.state.cursorCoords[1]]});
    //             }
    //             break;
    //         }
    //         case "ArrowUp": {
    //             if (this.state.cursorCoords[1] > 0) {
    //                 this.setState({ cursorCoords: [this.state.cursorCoords[0], this.state.cursorCoords[1] - 1]});
    //                 if (this.state.lines[this.state.cursorCoords[1] - 1].length <= this.state.cursorCoords[0]) {
    //                     this.setState({ cursorCoords: [this.state.lines[this.state.cursorCoords[1] - 1].length, this.state.cursorCoords[1] - 1] })
    //                 }
    //             }
    //             break;
    //         }
    //         case "ArrowDown": {
    //             if (this.state.cursorCoords[1] < this.state.lineNumbers.length - 1) {
    //                 this.setState({ cursorCoords: [this.state.cursorCoords[0], this.state.cursorCoords[1] + 1]});
    //                 if (this.state.lines[this.state.cursorCoords[1] + 1].length <= this.state.cursorCoords[0]) {
    //                     this.setState({ cursorCoords: [this.state.lines[this.state.cursorCoords[1] + 1].length, this.state.cursorCoords[1] + 1]});
    //                 }
    //             } 
    //             break;
    //         }
    //     }
    // }

    render() {
        return (
            <>
                <div>
                    <h1>--- Code Area ---</h1>
                    <button onClick={() => console.log(this.state.lines)}>
                        View lines state
                    </button>
                    {/* <div className='code-area-wrapper'>
                        {
                            this.state.lineNumbers.map((lineNumber: number, yCoord: number) => {
                                return (
                                    <div className='line'>
                                        <div className='line-number-wrapper'>
                                            <span className='line-number'>
                                                { lineNumber }
                                            </span>
                                        </div>
                                        <div className='line-code-wrapper'>
                                            { this.state.lines[lineNumber - 1].map((char: string, xCoord: number) => {
                                                return (
                                                    <div style={{ display: 'flex' }}>
                                                        <div style={{ display: 'flex'}}>
                                                            <span className='line-code'>
                                                                { char }
                                                            </span>
                                                            { this.state.cursorCoords[0] === xCoord && this.state.cursorCoords[1] === yCoord ? (
                                                                <span className='cursor'></span>
                                                            ) : null }
                                                        </div>
                                                        { this.state.cursorCoords[0] === this.state.lines[yCoord].length && xCoord + 1 === this.state.lines[yCoord].length &&  this.state.cursorCoords[1] === yCoord ? (
                                                            <span className='cursor' style={{ position: 'relative' }}></span>
                                                        ) : null }
                                                    </div>
                                                )
                                            })}
                                            { this.state.lines[lineNumber - 1].length === 0 && this.state.cursorCoords[1] === yCoord ? (
                                                <span className='cursor' style={{ position: 'relative' }}></span>
                                            ) : null }
                                        </div>
                                    </div>
                                )
                            })
                        }
                    </div> */}
                </div>
            </>
        )
    }
}

export default CodeArea;