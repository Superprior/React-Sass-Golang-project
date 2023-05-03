import React, { ChangeEvent, useState, useRef, useEffect } from 'react';
import './styles/register-window.sass'

interface Props {
	closeWindow: Function
}

function RegisterWindow(props: Props): JSX.Element {
	const [username, setUsername] = useState<string>("");
	const [email, setEmail] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	const _ref = useRef(null);

	useEffect(() => {
		window.onclick = (event: any) => {
			if (event.target.contains(_ref.current) && event.target !== _ref.current) {
				props.closeWindow(false);
			} 
		}
	}, []);

	const handleSubmit = async () => {
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

				props.closeWindow(false)
				alert('Registered!')
            } else {
                alert('Failed!');
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
                        placeholder='Username'
                        id='username' 
                        value={username} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<input 
                        type='email'
                        placeholder='Email'
                        id='email' 
                        value={email} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
					<input 
                        type='password'
                        placeholder='Password'
                        id='password' 
                        value={password} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className='form-field'
                    />
				</div>
				<div className='register-buttons-wrapper'>
					<button 
						type='submit'
						onClick={() => handleSubmit() }
						className='register-submit-button'
					>
						Register
					</button>
					<button
						onClick={() => props.closeWindow(false) }
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
