import { useState, useEffect } from 'react'
import { CodeLine, CodeCharacter } from '../interfaces';

interface Props {
    /**
     * Parameters of the useCodeSample custom hook.
     * 
     * @property {string} exampleId - A unique id of the example the user wants to retrieve
     * 
     * @interface
     */

    exampleId: string;
}

/**
 * The state of the useCodeSample custom React hook.
 * 
 * @property {string} status                - The current state of data fetching/parsing process.
 * @property {CodeLine[] | null} codeSample - Data that was already fetched and parsed.
 * @property {Error | string} error         - Storing possible errors.
 * 
 * @typedef {Object} State
 */
type State =
    | { status: 'idle', codeSample: null, error: null }
    | { status: 'loading', codeSample: null, error: null }
    | { status: 'success', codeSample: CodeLine[], error: null }
    | { status: 'error', codeSample: null, error: Error | string }


/**
 * useCodeSample is a custom React hook forconvenient data retrieving from the samples API.
 * 
 * @param {Props} props - The hook props.
 * @returns {State} - An object containing the state of a request. 
 * 
 * @function
 */
function useCodeSample(props: Props): State {
    const [state, setState] = useState<State>({ status: 'idle', codeSample: null, error: null });

    useEffect(() => {
        setState({ status: 'loading', codeSample: null, error: null });

        const fetchParseData = async () => {
            try {
                const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/samples`;
                const responseData = await fetch(url, {
                    method: 'GET',
                    mode: 'cors',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                if (responseData.status === 200) {
                    const samples = await responseData.json();
                    if (samples.length > 0) {
                        const sampleContent: string[] = samples[0].Content.split("\\n");
                        for (let i = 0; i < sampleContent.length; i++) {
                            sampleContent[i] = sampleContent[i].replace(/\\t/g, "    ");
                        }

                        const lines: CodeLine[] = []

                        for (let i = 0; i < sampleContent.length; i++) {
                            const line: CodeLine = { chars: [] }
                            for (let j = 0; j < sampleContent[i].length; j++) {
                                const char: CodeCharacter = { c: sampleContent[i][j], wasTyped: false, isHighlighted: false };
                                line.chars.push(char);
                            }
                            lines.push(line);
                        }
                        setState({ status: 'success', codeSample: lines, error: null })
                    } else {
                        setState({ status: 'error', codeSample: null, error: "" })
                    }
                } else {
                    setState({ status: 'error', codeSample: null, error: "Could not fulfill the request" })
                }
            } catch (error: any) {
                setState({ status: 'error', codeSample: null, error: error })
            }
        }

        fetchParseData();
    }, []);

    return state;
}

export default useCodeSample;