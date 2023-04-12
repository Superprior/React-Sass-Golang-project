import { useState, useEffect } from 'react'
import { CodeLine, CodeCharacter } from '../interfaces';
import { FetchState } from '../interfaces';


interface Props {
    /**
     * Parameters of the useCodeSample custom hook.
     * 
     * @property {string} exampleId - A unique id of the example the user wants to retrieve
     * 
     * @interface
     */

    sampleId: string;
}

/**
 * useCodeSample is a custom React hook forconvenient data retrieving from the samples API.
 * 
 * @param {Props} props - The hook props.
 * @returns {State} - An object containing the state of a request. 
 * 
 * @function
 */
function useCodeSample(props: Props): FetchState<CodeLine[]> {
    const [state, setState] = useState<FetchState<CodeLine[]>>({ status: 'idle', data: null, error: null });

    useEffect(() => {
        setState({ status: 'loading', data: null, error: null });

        const fetchParseData = async () => {
            try {
                const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/samples/${props.sampleId}`;
                const responseData = await fetch(url, {
                    method: 'GET',
                    mode: 'cors',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                
                if (responseData.status === 200) {
                    const sample = await responseData.json();
                    const lines: CodeLine[] = []

                    for (let i = 0; i < sample.Content.length; i++) {
                        const line: CodeLine = { chars: [] }
                        for (let j = 0; j < sample.Content[i].length; j++) {
                            const char: CodeCharacter = { c: sample.Content[i][j], wasTyped: false, isHighlighted: false };
                            line.chars.push(char);
                        }
                        lines.push(line);
                    }
                    setState({ status: 'success', data: lines, error: null })
                } else {
                    setState({ status: 'error', data: null, error: `Could not fulfill the request, code ${responseData.status}` })
                }
            } catch (error) {
                setState({ status: 'error', data: null, error: "Could not fetch data from API" })
            }
        }

        fetchParseData();
    }, []);

    return state;
}

export default useCodeSample;