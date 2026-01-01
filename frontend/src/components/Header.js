import React from 'react';
import './Header.css';

function Header() {
  return (
    <header className="header">
      <div className="header-content">
        <h1 className="logo">🔗 URL Shortener</h1>
        <p className="tagline">Fast, Simple, Powerful</p>
      </div>
    </header>
  );
}

export default Header;

