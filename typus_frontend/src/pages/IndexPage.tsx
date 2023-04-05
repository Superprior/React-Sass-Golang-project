import React, { useState } from 'react'
import CodeArea from '../components/CodeArea'
import './styles/index.sass'


interface Props {}

const IndexPage: React.FC<{}> = (props: Props) => {

    return (
        <>
            <CodeArea />
        </>
    )
}

export default IndexPage;