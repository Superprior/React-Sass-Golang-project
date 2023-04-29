import React, { useState } from 'react'
import './styles/header.sass'
import RegisterWindow from './RegisterWindow'
import SignInWindow from './SignInWindow';

		
function Header(): JSX.Element {
	const [isSignInWindowOpened, setIsSignInWindowOpened] = useState(false);
	const [isRegisterWindowOpened, setIsRegisterWindowOpened] = useState(false);

    return (
		<div className='page'>
			{ isSignInWindowOpened ? (
				<SignInWindow closeWindow={setIsSignInWindowOpened} />
			) : null }
			{ isRegisterWindowOpened ? (
				<RegisterWindow closeWindow={setIsRegisterWindowOpened} />
			) : null }
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
						onClick={() => setIsSignInWindowOpened(true)}
						className='header-button'
					>
						<h1 className='header-button-text'>
							SIGN IN
						</h1>
					</button>
					<button
						onClick={() => setIsRegisterWindowOpened(true)}
						className='header-button'
					>
						<h1 className='header-button-text'>
							REGISTER
						</h1>
					</button>
				</div>
			</div>
		</div>
    );
}

export default Header;
