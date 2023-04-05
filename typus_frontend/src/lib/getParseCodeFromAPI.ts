import { CodeLine } from "../interfaces";
import testCodeSample from "../data/testCodeSample";


interface Props {
    /**
     * Parameters of the getParseCodeFromAPI function.
     * 
     * @param {string} exampleId - A unique id of the example the user wants to retrieve
     * @param {boolean} isTest - Determines if the test sample is needed [dev purposes]
     * 
     */
    exampleId?: string;
    isTest?: boolean;
}

/**
 * This function asynchronically retrieves data from API or just returns testCodeSample.
 * 
 * @param {Props} props - Function parameters. 
 * @returns {CodeLine[]} - An array of CodeLines from API or the hardcoded one, depending on the isTest param.
 */
const getParseCodeFromAPI = async (props: Props): Promise<any> => {
    if (props.isTest) {
        let sample: CodeLine[] = [];

        for (let i = 0; i < testCodeSample.length; i++) {
            let codeLine: CodeLine = { chars: [] };
            for (let j = 0; j < testCodeSample[i].length; j++) {
                codeLine.chars.push({ c: testCodeSample[i][j], wasTyped: false, isHighlighted: false });
            }
            sample.push(codeLine);
        }
        return sample;
    } else {
        // TODO: fetch data from API
    }
}

export default getParseCodeFromAPI
