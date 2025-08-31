import React, { useState } from 'react';
import './commandbar.css';

export function CommandBarField({ value, onChange, placeholder = 'Type a command...' }: { value: string; onChange: (v: string) => void; placeholder?: string }) {
  return (
    <input
      className="commandbar-field"
      type="text"
      placeholder={placeholder}
      value={value}
      onChange={e => onChange(e.target.value)}
    />
  );
}
