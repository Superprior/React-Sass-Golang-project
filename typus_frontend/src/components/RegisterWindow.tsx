import React, { ChangeEvent, useState, useRef, useEffect } from 'react';
import './styles/register-window.sass'

interface Props {
	closeWindow: Function
	switchWindows: Function
}

function RegisterWindow(props: Props): JSX.Element {

	// Fields states
	const [username, setUsername] = useState<string>("");
	const [email, setEmail] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	// Field errors states
	const [usernameError, setUsernameError] = useState<string>("");
	const [emailError, setEmailError] = useState<string>("");
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
		if (email === "") { 
			setEmailError("Email cannot be empty"); 
			anyErrors = true; 
		} else { setEmailError(""); }
		if (password === "") { 
			setPasswordError("Password cannot be empty"); 
			anyErrors = true; 
		} else { setPasswordError(""); }

		if (anyErrors) return;

		try {
            const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/auth/register/`;
            const response = await fetch(url, {
                method: 'POST',
                body: JSON.stringify({
                    "Username": username,
					"Email": email,
                    "Password": password,
                })
            });

            if (response.status === 201) {
                const responseJSON = await response.json();
				setError("");
				props.closeWindow(false)
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
			case "email": {
				setEmail(value);
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
			<div className='register-wrapper' ref={_ref}>
				<h1 className='register-text'>Register</h1>
				<div className='register-form'>
					<input 
                        type='text'
						autoComplete='off'
                        placeholder='Username'
                        id='username' 
                        value={username} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<p className='form-error'>{ usernameError }&nbsp;</p>
					<input 
                        type='email'
						autoComplete='off'
                        placeholder='Email'
                        id='email' 
                        value={email} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<p className='form-error'>{ emailError }&nbsp;</p>
					<input 
                        type='password'
						autoComplete='off'
                        placeholder='Password'
                        id='password' 
                        value={password} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
				</div>
				<p className='form-error'>{ passwordError || error }&nbsp;</p>
				<div className='register-buttons-wrapper'>
					<button 
						type='submit'
						onClick={() => handleSubmit() }
						className='register-submit-button'
					>
						Register
					</button>
					<button
						onClick={() => props.switchWindows() }
						className='register-goto-signin-button'
					>
						Sign in
					</button>
				</div>
			</div>
		</div>
	);
}

export default RegisterWindow;
