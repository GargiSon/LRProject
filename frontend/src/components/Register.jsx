import React, { useState } from 'react';
import '../styles/Register.css';
import illustration from '../assets/register.png';

const Register = () => {
  const [msg, setMsg] = useState('');
  const [msgClass, setMsgClass] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMsg('Creating account...');
    setMsgClass('msg');

    const formData = new FormData(e.target);
    const payload = {
      firstname: formData.get('firstname'),
      lastname: formData.get('lastname'),
      email: formData.get('email'),
      password: formData.get('password')
    };

    try {
      const res = await fetch('/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });
      const data = await res.json();

      if (!res.ok) {
        setMsg(data.Description || data.message || 'Registration failed');
        setMsgClass('msg error');
      } else {
        setMsg('Success! Account created. UID: ' + (data.uid || ''));
        setMsgClass('msg success');
        e.target.reset();
      }
    } catch (err) {
      setMsg('Network error. Try again.');
      setMsgClass('msg error');
    }
  };

  return (
    <div className="register-container">
      <div className="register-left">
        <div className="register-card">
          <h2>Create Account</h2>
          <form onSubmit={handleSubmit}>
            <label>
              First Name
              <input type="text" name="firstname" required />
            </label>
            <label>
              Last Name
              <input type="text" name="lastname" required />
            </label>
            <label>
              Email
              <input type="email" name="email" required />
            </label>
            <label>
              Password
              <input type="password" name="password" required />
            </label>

            <p className="agreement">
              By clicking <strong>"Create Account"</strong>, you agree to our terms and policies.
            </p>

            <button type="submit">Create Account</button>
            <div id="msg" className={msgClass}>{msg}</div>

            <div className="login-link">
              <a href="/login">Already have an account? Login</a>
            </div>
          </form>
        </div>
      </div>
      <div className="register-right">
        <img src={illustration} alt="Registration Illustration" />
      </div>
    </div>
  );
};

export default Register;
