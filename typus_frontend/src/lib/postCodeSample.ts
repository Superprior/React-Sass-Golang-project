import React from 'react'
import { FetchState } from '../interfaces'


interface Props {
    title: string;
    langSlug: string;
    content: string;
}

const postCodeSample = async (props: Props): Promise<boolean> => {
    try {
        console.log(`${props.title} ${props.langSlug} ${props.content}`)
        const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/samples`;
        const responseData = await fetch(url, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "title": props.title,
                "langSlug": props.langSlug,
                "content": props.content
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