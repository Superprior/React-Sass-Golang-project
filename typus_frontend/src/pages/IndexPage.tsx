import React from 'react'
import CodeArea from '../components/CodeArea'
import CodeAreaHeader from '../components/CodeAreaHeader'
import './styles/index.sass'


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {

    return (
        <>
            <h1>Index Page</h1>
            <button 
                onClick={() => window.location.replace(`${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:3000/samples/`)}
            >
                Code samples list
            </button>
        </>
    );
}

export default IndexPage;