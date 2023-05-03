import React, { useState, useEffect } from 'react'
import './styles/header.sass'
import RegisterWindow from './RegisterWindow'
import SignInWindow from './SignInWindow';

		
function Header(): JSX.Element {
	const [isSignInWindowOpened, setIsSignInWindowOpened] = useState<boolean>(false);
	const [isRegisterWindowOpened, setIsRegisterWindowOpened] = useState<boolean>(false);
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

	useEffect(() => {
		if (localStorage.getItem("token") !== null) {
			setIsLoggedIn(true);
		} 
	}, []);

	const switchWindows = () => {
		setIsSignInWindowOpened(!isSignInWindowOpened);
		setIsRegisterWindowOpened(!isRegisterWindowOpened);
	}

    return (
		<div className='page'>
			{ isSignInWindowOpened ? (
				<SignInWindow closeWindow={setIsSignInWindowOpened} switchWindows={switchWindows}/>
			) : null }
			{ isRegisterWindowOpened ? (
				<RegisterWindow closeWindow={setIsRegisterWindowOpened} switchWindows={switchWindows}/>
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
					{ isLoggedIn ? (
						<button
							onClick={() => { localStorage.removeItem("token"); window.location.reload(); }}
							className='header-button' 
						>
							<h1 className='header-button-text'>
								LOG OUT
							</h1>
						</button>
					) : (
						<div style={{ display: 'flex' }}>
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
					)}
				</div>
			</div>
		</div>
    );
}

export default Header;
