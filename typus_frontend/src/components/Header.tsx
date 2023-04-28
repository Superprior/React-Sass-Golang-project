import React from 'react'
import './styles/header.sass'

function Header(): JSX.Element {
    return (
        <div className='header-wrapper'>
            <a 
                href={`${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:3000/`} 
                className='logo-button'
            >
                <h1 className='logo-text'>
                    Typus
                </h1>
            </a>
            <div className='header-buttons'>
                <button
                onClick={() => {}}
                    className='header-button'
                >
                    <h1 className='header-button-text'>
                        sign in
                    </h1>
                </button>
                <button
                onClick={() => {}}
                    className='header-button'
                >
                    <h1 className='header-button-text'>
                        register
                    </h1>
                </button>
            </div>
        </div>
    );
}

export default Header;