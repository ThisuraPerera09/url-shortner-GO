import React, { useState } from 'react';
import { api } from '../services/api';
import './URLShortener.css';

// Helper to get base URL (without /api)
const getBaseURL = () => {
  const apiUrl = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';
  // Remove /api from the end
  return apiUrl.replace(/\/api$/, '');
};

function URLShortener({ onURLCreated }) {
  const [longUrl, setLongUrl] = useState('');
  const [customCode, setCustomCode] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [copied, setCopied] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setShortUrl('');
    setCopied(false);

    try {
      const data = await api.shortenURL(longUrl, customCode);
      // Construct URL using same logic as URLList to ensure consistency
      // This also fixes any double-slash issues from backend
      const shortUrlToDisplay = `${getBaseURL()}/${data.short_code}`;
      setShortUrl(shortUrlToDisplay);
      setLongUrl('');
      setCustomCode('');
      
      // Save to localStorage for "My URLs" tracking
      const myUrls = JSON.parse(localStorage.getItem('myUrls') || '[]');
      if (!myUrls.includes(data.short_code)) {
        myUrls.push(data.short_code);
        localStorage.setItem('myUrls', JSON.stringify(myUrls));
      }
      
      if (onURLCreated) {
        onURLCreated();
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(shortUrl);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="url-shortener">
      <h2>✨ Shorten Your URL</h2>
      <p className="subtitle">Create short, memorable links in seconds</p>

      <form onSubmit={handleSubmit} className="form">
        <div className="input-group">
          <label>Long URL *</label>
          <input
            type="url"
            value={longUrl}
            onChange={(e) => setLongUrl(e.target.value)}
            placeholder="https://example.com/very/long/url/path"
            required
            className="input"
          />
        </div>

        <div className="input-group">
          <label>Custom Short Code (optional)</label>
          <input
            type="text"
            value={customCode}
            onChange={(e) => setCustomCode(e.target.value)}
            placeholder="my-custom-link"
            className="input"
            pattern="[a-zA-Z0-9-_]+"
            title="Only letters, numbers, hyphens, and underscores"
          />
          <small>Leave empty for a random code</small>
        </div>

        <button
          type="submit"
          disabled={loading}
          className="submit-button"
        >
          {loading ? '⏳ Shortening...' : '🚀 Shorten URL'}
        </button>
      </form>

      {error && (
        <div className="error-message">
          ❌ {error}
        </div>
      )}

      {shortUrl && (
        <div className="result">
          <h3>✅ Success!</h3>
          <div className="url-display">
            <a
              href={shortUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="short-link"
            >
              {shortUrl}
            </a>
            <button onClick={copyToClipboard} className="copy-button">
              {copied ? '✓ Copied!' : '📋 Copy'}
            </button>
          </div>
          <p className="success-note">Your short URL is ready to share!</p>
        </div>
      )}
    </div>
  );
}

export default URLShortener;

