import React from 'react'
import CodeArea from '../components/CodeArea'
import CodeAreaHeader from '../components/CodeAreaHeader'
import './styles/index.sass'
import { useParams } from 'react-router-dom'


interface Props {}

function SamplePage(props: Props): JSX.Element {
    const { id } = useParams();


    return (
        <>
            { id ? (
                <>
                    <CodeAreaHeader />
                    <CodeArea sampleId={id} />
                </>
            ) : null }
        </>
    );
}

export default SamplePage;