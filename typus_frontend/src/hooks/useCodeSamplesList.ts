import React, { useState, useEffect } from 'react'
import { FetchState, SampleCard } from '../interfaces';


function useCodeSamplesList(): FetchState<SampleCard[]> {
    const [state, setState] = useState<FetchState<SampleCard[]>>({ status: 'idle', data: null, error: null});

    useEffect(() => {
        setState({ status: 'loading', data: null, error: null })

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
                    const cards: SampleCard[] = []

                    for (let i = 0; i < samples.length; i++) {
                        cards.push({ sampleId: samples[i].ID.toString(), title: samples[i].Title, langSlug: samples[i].LangSlug })
                    }
                    setState({ status: 'success', data: cards, error: null })
                } else {
                    setState({ status: 'error', data: null, error: `Could not fulfill the request, code ${responseData.status}`})
                }
            } catch (error) {
                setState({ status: 'error', data: null, error: "Could not fetch data from API" })
            }
        }

        fetchParseData();
    }, []);

    return state;
}

export default useCodeSamplesList;