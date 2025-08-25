import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import '../styles/Home.css';

const Home = () => {
  // useEffect(() => {
  //   // Reload if coming back from bfcache
  //   const handlePageShow = (event) => {
  //     if (event.persisted) {
  //       window.location.reload();
  //     }
  //   };

  //   window.addEventListener('pageshow', handlePageShow);
  //   return () => {
  //     window.removeEventListener('pageshow', handlePageShow);
  //   };
  // }, []);

  return (
    <div className="home-container">
      <h1>Welcome!</h1>
      <Link to="/logout" className="logout-link">Logout</Link>
    </div>
  );
};

export default Home;
