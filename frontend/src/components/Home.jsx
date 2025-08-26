import React from "react";
import { Link } from "react-router-dom";
import logo from "../assets/login.png";
import boyImage from "../assets/boy.png"; 
import "../styles/Home.css";
import backgroundVideo from "../assets/LP5.mp4";

const Home = () => {
  return (
    <div className="home-container">
      <video autoPlay loop muted className="background-video">
        <source src={backgroundVideo} type="video/mp4" />
        Your browser does not support the video tag.
      </video>

      <nav className="navbar">
        <div className="nav-left">
          <img src={logo} alt="Skillverse Logo" className="nav-logo" />
          <span className="nav-title">Skillverse</span>
        </div>

        <div className="nav-links">
          <Link to="/home">Home</Link>
          <Link to="/about">About</Link>
          <Link to="/you">You</Link>
          <Link to="/contact">Contact</Link>
        </div>

        <div className="nav-right">
          <Link to="/logout" className="logout-btn">Logout</Link>
        </div>
      </nav>

      <div className="home-content">
        <div className="hero-card">
          <div className="glow-circle"></div>
          <img src={boyImage} alt="Welcome to Skillverse" className="hero-image" />
        </div>
      </div>
    </div>
  );
};

export default Home;
