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
            <button
                onClick={() => window.location.replace(`${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:3000/register/`)}
                className='register-button'
            >
                <h1 className='register-button-text'>
                    Register
                </h1>
            </button>
        </div>
    );
}

export default Header;