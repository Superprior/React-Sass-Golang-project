import React from 'react'
import CodeArea from '../components/CodeArea'
import Header from '../components/Header'
import './styles/index.sass'


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {

    return (
        <>
            <Header />
            <CodeArea sampleId='1' />
        </>
    );
}

export default IndexPage;