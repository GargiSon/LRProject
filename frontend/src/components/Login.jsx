import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import '../styles/Login.css';
import illustration from '../assets/login.png';

const Login = () => {
  const [msg, setMsg] = useState('');
  const [msgClass, setMsgClass] = useState('');
  const location = useLocation();

  useEffect(() => {
    const params = new URLSearchParams(location.search);
    if (params.get("status") === "loggedout") {
      setMsg("Logout successful!");
      setMsgClass("msg success");
    }
  }, [location]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMsg('Logging in...');
    setMsgClass('msg');

    const formData = new FormData(e.target);
    const payload = {
      email: formData.get('email'),
      password: formData.get('password')
    };

    try {
      const res = await fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
        credentials: "include",
      });
      const data = await res.json();

      if (!res.ok) {
        setMsg(data.Description || data.message || 'Login failed');
        setMsgClass('msg error');
      } else {
        setMsg('Login successful!');
        setMsgClass('msg success');
        e.target.reset();
        window.location.href = '/home';
      }
    } catch (err) {
      setMsg('Network error. Try again.');
      setMsgClass('msg error');
    }
  };

  return (
    <div className="login-container">
      <div className="login-left">
        <img src={illustration} alt="Login Illustration" />
      </div>

      <div className="login-right">
        <div className="login-card">
          <h2>Welcome back!</h2>
          <div className="login-header">
            <p>We're so excited to see you again</p>
          </div>

          <form onSubmit={handleSubmit}>
            <input type="email" name="email" placeholder="Email" required />
            <input type="password" name="password" placeholder="Password" required />

            <div className="form-links">
              <a href="/forgot">Forgot Password?</a>
            </div>

            <button type="submit">Login</button>
            {msg && <div id="msg" className={msgClass}>{msg}</div>}

            <div style={{ marginTop: '10px', textAlign: 'center' }}>
              <a href="/register">Need an Account? Register</a>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
