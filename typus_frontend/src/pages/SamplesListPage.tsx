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
                            <div 
                                style={{ marginBottom: '20px', backgroundColor: 'gray', width: '500px', cursor: 'pointer' }}
                                onClick={() => window.location.replace(`${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:3000/samples/${card.sampleId}`)}
                                key={card.sampleId}
                            >
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
