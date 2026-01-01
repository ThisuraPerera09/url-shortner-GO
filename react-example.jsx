// Example React Component for URL Shortener
// Copy this into your React project (e.g., src/components/URLShortener.jsx)

import React, { useState } from 'react';

const API_BASE = 'http://localhost:8080/api';

function URLShortener() {
  const [longUrl, setLongUrl] = useState('');
  const [customCode, setCustomCode] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [stats, setStats] = useState(null);

  const handleShorten = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setShortUrl('');

    try {
      const response = await fetch(`${API_BASE}/shorten`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          url: longUrl,
          custom_code: customCode || undefined,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || 'Failed to shorten URL');
      }

      setShortUrl(data.short_url);
      
      // Automatically get stats
      getStats(data.short_code);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const getStats = async (shortCode) => {
    try {
      const response = await fetch(`${API_BASE}/stats/${shortCode}`);
      const data = await response.json();
      
      if (response.ok) {
        setStats(data);
      }
    } catch (err) {
      console.error('Failed to fetch stats:', err);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(shortUrl);
    alert('Copied to clipboard!');
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>🔗 URL Shortener</h1>
      
      <form onSubmit={handleShorten} style={styles.form}>
        <div style={styles.inputGroup}>
          <label style={styles.label}>Long URL *</label>
          <input
            type="url"
            value={longUrl}
            onChange={(e) => setLongUrl(e.target.value)}
            placeholder="https://example.com/very/long/url"
            required
            style={styles.input}
          />
        </div>

        <div style={styles.inputGroup}>
          <label style={styles.label}>Custom Short Code (optional)</label>
          <input
            type="text"
            value={customCode}
            onChange={(e) => setCustomCode(e.target.value)}
            placeholder="my-custom-link"
            style={styles.input}
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          style={{
            ...styles.button,
            opacity: loading ? 0.6 : 1,
            cursor: loading ? 'not-allowed' : 'pointer',
          }}
        >
          {loading ? 'Shortening...' : 'Shorten URL'}
        </button>
      </form>

      {error && (
        <div style={styles.error}>
          ❌ {error}
        </div>
      )}

      {shortUrl && (
        <div style={styles.result}>
          <h2 style={styles.resultTitle}>✅ Success!</h2>
          <div style={styles.urlDisplay}>
            <a
              href={shortUrl}
              target="_blank"
              rel="noopener noreferrer"
              style={styles.link}
            >
              {shortUrl}
            </a>
            <button onClick={copyToClipboard} style={styles.copyButton}>
              📋 Copy
            </button>
          </div>

          {stats && (
            <div style={styles.stats}>
              <h3 style={styles.statsTitle}>📊 Statistics</h3>
              <div style={styles.statsGrid}>
                <div style={styles.statItem}>
                  <span style={styles.statLabel}>Clicks:</span>
                  <span style={styles.statValue}>{stats.clicks}</span>
                </div>
                <div style={styles.statItem}>
                  <span style={styles.statLabel}>Created:</span>
                  <span style={styles.statValue}>
                    {new Date(stats.created_at).toLocaleDateString()}
                  </span>
                </div>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

// Simple inline styles (you can move these to CSS modules or styled-components)
const styles = {
  container: {
    maxWidth: '600px',
    margin: '50px auto',
    padding: '30px',
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif',
    backgroundColor: '#f5f5f5',
    borderRadius: '10px',
    boxShadow: '0 2px 10px rgba(0,0,0,0.1)',
  },
  title: {
    textAlign: 'center',
    color: '#333',
    marginBottom: '30px',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',
  },
  inputGroup: {
    display: 'flex',
    flexDirection: 'column',
    gap: '8px',
  },
  label: {
    fontWeight: '600',
    color: '#555',
    fontSize: '14px',
  },
  input: {
    padding: '12px',
    fontSize: '16px',
    border: '2px solid #ddd',
    borderRadius: '6px',
    transition: 'border-color 0.3s',
    outline: 'none',
  },
  button: {
    padding: '14px',
    fontSize: '16px',
    fontWeight: '600',
    color: 'white',
    backgroundColor: '#4CAF50',
    border: 'none',
    borderRadius: '6px',
    transition: 'background-color 0.3s',
  },
  error: {
    marginTop: '20px',
    padding: '15px',
    backgroundColor: '#ffebee',
    color: '#c62828',
    borderRadius: '6px',
    borderLeft: '4px solid #c62828',
  },
  result: {
    marginTop: '30px',
    padding: '20px',
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 1px 5px rgba(0,0,0,0.1)',
  },
  resultTitle: {
    color: '#4CAF50',
    marginBottom: '15px',
    fontSize: '20px',
  },
  urlDisplay: {
    display: 'flex',
    gap: '10px',
    alignItems: 'center',
    marginBottom: '20px',
  },
  link: {
    flex: 1,
    padding: '12px',
    backgroundColor: '#f0f0f0',
    borderRadius: '6px',
    color: '#2196F3',
    textDecoration: 'none',
    fontWeight: '500',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
  },
  copyButton: {
    padding: '12px 20px',
    backgroundColor: '#2196F3',
    color: 'white',
    border: 'none',
    borderRadius: '6px',
    cursor: 'pointer',
    fontWeight: '500',
  },
  stats: {
    borderTop: '1px solid #eee',
    paddingTop: '15px',
  },
  statsTitle: {
    fontSize: '16px',
    color: '#555',
    marginBottom: '10px',
  },
  statsGrid: {
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
    gap: '15px',
  },
  statItem: {
    display: 'flex',
    flexDirection: 'column',
    gap: '5px',
  },
  statLabel: {
    fontSize: '12px',
    color: '#888',
    textTransform: 'uppercase',
    fontWeight: '600',
  },
  statValue: {
    fontSize: '18px',
    color: '#333',
    fontWeight: '700',
  },
};

export default URLShortener;

