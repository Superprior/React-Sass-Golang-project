import React from 'react'
import { FetchState } from '../interfaces'


interface Props {
    title: string;
    langSlug: string;
    content: string;
}

const postCodeSample = async (props: Props): Promise<boolean> => {
    try {
        const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/samples/`;
        const responseData = await fetch(url, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "Title": props.title,
                "LangSlug": props.langSlug,
                "Content": props.content
            })
        });
        if (responseData.status === 201) {
            return true;
        } else {
            /**
             * @todo Handle errors
             */
            return false;
        }
    } catch (err) {
        return false;
    }
}

export default postCodeSample;