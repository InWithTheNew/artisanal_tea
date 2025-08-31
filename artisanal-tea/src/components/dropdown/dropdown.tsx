import React, { useState, useRef, useEffect } from 'react';
import './dropdown.css';

export interface DropdownOption {
  label: string;
  value: string;
}

export interface DropdownProps {
  placeholder?: string;
  options: DropdownOption[];
  value?: string;
  onSelect?: (value: string) => void;
}

export function Dropdown({ placeholder = 'Select...', options, value, onSelect }: DropdownProps) {
  const [open, setOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!open) return;
    function handleClickOutside(event: MouseEvent) {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setOpen(false);
      }
    }
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [open]);

  const handleSelect = (v: string) => {
    if (onSelect) onSelect(v);
    setOpen(false);
  };

  return (
    <div className="dropdown" ref={dropdownRef}>
      <button className="dropdown-toggle" onClick={() => setOpen((o) => !o)}>
        {value ? options.find(o => o.value === value)?.label : placeholder}
        <span className="dropdown-arrow">▼</span>
      </button>
      {open && (
        <ul className="dropdown-menu">
          {options.map(option => (
            <li
              key={option.value}
              className={option.value === value ? 'selected' : ''}
              onClick={() => handleSelect(option.value)}
            >
              {option.label}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}