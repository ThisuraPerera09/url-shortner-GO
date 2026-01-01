import React, { useState, useEffect } from 'react';
import { api } from '../services/api';
import './Stats.css';

function Stats({ selectedURL }) {
  const [shortCode, setShortCode] = useState('');
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (selectedURL) {
      setShortCode(selectedURL.short_code);
      fetchStats(selectedURL.short_code);
    }
  }, [selectedURL]);

  const fetchStats = async (code) => {
    if (!code) return;

    setLoading(true);
    setError('');

    try {
      const data = await api.getStats(code);
      setStats(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    fetchStats(shortCode);
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'Never';
    return new Date(dateString).toLocaleString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  return (
    <div className="stats">
      <h2>📊 URL Analytics</h2>
      <p className="subtitle">View detailed statistics for your short URLs</p>

      <form onSubmit={handleSubmit} className="stats-form">
        <input
          type="text"
          value={shortCode}
          onChange={(e) => setShortCode(e.target.value)}
          placeholder="Enter short code (e.g., abc123)"
          className="stats-input"
        />
        <button type="submit" disabled={loading} className="stats-button">
          {loading ? '⏳ Loading...' : '🔍 Get Stats'}
        </button>
      </form>

      {error && (
        <div className="error-message">
          ❌ {error}
        </div>
      )}

      {loading && (
        <div className="loading">
          <div className="spinner"></div>
          <p>Fetching statistics...</p>
        </div>
      )}

      {stats && !loading && (
        <div className="stats-result">
          <div className="stats-header">
            <h3>Statistics for: {stats.short_code}</h3>
            <div className="url-badge">
              <a
                href={`http://localhost:8080/${stats.short_code}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                localhost:8080/{stats.short_code} ↗
              </a>
            </div>
          </div>

          <div className="stats-grid">
            <div className="stat-card primary">
              <div className="stat-icon">👆</div>
              <div className="stat-value">{stats.clicks}</div>
              <div className="stat-label">Total Clicks</div>
            </div>

            <div className="stat-card">
              <div className="stat-icon">📅</div>
              <div className="stat-value-small">{formatDate(stats.created_at)}</div>
              <div className="stat-label">Created On</div>
            </div>

            <div className="stat-card">
              <div className="stat-icon">🕐</div>
              <div className="stat-value-small">{formatDate(stats.last_accessed)}</div>
              <div className="stat-label">Last Accessed</div>
            </div>
          </div>

          <div className="original-url-section">
            <h4>🔗 Original URL</h4>
            <a
              href={stats.original_url}
              target="_blank"
              rel="noopener noreferrer"
              className="original-url-link"
            >
              {stats.original_url}
            </a>
          </div>

          <div className="stats-insights">
            <h4>📈 Insights</h4>
            <div className="insights-grid">
              <div className="insight">
                <span className="insight-label">Average daily clicks:</span>
                <span className="insight-value">
                  {calculateDailyAverage(stats.created_at, stats.clicks)}
                </span>
              </div>
              <div className="insight">
                <span className="insight-label">Age:</span>
                <span className="insight-value">
                  {calculateAge(stats.created_at)}
                </span>
              </div>
              <div className="insight">
                <span className="insight-label">Status:</span>
                <span className={`insight-value ${stats.clicks > 0 ? 'active' : 'inactive'}`}>
                  {stats.clicks > 0 ? '🟢 Active' : '🔵 Unused'}
                </span>
              </div>
            </div>
          </div>
        </div>
      )}

      {!stats && !loading && !error && (
        <div className="empty-stats">
          <div className="empty-icon">📊</div>
          <h3>Enter a short code to view statistics</h3>
          <p>You can find short codes in the "My URLs" tab</p>
        </div>
      )}
    </div>
  );
}

function calculateDailyAverage(createdAt, clicks) {
  const created = new Date(createdAt);
  const now = new Date();
  const daysOld = Math.max(1, Math.floor((now - created) / (1000 * 60 * 60 * 24)));
  const average = (clicks / daysOld).toFixed(1);
  return `${average} clicks/day`;
}

function calculateAge(createdAt) {
  const created = new Date(createdAt);
  const now = new Date();
  const diffMs = now - created;
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));
  const diffHours = Math.floor(diffMs / (1000 * 60 * 60));
  const diffMinutes = Math.floor(diffMs / (1000 * 60));

  if (diffDays > 0) return `${diffDays} day${diffDays > 1 ? 's' : ''}`;
  if (diffHours > 0) return `${diffHours} hour${diffHours > 1 ? 's' : ''}`;
  return `${diffMinutes} minute${diffMinutes > 1 ? 's' : ''}`;
}

export default Stats;

