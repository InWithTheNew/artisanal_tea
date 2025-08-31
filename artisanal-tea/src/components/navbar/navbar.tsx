import React from 'react';
import './navbar.css';
import logo from '../../logo.svg';

export function Navbar() {
  return (
    <nav className="navbar">
      <img src={logo} alt="Logo" className="navbar-logo" />
    </nav>
  );
}
