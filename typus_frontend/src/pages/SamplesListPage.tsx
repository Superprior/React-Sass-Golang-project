import React, { useEffect, useState } from 'react'
import useCodeSamplesList from '../hooks/useCodeSamplesList';
import { SampleCard } from '../interfaces';
import Header from '../components/Header';
import SampleListCard from '../components/SampleCard';
import './styles/samples-list.sass';

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
		<div className='samples-list-page-wrapper'>
			<Header />
			<div className='samples-list-page-body'>
				<h1 className='heading-text'>All code samples</h1>
				<hr className='divider' />
				<div className='samples-list-page-top'>
					<h1 className='sample-title'>Title</h1>
					<h2 className='sample-language'>Language</h2>
				</div>
				{ status === 'success' ? (
						cards.map((card: SampleCard) => {
							return (
								<SampleListCard key={card.sampleId} sampleId={card.sampleId} title={card.title} language={card.langSlug} />
							)}
						)
				) : null }
			</div>
        </div>       
    )
}

export default SamplesPage;
