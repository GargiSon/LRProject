import React from "react";
import Navbar from "./NavBar.jsx"; 
import boyImage from "../assets/boy.png"; 
import "../styles/Home.css";
import backgroundVideo from "../assets/LP5.mp4";
import Footer from "./Footer.jsx";

const Home = () => {
  return (
    <div className="home-container">
      <video autoPlay loop muted className="background-video">
        <source src={backgroundVideo} type="video/mp4" />
        Your browser does not support the video tag.
      </video>

      <Navbar />

      <div className="home-content">
        <div className="hero-card">
          <div className="glow-circle"></div>
          <img src={boyImage} alt="Welcome to Skillverse" className="hero-image" />
        </div>
      </div>
      <Footer/>
    </div>
  );
};

export default Home;
