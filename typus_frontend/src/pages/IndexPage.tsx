import React from 'react'
import CodeArea from '../components/CodeArea'
import CodeAreaHeader from '../components/CodeAreaHeader'
import './styles/index.sass'
import Header from '../components/Header'


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {

    return (
        <div className='index-page-wrapper'>
            <Header />
        </div>
    );
}

export default IndexPage;
