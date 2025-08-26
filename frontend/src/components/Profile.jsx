import React, { useEffect, useState } from "react";
import "../styles/Profile.css";
import Navbar from "./NavBar.jsx";
import illustration from "../assets/Profile.png";
import Footer from "./Footer.jsx";

const Profile = () => {
  const [profile, setProfile] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/profile", {
      credentials: "include",
    })
      .then((res) => {
        if (res.status === 401) {
          window.location.href = "/login";
          return null;
        }
        return res.json();
      })
      .then((data) => {
        console.log("Profile API response:", data);
        setProfile(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching profile:", err);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Loading...</p>;
  if (!profile || !profile.FirstName) return <p>No profile data available</p>;

  return (
    <div className="profile-page">
      <Navbar />
      <main className="profile-main">
        <div className="profile-container">
          <div className="profile-image">
            <img src={illustration} alt="Profile" />
          </div>

          <div className="profile-details">
            <p><strong>FirstName:</strong> {profile.FirstName}</p>
            <p><strong>LastName:</strong> {profile.LastName}</p>
            <p><strong>Email:</strong> {profile.Email}</p>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  );
};

export default Profile;
