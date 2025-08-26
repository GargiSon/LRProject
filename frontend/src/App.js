import './App.css';
import {BrowserRouter as Router, Route, Routes, Navigate} from 'react-router-dom';
import Register from './components/Register';
import Login from './components/Login';
import Home from './components/Home';
import ForgotPassword from './components/ForgotPassword';
import ResetPassword from "./components/ResetPassword";
import Logout from './components/Logout';
import ProtectedRoute from './components/ProtectedRoute';
import Contact from './components/Contact';
import About from './components/About';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Navigate to="/register" />} />
        <Route path="/login" element={<Login />} />
        <Route path='/logout' element={<Logout/>} />
        <Route path="/register" element={<Register />} />
        <Route path="/home" element={<ProtectedRoute> <Home/> </ProtectedRoute>} />
        <Route path="/contact" element={<ProtectedRoute> <Contact/> </ProtectedRoute>} />
        <Route path="/about" element={<ProtectedRoute> <About/> </ProtectedRoute>} />
        <Route path="/forgot" element={<ForgotPassword />} />
        <Route path="/reset" element={<ResetPassword />} />
      </Routes>
    </Router>
  );
}

export default App;