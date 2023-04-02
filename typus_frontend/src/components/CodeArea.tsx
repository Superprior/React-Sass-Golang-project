import React from 'react'
import './styles/code-area.sass'

interface Props {}

interface State {
    lineNumbers: number[];
    lines: string[];
}


class CodeArea extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            lineNumbers: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14],
            lines: [
                "import hashlib",
                "",
                "def hash_file(filename):",
                "\th = hashlib.sha1()",
                "\twith open(filename,'rb') as file:",
                "\t\tchunk = 0",
                "\t\twhile chunk != b'':",
                "\t\t\tchunk = file.read(1024)",
                "\t\t\th.update(chunk)",
                "",
                "\treturn h.hexdigest()",
                "",
                "message = hash_file(\"track1.mp3\")",
                "print(message)"
            ],
        }
    }

    componentDidMount(): void {
        document.addEventListener("keydown", this.handleCursorMovement); 
    }

    handleCursorMovement = (event: KeyboardEvent): void => {
        switch (event.key) {
            case "ArrowRight": {
                throw new Error('Not implemented yet');
                break;
            }
            case "ArrowLeft": {
                throw new Error('Not implemented yet');
                break;
            }
            case "ArrowUp": {
                throw new Error('Not implemented yet');
                break;
            }
            case "ArrowDown": {
                throw new Error('Not implemented yet');
                break;
            }
        }
    }

    render() {
        return (
            <>
                <div>
                    <div className='code-area-wrapper'>
                        <span className='cursor'></span>
                        {
                            this.state.lineNumbers.map((lineNumber: number) => {
                                return (
                                    <div className='line'>
                                        <div className='line-number-wrapper'>
                                            <span className='line-number'>
                                                { lineNumber }
                                            </span>
                                        </div>
                                        <div className='line-code-wrapper'>
                                            <span className='line-code'>
                                                { this.state.lines[lineNumber-1] }
                                            </span>
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