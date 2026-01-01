import React, { useState, useEffect } from 'react';
import { api } from '../services/api';
import './URLList.css';

function URLList({ refresh, onViewStats }) {
  const [urls, setUrls] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [deletingId, setDeletingId] = useState(null);

  useEffect(() => {
    fetchURLs();
  }, [refresh]);

  const fetchURLs = async () => {
    setLoading(true);
    setError('');
    
    try {
      const data = await api.listURLs();
      setUrls(data.urls || []);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (shortCode) => {
    if (!window.confirm('Are you sure you want to delete this URL?')) {
      return;
    }

    setDeletingId(shortCode);
    try {
      await api.deleteURL(shortCode);
      setUrls(urls.filter(url => url.short_code !== shortCode));
    } catch (err) {
      alert('Failed to delete: ' + err.message);
    } finally {
      setDeletingId(null);
    }
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    alert('Copied to clipboard!');
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  if (loading) {
    return (
      <div className="url-list">
        <h2>📋 My URLs</h2>
        <div className="loading">
          <div className="spinner"></div>
          <p>Loading your URLs...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="url-list">
        <h2>📋 My URLs</h2>
        <div className="error-message">
          ❌ {error}
        </div>
        <button onClick={fetchURLs} className="retry-button">
          🔄 Retry
        </button>
      </div>
    );
  }

  if (urls.length === 0) {
    return (
      <div className="url-list">
        <h2>📋 My URLs</h2>
        <div className="empty-state">
          <div className="empty-icon">🔗</div>
          <h3>No URLs yet</h3>
          <p>Create your first short URL to get started!</p>
        </div>
      </div>
    );
  }

  return (
    <div className="url-list">
      <div className="list-header">
        <h2>📋 My URLs</h2>
        <button onClick={fetchURLs} className="refresh-button">
          🔄 Refresh
        </button>
      </div>

      <div className="urls-grid">
        {urls.map((url) => (
          <div key={url.short_code} className="url-card">
            <div className="url-card-header">
              <span className="short-code">{url.short_code}</span>
              <span className="clicks-badge">{url.clicks} clicks</span>
            </div>

            <div className="url-card-body">
              <div className="url-info">
                <label>Short URL:</label>
                <div className="url-value">
                  <a
                    href={`http://localhost:8080/${url.short_code}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="url-link"
                  >
                    localhost:8080/{url.short_code}
                  </a>
                  <button
                    onClick={() => copyToClipboard(`http://localhost:8080/${url.short_code}`)}
                    className="mini-button"
                    title="Copy"
                  >
                    📋
                  </button>
                </div>
              </div>

              <div className="url-info">
                <label>Original URL:</label>
                <div className="url-value original">
                  <a
                    href={url.original_url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="url-link"
                    title={url.original_url}
                  >
                    {url.original_url.length > 50
                      ? url.original_url.substring(0, 50) + '...'
                      : url.original_url}
                  </a>
                </div>
              </div>

              <div className="url-meta">
                <span>📅 {formatDate(url.created_at)}</span>
              </div>
            </div>

            <div className="url-card-footer">
              <button
                onClick={() => onViewStats && onViewStats(url)}
                className="stats-button"
              >
                📊 View Stats
              </button>
              <button
                onClick={() => handleDelete(url.short_code)}
                disabled={deletingId === url.short_code}
                className="delete-button"
              >
                {deletingId === url.short_code ? '⏳' : '🗑️'} Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default URLList;

