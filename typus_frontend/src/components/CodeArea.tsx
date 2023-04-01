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
                "\n",
                "def hash_file(filename):",
                "\th = hashlib.sha1()",
                "\twith open(filename,'rb') as file:",
                "\t\tchunk = 0",
                "\t\twhile chunk != b'':",
                "\t\t\tchunk = file.read(1024)",
                "\t\t\th.update(chunk)",
                "\n",
                "\treturn h.hexdigest()",
                "\n",
                "message = hash_file(\"track1.mp3\")",
                "print(message)"
            ],
        }
    }

    render() {
        return (
            <>
                <div>
                    <h1>This is the code area</h1>
                    <div className='code-area-wrapper'>
                        <div className='line-numbering'>
                            {
                                this.state.lineNumbers.map((lineNumber: number) => {
                                    return (
                                        <pre>
                                            <span className='line-number'>
                                                { lineNumber }
                                            </span>
                                        </pre>
                                    )
                                })
                            }
                        </div>
                        <div className='code-area'>
                            {
                                this.state.lines.map((line: string) => {
                                    return (
                                        <pre>
                                            <span className='code-line'>
                                                { line }
                                            </span>
                                        </pre>
                                    )
                                })
                            }
                        </div>
                    </div>
                </div>
            </>
        )
    }
}

export default CodeArea;