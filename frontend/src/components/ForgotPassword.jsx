import React, { useState } from 'react';
import '../styles/ForgotPassword.css';
import forgotImage from '../assets/forgot.png';

const ForgotPassword = () => {
  const [msg, setMsg] = useState('');
  const [msgClass, setMsgClass] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMsg('');
    setMsgClass('');

    const email = e.target.email.value;

    try {
      const res = await fetch('/forgot', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },  //Because go will accept email in this form only
        body: new URLSearchParams({ email })
      });

      const text = await res.text();

      if (!res.ok) {
        setMsg(text || 'Unable to send reset link');
        setMsgClass('msg error');
      } else {
        setMsg(text || 'If an account exists, a reset link has been sent.');
        setMsgClass('msg success');
        e.target.reset();
      }
    } catch (err) {
      setMsg('Network error. Try again.');
      setMsgClass('msg error');
    }
  };

  return (
    <div className="forgot-container">
      <div className="forgot-left">
        <img src={forgotImage} alt="Forgot password" />
      </div>

      <div className="forgot-right">
        <form onSubmit={handleSubmit} className="forgot-form">
          <h2>Forgot Password</h2>
          <p>Enter your email and we'll send you a link to reset your password.</p>

          <input
            type="email"
            name="email"
            placeholder="Enter your email"
            required
          />

          {msg && <div className={msgClass}>{msg}</div>}

          <button type="submit" className="submit-btn">Submit</button>

          <div className="form-links">
            <a href="/login">Back to Login</a>
            <span> | </span>
            <a href="/register">Don't have an account? Register</a>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ForgotPassword;
