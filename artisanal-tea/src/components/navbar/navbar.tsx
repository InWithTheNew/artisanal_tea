import React from 'react';
import './navbar.css';
import logo from '../../logo.svg';
import { UserMenu } from './UserMenu';

export interface NavbarProps {
  user: { name?: string; picture?: string } | null;
  onLogout: () => void;
}

export function Navbar({ user, onLogout }: NavbarProps) {
  return (
    <nav className="navbar">
      <img src={logo} alt="Logo" className="navbar-logo" />
      <div className="navbar-right">
        <UserMenu user={user} onLogout={onLogout} />
      </div>
    </nav>
  );
}
