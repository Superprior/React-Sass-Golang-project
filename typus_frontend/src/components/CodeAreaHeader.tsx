import React from 'react'
import { BiReset } from 'react-icons/bi'
import './styles/code-area-header.sass'


interface Props {}

const CodeAreaHeader: React.FC<{}> = (props: Props): JSX.Element => {
    return (
        <div className='code-area-header-wrapper'>
            <button className='reset-button' onClick={() => {}}>
                <BiReset size='24px' color='#B9B9B9' />
            </button>
        </div>
    );
}

export default CodeAreaHeader;