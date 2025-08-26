import React from "react";
import { Link } from "react-router-dom";
import logo from "../assets/login.png";
import "../styles/NavBar.css";

const Navbar = () => {
  return (
    <nav className="navbar">
      <div className="nav-left">
        <img src={logo} alt="Skillverse Logo" className="nav-logo" />
        <span className="nav-title">Skillverse</span>
      </div>

      <div className="nav-links">
        <Link to="/home">Home</Link>
        <Link to="/about">About</Link>
        <Link to="/profile">You</Link>
        <Link to="/contact">Contact</Link>
      </div>

      <div className="nav-right">
        <Link to="/logout" className="logout-btn">Logout</Link>
      </div>
    </nav>
  );
};

export default Navbar;
