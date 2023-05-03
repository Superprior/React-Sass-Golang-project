
import React, { ChangeEvent, useState, useRef, useEffect } from 'react';
import './styles/signin-window.sass'
    
interface Props {
	closeWindow: Function
}

function SignInWindow(props: Props): JSX.Element {
	const [username, setUsername] = useState<string>("");
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
			case "password": {
				setPassword(value);
				return;
			}
		}
	}

	return (
		<div className='dark-bg'>
			<div className='signin-wrapper' ref={_ref}>
				<h1 className='login-text'>Sign in</h1>
				<div className='login-form'>
					<input 
                        type='text'
                        placeholder='Username'
                        id='username' 
                        value={username} 
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
				<div className='login-submit-button-wrapper'>
					<button 
						type='submit'
						onClick={() => handleSubmit() }
						className='login-submit-button'
					>
						Sign in
					</button>
				</div>
			</div>
		</div>
	);
}

export default SignInWindow;
