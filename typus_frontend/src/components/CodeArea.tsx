import React, { useState, useEffect, useRef } from 'react'
import './styles/code-area.sass'
import { CodeCharacter, CodeLine, Cursor } from '../interfaces';
import isCodeSymbol from '../lib/isCodeSymbol';
import useCodeSample from '../hooks/useCodeSample';


interface Props {
    sampleId: string;
}


function CodeArea(props: Props): JSX.Element {
    const { status, codeSample, error } = useCodeSample({ sampleId: props.sampleId });
    const [lines, _setLines] = useState<CodeLine[]>([]);
    const [csr, _setCsr] = useState<Cursor>({ x: 0, y: 0});

    const lnsRef = useRef(lines);
    const csrRef = useRef(csr);

    const setLines = (data: CodeLine[]): void => {
        lnsRef.current = data;
        _setLines(data);
    }
    const setCsr = (data: Cursor): void => {
        csrRef.current = data;
        _setCsr(data);
    }

    useEffect(() => {
        if (status === 'success') {
            setLines(codeSample);
        }
        document.addEventListener("keydown", handleKeyboard);

        return () => {
            document.removeEventListener("keydown", handleKeyboard);
        }
    }, [codeSample]);

    const handleKeyboard = (event: KeyboardEvent): void => {
        if (isCodeSymbol(event.key) && csrRef.current.x < lnsRef.current[csrRef.current.y].chars.length) {
            const currentSymbolToType = lnsRef.current[csrRef.current.y].chars[csrRef.current.x].c;
            if (event.key === currentSymbolToType) {
                lnsRef.current[csrRef.current.y].chars[csrRef.current.x].wasTyped = true;
                setCsr({x: csrRef.current.x + 1, y: csrRef.current.y });
            }   
            return;
        } else if (event.key === "Enter" && csrRef.current.y < lnsRef.current.length - 1 && 
                    csrRef.current.x === lnsRef.current[csrRef.current.y].chars.length) {
            setCsr({ x: 0, y: csrRef.current.y + 1 });
            return;
        }  
    }


    return (
        <>
            <div className='code-area-wrapper'>
                {
                    lines?.map((line: CodeLine, lineNumber: number) => {
                        return (
                            <div className='line' key={lineNumber}>
                                <div className='line-number-wrapper' style={{ paddingTop: `${lineNumber ? 0 : '10px'}` }}>
                                    <span className='line-number'>
                                        { lineNumber + 1 }
                                    </span>
                                </div>
                                { error ? (
                                    <h1>An error ocurred: { error }</h1>
                                ) : (
                                    <div className='line-code-wrapper' style={{ paddingTop: `${lineNumber ? 0 : '10px'}` }}>
                                        {    
                                            lines[lineNumber].chars.map((char: CodeCharacter, charIndex: number) => {
                                                return (
                                                    <div style={{ display: 'flex' }} key={`${lineNumber}:${charIndex}`}>
                                                        <div style={{ display: 'flex'}}>
                                                            <span className='line-code' style={{ opacity: `${char.wasTyped ? '1' : '0.5'}` }}>
                                                                { char.c }
                                                            </span>
                                                            { csr.x === charIndex && csr.y === lineNumber ? (
                                                                <span className='cursor'></span>
                                                            ) : null }
                                                        </div>
                                                        { csr.x === lines[lineNumber].chars.length && 
                                                            charIndex + 1 === lines[lineNumber].chars.length &&
                                                            csr.y === lineNumber ? (
                                                            <span className='cursor' style={{ position: 'relative' }}></span>
                                                        ) : null }
                                                    </div>
                                                )
                                            })
                                        }
                                        { lines[lineNumber].chars.length === 0 && csr.y === lineNumber ? (
                                            <span className='cursor' style={{ position: 'relative' }}></span>
                                        ) : null }
                                    </div>
                                )}
                            </div>
                        )
                    })
                }
                <div className='filler'>
                    <div className='filler-line-numbers'></div>
                    <div className='filler-code'></div>
                </div>
            </div>
        </>
    )
}

export default CodeArea;