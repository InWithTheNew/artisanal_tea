import React, { useState } from 'react';
import './searchbar.css';

export interface SearchbarProps {
  placeholder?: string;
  onSearch?: (query: string) => void;
}

export function Searchbar({ placeholder = 'Search...', onSearch }: SearchbarProps) {
  const [query, setQuery] = useState('');

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (onSearch) {
      onSearch(query);
    }
  };

  return (
    <form className="searchbar" onSubmit={handleSubmit}>
      <input
        type="text"
        className="searchbar-input"
        placeholder={placeholder}
        value={query}
        onChange={handleInputChange}
      />
      <button className="searchbar-button" type="submit">🔍</button>
    </form>
  );
}
