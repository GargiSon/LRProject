import React from "react";
import "../styles/Contact.css";
import Navbar from "./NavBar.jsx"; 
import Footer from "./Footer.jsx";

const Contact = () => {
  return (
    <>
      <Navbar />
      <div className="contact-container">
        
        <h1 className="contact-title">Contact Us</h1>
        <p className="contact-subtitle">
          We'd love to hear from you. Fill out the form and we’ll get back to you
          as soon as possible.
        </p>

        <div className="contact-content">
          <div className="contact-info">
            <div className="info-box">
              <span className="info-icon">🏠</span>
              <div>
                <h3>Address</h3>
                <p>4671 Sugar Camp Road, Owatonna, Minnesota, 55060</p>
              </div>
            </div>

            <div className="info-box">
              <span className="info-icon">📞</span>
              <div>
                <h3>Phone</h3>
                <p>561-456-2321</p>
              </div>
            </div>

            <div className="info-box">
              <span className="info-icon">📧</span>
              <div>
                <h3>Email</h3>
                <p>example@email.com</p>
              </div>
            </div>
          </div>

          <div className="contact-form">
            <h2>Send Message</h2>
            <form>
              <input type="text" placeholder="Full Name" required />
              <input type="email" placeholder="Email" required />
              <textarea placeholder="Type your Message..." rows="5" required />
              <button type="submit">Send</button>
            </form>
          </div>
        </div>
      </div> 

      <Footer/>
    </>
  );
};

export default Contact;
