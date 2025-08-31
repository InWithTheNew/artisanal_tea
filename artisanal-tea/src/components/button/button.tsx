import React from 'react';
import './button.css';

export interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  type?: 'button' | 'submit' | 'reset';
  disabled?: boolean;
}

export function Button({ children, onClick, type = 'button', disabled = false }: ButtonProps) {
  return (
    <button className="custom-button" onClick={onClick} type={type} disabled={disabled}>
      {children}
    </button>
  );
}
