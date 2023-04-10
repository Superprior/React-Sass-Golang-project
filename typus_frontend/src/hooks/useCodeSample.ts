import { useState, useEffect } from 'react'
import { CodeLine, CodeSamples } from '../interfaces';


interface Props {
    /**
     * Parameters of the useCodeSample custom hook.
     * 
     * @param {string}  exampleId - A unique id of the example the user wants to retrieve
     * @param {boolean} isTest    - Determines if the test sample is needed [dev purposes]
     * 
     */
    exampleId: string;
}

type State =
    | { status: 'idle', codeSample: null, error: null }
    | { status: 'loading', codeSample: null, error: null }
    | { status: 'success', codeSample: CodeLine[], error: null }
    | { status: 'error', codeSample: null, error: Error }



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
                        const sample = samples[0];
                        console.log(sample);
                    }
                    setState({ status: 'success', codeSample: [], error: null })
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