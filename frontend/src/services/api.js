const API_BASE = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

export const api = {
  // Shorten a URL
  shortenURL: async (url, customCode = '') => {
    const response = await fetch(`${API_BASE}/shorten`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        url,
        custom_code: customCode || undefined,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || 'Failed to shorten URL');
    }

    return data;
  },

  // Get URL statistics
  getStats: async (shortCode) => {
    const response = await fetch(`${API_BASE}/stats/${shortCode}`);
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || 'Failed to fetch stats');
    }

    return data;
  },

  // List all URLs
  listURLs: async (limit = 50, offset = 0) => {
    const response = await fetch(`${API_BASE}/urls?limit=${limit}&offset=${offset}`);
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || 'Failed to fetch URLs');
    }

    return data;
  },

  // Delete a URL
  deleteURL: async (shortCode) => {
    const response = await fetch(`${API_BASE}/urls/${shortCode}`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || 'Failed to delete URL');
    }

    return true;
  },

  // Check API health
  checkHealth: async () => {
    const response = await fetch(`${API_BASE}/health`);
    return response.ok;
  },
};

export default api;

