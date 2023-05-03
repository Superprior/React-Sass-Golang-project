
import React, { ChangeEvent, useState, useRef, useEffect } from 'react';
import './styles/signin-window.sass'
    
interface Props {
	closeWindow: Function
	switchWindows: Function
}

function SignInWindow(props: Props): JSX.Element {

	// Fields states
	const [username, setUsername] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	// Field errors states
	const [usernameError, setUsernameError] = useState<string>("");
	const [passwordError, setPasswordError] = useState<string>("");

	// General error state
	const [error, setError] = useState<string>("");

	const _ref = useRef(null);

	useEffect(() => {
		window.onclick = (event: any) => {
			if (event.target.contains(_ref.current) && event.target !== _ref.current) {
				props.closeWindow(false);
			} 
		}
	}, []);

	const handleSubmit = async () => {
		let anyErrors = false;

		if (username === "") { 
			setUsernameError("Username cannot be empty"); 
			anyErrors = true; 
		} else { setUsernameError(""); }
		if (password === "") { 
			setPasswordError("Password cannot be empty"); 
			anyErrors = true; 
		} else { setPasswordError(""); }

		if (anyErrors) return;

		try {
            const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/auth/login/`;
            const response = await fetch(url, {
                method: 'POST',
                body: JSON.stringify({
                    "Username": username,
                    "Password": password,
                })
            });

            if (response.status === 200) {
                const responseJSON = await response.json();

                localStorage.setItem("token", responseJSON.token);
				props.closeWindow(false)
				window.location.replace(`${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:3000/samples/`);
            } else {
				const responseJSON = await response.json();
				if (responseJSON.Error !== null) {
					setError(responseJSON.Error);
				}
            }
        } catch (err: any) {
            console.log(err);
        }
	}

	const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
		const { id, value } = e.target;

		switch (id) {
			case "username": {
				setUsername(value);
				return;
			}
			case "password": {
				setPassword(value);
				return;
			}
		}
	}

	return (
		<div className='dark-bg'>
			<div className='signin-wrapper' ref={_ref}>
				<h1 className='signin-text'>Sign in</h1>
				<div className='signin-form'>
					<input 
                        type='text'
                        placeholder='Username'
                        id='username' 
                        value={username} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<p className='form-error'>{ usernameError }&nbsp;</p>
					<input 
                        type='password'
                        placeholder='Password'
						id='password'
                        value={password} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<p className='form-error'>{ passwordError || error }&nbsp;</p>
				</div>
				<div className='signin-buttons-wrapper'>
					<button 
						type='submit'
						onClick={() => handleSubmit() }
						className='signin-submit-button'
					>
						Sign in
					</button>
					<button
						onClick={() => props.switchWindows() }
						className='signin-goto-register-button'
					>
						Register
					</button>
				</div>
			</div>
		</div>
	);
}

export default SignInWindow;
