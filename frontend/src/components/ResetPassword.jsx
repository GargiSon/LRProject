import React, { useState } from "react";
import "../styles/ResetPassword.css";
import resetImage from "../assets/reset.png";

const ResetPassword = () => {
  const [msg, setMsg] = useState("");
  const [msgClass, setMsgClass] = useState("");

  const token = new URLSearchParams(window.location.search).get("vtoken") ||
                new URLSearchParams(window.location.search).get("token");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMsg("Resetting password...");
    setMsgClass("msg");

    const formData = new FormData(e.target);
    const password = formData.get("password");
    const confirmPassword = formData.get("confirmPassword");

    try {
      const res = await fetch("/reset", {
        method: "POST",
        body: new URLSearchParams({
          token,
          password,
          confirmPassword
        }),
      });

      if (!res.ok) {
        const text = await res.text();
        setMsg(text || "Password reset failed.");
        setMsgClass("msg error");
      } else {
        setMsg("Password reset successful! Redirecting...");
        setMsgClass("msg success");
        setTimeout(() => {
          window.location.href = "/login";
        }, 1500);
      }
    } catch (err) {
      setMsg("Network error. Try again.");
      setMsgClass("msg error");
    }
  };

  return (
    <div className="reset-container">
      <div className="reset-left">
        <img src={resetImage} alt="Reset Password" />
      </div>
      <div className="reset-right">
        <form className="reset-form" onSubmit={handleSubmit}>
          <h2>Reset Your Password</h2>
          <p>Enter your new password below</p>

          <input
            type="password"
            name="password"
            placeholder="New Password"
            required
          />
          <input
            type="password"
            name="confirmPassword"
            placeholder="Confirm Password"
            required
          />

          <div id="msg" className={msgClass}>{msg}</div>

          <button type="submit" className="submit-btn">Change Password</button>
        </form>
      </div>
    </div>
  );
};

export default ResetPassword;
