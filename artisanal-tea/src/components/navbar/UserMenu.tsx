import React, { useState, useRef, useEffect } from 'react';

import { Button } from '../button/button';

interface UserMenuProps {
  user: { name?: string; picture?: string } | null;
  onLogout: () => void;
}

export function UserMenu({ user, onLogout }: UserMenuProps) {

  const [open, setOpen] = useState(false);
  const menuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!open) return;
    function handleClickOutside(event: MouseEvent) {
      if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
        setOpen(false);
      }
    }
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [open]);

  if (!user) return null;
  return (
    <div className="user-menu" ref={menuRef}>
      <span className="user-greet">Welcome, {user.name}</span>
      <div className="user-avatar-dropdown">
        {user.picture && (
          <img
            src={user.picture}
            alt="avatar"
            className="user-avatar"
            onClick={() => setOpen(o => !o)}
            style={{ cursor: 'pointer' }}
          />
        )}
        {open && (
          <ul className="user-dropdown-menu">
            <li onClick={onLogout}>Logout</li>
          </ul>
        )}
      </div>
    </div>
  );
}
