import { useState, useEffect } from 'react'
import { CodeLine } from '../interfaces';
import testCodeSample from '../data/testCodeSample';


interface Props {
    /**
     * Parameters of the useCodeSample custom hook.
     * 
     * @param {string}  exampleId - A unique id of the example the user wants to retrieve
     * @param {boolean} isTest    - Determines if the test sample is needed [dev purposes]
     * 
     */
    exampleId?: string;
    isTest?: boolean;
}

const useCodeSample = (props: Props): CodeLine[] => {
    const [codeSample, setCodeSample] = useState<CodeLine[]>([]);

    useEffect(() => {
        if (props.isTest) {
            let sample: CodeLine[] = [];

            for (let i = 0; i < testCodeSample.length; i++) {
                let codeLine: CodeLine = { chars: [] };
                for (let j = 0; j < testCodeSample[i].length; j++) {
                    codeLine.chars.push({ c: testCodeSample[i][j], wasTyped: false, isHighlighted: false });
                }
                sample.push(codeLine);
            }
            setCodeSample(sample);
        } else {
            // TODO: fetch data from API
        }
    }, []);
    return codeSample;
}

export default useCodeSample;