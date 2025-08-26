import React from "react";
import "../styles/About.css";

const About = () => {
  return (
    <div className="about-container">
      <section className="about-hero">
        <h1>About SkillVerse</h1>
        <p>Empowering learners, mentors, and innovators to build the future of skills.</p>
      </section>

      <section className="about-section">
        <div className="about-img">
          <img
            src="https://images.unsplash.com/photo-1521737604893-d14cc237f11d"
            alt="Who We Are"
          />
        </div>
        <div className="about-text">
          <h2>Who We Are</h2>
          <p>
            SkillVerse connects passionate learners and mentors in one place. 
            Whether you’re just starting out or sharing expertise, this is where curiosity meets opportunity.
          </p>
        </div>
      </section>

      <section className="about-section reverse">
        <div className="about-img">
          <img
            src="https://images.unsplash.com/photo-1522202176988-66273c2fd55f"
            alt="Mission"
          />
        </div>
        <div className="about-text">
          <h2>Our Mission</h2>
          <p>
            To empower individuals by providing a space to learn, practice, and showcase skills in a supportive community.
          </p>
        </div>
      </section>

      <section className="about-section">
        <div className="about-img">
          <img
            src="https://images.unsplash.com/photo-1529333166437-7750a6dd5a70"
            alt="Vision"
          />
        </div>
        <div className="about-text">
          <h2>Our Vision</h2>
          <p>
            To become a global hub for skill development — where anyone, anywhere can learn, teach, and grow together.
          </p>
        </div>
      </section>

      <section className="about-section reverse">
        <div className="about-img">
          <img
            src="https://images.unsplash.com/photo-1504384308090-c894fdcc538d"
            alt="Different"
          />
        </div>
        <div className="about-text">
          <h2>What Makes Us Different</h2>
          <div className="features">
            <div className="feature-card">🌍 Community-driven</div>
            <div className="feature-card">🚀 Practical Skills</div>
            <div className="feature-card">🤝 Collaborative Learning</div>
            <div className="feature-card">🎯 Personalized Growth</div>
          </div>
        </div>
      </section>

      <section className="about-cta">
        <h2>✨ Join Us</h2>
        <p>Explore, learn, and share your skills in the SkillVerse. 🚀</p>
        <button>Get Started</button>
      </section>
    </div>
  );
};

export default About;
