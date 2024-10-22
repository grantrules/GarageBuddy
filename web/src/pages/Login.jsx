import React from 'react';

import './Login.css';

function Login() {

    const [values, setValues] = React.useState({ email: '', password: '' });

    const handleChange = name => (event) => { setValues({ ...values, [name]: event.target.value }); };

    const submitForm = (event) => {
        event.preventDefault();
        console.log(values);
    }
    
  return (
    <div className="Home">
        <h1>Login</h1>
        <form onSubmit={submitForm}>
            <input type="text" onChange={handleChange('email')} placeholder="Username" />
            <input type="password" onChange={handleChange('password')} placeholder="Password" />
            <button>Login</button>
        </form>

    </div>
  );
}

export default Login;
