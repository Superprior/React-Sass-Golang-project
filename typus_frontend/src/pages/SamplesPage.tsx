import React, { useEffect, useState } from 'react'
import useCodeSamplesList from '../hooks/useCodeSamplesList';
import { SampleCard } from '../interfaces';


interface Props {}

function SamplesPage(props: Props): JSX.Element {
    const { status, data, error } = useCodeSamplesList();
    const [cards, setCards] = useState<SampleCard[]>([]);

    useEffect(() => {
        if (status === 'success') {
            setCards(data);
        }
    }, [data]);

    return (
        <>
            <h1>All code samples</h1>
            { status === 'success' ? (
                    cards.map((card: SampleCard) => {
                        return (
                            <div style={{ paddingBottom: '20px' }}>
                                <div style={{ display: 'flex' }}>
                                    <p>{ card.sampleId }</p>
                                    <p style={{ paddingLeft: '30px' }}>{ card.langSlug }</p>
                                </div>
                                <h2 style={{ margin: '0' }}>{ card.title }</h2>
                            </div>
                        )}
                    )
            ) : null }
        </>       
    )
}

export default SamplesPage;
