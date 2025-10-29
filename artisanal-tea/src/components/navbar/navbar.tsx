import React from 'react';
import './navbar.css';
import logo from '../../logo.svg';
import { UserMenu } from './UserMenu';

export interface NavbarProps {
  showAdminButton?: boolean;
  user: { name?: string; picture?: string } | null;
  onLogout: () => void;
  onAdminClick?: () => void
}

export function Navbar({ showAdminButton, user, onLogout, onAdminClick }: NavbarProps) {
  return (
    <nav className="navbar">
      <img src={logo} alt="Logo" className="navbar-logo" />
      <div className="navbar-right">
        {showAdminButton && (
          <button className="admin-button" onClick={onAdminClick}>
            Admin
          </button>
        )}
        <UserMenu user={user} onLogout={onLogout} />
      </div>
    </nav>
  );
}
