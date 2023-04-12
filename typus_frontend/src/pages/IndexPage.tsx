import React from 'react'
import CodeArea from '../components/CodeArea'
import CodeAreaHeader from '../components/CodeAreaHeader'
import './styles/index.sass'


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {

    return (
        <>
            <CodeAreaHeader />
            <CodeArea sampleId='1' />
        </>
    );
}

export default IndexPage;