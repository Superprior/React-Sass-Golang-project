import React from 'react'
import './styles/code-area.sass'

interface Props {}

interface State {
    lineNumbers: number[];
    lines: string[][];
    cursorCoords: [number, number];
}


class CodeArea extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            lineNumbers: [1, 2, 3, 4], //, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14],
            cursorCoords: [0, 0],
            lines: [
                ["i", "m", "p", "o", "r", "t", " ", "h", "a", "s", "h", "l", "i", "b"],
                [],
                ["d", "e", "f", " ", "h", "a", "s", "h", "_", "f", "i", "l", "e", "(", "f", "i", "l", "e", "n", "a", "m", "e", ")", ":"],
                [" ", " ", " ", " ", "h", " ", "=", " ", "h", "a", "s", "h", "l", "i", "b", ".", "s", "h", "a", "1", "(", ")"]
                // "\twith open(filename,'rb') as file:",
                // "\t\tchunk = 0",
                // "\t\twhile chunk != b'':",
                // "\t\t\tchunk = file.read(1024)",
                // "\t\t\th.update(chunk)",
                // "",
                // "\treturn h.hexdigest()",
                // "",
                // "message = hash_file(\"track1.mp3\")",
                // "print(message)"
            ],
        }
    }

    componentDidMount(): void {
        document.addEventListener("keydown", this.handleCursorMovement); 
    }

    handleCursorMovement = (event: KeyboardEvent): void => {
        switch (event.key) {
            case "ArrowRight": {
                if (this.state.cursorCoords[0] < this.state.lines[this.state.cursorCoords[1]].length) {
                    this.setState({ cursorCoords: [this.state.cursorCoords[0] + 1, this.state.cursorCoords[1]]});
                }
                break;
            }
            case "ArrowLeft": {
                if (this.state.cursorCoords[0] > 0) {
                    this.setState({ cursorCoords: [this.state.cursorCoords[0] - 1, this.state.cursorCoords[1]]});
                }
                break;
            }
            case "ArrowUp": {
                if (this.state.cursorCoords[1] > 0) {
                    this.setState({ cursorCoords: [this.state.cursorCoords[0], this.state.cursorCoords[1] - 1]});
                    if (this.state.lines[this.state.cursorCoords[1] - 1].length <= this.state.cursorCoords[0]) {
                        this.setState({ cursorCoords: [this.state.lines[this.state.cursorCoords[1] - 1].length, this.state.cursorCoords[1] - 1] })
                    }
                }
                break;
            }
            case "ArrowDown": {
                if (this.state.cursorCoords[1] < this.state.lineNumbers.length - 1) {
                    this.setState({ cursorCoords: [this.state.cursorCoords[0], this.state.cursorCoords[1] + 1]});
                    if (this.state.lines[this.state.cursorCoords[1] + 1].length <= this.state.cursorCoords[0]) {
                        this.setState({ cursorCoords: [this.state.lines[this.state.cursorCoords[1] + 1].length, this.state.cursorCoords[1] + 1]});
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
                    </div>
                </div>
            </>
        )
    }
}

export default CodeArea;